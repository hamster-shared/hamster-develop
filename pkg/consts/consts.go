package consts

const SecretKey = "93dfe293a9c897c795a7e4ee737e5734"

type TemplateType int

const (
	Solidity TemplateType = iota + 1
	Ink
	Move
	Vue
	Nuxt
	Next
	Vite
	Angular
)

type WorkflowType int

const (
	Check WorkflowType = iota + 1
	Build
	Deploy
)

type FrontendDeployType int

const (
	IPFS FrontendDeployType = iota + 1
	CONTAINER
)

// jian-guo
//const ClientSecrets = "a20f12e3ab2bd7228476d167a54a7ed279121955"

// 34.232.105.81
//const ClientSecrets = "2923a2870379c3b7237d2703852828017ca1de9b"

// https://develop.alpha.hamsternet.io/
//const ClientSecrets = "c99eef44205a6dfe975a62556f0601957dc3df9c"

// https://develop.test.hamsternet.io/
const ClientSecrets = "968331f48983b1521c8cb58ba78db313bb0143ce"

// test
const (
	AppsClientId      = "Iv1.6d9972fa6afd1c02"
	AppsClientSecrets = "90bb54dd864a215b860b933f705801f043e287a2"
)

// al
//const (
//	AppsClientId      = "Iv1.84a628b1689aab9d"
//	AppsClientSecrets = "9c5ffca3481fd02c6520e57486bd7948338089d0"
//)

type ProjectType uint

const (
	CONTRACT ProjectType = iota + 1
	FRONTEND
	BLOCKCHAIN
)

const (
	TemplateOwner    = "hamster-template"
	TemplateRepoName = "truffle-frame"
	TemplateUrl      = "https://github.com/hamster-template/truffle-frame.git"
)

type ProjectFrameType uint

// project frame type
const (
	Evm uint = iota + 1
	Aptos
	Ton
	StarkWare
)

const IpfsUrl = "http://183.66.65.247:32509"

const (
	STATUS_RUNNING = 1
	STATUS_SUCCESS = 2
	STATUS_FAIL    = 3
)

const RepositoryDir = "repository"

const (
	PIPELINE_DIR_NAME = "pipelines"
	JOB_DIR_NAME      = "jobs"
	WORK_DIR_NAME     = "workdir"
)

const DockerHubName = "registry.g.develop.hamsternet.io"
const Gateway = "authright.sh.newtouch.com"
