package GoCycle

var p *Program
func SetGlobal(pr *Program) {
	p = pr
}

func GetGlobal() *Program {
	return p
}