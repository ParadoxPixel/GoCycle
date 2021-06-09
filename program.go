package GoCycle

type Program struct {
	load []func()
	start []func()
	stop []func()
	err chan error
}

func NewProgram() *Program {
	return &Program{
		load:  []func(){},
		start: []func(){},
		stop:  []func(){},
		err:   nil,
	}
}

func (p *Program) AddLoader(loader... func()) {
	p.load = append(p.load, loader...)
}

func (p *Program) AddStart(start... func()) {
	p.start = append(p.start, start...)
}

func (p *Program) AddStop(stop... func()) {
	p.stop = append(p.stop, stop...)
}

func (p *Program) Start() {
	if p.err != nil {
		return
	}

	p.err = make(chan error)
	for _,f := range p.load {
		f()
	}

	for _,f := range p.start {
		go f()
	}

	<- p.err
	for _,f := range p.stop {
		f()
	}
}

func (p *Program) Stop() {
	if p.err == nil {
		return
	}

	p.err <- nil
	close(p.err)
	p.err = nil
}