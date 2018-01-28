package mockingaround

import "fmt"

type Something struct {
	counter int
}

func (s *Something) Increment() error {
	s.counter++
	return nil
}

// Something Interface is something for something to implement
type SomethingInterface interface {
	Increment() error
}

func actionOnSomething(s SomethingInterface) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	if err := s.Increment(); err != nil {
		panic("arrghhhh")
	}
}
