package checker

// operations
const (
	PLUS  = "PLUS"
	EQUAL = "EQ"
	LT    = "LT"
	GT    = "GT"
	MINUS = "MINUS"
	TIMES = "TIMES"
	AND   = "AND"
	OR    = "OR"
	PRINT = "PRINT"
)

// variable types
const (
	INT_TYPE     = "Int"
	STRING_TYPE  = "String"
	BOOL_TYPE    = "Bool"
	NOTHING_TYPE = "Nothing"
	SET_TYPE     = "Set"
	VAL_TYPE     = "Val"
)

type Signature struct {
	Return string
	Params []string // list of types
}

type Methods map[string]Signature

// type methods
var TypeTable = map[string]Methods{
	SET_TYPE: {
		PLUS: {
			Return: SET_TYPE,
			Params: []string{VAL_TYPE, SET_TYPE},
		},
	},
	VAL_TYPE: {
		PLUS: {
			Return: SET_TYPE,
			Params: []string{SET_TYPE, VAL_TYPE},
		},
		PRINT: {
			Return: NOTHING_TYPE,
			Params: []string{},
		},
	},
	STRING_TYPE: {
		PLUS:  {STRING_TYPE, []string{STRING_TYPE}},
		PRINT: {NOTHING_TYPE, []string{}}},
	BOOL_TYPE: {
		AND:   {BOOL_TYPE, []string{BOOL_TYPE}},
		OR:    {BOOL_TYPE, []string{BOOL_TYPE}},
		PRINT: {NOTHING_TYPE, []string{}}}}

type Environment struct {
	Vals  map[string]string    // map identifier to type
	Funcs map[string]Signature // map function name to return type
	Types map[string]bool      // track valid types
}

var env Environment // set global

func IsBuiltin(name string) bool {
	return name == "PRINT"
}

func NewEnvironment() Environment {
	return Environment{Vals: map[string]string{}, Funcs: map[string]Signature{}, Types: map[string]bool{}}
}

func MethodExist(kind, method string) bool {
	methods, ok := TypeTable[kind]
	if !ok {
		return false
	}

	_, ok = methods[method]
	return ok
}

func GetMethod(kind, method string) (Signature, bool) {
	methods, ok := TypeTable[kind]
	if !ok {
		return Signature{}, false
	}

	sig, ok := methods[method]
	return sig, ok
}

func (e *Environment) Set(name, kind string) {
	e.Vals[name] = kind
}

func SetFunctionSignature(name string, sig Signature) {
	env.Funcs[name] = sig
}

func GetFunctionSignature(name string) (Signature, bool) {
	kind, ok := env.Funcs[name]
	return kind, ok
}

func (e *Environment) Get(name string) (string, bool) {
	kind, ok := e.Vals[name]
	return kind, ok
}

func (e *Environment) IdentExist(kind string) bool {
	_, ok := e.Vals[kind]
	return ok
}

func GetIdentType(name string) (string, bool) {
	kind, ok := env.Vals[name]
	return kind, ok
}

func (e *Environment) TypeExist(kind string) bool {
	_, ok := e.Types[kind]
	return ok
}
