package main

func main() {
  r := Rubyist{}
  g := Gopher{}

  r.SayIt()
  g.SayIt()

  Bored()
}

func main() {
  r := Rubyist{}
  g := Gopher{}

  go r.SayIt()
  go g.SayIt()

  Bored()
}
