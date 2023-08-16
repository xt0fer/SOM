package smog

import (
	"fmt"
	"log"
)

/**
* Some ideas on how to break the log-jam
*
* EVERYTHING is an Object
* based on Class, which has a name, field templates (names), and methods (with names)
* a Method is named, and has an array of bytecodes

// SAbstractObject.som     xSClass.som              xSMethod.som             xSString.som
// xSArray.som              xSDouble.som             xSObject.som             xSSymbol.som
// xSBlock.som              xSInteger.som            xSPrimitive.som

// The Data Model (Objects everywhere)

*
*/

// type Sender interface {
// 	send(selectorString string, arguments []Object, universe *Universe, interpreter *Interpreter)
// 	sendDoesNotUnderstand(selector string, universe *Universe, interpreter *Interpreter)
// 	sendUnknownGlobal(globalName Object, universe *Universe, interpreter *Interpreter)
// 	sendEscapedBlock(block Object, universe *Universe, interpreter *Interpreter)
// }

type ObjectInterface interface {
	Send(selectorString string, arguments []ObjectInterface, universe *Universe, interpreter *Interpreter)
	SendDoesNotUnderstand(selector string, universe *Universe, interpreter *Interpreter)
	SendUnknownGlobal(globalName ObjectInterface, universe *Universe, interpreter *Interpreter)
	SendEscapedBlock(block ObjectInterface, universe *Universe, interpreter *Interpreter)
	//
	SomClass() ClassInterface                                 // get the class of this object
	SetSomClass(aSClass ClassInterface)                       // set the class of this object
	InitializeWith(numberOfFields int32, obj ObjectInterface) // create object with N slots in it's Fields array
}

type ClassInterface interface {
	Initialize(aUniverse *Universe)                                       // init this class object in the Universe
	InitializeIn(numberOfFields int32, aUniverse *Universe)               // init this class with N fields in the Universe
	SuperClass() ClassInterface                                           // get superclass
	SetSuperClass(aSClass ClassInterface)                                 // set superclass
	HasSuperClass() bool                                                  // confirm superclass
	Name() *Symbol                                                        // get classname
	SetName(aSymbol Symbol)                                               // set classname
	InstanceFields() []ObjectInterface                                    // get array of instance field names? or objects?
	SetInstanceFields(aSArray ArrayInterface)                             // set array of instance fields
	InstanceInvokables() ArrayInterface                                   // get array of invokables (methods and blocks?)
	setInstanceInvokables(aSArray ArrayInterface)                         // set the invokables array
	NumberOfInstanceInvokables() int32                                    // get size of invokables array
	InstanceInvokable(idx int32) ObjectInterface                          // get invokable at idx
	InstanceInvokablePut(idx int32, aSInvokable Invokable)                // set invokable at idx
	LookupInvokable(signature *String) ObjectInterface                    // get invokable by symbol name
	LookupFieldIndex(fieldName ObjectInterface) int32                     // get index of invokable by name
	AddInstanceInvokable(value ObjectInterface)                           // add an Invokable method to array
	AddInstancePrimitive(value ObjectInterface)                           // add a Primitive to Invokable array
	AddInstancePrimitiveWarn(value ObjectInterface, suppressWarning bool) // same as above, w|w/o error report(?)
	InstanceFieldName(index int32) string                                 // get name of instance variable at index
	NumberOfInstanceFields() int32                                        // size of instanceFields array
	NumberOfSuperInstanceFields() int32                                   // size of superclass' instanceFields array
	HasPrimitives() bool                                                  // class contains primitives for some methods
	LoadPrimitives()                                                      // "load" primitives from what?
	DebugString() string                                                  // print what class is named on debug output
	// new: universe);
	// new: numberOfFields in: universe;
}

type ArrayInterface interface {
	initializeWithAnd(length int32, object ObjectInterface)        // make new SomArray with element type object
	somClassIn(universe *Universe) ClassInterface                  // get class of the array element type
	indexableField(idx int32) ObjectInterface                      // get element at:
	indexableFieldPut(idx int32, value ObjectInterface)            // set element at:
	numberOfIndexableFields() int32                                // get size of SomArray
	copyAndExtendWithInt(value ArrayInterface, universe *Universe) // add a ???
	copyIndexableFieldsTo(destination ArrayInterface)              // copy SomArray to new destination
	debugString() string                                           // print array to debug out
	// new: length with: nilObject = (
}

// The Data Model (Objects everywhere)

type Object struct {
	Fields  []*Object // local vars (any object) index of field is same as index of Class.InstanceFields
	Nfields int32
	Clazz   *Class
}

type Class struct {
	Object
	Universe           *Universe   // where it is defined, as a singleton
	SuperClass         *Class      // immediate superclass of this class
	Name               *Symbol     // name(string) of the class
	InstanceInvokables []Invokable // all the pretty horses (all the Methods, Primitives, ...???)
	InstanceFields     []*Symbol   // template for InstanceFields, the index of the Name is the index within the Object.
}

type Symbol struct { // used for SomSymbol as well as model string
	Object
	Name      string // className, instanceFieldName, globalName, methodSignature, primitive(?)
	NumOfArgs int32
}

type Array struct { // used for a SomArray data structure, not used within the Data Model
	Object
	Elements []*Object
}
type String struct { // used to model a SomString
	Object
	Value string
}
type Integer struct {
	Object
	Value int32
}
type Double struct {
	Object
	Value float64
}

type Method struct {
	Array                          // used to store the method local objects
	Signature        *Symbol       // symbol with method signature in it
	Holder           *Object       // what Class is this attached to?
	Bytecodes        []byte        // bytecode array, code to be run when method invoked.
	Literals         []interface{} // array of symbols as literals #()
	NumLocals        int32         // number of local objects
	MaxStackElements int32         // limit on Stack??
}

// For instance, in usage,
// bootstrapMethod := self newMethod: (self symbolFor: 'bootstrap')
//      bc: #(#halt) literals: #() numLocals: 0 maxStack: 2.

type Primitive struct {
	Object
	Signature *Symbol // something like "of:at:argN:"
	Holder    *Object // what Class is this attached to?
	IsEmpty   bool    // not sure
	Operation *Block  // Is this the code run for primitive?
}

type Block struct { // not sure what these are just yet
	Object
	Method     *Method // method which implements the bytecodes
	Context    *Frame  // this seems to be the Universe
	BlockClass *Class  // which block class?
}

func NewObject(n int32, with *Object) *Object {
	so := &Object{}
	so.Fields = make([]*Object, n)
	so.Nfields = n
	so.initializeWith(n, with)
	return so
}

func (so *Object) initializeWith(numberOfFields int32, obj *Object) {
	for i := range so.Fields {
		so.Fields[i] = obj
	}
	//so.Clazz = init??
}

func (receiver *Object) Send(selectorString string, arguments []Object, universe *Universe, interpreter *Interpreter) {
}
func (receiver *Object) SendDoesNotUnderstand(selector string, universe *Universe, interpreter *Interpreter) {
}
func (receiver *Object) SendUnknownGlobal(globalName Object, universe *Universe, interpreter *Interpreter) {
}
func (receiver *Object) SendEscapedBlock(block Object, universe *Universe, interpreter *Interpreter) {
}

// func (so *Object) SomClass() *Class {
// 	return so.Clazz
// }

//	func (so *Object) SetSomClass(aSClass *Class) {
//		so.Clazz = aSClass
//	}
func (so *Object) SomClassIn(u *Universe) *Class {
	return so.Clazz
}
func (so *Object) FieldName(index int32) *Symbol {
	return so.Clazz.InstanceFields[index]
}
func (so *Object) FieldIndex(name string) int32 {
	for i, v := range so.Clazz.InstanceFields {
		if v.Name == name {
			return int32(i)
		}
	}
	return -1
}

// func (so *Object) Field(index int32) *Object {
// 	return so.Fields[index]
// }
// func (so *Object) SetField(index int32, value *Object) {
// 	so.Fields[index] = value
// }

// ??
type Invokable *Object

func NewClass(numberOfFields int32, u *Universe) *Class {
	sc := &Class{}
	//sc.Object = &Object{}
	return sc
}

func (sc *Class) InitializeIn(numberOfFields int32, u *Universe) {
	sc.Universe = u
	sc.Object.initializeWith(numberOfFields, u.NilObject)
}
func (sc *Class) setSuperClass(nc *Class) {
	sc.SuperClass = nc
}

// setName
// func (sc *Class) SetName(sym *Symbol) {
// 	sc.Name = sym
// }
// func (sc *Class) GetName() *Symbol {
// 	return sc.Name
// }

// SetInstancesFields
func (sc *Class) SetInstancesFields(size int32) {
	sc.InstanceFields = make([]*Symbol, size)
}

// SetInstanceInvokables
func (sc *Class) SetInstanceInvokables(size int32) {
	sc.InstanceInvokables = make([]Invokable, size)
}

// NumberOfInstanceFields
func (sc *Class) NumberOfInstanceFields() int32 {
	return int32(len(sc.InstanceFields))
}
func NewSymbol(value string, n int32) *Symbol {
	ss := &Symbol{}
	ss.Name = value
	return ss
}
func (S *Symbol) SomClassIn(u *Universe) *Class {
	return u.SymbolClass
}

func NewString(aString string) *String {
	s := &String{}
	s.Initialize(aString)
	return s
}

func (s *String) Initialize(aString string) {
	s.Value = aString
}

func (S *String) StringValue() string { return S.Value }

// "For using in debugging tools such as the Diassembler"
func (S *String) DebugString() string {
	t := "String(" + S.Value + ")"
	return t
}

func (S *String) SomClassIn(u *Universe) *Class {
	return u.StringClass
}

func NewInteger(n int32) *Integer {
	s := &Integer{}
	s.Initialize(n)
	return s
}

func (s *Integer) Initialize(n int32) {
	s.Value = n
}

func (i *Integer) IntegerValue() int32 { return i.Value }

// "For using in debugging tools such as the Diassembler"
func (i *Integer) DebugString() string {
	t := "Integer(" + fmt.Sprintf("%v", i.Value) + ")"
	return t
}

func (i *Integer) SomClassIn(u *Universe) *Class {
	return u.IntegerClass
}

func NewDouble(n float64) *Double {
	s := &Double{}
	s.Initialize(n)
	return s
}

func (d *Double) Initialize(n float64) {
	d.Value = n
}

func (d *Double) DoubleValue() float64 { return d.Value }

// "For using in debugging tools such as the Diassembler"
func (d *Double) DebugString() string {
	t := "Double(" + fmt.Sprintf("%v", d.Value) + ")"
	return t
}

func (i *Double) SomClassIn(u *Universe) *Class {
	return u.DoubleClass
}

// METHOD

func NewMethod() *Method {
	m := &Method{}
	return m
}

// array of literals, NQR
func (m *Method) InitializeWith(sym *Symbol, bcArray []byte, literalsArray []interface{},
	numLocals int32, maxStack int32) {
	m.Signature = sym
	m.Bytecodes = bcArray
	m.Literals = literalsArray
	m.NumLocals = numLocals
	m.MaxStackElements = maxStack
}

// func (m *Method) GetBytecode(index int32) int32 {
// 	return int32(m.Bytecodes[index])
// }

func (m *Method) IsPrimitive() bool {
	return false
}

// func (m *Method) NumberOfLocals() int32 {
// 	return m.NumOfLocals
// }

// func (m *Method) MaximumNumberOfStackElements() int32 {
// 	return m.MaximumStackElements
// }

// func (m *Method) Signature() *Symbol {
// 	return m.Sig
// }

// func (m *Method) Holder() *Object {
// 	return m.Hold
// }

func (m *Method) SetHolder(h *Object) {
	m.Holder = h

	// literals == nil ifTrue: [ ^ self ].
	if m.Literals == nil {
		return
	}

	panic("Invokables do not have the correct Holder")
	// "Make sure all nested invokables have the same holder"
	//
	//	literals do: [:l |
	//	  (l class == SMethod or: [l class == SPrimitive]) ifTrue: [
	//	    l holder: value ] ]
}

// "Get the constant associated to a given bytecode index"
func (m *Method) Constant(bytecodeIndex int32) *Object {
	return m.Literals[m.Bytecodes[bytecodeIndex+1]].(*Object)
}

func (m *Method) NumberOfArguments() int32 {
	return m.Signature.NumOfArgs
}

func (m *Method) NumberOfBytecodes() int32 {
	return int32(len(m.Bytecodes))
}

func (m *Method) BytecodeAt(index int32) int32 {
	return int32(m.Bytecodes[index])
}

// invoke: frame using: interpreter = (
//
//	| newFrame |
//	newFrame := interpreter pushNewFrame: self.
//	newFrame copyArgumentsFrom: frame
//
// )
//
//	"Allocate and push a new frame on the interpreter stack"
func (m *Method) InvokeUsing(frame *Frame, interpreter *Interpreter) {
	// newFrame := interpreter.PushNewFrame(m)

	// newFrame.CopyArgumentsFrom(frame)
}

func (m *Method) SomClassIn(u *Universe) *Class {
	return u.MethodClass
}

// "For using in debugging tools such as the Diassembler"
// debugString = ( ^ 'SMethod(' + holder name + '>>#' + signature string + ')' )
func (m *Method) DebugString() {
	log.Printf("Method(%s>>#%s)\n", m.Holder.Clazz.Name.Name, m.Signature.Name)
}

// BLOCK
// Method     *Method   // method which implements the bytecodes
// Context    *Universe // this seems to be the Universe
// BlockClass *Class    // which block class?

// new: aSMethod in: aContext with: aBlockClass = (
func NewBlock(aMethod *Method, aContext *Frame, aBlockClass *Class) *Block {
	nb := &Block{}
	nb.Method = aMethod
	nb.Context = aContext
	nb.BlockClass = aBlockClass
	return nb
}

// 	| method context blockClass |

func (nb *Block) Initialize(aMethod *Method, aContext *Frame, aBlockClass *Class) {
	nb.Method = aMethod
	nb.Context = aContext
	nb.BlockClass = aBlockClass
}
func (nb *Block) GetMethod() *Method {
	return nb.Method
}
func (nb *Block) GetContext() *Frame {
	return nb.Context
}

func (nb *Block) SomClassIn() *Class {
	return nb.BlockClass
}

// "For using in debugging tools such as the Diassembler"
func (m *Block) DebugString() {
	log.Printf("Block(%s>>#%s)\n", m.Method.Signature.Name)
}

// 	evaluationPrimitive: numberOfArguments in: universe = (
// 	  ^ SPrimitive new: (self computeSignatureString: numberOfArguments)
// 					in: universe
// 				  with: [:frame :interp |
// 		  | rcvr context newFrame |
// 		  "Get the block (the receiver) from the stack"
// 		  rcvr := frame stackElement: numberOfArguments - 1.

// 		  "Get the context of the block"
// 		  context := rcvr context.

// 		  "Push a new frame and set its context to be the one specified in
// 		   the block"
// 		  newFrame := interp pushNewFrame: rcvr method with: context.
// 		  newFrame copyArgumentsFrom: frame ]
// 	)

func (nb *Block) ComputeSignatureString(nArgs int32) string {
	signatureString := "value"
	if nArgs > 1 {
		signatureString = signatureString + ":"
	}
	for i := 2; i < int(nArgs); i++ {
		signatureString = signatureString + "with:"
	}
	return signatureString
}
