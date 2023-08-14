package smog

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

type Object struct {
	Fields []*Object // local vars (any object) index of field is same as index of Class.InstanceFields
	Clazz  *Class
}

type Class struct {
	Universe       *Universe   // where it is defined, as a singleton
	SuperClass     *Class      // immediate superclass of this class
	Name           *Symbol     // name(string) of the class
	InstanceInvokables []Invokable // all the pretty horses (all the Methods, Primitives, ...???)
	InstanceFields []*Symbols       // template for InstanceFields, the index of the Name is the index within the Object.
}

type Symbol struct { // used for SomSymbol as well as model string
	Name string // className, instanceFieldName, globalName, methodSignature, primitive(?)
}

type Array struct { // used for a SomArray data structure, not used within the Data Model
	Fields []*Object
}
type String struct { // used to model a SomString
	Name string
}
type Integer struct {
	Value int32
}
type Double struct {
	Value float32
}

type Method struct {
	Signature *Symbol // symbol with method signature in it
	Holder *Object
    Bytecodes []byte // bytecode array, code to be run when method invoked.
	Literals []string // array of symbols as literals #()
    NumberOfLocals int32 // number of local objects
	MaximumNumberOfStackElements int32 // limit on Stack??

}
// For instance, in usage,
// bootstrapMethod := self newMethod: (self symbolFor: 'bootstrap')
//      bc: #(#halt) literals: #() numLocals: 0 maxStack: 2.

type Primitive struct {
	Signature *Symbol
	Holder *Object
	IsEmpty bool
	Operation // Is this the code
}
type Block struct { // not sure what these are just yet
	Method *Method
	Context *Context
	blockClass *Class
}


*
*/

// type Sender interface {
// 	send(selectorString string, arguments []Object, universe *Universe, interpreter *Interpreter)
// 	sendDoesNotUnderstand(selector string, universe *Universe, interpreter *Interpreter)
// 	sendUnknownGlobal(globalName Object, universe *Universe, interpreter *Interpreter)
// 	sendEscapedBlock(block Object, universe *Universe, interpreter *Interpreter)
// }

type ObjectInterface interface {
	send(selectorString string, arguments []ObjectInterface, universe *Universe, interpreter *Interpreter)
	sendDoesNotUnderstand(selector string, universe *Universe, interpreter *Interpreter)
	sendUnknownGlobal(globalName ObjectInterface, universe *Universe, interpreter *Interpreter)
	sendEscapedBlock(block ObjectInterface, universe *Universe, interpreter *Interpreter)
	//
	somClass() ClassInterface                                 // get the class of this object
	setSomClass(aSClass ClassInterface)                       // set the class of this object
	initializeWith(numberOfFields int32, obj ObjectInterface) // create object with N slots in it's Fields array
}

type ClassInterface interface {
	initialize(aUniverse *Universe)                                       // init this class object in the Universe
	initializeIn(numberOfFields int32, aUniverse *Universe)               // init this class with N fields in the Universe
	superClass() ClassInterface                                           // get superclass
	setSuperClass(aSClass ClassInterface)                                 // set superclass
	hasSuperClass() bool                                                  // confirm superclass
	name() *Symbol                                                       // get classname
	setName(aSymbol Symbol)                                             // set classname
	instanceFields() []ObjectInterface                                    // get array of instance field names? or objects?
	setInstanceFields(aSArray ArrayInterface)                             // set array of instance fields
	instanceInvokables() ArrayInterface                                   // get array of invokables (methods and blocks?)
	setInstanceInvokables(aSArray ArrayInterface)                         // set the invokables array
	numberOfInstanceInvokables() int32                                    // get size of invokables array
	instanceInvokable(idx int32) ObjectInterface                          // get invokable at idx
	instanceInvokablePut(idx int32, aSInvokable Invokable)                // set invokable at idx
	lookupInvokable(signature *String) ObjectInterface                   // get invokable by symbol name
	lookupFieldIndex(fieldName ObjectInterface) int32                     // get index of invokable by name
	addInstanceInvokable(value ObjectInterface)                           // add an Invokable method to array
	addInstancePrimitive(value ObjectInterface)                           // add a Primitive to Invokable array
	addInstancePrimitiveWarn(value ObjectInterface, suppressWarning bool) // same as above, w|w/o error report(?)
	instanceFieldName(index int32) string                                 // get name of instance variable at index
	numberOfInstanceFields() int32                                        // size of instanceFields array
	numberOfSuperInstanceFields() int32                                   // size of superclass' instanceFields array
	hasPrimitives() bool                                                  // class contains primitives for some methods
	loadPrimitives()                                                      // "load" primitives from what?
	debugString() string                                                  // print what class is named on debug output
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
	Fields []*Object // local vars (any object) index of field is same as index of Class.InstanceFields
	Clazz  *Class
}

type Class struct {
	Universe       *Universe   // where it is defined, as a singleton
	SuperClass     *Class      // immediate superclass of this class
	Name           *Symbol     // name(string) of the class
	InstanceInvokables []Invokable // all the pretty horses (all the Methods, Primitives, ...???)
	InstanceFields []*Symbol       // template for InstanceFields, the index of the Name is the index within the Object.
}

type Symbol struct { // used for SomSymbol as well as model string
	Name string // className, instanceFieldName, globalName, methodSignature, primitive(?)
}

type Array struct { // used for a SomArray data structure, not used within the Data Model
	Fields []*Object
}
type String struct { // used to model a SomString
	String string
}
type Integer struct {
	Value int32
}
type Double struct {
	Value float32
}

type Method struct {
	Signature *Symbol // symbol with method signature in it
	Holder *Object
    Bytecodes []byte // bytecode array, code to be run when method invoked.
	Literals []string // array of symbols as literals #()
    NumberOfLocals int32 // number of local objects
	MaximumNumberOfStackElements int32 // limit on Stack??

}
// For instance, in usage,
// bootstrapMethod := self newMethod: (self symbolFor: 'bootstrap')
//      bc: #(#halt) literals: #() numLocals: 0 maxStack: 2.

type Primitive struct {
	Signature *Symbol
	Holder *Object
	IsEmpty bool
	Operation // Is this the code
}
type Block struct { // not sure what these are just yet
	Method *Method	// method which implements the bytecodes
	Context *Universe // this seems to be the Universe
	blockClass *Class // which block class?
}

// SSymbol = SString (
// 	| numSignatureArguments |

// SObject = SAbstractObject (
// 	| fields clazz |

// type SObject struct {
// 	Object
// 	Fields []*Object
// 	Clazz  Class
// }

func NewSObject(n int32, with *Object) *Object {
	so := &Object{}
	so.Fields = make([]*Object, n)
	so.initializeWith(n, with)
	return so
}

func (so *Object) initializeWith(numberOfFields int32, obj *Object) {
	for i := range so.Fields {
		so.Fields[i] = obj
	}
	//so.Clazz =
}

func (self *Object) send(selectorString string, arguments []Object, universe *Universe, interpreter *Interpreter) {
}
func (self *Object) sendDoesNotUnderstand(selector string, universe *Universe, interpreter *Interpreter) {
}
func (self *Object) sendUnknownGlobal(globalName Object, universe *Universe, interpreter *Interpreter) {
}
func (self *Object) sendEscapedBlock(block Object, universe *Universe, interpreter *Interpreter) {
}

func (so *Object) somClass() *Class {
	return so.Clazz
}

func (so *Object) setSomClass(aSClass *Class) {
	so.Clazz = aSClass
}
func (so *Object) somClassIn(u *Universe) *Class {
	return so.Clazz
}
func (so *Object) fieldName(index int32) string {
	return "" //so.Clazz.fieldName(index)
}
func (so *Object) fieldIndex(name string) int32 {
	return 0 //so.Clazz.fieldIndex(name)
}
func (so *Object) field(index int32) *Object {
	return so.Fields[index]
}
func (so *Object) fieldPut(index int32, value *Object) {
	so.Fields[index] = value
}

// ??
type Invokable *Object

// type SClass struct {
// 	Object
// 	Universe       *Universe
// 	SuperClass     *SClass
// 	Name           *SSymbol
// 	InstanceFields []Invokable
// }

func NewSClass(numberOfFields int32, u *Universe) *Class {
	sc := &Class{}
	sc.Object = &Object{}
	return sc
}

func (sc *Class) initializeIn(numberOfFields int32, u *Universe) {
	sc.Universe = u
	sc.Object.initializeWith(numberOfFields, u.NilObject)
}

// SSymbol = SString (
//
//	| numSignatureArguments |
// type SSymbol struct {
// 	SString
// 	NumSignatureArguments int32
// }

func NewSymbol(value string, n int32) *Symbol {
	ss := &Symbol{}
	ss.Name = *NewString(value)
	return ss
}

// SString = SAbstractObject (
// type SString struct {
// 	SObject
// 	S string
// }

func NewString(aString string) *String {
	s := &String{}
	s.initialize(aString)
	return s
}

func (s *String) initialize(aString string) {
	s.String = aString
}

func (S *String) string() string { return S.S }

// "For using in debugging tools such as the Diassembler"
func (S *String) debugString() string {
	t := "String(" + S.String + ")"
	return t
}

func (S *String) somClassIn(u *Universe) *Class {
	return u.StringClass
}
