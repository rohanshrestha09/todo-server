package enums

type Provider int

const (
	Facebook Provider = iota
	Google
)

type Model string

const (
	USER    Model = "user"
	PROJECT Model = "project"
	REQUEST Model = "request"
)

type TodoStatus string

const (
	IN_PROGRESS TodoStatus = "IN_PROGRESS"
	COMPLETED   TodoStatus = "COMPLETED"
)

type FileType string

const (
	IMAGE FileType = "image"
)

type FileDIR string

const (
	USER_DIR FileDIR = "users/"
	TODO_DIR FileDIR = "todos/"
)
