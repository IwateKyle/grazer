package grazer

import "fmt"

type Processor[Out any, In any] struct {
	steps    []func(out *Out, in *In) error
	Err      error
	Response *Out
	Request  *In
}

func NewProcessor[Out any, In any](response *Out, request *In) *Processor[Out, In] {
	fmt.Println("**** : Response : ", response)
	return &Processor[Out, In]{
		Response: response,
		Request:  request,
	}
}

func (p *Processor[Out, In]) Use(step func(response *Out, request *In) error) {
	p.steps = append(p.steps, step)
}

func (p *Processor[Out, In]) Process() {
	for i, step := range p.steps {

		fmt.Println(i, " : ", p.Response, p.Request)
		err := step(p.Response, p.Request)
		if err != nil {
			p.Err = err
			break
		}
	}
}
