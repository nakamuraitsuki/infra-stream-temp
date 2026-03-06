package ffmpeg

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

const (
	// TODO: UPLOAD_WORKER_POOL_SIZE は環境変数などで設定できるようにすることも検討
	UPLOAD_WORKER_POOL_SIZE   = 4
	PATHS_CHANNEL_BUFFER_SIZE = UPLOAD_WORKER_POOL_SIZE * 2 // ワーカーが待機する分のバッファを確保
)

func (t *ffmpegTranscoder) Transcode(
	ctx context.Context,
	sourceKey string,
	streamKey string,
) error {
	tmpDir, err := os.MkdirTemp("", "transcode-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	sourcePath := filepath.Join(tmpDir, "source.mp4")
	if err := t.downloadSource(ctx, sourceKey, sourcePath); err != nil {
		return err
	}

	pathCh := make(chan string, PATHS_CHANNEL_BUFFER_SIZE)
	eg, gCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		log.Println("Starting worker pool...")
		err := t.workerPool(gCtx, UPLOAD_WORKER_POOL_SIZE, pathCh, streamKey)
		log.Println("Worker pool stopped", err)
		return nil // workerPool内でエラーはログに出力しているので、ここでは常にnilを返す
	})

	eg.Go(func() error {
		// この関数が終わった時点で Ch に詰まれたものだけを処理する
		defer close(pathCh)

		playlistPath := filepath.Join(tmpDir, "index.m3u8")
		segmentPattern := filepath.Join(tmpDir, "segment_%03d.ts")
		cmd := exec.CommandContext(gCtx, "ffmpeg",
			"-threads", "1", // 本当は最適割当をしたいが、とりあえず制限
			"-i", sourcePath,
			"-c:v", "libx264", "-c:a", "aac",
			"-f", "hls",
			"-hls_time", "6",
			"-hls_playlist_type", "vod",
			"-hls_flags", "temp_file+independent_segments",
			"-hls_segment_filename", segmentPattern,
			playlistPath,
		)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("ffmpeg error: %w", err)
		}

		// まず .ts をすべて探す
		files, err := filepath.Glob(filepath.Join(tmpDir, "*.ts"))
		if err != nil {
			return err
		}

		for _, f := range files {
			select {
			case pathCh <- f:
			case <-gCtx.Done():
				return err
			}
		}

		select {
		case pathCh <- playlistPath:
		case <-gCtx.Done():
			return err
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		// context が終了されている可能性を考慮
		deleteCtx := context.WithoutCancel(ctx)
		_ = t.storage.DeleteStream(deleteCtx, streamKey)
		return err
	}

	return nil
}

func (t *ffmpegTranscoder) downloadSource(ctx context.Context, key, destPath string) error {
	srcReader, err := t.storage.GetSource(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to get source: %w", err)
	}
	defer srcReader.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create local source file: %w", err)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcReader); err != nil {
		return fmt.Errorf("failed to save source to local: %w", err)
	}

	return nil
}
