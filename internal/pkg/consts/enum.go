package consts

type RunType string

const (
	FromServer RunType = "server"
	FromAgent  RunType = "agent"
)

func (e RunType) String() string {
	return string(e)
}
