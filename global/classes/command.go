package classes

const (
	TypeSequential = 0
	TypeParallel   = 1
	TypeWatch      = 2
)

type Command struct {
	Type int
	Run  []string
}
