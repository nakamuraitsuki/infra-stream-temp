package video

type PlaybackInfo struct {
	StreamKey string // URL 組み立てはHandler側で行う（リダイレクト）
	MIMEType  string
}