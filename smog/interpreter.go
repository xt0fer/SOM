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
	stackPointer *OOObject
	bytecodeIndex int 
	localOffset int 
	method *OOObject
	context *OOObject
	previousFrame *OOObject
	stack *OOObject
}

func NewFrame() *Frame {
	f := &Frame{}

	return f
}