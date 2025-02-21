package internal

type Environment struct {
	Enclosing *Environment
	Variable  map[string]any
}

func NewEnvironment(enclosing Environment) *Environment {
	return &Environment{
		Enclosing: &enclosing,
		Variable:  make(map[string]any),
	}
}

func (env *Environment) Get(name string) any {
  if value, ok := env.Variable[name]; ok && value != nil {
		return value
	}

	if env.Enclosing != nil {
		return env.Enclosing.Get(name)
	}

	panic("Undefined variable " + name)
}

func (env *Environment) Define(name string, value any) {
	env.Variable[name] = value
}

func (env *Environment) Assign(name string, value any) {
	if _, exists := env.Variable[name]; exists {
		env.Variable[name] = value
		return
	}

	if env.Enclosing != nil {
		env.Enclosing.Assign(name, value)
		return
	}
	panic("Undefined variable " + name)
}

var env = Environment{
	Variable: make(map[string]any),
}
