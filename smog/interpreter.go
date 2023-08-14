package smog

type Interpreter struct {
	universe *Universe
	frame    *Frame
}

func NewInterpreter(u *Universe) *Interpreter {
	ii := &Interpreter{}
	ii.universe = u
	return ii
}

// "
// Frame layout:
// +-----------------+
// | Arguments       | 1
// +-----------------+
// | Local Variables | <-- localOffset
// +-----------------+
// | Stack           | <-- stackPointer
// | ...             |
// +-----------------+
// "
// |
//   "Points at the top element"
//   stackPointer
//   bytecodeIndex

//   "the offset at which local variables start"
//   localOffset

//   method
//   context
//   previousFrame
//   stack
// |

type Frame struct {
	stackPointer  Object
	bytecodeIndex int
	localOffset   int
	method        Object
	context       Object
	previousFrame Object
	stack         Object
}

func NewFrame() *Frame {
	f := &Frame{}

	return f
}
