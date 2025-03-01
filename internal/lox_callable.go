package internal

import "time"

type LoxCallable interface {
	Call(i *Interpreter, arguments *[]any) any
	Arity() int
}

type GlobalClock struct{}

func (c *GlobalClock) Call(i *Interpreter, arguments *[]any) any {
	return time.Now()
}

func (c *GlobalClock) Arity() int {
	return 0
}

func (c *GlobalClock) String() string {
  return "<native fn>"
}


