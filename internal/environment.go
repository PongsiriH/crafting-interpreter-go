package internal

type Environment struct {
  Variable map[string]any
}

func (env *Environment) Get(name string) any {
  value, ok := env.Variable[name]
  if !ok {
    panic("Undefined variable " + name)
  }
  return value
}

func (env *Environment) Define(name string, value any) {
  env.Variable[name] = value
}

var env = Environment{
  Variable: make(map[string]any),
}
