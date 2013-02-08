package main

type Fanboy interface {
	Rant()
}

type Rubyist struct {}
type Gopher struct {}

func (r Rubyist) Rant() string {
	"Test all the fucking time"
}

func (g Gopher) Rant() string {
	"Concurrency is king"
}

func (f Fanboy) SayIt() {
	f.Rant()
}

func main() {
}
