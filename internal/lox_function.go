package internal

type LoxFunction struct {
	Declaration FunctionStmt
	Closure     *Environment
}

func NewLoxFunction(declaration FunctionStmt) LoxFunction {
	return LoxFunction{
		Declaration: declaration,
	}
}

func (lx *LoxFunction) Call(i *Interpreter, args *[]any) any {
	env := i.globalEnv
	for j := 0; j < len(lx.Declaration.Params); j++ {
		env.Define(lx.Declaration.Params[j].Lexeme, (*args)[j])
	}
	lx.Declaration.Body.Apply(i)
  return nil
}

func (lx *LoxFunction) Arity() int {
  return len(lx.Declaration.Params)
}
