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
)

// jian-guo
//const ClientSecrets = "a20f12e3ab2bd7228476d167a54a7ed279121955"

// 34.232.105.81
const ClientSecrets = "2923a2870379c3b7237d2703852828017ca1de9b"
