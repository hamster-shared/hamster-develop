package consts

const (
	PIPELINE_DIR_NAME       = "pipelines"
	JOB_DIR_NAME            = "jobs"
	JOB_DETAIL_DIR_NAME     = "job-details"
	JOB_DETAIL_LOG_DIR_NAME = "job-details-log"
)

const (
	LANG_EN = "en"
	LANG_ZH = "zh"
)

const (
	TRIGGER_MODE = "Manual trigger"
)

const (
	ArtifactoryName = "/artifactory"
	ArtifactoryDir  = PIPELINE_DIR_NAME + "/" + JOB_DIR_NAME
)

const (
	IpfsUploadUrl = "https://api.ipfs-gateway.cloud/upload"
	CarVersion    = 1
)

var InkUrlMap = map[string]string{
	"Local":   "ws://127.0.0.1:9944",
	"Rococo":  "wss://rococo-contracts-rpc.polkadot.io",
	"Shibuya": "wss://rpc.shibuya.astar.network",
	"Shiden":  "wss://rpc.shiden.astar.network",
}
