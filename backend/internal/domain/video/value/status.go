package value

type Status string

const (
	StatusUploaded   Status = "uploaded"   // アップロード完了
	StatusQueued     Status = "queued"     // 非同期処理待ち
	StatusProcessing Status = "processing" // 非同期処理中
	StatusReady      Status = "ready"      // 再生可能
	StatusFailed     Status = "failed"     // 処理失敗
)
