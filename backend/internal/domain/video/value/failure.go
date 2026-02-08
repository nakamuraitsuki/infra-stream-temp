package value

type FailureReason string

const (
	FailureTranscode FailureReason = "transcode_failure" // トランスコード失敗
	FailureUnknown   FailureReason = "unknown_failure"   // 不明な失敗
)