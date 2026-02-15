package ffmpeg

import (
	"context"
	"fmt"
	"io"
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

	playlistPath := filepath.Join(tmpDir, "index.m3u8")
	pathCh := make(chan string, PATHS_CHANNEL_BUFFER_SIZE)
	eg, gCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return t.workerPool(gCtx, UPLOAD_WORKER_POOL_SIZE, pathCh, streamKey)
	})

	eg.Go(func() error {
		// この関数が終わった時点で Ch に詰まれたものだけを処理する
		defer close(pathCh)

		vEg, vCtx := errgroup.WithContext(gCtx)
		vCtx, cancel := context.WithCancel(vCtx)
		defer cancel()

		vEg.Go(func() error {
			return t.watchAndQueue(vCtx, tmpDir, pathCh)
		})

		vEg.Go(func() error {
			defer cancel() // ffmpegが終わったらwatcherの loop も終了させる
			segmentPattern := filepath.Join(tmpDir, "segment_%03d.ts")

			cmd := exec.CommandContext(gCtx, "ffmpeg",
				"-i", sourcePath,
				"-c:v", "libx264", "-c:a", "aac",
				"-f", "hls",
				"-hls_time", "10",
				"-hls_playlist_type", "vod",
				"-hls_flags", "temp_file",
				"-hls_segment_filename", segmentPattern,
				playlistPath,
			)
			if out, err := cmd.CombinedOutput(); err != nil {
				return fmt.Errorf("ffmpeg error: %w, output: %s", err, string(out))
			}

			return nil
		})

		return vEg.Wait()
	})

	if err := eg.Wait(); err != nil {
		// context が終了されている可能性を考慮
		deleteCtx := context.WithoutCancel(ctx)
		_ = t.storage.DeleteStream(deleteCtx, streamKey)
		return err
	}

	err = t.uploadFile(ctx, playlistPath, streamKey)
	if err != nil {
		// context が終了されている可能性を考慮
		deleteCtx := context.WithoutCancel(ctx)
		_ = t.storage.DeleteStream(deleteCtx, streamKey)
		return fmt.Errorf("failed to upload playlist: %w", err)
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
