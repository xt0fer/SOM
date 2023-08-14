package smog

import "os"

// | symbolTable globals classPath dumpBytecodes interpreter

// avoidExit
// lastExitCode
// exitBlock

// nilObject
// trueObject
// falseObject

// objectClass
// classClass
// metaclassClass

// nilClass
// integerClass
// arrayClass
// methodClass
// symbolClass
// primClass
// stringClass
// systemClass
// blockClass
// doubleClass

// trueClass
// falseClass
// |

// initialize = (
// symbolTable := Dictionary new.
// globals := Dictionary new.
// interpreter := Interpreter new: self.
// dumpBytecodes := false.
// avoidExit := false
// )

type ObjToObjMap map[Object]Object

type Universe struct {
	symbolTable   ObjToObjMap
	globals       ObjToObjMap
	interpreter   *Interpreter
	dumpBytecodes bool
	avoidExit     bool

	//
	NilObject Object
}

func NewUniverse() *Universe {
	nu := &Universe{}
	return nu
}

func (u *Universe) initialize() {

	u.NilObject = SObject{}
}

// UNIVERSE

func (u *Universe) Exit(code int) {
	os.Exit(code)
}
func (u *Universe) Interpret(args []string) {

}
