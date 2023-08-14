package smog

import "os"

type stringToObjMap map[string]*Object

type Universe struct {
	symbolTable   map[string]*Symbol
	globals       map[string]*Object
	interpreter   *Interpreter
	dumpBytecodes bool
	avoidExit     bool

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
	StringClass    *Class
	SymbolClass    *Class

	TrueClass  *Class
	FalseClass *Class
}

func NewUniverse() *Universe {
	nu := &Universe{}
	return nu
}

func (u *Universe) initialize() {
	u.symbolTable = make(map[string]*Symbol)
	u.globals = make(map[string]*Object)

	u.NilObject = &Object{}
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
