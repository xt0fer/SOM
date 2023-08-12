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

// SObject
type SObject struct {
	OOObject
	Fields []*OOObject
	Clazz  *SClass
}

// SObject = SAbstractObject (
// 	| fields clazz |

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
