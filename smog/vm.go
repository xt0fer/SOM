package smog

import (
	"fmt"
	"log"
	"os"
)

type Universe struct {
	symbolTable   map[string]*Symbol
	Globals       map[*Symbol]*Object
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
	u.Globals = make(map[*Symbol]*Object)
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
	result.Clazz = NewClass(0, u)
	result.Clazz.Clazz = result
	return result
}

func (u *Universe) NewSystemClass() *Class {
	systemClass := NewClass(0, u)
	systemClass.Clazz = NewClass(0, u)
	systemClass.Clazz.Clazz = u.MetaclassClass
	return systemClass
}

// initializeSystemClass: systemClass superClass: superClass name: name = (
func (u *Universe) InitializeSystemClass(systemClass *Class, superClass *Class, name string) {
	// "Initialize the superclass hierarchy"
	if superClass != nil {
		systemClass.setSuperClass(superClass)
		systemClass.Clazz.setSuperClass(superClass.Clazz)
	} else {
		systemClass.Clazz.setSuperClass(u.ClassClass)
	}

	// "Initialize the array of instance fields"
	systemClass.SetInstancesFields(0)
	systemClass.Clazz.SetInstancesFields(0)

	// "Initialize the array of instance invokables"
	//
	systemClass.SetInstanceInvokables(0)
	systemClass.Clazz.SetInstanceInvokables(0)

	// "Initialize the name of the system class"
	//
	systemClass.Name = u.symbolFor(name)
	systemClass.Clazz.Name = u.symbolFor(name + " class")

	// "Insert the system class into the dictionary of globals"
	u.setGlobal(systemClass.Name, &systemClass.Object)

}

func (u *Universe) LoadSystemClass(sc *Class) {
	// | result |
	// "Load the system class"
	// result := self loadClass: systemClass name into: systemClass.

	// "Load primitives if necessary"
	// result hasPrimitives ifTrue: [
	//   result loadPrimitives ].
}

func (u *Universe) LoadClass(name *Symbol, sc *Class) *Class {
	//     "Try loading the class from all different paths"
	panic("No class loaded for " + name.Name)
	//return nil
}

// loadClass: name into: systemClass = (
//     "Try loading the class from all different paths"
//     classPath do: [:cpEntry |
//       | result |
//       "Load the class from a file and return the loaded class"
//       result := SourcecodeCompiler compileClass: cpEntry name: name string into: systemClass in: self.

//       (result notNil and: dumpBytecodes) ifTrue: [
//         Disassembler dump: result somClass in: self.
//         Disassembler dump: result in: self ].

//       result ifNotNil: [ ^ result ] ].

//     "The class could not be found."
//     ^ nil
//   )

func (u *Universe) setGlobal(sym *Symbol, obj *Object) {
	u.Globals[sym] = obj
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
	u.NilObject.Clazz = u.NilClass
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
	u.BlockClass.Name = u.symbolFor("Block")

	//     "Setup the true and false objects"
	trueSymbol := u.symbolFor("True")
	//     trueClass := self loadClass: trueSymbol.
	u.TrueClass = NewClass(0, u)
	u.TrueClass.Name = trueSymbol
	u.TrueObject = u.NewInstance(u.TrueClass)
	falseSymbol := u.symbolFor("False")
	//     falseClass := self loadClass: falseSymbol.
	u.FalseClass = NewClass(0, u)
	u.FalseClass.Name = falseSymbol
	u.FalseObject = u.NewInstance(u.FalseClass)

	//     "Load the system class and create an instance of it"
	u.SystemClass = NewClass(0, u)
	u.SystemClass.Name = u.symbolFor("System")
	u.systemObject = u.NewInstance(u.SystemClass)

	//     "Put special objects and classes into the dictionary of globals"
	u.setGlobal(u.symbolFor("nil"), u.NilObject)
	u.setGlobal(u.symbolFor("true"), u.TrueObject)
	u.setGlobal(u.symbolFor("false"), u.FalseObject)
	u.setGlobal(u.symbolFor("system"), u.systemObject)
	u.setGlobal(u.symbolFor("System"), &u.SystemClass.Object)
	u.setGlobal(u.symbolFor("Block"), &u.BlockClass.Object)
	u.setGlobal(trueSymbol, &u.TrueClass.Object)
	u.setGlobal(falseSymbol, &u.FalseClass.Object)
	return u.systemObject
}

func (u *Universe) NewInstance(c *Class) *Object {
	result := NewObject(c.NumberOfInstanceFields(), u.NilObject)
	result.Clazz = c
	return result
}

// newBlock: method with: context numArgs: arguments = (
//
//	  ^ SBlock new: method in: context with: (self blockClass: arguments)
//	)
func (u *Universe) NewBlock(m *Method, context *Frame, numArgs int32) *Block {
	result := NewBlock(m, context, u.blockClass(numArgs))
	return result
}

func (u *Universe) blockClass(numArgs int32) *Class {
	name := u.symbolFor("Block" + string(numArgs))
	if g, ok := u.Globals[name]; ok {
		return g.Clazz
	}
	result := u.LoadClass(name, nil)
	// "Add the appropriate value primitive to the block class"
	//result.AddInstancePrimitive()
	//       (SBlock evaluationPrimitive: numberOfArguments in: self).

	u.Globals[name] = &result.Object
	return result

}

// blockClass: numberOfArguments = (
//     | name result |
//     "Determine the name of the block class with the given number of arguments"
//     name := self symbolFor: 'Block' + numberOfArguments.

//     "Lookup the block class in the dictionary of globals and return it"
//     (self hasGlobal: name) ifTrue: [
//       ^ self global: name ].

//     result := self loadClass: name into: nil.

//     "Add the appropriate value primitive to the block class"
//     result addInstancePrimitive:
//       (SBlock evaluationPrimitive: numberOfArguments in: self).

//     self global: name put: result.
//     ^ result
//   )

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
	Context       *Frame
	PreviousFrame *Frame
	Stack         []interface{}
}

func NewFrame() *Frame {
	f := &Frame{}
	return f
}

func (f *Frame) Initialize(aNil *Object, prevFrame *Frame, contextFrame *Frame, aMethod *Method, maxStack int32) {
	f.PreviousFrame = prevFrame
	f.Context = contextFrame
	f.Method = aMethod
	f.Stack = make([]interface{}, maxStack)
	f.ResetStackPointer()
	f.BytecodeIndex = 1 // should be Zero?
}

//func (f *Frame) GetPreviousFrame() *Frame { return f.PreviousFrame }

func (f *Frame) ClearPreviousFrame()    { f.PreviousFrame = nil }
func (f *Frame) HasPreviousFrame() bool { return f.PreviousFrame != nil }
func (f *Frame) IsBootstrapFrame() bool { return !f.HasPreviousFrame() }

// func (f *Frame) GetContext() *Frame {
// 	return f.Context
// }

func (f *Frame) HasContext() bool { return f.Context != nil }

// "Get the context frame at the given level"
func (f *Frame) ContextAt(level int32) *Frame {
	frame := f
	// "Iterate through the context chain until the given level is reached"
	for level > 0 {
		// "Get the context of the current frame"
		frame = f.Context
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
		frame = frame.Context
	}
	return frame
}

// func (f *Frame) GetMethod() *Method {
// 	return f.Method
// }

// "Pop an object from the expression stack and return it"
func (f *Frame) Pop() *Object {
	sp := f.StackPointer
	f.StackPointer -= 1
	return f.Stack[sp].(*Object)
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
	f.StackPointer = f.LocalOffset + f.Method.NumLocals - 1
}

//	"Get the current bytecode index for this frame"
//
// OR just call f.BytecodeIndex
// func (f *Frame) GetBytecodeIndex() int32 {
// 	return f.BytecodeIndex
// }

// "Set the current bytecode index for this frame"
func (f *Frame) SetBytecodeIndex(index int32) {
	f.BytecodeIndex = index
}

// "Get the stack element with the given index
//
//	(an index of zero yields the top element)"
// func (f *Frame) GetStackElement(index int32) *Object {
// 	return f.Stack[index]
// }

// "Set the stack element with the given index to the given value
//
//	(an index of zero yields the top element)"
func (f *Frame) PutStackElement(index int32, value *Object) {
	f.Stack[f.StackPointer-index] = value
}

// Locals
func (f *Frame) Local(index int32) *Object {
	return f.Stack[f.LocalOffset+index+1].(*Object)
}
func (f *Frame) PutLocal(index int32, value *Object) {
	f.Stack[f.LocalOffset+index-1] = value
}

func (f *Frame) LocalAt(index int32, level int32) *Object {
	return f.ContextAt(level).Local(index)
}
func (f *Frame) PutLocalAt(index int32, level int32, value *Object) {
	f.ContextAt(level).PutLocal(index, value)
}

// Arguments
func (f *Frame) Argument(index int32) *Object {
	return f.Stack[index].(*Object)
}
func (f *Frame) PutArgument(index int32, value *Object) {
	f.Stack[index] = value
}

func (f *Frame) ArgumentAt(index int32, level int32) *Object {
	return f.ContextAt(level).Argument(index)
}
func (f *Frame) PutArgumentAt(index int32, level int32, value *Object) {
	f.ContextAt(level).PutArgument(index, value)
}

// "copy arguments from frame:
//   - arguments are at the top of the stack of frame.
//   - copy them into the argument area of the current frame"
func (f *Frame) CopyArgumentsFrom(frame *Frame) {
	numArgs := len(f.Method.Array.Fields)
	for i := 0; i < numArgs-1; i++ {
		f.Stack[i+1] = frame.Stack[int32(numArgs-1-i)]
	}
}

func (f *Frame) PrintStackTrace() {
	//     "Print a stack trace starting in this frame"
	if f.HasPreviousFrame() {
		f.PreviousFrame.PrintStackTrace()
	}
	className := f.Method.Holder.Clazz.Name
	methodName := f.Method.Signature.Name
	log.Printf("%s>>#%s @bi: %d\n", className, methodName, f.BytecodeIndex)
}

// INTERPRETER

func (p *Interpreter) DoDup() {
	p.frame.Push(p.frame.Stack[0].(*Object))
}

// doPushLocal: bytecodeIndex = (
//
//	frame push: (
//	    frame local: (frame method bytecode: bytecodeIndex + 1)
//	             at: (frame method bytecode: bytecodeIndex + 2))
//
// )
func (p *Interpreter) DoPushLocal(bytecodeIndex int32) {
	p.frame.Push(
		p.frame.LocalAt(int32(p.frame.Method.Bytecodes[p.frame.BytecodeIndex+1]),
			int32(p.frame.Method.Bytecodes[p.frame.BytecodeIndex+2])))
}

// doPushArgument: bytecodeIndex = (
//
//	frame push: (
//	    frame argument: (frame method bytecode: bytecodeIndex + 1)
//	                at: (frame method bytecode: bytecodeIndex + 2))
//
// )
func (p *Interpreter) DoPushArgument(bytecodeIndex int32) {
	p.frame.Push(
		p.frame.ArgumentAt(int32(p.frame.Method.Bytecodes[p.frame.BytecodeIndex+1]),
			int32(p.frame.Method.Bytecodes[p.frame.BytecodeIndex+2])))
}

//   doPushField: bytecodeIndex = (
//     | fieldIndex |
//     fieldIndex := frame method bytecode: bytecodeIndex + 1.

//	"Push the field with the computed index onto the stack"
//	frame push: (self getSelf field: fieldIndex)
//
// )
func (p *Interpreter) DoPushField(bytecodeIndex int32) {
	fieldIndex := p.frame.Method.Bytecodes[p.frame.BytecodeIndex+1]
	p.frame.Push(p.GetSelf().Fields[fieldIndex])
}

//     blockMethod := frame method constant: bytecodeIndex.

//     frame push: (
//         universe newBlock: blockMethod
//                      with: frame
//                   numArgs: blockMethod numberOfArguments)
//   )

// "Push a new block with the current frame as context onto the stack"
func (p *Interpreter) DoPushBlock(bytecodeIndex int32) {
	var item interface{}
	item = p.frame.Method.Constant(bytecodeIndex)
	blockMethod, ok := item.(*Method)
	if !ok {
		fmt.Printf("method type is wrong!")
		os.Exit(1)
	}
	p.frame.Push(&(p.universe.NewBlock(blockMethod, p.frame, p.frame.Method.NumberOfArguments()).BlockClass.Object))
}

//   doPushConstant: bytecodeIndex = (
//     frame push: (frame method constant: bytecodeIndex)
//   )

//   doPushGlobal: bytecodeIndex = (
//     | globalName global |
//     globalName := frame method constant: bytecodeIndex.

//     "Get the global from the universe"
//     global := universe global: globalName.

//     global ~= nil
//       ifTrue: [ frame push: global  ]
//       ifFalse: [
//         "Send 'unknownGlobal:' to self"
//         self getSelf sendUnknownGlobal: globalName in: universe using: self ]
//   )

// doPop = (
//
//	frame pop
//
// )
func (p *Interpreter) DoPop() *Object {
	return p.frame.Pop()
}

// doPopLocal: bytecodeIndex = (
//
//	frame local: (frame method bytecode: bytecodeIndex + 1)
//	         at: (frame method bytecode: bytecodeIndex + 2)
//	        put: frame pop
//
// )
func (p *Interpreter) DoPopLocal(bytecodeIndex int32) {
	p.frame.PutLocalAt(int32(p.frame.Method.Bytecodes[p.frame.BytecodeIndex+1]),
		int32(p.frame.Method.Bytecodes[p.frame.BytecodeIndex+2]),
		p.frame.Pop())
}

// doPopArgument: bytecodeIndex = (
//
//	frame argument: (frame method bytecode: bytecodeIndex + 1)
//	            at: (frame method bytecode: bytecodeIndex + 2)
//	           put: frame pop
//
// )
func (p *Interpreter) DoPopArgument(bytecodeIndex int32) {
	p.frame.PutArgumentAt(int32(p.frame.Method.Bytecodes[p.frame.BytecodeIndex+1]),
		int32(p.frame.Method.Bytecodes[p.frame.BytecodeIndex+2]),
		p.frame.Pop())
}

//   doPopField: bytecodeIndex = (
//     | fieldIndex |
//     fieldIndex := frame method bytecode: bytecodeIndex + 1.

//	"Set the field with the computed index to the value popped from the stack"
//	self getSelf field: fieldIndex put: frame pop
//
// )
func (p *Interpreter) DoPopField(bytecodeIndex int32) {
	fieldIndex := p.frame.Method.Bytecodes[bytecodeIndex+1]
	p.GetSelf().Fields[fieldIndex] = p.frame.Pop()
}

// doSend: bytecodeIndex = (
//
//	| signature numberOfArguments receiver |
//	signature := frame method constant: bytecodeIndex.
//	numberOfArguments := signature numberOfSignatureArguments.
//	receiver := frame stackElement: numberOfArguments - 1.
//	self send: signature rcvrClass: (receiver somClassIn: universe)
//
// )
func (p *Interpreter) DoSend(bytecodeIndex int32) {
	signature := p.frame.Method.Constant(bytecodeIndex)
	nargs := signature.Nfields
	receiver := p.frame.Stack[nargs-1]
	p.Send(signature.Name, receiver.SomClassIn(p.universe))
}

// send: selector rcvrClass: receiverClass = (
//
//	  | invokable |
//	  invokable := receiverClass lookupInvokable: selector.
//	  self activate: invokable orDnu: selector
//	)
func (p *Interpreter) Send(selector string, receiverClass *Class) {
	invokable := receiverClass.LookupInvokable(selector)
	p.ActivateOrDNU(invokable, selector)
}

//   doSuperSend: bytecodeIndex = (
//     | signature holderSuper invokable |
//     signature := frame method constant: bytecodeIndex.

//     "Send the message
//      Lookup the invokable with the given signature"
//     holderSuper := frame method holder superClass.
//     invokable := holderSuper lookupInvokable: signature.

//     self activate: invokable orDnu: signature
//   )

//   doReturnLocal = (
//     | result |
//     result := frame pop.

//     "Pop the top frame and push the result"
//     self popFrameAndPushResult: result
//   )

//   doReturnNonLocal = (
//     | result context |
//     result := frame pop.

//     "Compute the context for the non-local return"
//     context := frame outerContext.

//     "Make sure the block context is still on the stack"
//     context hasPreviousFrame ifFalse: [
//       | block sender method numArgs |
//       "Try to recover by sending 'escapedBlock:' to the sending object
//        this can get a bit nasty when using nested blocks. In this case
//        the 'sender' will be the surrounding block and not the object
//        that actually sent the 'value' message."
//       block := frame argument: 1 at: 0.
//       sender := frame previousFrame outerContext argument: 1 at: 0.

//       "pop the frame of the currently executing block..."
//       self popFrame.

//       "pop old arguments from stack"
//       method := frame method.
//       numArgs := method numberOfArguments.
//       numArgs timesRepeat: [ frame pop ].

//       "... and execute the escapedBlock message instead"
//       sender sendEscapedBlock: block in: universe using: self.
//       ^ self ].

//     "Unwind the frames"
//     [frame ~= context] whileTrue: [
//       self popFrame ].

//     self popFrameAndPushResult: result
//   )

func (p *Interpreter) Start() {
	//     [true] whileTrue: [
	for {
		//       | bytecodeIndex bytecode bytecodeLength nextBytecodeIndex result |
		//       bytecodeIndex := frame bytecodeIndex.
		//       bytecode := frame method bytecode: bytecodeIndex.
		//       bytecodeLength := Bytecodes length: bytecode.
		//       nextBytecodeIndex := bytecodeIndex + bytecodeLength.
		//       frame bytecodeIndex: nextBytecodeIndex.

		//       result := self dispatch: bytecode idx: bytecodeIndex.
		//       result ~= nil
		//         ifTrue: [ ^ result ] ]
		//   )
	}
}

func (p *Interpreter) Dispatch(bytecode byte, index byte) {
	//   dispatch: bytecode idx: bytecodeIndex = (
	//     bytecode == #halt ifTrue: [
	//       ^ frame stackElement: 0 ].

	//     bytecode == #dup ifTrue: [
	//       self doDup.
	//       ^ nil ].

	//     bytecode == #pushLocal ifTrue: [
	//       self doPushLocal: bytecodeIndex.
	//       ^ nil ].

	//     bytecode == #pushArgument ifTrue: [
	//       self doPushArgument: bytecodeIndex.
	//       ^ nil ].

	//     bytecode == #pushField ifTrue: [
	//       self doPushField: bytecodeIndex.
	//       ^ nil ].

	//     bytecode == #pushBlock ifTrue: [
	//       self doPushBlock: bytecodeIndex.
	//       ^ nil ].

	//     bytecode == #pushConstant ifTrue: [
	//       self doPushConstant: bytecodeIndex.
	//       ^ nil ].

	//     bytecode == #pushGlobal ifTrue: [
	//       self doPushGlobal: bytecodeIndex.
	//       ^ nil ].

	//     bytecode == #pop ifTrue: [
	//       self doPop.
	//       ^ nil ].

	//     bytecode == #popLocal ifTrue: [
	//       self doPopLocal: bytecodeIndex.
	//       ^ nil ].

	//     bytecode == #popArgument ifTrue: [
	//       self doPopArgument: bytecodeIndex.
	//       ^ nil ].

	//     bytecode == #popField ifTrue: [
	//       self doPopField: bytecodeIndex.
	//       ^ nil ].

	//     bytecode == #send ifTrue: [
	//       self doSend: bytecodeIndex.
	//       ^ nil ].

	//     bytecode == #superSend ifTrue: [
	//       self doSuperSend: bytecodeIndex.
	//       ^ nil ].

	//     bytecode == #returnLocal ifTrue: [
	//       self doReturnLocal.
	//       ^ nil ].

	//     bytecode == #returnNonLocal ifTrue: [
	//       self doReturnNonLocal.
	//       ^ nil ].

	//     self error: 'Unknown bytecode' + bytecode asString
	log.Println("Unknown Bytecode " + string(bytecode))
}

//   pushNewFrame: method with: contextFrame = (
//     frame := universe newFrame: frame with: method with: contextFrame.
//     ^ frame
//   )

//   pushNewFrame: method = (
//     ^ self pushNewFrame: method with: nil
//   )

//   frame = (
//     ^ frame
//   )

//   method = (
//     ^ frame method
//   )

// getSelf = (
//
//	"Get the self object from the interpreter"
//	^ frame outerContext argument: 1 at: 0
//
// )
func (p *Interpreter) GetSelf() *Object {
	return p.frame.OuterContext().ArgumentAt(1, 0)
}

//   send: selector rcvrClass: receiverClass = (
//     | invokable |
//     invokable := receiverClass lookupInvokable: selector.
//     self activate: invokable orDnu: selector
//   )

//   activate: invokable orDnu: signature = (
//     invokable ~= nil
//         ifTrue: [
//           "Invoke the invokable in the current frame"
//           invokable invoke: frame using: self ]
//         ifFalse: [
//           | numberOfArguments receiver |
//           numberOfArguments := signature numberOfSignatureArguments.
//           receiver := frame stackElement: numberOfArguments - 1.
//           receiver sendDoesNotUnderstand: signature in: universe using: self ]
//   )

//   popFrame = (
//     | result |
//     "Save a reference to the top frame"
//     result := frame.

//     "Pop the top frame from the frame stack"
//     frame := frame previousFrame.

//     "Destroy the previous pointer on the old top frame"
//     result clearPreviousFrame.

//     "Return the popped frame"
//     ^ result
//   )

//   popFrameAndPushResult: result = (
//     | numberOfArguments |
//     "Pop the top frame from the interpreter frame stack and
//      get the number of arguments"
//     numberOfArguments := self popFrame method numberOfArguments.

//     "Pop the arguments"
//     numberOfArguments
//       timesRepeat: [ frame pop ].

//     frame push: result
//   )
