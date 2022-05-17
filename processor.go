package grazer

import "fmt"

type Processor[Out any, In any] struct {
	steps    []func(out *Out, in *In) error
	err      error
	response *Out
	request  *In
}

func NewProcessor[Out any, In any](response *Out, request *In) *processor[Out, In] {
	fmt.Println("**** : response : ", response)
	return &processor[Out, In]{
		response: response,
		request:  request,
	}
}

//func (p *processor) use(step processStep) {
func (p *processor[Out, In]) Use(step func(response *Out, request *In) error) {
	p.steps = append(p.steps, step)
}

//func (p *processor) process(response *RegistrationResponse, form *RegistrationForm) {
func (p *processor[Out, In]) Process() {
	for i, step := range p.steps {

		fmt.Println(i, " : ", p.response, p.request)
		err := step(p.response, p.request)
		if err != nil {
			p.err = err
			break
		}
	}
}
