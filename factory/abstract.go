package factory

type A struct {
}

type B struct {
}

type AbstractFactory interface {
	CreatA() A
	CreatB() B
}

type Factory struct {
}

func (f *Factory) CreatA() A {
	panic("implement me")
}

func (f *Factory) CreatB() B {
	panic("implement me")
}
