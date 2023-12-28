package pkg1

type Pkg1 interface {
	SayHello()
}

type pkg1 struct {
}

func NewPkg1() Pkg1 {
	return &pkg1{}
}

func (p *pkg1) SayHello() {
	println("Hello from pkg1")
}
