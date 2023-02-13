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

// jian-guo
//const ClientSecrets = "a20f12e3ab2bd7228476d167a54a7ed279121955"

// 34.232.105.81
//const ClientSecrets = "2923a2870379c3b7237d2703852828017ca1de9b"

// https://develop.alpha.hamsternet.io/
//const ClientSecrets = "c99eef44205a6dfe975a62556f0601957dc3df9c"

// https://develop.test.hamsternet.io/
const ClientSecrets = "968331f48983b1521c8cb58ba78db313bb0143ce"

type ProjectType uint

const (
	CONTRACT ProjectType = iota + 1
	FRONTEND
	BLOCKCHAIN
)

const (
	TemplateOwner    = "hamster-template"
	TemplateRepoName = "truffle-frame"
)

const IpfsUrl = "http://183.66.65.247:32509"
