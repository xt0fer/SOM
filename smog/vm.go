package smog

import (
	"log"
	"os"
)

type Universe struct {
	symbolTable   map[string]*Symbol
	globals       map[*Symbol]*Object
	interpreter   *Interpreter
	dumpBytecodes bool
	avoidExit     bool

	systemObject *Object // main Universe??
	//
	NilObject   *Object
	TrueObject  *Object
	FalseObject *Object

	//
	NilClass       *Class
	ObjectClass    *Class
	ClassClass     *Class
	SystemClass    *Class
	MetaclassClass *Class
	BlockClass     *Class
	MethodClass    *Class
	PrimitiveClass *Class
	ArrayClass     *Class
	DoubleClass    *Class
	IntegerClass   *Class
	StringClass    *Class
	SymbolClass    *Class

	TrueClass  *Class
	FalseClass *Class
}

func NewUniverse() *Universe {
	nu := &Universe{}
	nu.initialize()
	return nu
}

func (u *Universe) initialize() {
	u.symbolTable = make(map[string]*Symbol)
	u.globals = make(map[*Symbol]*Object)
	u.dumpBytecodes = false
	u.avoidExit = false
	u.interpreter = NewInterpreter(u)
	u.initializeObjectSystem()

}

// UNIVERSE

func (u *Universe) Exit(code int) {
	os.Exit(code)
}
func (u *Universe) Interpret(args []string) {

}

func (u *Universe) symbolFor(name string) *Symbol {
	if result, ok := u.symbolTable[name]; ok {
		return result
	}
	ns := NewSymbol(name, 0)
	u.symbolTable[name] = ns
	return ns
}

func (u *Universe) NewMetaclassClass() *Class {
	result := NewClass(0, u)
	result.SetSomClass(NewClass(0, u))
	result.SomClass().SetSomClass(result)
	return result
}

func (u *Universe) NewSystemClass() *Class {
	systemClass := NewClass(0, u)
	systemClass.SetSomClass(NewClass(0, u))
	systemClass.SomClass().SetSomClass(u.MetaclassClass)
	return systemClass
}

// initializeSystemClass: systemClass superClass: superClass name: name = (
func (u *Universe) InitializeSystemClass(systemClass *Class, superClass *Class, name string) {
	// "Initialize the superclass hierarchy"
	if superClass != nil {
		systemClass.setSuperClass(superClass)
		systemClass.SomClass().setSuperClass(superClass.SomClass())
	} else {
		systemClass.SomClass().setSuperClass(u.ClassClass)
	}

	// "Initialize the array of instance fields"
	systemClass.SetInstancesFields(0)
	systemClass.SomClass().SetInstancesFields(0)

	// "Initialize the array of instance invokables"
	//
	systemClass.SetInstanceInvokables(0)
	systemClass.SomClass().SetInstanceInvokables(0)

	// "Initialize the name of the system class"
	//
	systemClass.SetName(u.symbolFor(name))
	systemClass.SomClass().SetName(u.symbolFor(name + " class"))

	// "Insert the system class into the dictionary of globals"
	u.setGlobal(systemClass.Name, systemClass.Object)

}

func (u *Universe) LoadSystemClass(sc *Class) {
	// | result |
	// "Load the system class"
	// result := self loadClass: systemClass name into: systemClass.

	// "Load primitives if necessary"
	// result hasPrimitives ifTrue: [
	//   result loadPrimitives ].
}

func (u *Universe) setGlobal(sym *Symbol, obj *Object) {
	u.globals[sym] = obj
}

func (u *Universe) initializeObjectSystem() *Object {
	//     | trueSymbol falseSymbol systemObject |

	// "Allocate the nil object"
	u.NilObject = &Object{}

	// "Allocate the Metaclass classes"
	u.MetaclassClass = u.NewMetaclassClass()
	//     "Allocate the rest of the system classes"
	u.ObjectClass = u.NewSystemClass()
	u.NilClass = u.NewSystemClass()
	u.ClassClass = u.NewSystemClass()
	u.ArrayClass = u.NewSystemClass()
	u.SymbolClass = u.NewSystemClass()
	u.MethodClass = u.NewSystemClass()
	u.IntegerClass = u.NewSystemClass()
	u.PrimitiveClass = u.NewSystemClass()
	u.StringClass = u.NewSystemClass()
	u.DoubleClass = u.NewSystemClass()
	u.SystemClass = u.NewSystemClass()

	//     "Setup the class reference for the nil object"
	u.NilObject.SetSomClass(u.NilClass)
	//     "Initialize the system classes."
	u.InitializeSystemClass(u.ObjectClass, u.NilClass, "Object")
	u.InitializeSystemClass(u.ClassClass, u.ObjectClass, "Class")
	u.InitializeSystemClass(u.MetaclassClass, u.ClassClass, "Metaclass")
	u.InitializeSystemClass(u.NilClass, u.ObjectClass, "Nil")
	u.InitializeSystemClass(u.ArrayClass, u.ObjectClass, "Array")
	u.InitializeSystemClass(u.MethodClass, u.ArrayClass, "Method")
	u.InitializeSystemClass(u.StringClass, u.ObjectClass, "String")
	u.InitializeSystemClass(u.SymbolClass, u.StringClass, "Symbol")
	u.InitializeSystemClass(u.IntegerClass, u.ObjectClass, "Integer")
	u.InitializeSystemClass(u.PrimitiveClass, u.ObjectClass, "Primitive")
	u.InitializeSystemClass(u.DoubleClass, u.ObjectClass, "Double")
	u.InitializeSystemClass(u.SystemClass, u.ObjectClass, "System")

	//     "Load methods and fields into the system classes"
	u.LoadSystemClass(u.ObjectClass)
	u.LoadSystemClass(u.ClassClass)
	u.LoadSystemClass(u.MetaclassClass)
	u.LoadSystemClass(u.NilClass)
	u.LoadSystemClass(u.ArrayClass)
	u.LoadSystemClass(u.MethodClass)
	u.LoadSystemClass(u.SymbolClass)
	u.LoadSystemClass(u.IntegerClass)
	u.LoadSystemClass(u.PrimitiveClass)
	u.LoadSystemClass(u.StringClass)
	u.LoadSystemClass(u.DoubleClass)
	u.LoadSystemClass(u.SystemClass)

	//     "Fix up objectClass"
	//     objectClass superClass: nilObject.
	u.ObjectClass.setSuperClass(u.NilClass)
	//     "Load the generic block class"
	//     blockClass := self loadClass: (self symbolFor: 'Block').

	//u.BlockClass.LoadClass(u.symbolFor("Block"))
	u.BlockClass = NewClass(0, u)
	u.BlockClass.SetName(u.symbolFor("Block"))

	//     "Setup the true and false objects"
	trueSymbol := u.symbolFor("True")
	//     trueClass := self loadClass: trueSymbol.
	u.TrueClass = NewClass(0, u)
	u.TrueClass.SetName(trueSymbol)
	u.TrueObject = u.NewInstance(u.TrueClass)
	falseSymbol := u.symbolFor("False")
	//     falseClass := self loadClass: falseSymbol.
	u.FalseClass = NewClass(0, u)
	u.FalseClass.SetName(falseSymbol)
	u.FalseObject = u.NewInstance(u.FalseClass)

	//     "Load the system class and create an instance of it"
	u.SystemClass = NewClass(0, u)
	u.SystemClass.SetName(u.symbolFor("System"))
	u.systemObject = u.NewInstance(u.SystemClass)

	//     "Put special objects and classes into the dictionary of globals"
	u.setGlobal(u.symbolFor("nil"), u.NilObject)
	u.setGlobal(u.symbolFor("true"), u.TrueObject)
	u.setGlobal(u.symbolFor("false"), u.FalseObject)
	u.setGlobal(u.symbolFor("system"), u.systemObject)
	u.setGlobal(u.symbolFor("System"), u.SystemClass.Object)
	u.setGlobal(u.symbolFor("Block"), u.BlockClass.Object)
	u.setGlobal(trueSymbol, u.TrueClass.Object)
	u.setGlobal(falseSymbol, u.FalseClass.Object)
	return u.systemObject
}

func (u *Universe) NewInstance(c *Class) *Object {
	result := NewObject(c.NumberOfInstanceFields(), u.NilObject)
	result.SetSomClass(c)
	return result
}

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
	StackPointer  int32
	BytecodeIndex int32
	LocalOffset   int32
	Method        *Method
	ContextObj    *Frame
	PreviousFrame *Frame
	Stack         []*Object
}

func NewFrame() *Frame {
	f := &Frame{}
	return f
}

func (f *Frame) Initialize(aNil *Object, prevFrame *Frame, contextFrame *Frame, aMethod *Method, maxStack int32) {
	f.PreviousFrame = prevFrame
	f.ContextObj = contextFrame
	f.Method = aMethod
	f.Stack = make([]*Object, maxStack)
	f.ResetStackPointer()
	f.BytecodeIndex = 1 // should be Zero?
}

func (f *Frame) GetPreviousFrame() *Frame { return f.PreviousFrame }

func (f *Frame) ClearPreviousFrame()    { f.PreviousFrame = nil }
func (f *Frame) HasPreviousFrame() bool { return f.PreviousFrame != nil }
func (f *Frame) IsBootstrapFrame() bool { return !f.HasPreviousFrame() }

func (f *Frame) GetContext() *Frame {
	return f.ContextObj
}

func (f *Frame) HasContext() bool { return f.ContextObj != nil }

// "Get the context frame at the given level"
func (f *Frame) GetContextAt(level int32) *Frame {
	frame := f
	// "Iterate through the context chain until the given level is reached"
	for level > 0 {
		// "Get the context of the current frame"
		frame = f.GetContext()
		// "Go to the next level"
		level = level - 1
	}
	return frame
}

// "Compute the outer context of this frame"
func (f *Frame) OuterContext() *Frame {
	frame := f
	//     "Iterate through the context chain until null is reached"
	for frame.HasContext() {
		frame = frame.GetContext()
	}
	return frame
}

func (f *Frame) GetMethod() *Method {
	return f.Method
}

// "Pop an object from the expression stack and return it"
func (f *Frame) Pop() *Object {
	sp := f.StackPointer
	f.StackPointer -= 1
	return f.Stack[sp]
}

// "Push an object onto the expression stack"
func (f *Frame) Push(obj *Object) {
	sp := f.StackPointer + 1
	f.Stack[sp] = obj
	f.StackPointer = sp
}

func (f *Frame) ResetStackPointer() {
	// "arguments are stored in front of local variables"
	f.LocalOffset = int32(len(f.Method.Array.Fields) + 1)
	// "Set the stack pointer to its initial value thereby clearing the stack"
	f.StackPointer = f.LocalOffset + f.Method.NumberOfLocals - 1
}

//	"Get the current bytecode index for this frame"
//
// OR just call f.BytecodeIndex
func (f *Frame) GetBytecodeIndex() int32 {
	return f.BytecodeIndex
}

// "Set the current bytecode index for this frame"
func (f *Frame) SetBytecodeIndex(index int32) {
	f.BytecodeIndex = index
}

// "Get the stack element with the given index
//
//	(an index of zero yields the top element)"
func (f *Frame) GetStackElement(index int32) *Object {
	return f.Stack[index]
}

// "Set the stack element with the given index to the given value
//
//	(an index of zero yields the top element)"
func (f *Frame) PutStackElement(index int32, value *Object) {
	f.Stack[f.StackPointer-index] = value
}

// Locals
func (f *Frame) GetLocal(index int32) *Object {
	return f.Stack[f.LocalOffset+index+1]
}
func (f *Frame) PutLocal(index int32, value *Object) {
	f.Stack[f.LocalOffset+index-1] = value
}

func (f *Frame) GetLocalAt(index int32, level int32) *Object {
	return f.GetContextAt(level).GetLocal(index)
}
func (f *Frame) PutLocalAt(index int32, level int32, value *Object) {
	f.GetContextAt(level).PutLocal(index, value)
}

// Arguments
func (f *Frame) GetArgument(index int32) *Object {
	return f.Stack[index]
}
func (f *Frame) PutArgument(index int32, value *Object) {
	f.Stack[index] = value
}

func (f *Frame) GetArgumentAt(index int32, level int32) *Object {
	return f.GetContextAt(level).GetArgument(index)
}
func (f *Frame) PutArgumentAt(index int32, level int32, value *Object) {
	f.GetContextAt(level).PutArgument(index, value)
}

//     "copy arguments from frame:
//      - arguments are at the top of the stack of frame.
//      - copy them into the argument area of the current frame"
func (f *Frame) CopyArgumentsFrom(frame *Frame) {
	numArgs := len(f.Method.Array.Fields)
	for i := 0; i < numArgs - 1; i++ {
		f.Stack[i+1] = frame.GetStackElement(int32(numArgs - 1 - i))
	}
}

func (f *Frame) PrintStackTrace() {
//     "Print a stack trace starting in this frame"
	if f.HasPreviousFrame() {
		f.PreviousFrame.PrintStackTrace()
	}
	className := f.Method.Holder.Clazz.Name
	methodName := f.Method.Signature.Name
	log.Printf("%s>>#%s @bi: %d\n",className,methodName,f.BytecodeIndex)
}

