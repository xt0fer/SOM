package smog

// OOObject - an attempt to make All objects in this project this struct tied to its interface
type OOObject struct {
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
	OOObject
	Fields []*OOObject
	Clazz  *SClass
}

// initialize: numberOfFields with: nilObject = (
//     fields := Array new: numberOfFields withAll: nilObject
//   )

//   somClass = (
//     ^ clazz
//   )

//   somClass: aSClass = (
//     clazz := aSClass
//   )

//   somClassIn: universe = (
//     ^ clazz
//   )

//   fieldName: index = (
//     "Get the name of the field with the given index"
//     ^ clazz instanceFieldName: index
//   )

//   fieldIndex: name = (
//     "Get the index for the field with the given name"
//     ^ clazz lookupFieldIndex: name
//   )

//   numberOfFields = (
//     "Get the number of fields in this object"
//     ^ fields length
//   )

//   field: index = (
//     "Get the field with the given index"
//     ^ fields at: index
//   )

//   field: index put: value = (
//     "Set the field with the given index to the given value"
//     fields at: index put: value
//   )

//   "For using in debugging tools such as the Diassembler"
//   debugString = ( ^ 'SObject(' + clazz name string + ')' )

//   ----

//   new: numberOfFields with: nilObject = (
//     ^ self new initialize: numberOfFields with: nilObject
//   )

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
