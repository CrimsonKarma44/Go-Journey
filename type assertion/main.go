package main

import "fmt"

type player struct {
	name string
}

func (p player) move() {
	fmt.Println(p.name)
}

func (p *player) moveNot() {
	fmt.Println(p.name)
}

type team interface {
	move()
}

type pointerTeam interface {
	moveNot()
}

func main() {
	var club team
	var clubPointer pointerTeam

	club = player{"Princewill"}
	clubPointer = &player{"Princewill"}

	t, ok := club.(player)
	fmt.Println(t.name, ok)
	v, ok := clubPointer.(*player)
	fmt.Println(v.name, ok)
}
