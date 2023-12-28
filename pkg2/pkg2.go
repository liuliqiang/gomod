package pkg2

type Pkg2 interface {
	SayHello()
}

type pkg2 struct {
}

func NewPkg2() Pkg2 {
	return &pkg2{}
}

func (p *pkg2) SayHello() {
	println("Hello from pkg1")
}
