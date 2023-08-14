package smog

/**
* Some ideas on how to break the log-jam
*
* EVERYTHING is an Object
* based on Class, which has a name, field templates (names), and methods (with names)
* a Method is named, and has an array of bytecodes

type Object struct {
	Fields []*Object // local vars (any object) index of field is same as index of Class.InstanceFields
	Clazz  *Class
}

type Class struct {
	Universe       *Universe   // where it is defined, as a singleton
	SuperClass     *Class      // immediate superclass of this class
	Name           *Symbol     // name(string) of the class
	InstanceInvokables []Invokable // all the pretty horses (all the Methods)
	InstanceFields []*Symbols       // template for InstanceFields, the index of the Name is the index within the Object.
}

type Symbol struct {
	Name string // className, instanceFieldName, globalName, methodSignature
}
*/
// type Sender interface {
// 	send(selectorString string, arguments []Object, universe *Universe, interpreter *Interpreter)
// 	sendDoesNotUnderstand(selector string, universe *Universe, interpreter *Interpreter)
// 	sendUnknownGlobal(globalName Object, universe *Universe, interpreter *Interpreter)
// 	sendEscapedBlock(block Object, universe *Universe, interpreter *Interpreter)
// }

type Object interface {
	send(selectorString string, arguments []Object, universe *Universe, interpreter *Interpreter)
	sendDoesNotUnderstand(selector string, universe *Universe, interpreter *Interpreter)
	sendUnknownGlobal(globalName Object, universe *Universe, interpreter *Interpreter)
	sendEscapedBlock(block Object, universe *Universe, interpreter *Interpreter)
	//
	somClass() Class
	setSomClass(aSClass Class)
	initializeWith(numberOfFields int32, obj Object)
}

type Class interface {
	initialize(aUniverse *Universe)
	initializeIn(numberOfFields int32, aUniverse *Universe)
	superClass() Class
	setSuperClass(aSClass Class)
	hasSuperClass() bool
	name() *SSymbol
	setName(aSSymbol SSymbol)
	instanceFields() []Object
	setInstanceFields(aSArray Array)
	instanceInvokables() Array
	setInstanceInvokables(aSArray Array)
	numberOfInstanceInvokables() int32
	instanceInvokable(idx int32) Object
	instanceInvokablePut(idx int32, aSInvokable Invokable)
	lookupInvokable(signature *SString) Object
	lookupFieldIndex(fieldName Object) int32
	addInstanceInvokable(value Object)
	addInstancePrimitive(value Object)
	addInstancePrimitiveWarn(value Object, suppressWarning bool)
	instanceFieldName(index int32) string
	numberOfInstanceFields() int32
	numberOfSuperInstanceFields() int32
	hasPrimitives() bool
	loadPrimitives()
	debugString() string
	// new: universe);
	// new: numberOfFields in: universe;
}

type Array interface {
	initializeWithAnd(length int32, object *Object)
	somClassIn(universe *Universe) *SClass
	indexableField(idx int32) *Object
	indexableFieldPut(idx int32, value *Object)
	numberOfIndexableFields() int32
	copyAndExtendWithInt(value *Array, universe *Universe)
	copyIndexableFieldsTo(destination *Array)
	debugString() string
	// new: length with: nilObject = (
}

// SSymbol = SString (
// 	| numSignatureArguments |

// SObject = SAbstractObject (
// 	| fields clazz |

type SObject struct {
	Object
	Fields []*Object
	Clazz  Class
}

func NewSObject(n int32, with *SObject) *SObject {
	so := &SObject{}
	so.Fields = make([]*Object, n)
	so.initializeWith(n, with)
	return so
}

func (so *SObject) initializeWith(numberOfFields int32, obj Object) {
	for i := range so.Fields {
		so.Fields[i] = &obj
	}
	//so.Clazz =
}

func (self *SObject) send(selectorString string, arguments []Object, universe *Universe, interpreter *Interpreter) {
}
func (self *SObject) sendDoesNotUnderstand(selector string, universe *Universe, interpreter *Interpreter) {
}
func (self *SObject) sendUnknownGlobal(globalName Object, universe *Universe, interpreter *Interpreter) {
}
func (self *SObject) sendEscapedBlock(block Object, universe *Universe, interpreter *Interpreter) {
}

func (so SObject) somClass() *Class {
	return &so.Clazz
}

func (so SObject) setSomClass(aSClass Class) {
	so.Clazz = aSClass
}
func (so *SObject) somClassIn(u *Universe) *SClass {
	return so.Clazz
}
func (so *SObject) fieldName(index int32) string {
	return so.Clazz.fieldName(index)
}
func (so *SObject) fieldIndex(name string) int32 {
	return so.Clazz.fieldIndex(name)
}
func (so *SObject) field(index int32) *Object {
	return so.Fields[index]
}
func (so *SObject) fieldPut(index int32, value *Object) {
	so.Fields[index] = value
}

// ??
type Invokable *Object

type SClass struct {
	*SObject
	Universe       *Universe
	SuperClass     *SClass
	Name           *SSymbol
	InstanceFields []Invokable
}

func NewSClass(numberOfFields int32, u *Universe) *SClass {
	sc := &SClass{}
	return sc
}

func (sc *SClass) initializeIn(numberOfFields int32, u *Universe) {
	sc.Universe = u
	sc.Object.initializeWith(numberOfFields, u.NilObject)
}

// SSymbol = SString (
//
//	| numSignatureArguments |
type SSymbol struct {
	SString
	NumSignatureArguments int32
}

func NewSymbol(value string, n int32) *SSymbol {
	ss := &SSymbol{}
	ss.SString = *NewString(value)
	ss.NumSignatureArguments = n
	return ss
}

// SString = SAbstractObject (
type SString struct {
	SObject
	S string
}

func NewString(aString string) *SString {
	s := &SString{}
	s.initialize(aString)
	return s
}

func (s *SString) initialize(aString string) {
	s.S = aString
}

func (S *SString) string() string { return S.S }

// "For using in debugging tools such as the Diassembler"
func (S *SString) debugString() string {
	t := "SString(" + S.S + ")"
	return t
}

func (S *SString) somClassIn(u *Universe) *SClass {
	return S.SObject.Clazz
}
