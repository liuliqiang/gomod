package pkg3

type Pkg3 interface {
	SayHello()
}

type pkg3 struct {
}

func NewPkg3() Pkg3 {
	return &pkg3{}
}

func (p *pkg3) SayHello() {
	println("Hello from pkg version3")
}
