package smog

// OOObject - an attempt to make All objects in this project this struct tied to its interface
type OOObject struct {
}

type Object interface {
	somClass() *SClass
	setSomClass(aSClass *SClass)
}

type Class interface {
	Object
	initialize(aUniverse *Universe)
	initializeIn(numberOfFields int32, aUniverse *Universe)
	superClass() *SClass
	setSuperClass(aSClass *SClass)
	hasSuperClass() bool
	name() *SSymbol
	setName(aSSymbol *SSymbol)
	instanceFields() []*Object
	setInstanceFields(aSArray *Array)
	instanceInvokables() *Array
	setInstanceInvokables(aSArray *Array)
	numberOfInstanceInvokables() int32
	instanceInvokable(idx int32) *Object
	instanceInvokablePut(idx int32, aSInvokable *SInvokable)
	lookupInvokable(signature *SString) *Object
	lookupFieldIndex(fieldName *Object) int32
	addInstanceInvokable(value *Object)
	addInstancePrimitive(value *Object)
	addInstancePrimitiveWarn(value *Object, suppressWarning bool)
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
	Class
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

type Sender interface {
	send(selectorString string, arguments []*OOObject, universe *Universe, interpreter *Interpreter)
	sendDoesNotUnderstand(selector string, universe *Universe, interpreter *Interpreter)
	sendUnknownGlobal(globalName *OOObject, universe *Universe, interpreter *Interpreter)
	sendEscapedBlock(block *OOObject, universe *Universe, interpreter *Interpreter)
}

// SSymbol = SString (
// 	| numSignatureArguments |

// SObject = SAbstractObject (
// 	| fields clazz |

type SObject struct {
	Fields []Object
	Clazz  *SClass
}

func NewSObject(n int32, with *OOObject) *SObject {
	so := &SObject{}
	so.Fields = make([]Object, n)

	return so
}

// initialize: numberOfFields with: nilObject = (
//     fields := Array new: numberOfFields withAll: nilObject
//   )

// somClass = (
//
//	^ clazz
//
// )
func (so *SObject) somClass() *SClass {
	return so.Clazz
}

// somClass: aSClass = (
//
//	clazz := aSClass
//
// )
func (so *SObject) setSomClass(aSClass *SClass) {
	so.Clazz = aSClass
}

// somClassIn: universe = (
//
//	^ clazz
//
// )
func (so *SObject) somClassIn(u *Universe) *SClass {
	return so.Clazz
}

// fieldName: index = (
//
//	"Get the name of the field with the given index"
//	^ clazz instanceFieldName: index
//
// )
func (so *SObject) fieldName(index int32) string {
	return so.Clazz.fieldName(index)
}

// fieldIndex: name = (
//
//	"Get the index for the field with the given name"
//	^ clazz lookupFieldIndex: name
//
// )
func (so *SObject) fieldIndex(name string) int32 {
	return so.Clazz.fieldIndex(name)
}

//   numberOfFields = (
//     "Get the number of fields in this object"
//     ^ fields length
//   )

// field: index = (
//
//	"Get the field with the given index"
//	^ fields at: index
//
// )
func (so *SObject) field(index int32) Object {
	return so.Fields[index]
}

// field: index put: value = (
//
//	"Set the field with the given index to the given value"
//	fields at: index put: value
//
// )
func (so *SObject) fieldPut(index int32, value Object) Object {
	so.Fields[index] = value
}

//   "For using in debugging tools such as the Diassembler"
//   debugString = ( ^ 'SObject(' + clazz name string + ')' )

//   ----

// ??
type Invokable *OOObject

// SClass = SObject (
//
//	| universe
//	  superClass
//	  name
//	  instanceInvokables instanceFields|
type SClass struct {
	SObject
	Universe       *Universe
	SuperClass     *SClass
	Name           *SSymbol
	InstanceFields []Invokable
}

// SSymbol = SString (
//
//	| numSignatureArguments |
type SSymbol struct {
	SString
	NumSignatureArguments int
}

// SString = SAbstractObject (
//
//	| string |
type SString struct {
	SObject
	S string
}

func (S *SString) string() string { return S.S }

// "For using in debugging tools such as the Diassembler"
func (S *SString) debugString() string {
	t := "SString(" + S.S + ")"
	return t
}

// somClassIn: universe = (
//
//	  ^ universe stringClass
//	)
func (S *SString) somClassIn(u *Universe) *SClass {
	return S.SObject.Clazz
}

// initializeWith: aString = (
func NewString(aString string) *SString {
	s := &SString{}
	s.S = aString
	return s
}
