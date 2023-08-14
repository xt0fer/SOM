package smog

import "os"

type stringToObjMap map[string]*Object

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
		//         systemClass superClass: superClass.
		//         systemClass somClass superClass: (superClass somClass) ]
		systemClass.setSuperClass(superClass)
		systemClass.SomClass().setSuperClass(superClass.SomClass())
	} else {
		//         systemClass somClass superClass: classClass ].
		systemClass.SomClass().setSuperClass(u.ClassClass)
	}

	// "Initialize the array of instance fields"
	//
	//	systemClass instanceFields: (self newArray: 0).
	systemClass.SetInstancesFields(0)
	// systemClass somClass instanceFields: (self newArray: 0).
	systemClass.SomClass().SetInstancesFields(0)

	// "Initialize the array of instance invokables"
	//
	//	systemClass instanceInvokables: (self newArray: 0).
	systemClass.SetInstanceInvokables(0)
	// systemClass somClass instanceInvokables: (self newArray: 0).
	systemClass.SomClass().SetInstanceInvokables(0)

	// "Initialize the name of the system class"
	//
	//	systemClass name: (self symbolFor: name).
	systemClass.SetName(u.symbolFor(name))
	// systemClass somClass name: (self symbolFor: name + ' class').
	systemClass.SomClass().SetName(u.symbolFor(name + " class"))

	// "Insert the system class into the dictionary of globals"
	//     self global: systemClass name put: systemClass.
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

func (u *Universe) initializeObjectSystem() {
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
	//     trueSymbol := self symbolFor: 'True'.
	trueSymbol := u.symbolFor("True")
	//     trueClass := self loadClass: trueSymbol.
	u.TrueClass = NewClass(0, u)
	u.TrueClass.SetName(trueSymbol)
	//     trueObject := self newInstance: trueClass.
	u.TrueObject = u.NewInstance(u.TrueClass)
	//     falseSymbol := self symbolFor: 'False'.
	falseSymbol := u.symbolFor("False")
	//     falseClass := self loadClass: falseSymbol.
	u.FalseClass = NewClass(0, u)
	u.FalseClass.SetName(falseSymbol)
	//     falseObject := self newInstance: falseClass.
	u.FalseObject = u.NewInstance(u.FalseClass)

	//     "Load the system class and create an instance of it"
	//     systemClass := self loadClass: (self symbolFor: 'System').
	u.SystemClass = NewClass(0, u)
	u.SystemClass.SetName(u.symbolFor("System"))
	//     systemObject := self newInstance: systemClass.
	u.systemObject = u.NewInstance(u.SystemClass)

	//     "Put special objects and classes into the dictionary of globals"
	//u.Iglobal: (self symbolFor: 'nil') put: nilObject.
	u.setGlobal(u.symbolFor("nil"), u.NilObject)
	// u.Iglobal: (self symbolFor: 'true') put: trueObject.
	u.setGlobal(u.symbolFor("true"), u.TrueObject)
	// u.Iglobal: (self symbolFor: 'false') put: falseObject.
	u.setGlobal(u.symbolFor("false"), u.FalseObject)
	// u.Iglobal: (self symbolFor: 'system') put: systemObject.
	// u.setGlobal("system", u.)
	// u.Iglobal: (self symbolFor: 'System') put: systemClass.
	u.setGlobal(u.symbolFor("System"), u.SystemClass.Object)
	// u.Iglobal: (self symbolFor: 'Block') put: blockClass.
	u.setGlobal(u.symbolFor("Block"), u.BlockClass.Object)
	// u.Iglobal: trueSymbol  put: trueClass.
	u.setGlobal(trueSymbol, u.TrueClass.Object)
	u.setGlobal(falseSymbol, u.FalseClass.Object)
	// u.Iglobal: falseSymbol put: falseClass.
	//     ^ systemObject

}

// newInstance: instanceClass = (
//     | result |
//     result := SObject new: instanceClass numberOfInstanceFields with: nilObject.
//     result somClass: instanceClass.

//	^ result
//
// )
func (u *Universe) NewInstance(c *Class) *Object {
	result := NewObject(c.NumberOfInstanceFields(), u.NilObject)
	result.SetSomClass(c)
	return result
}
