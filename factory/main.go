package main

import (
	"github.com/k0kubun/pp"
)

type Pet interface {
	Bark() string
}

type Dog struct {
}

func (d Dog) Bark() string {
	return "わんわん"
}

type Cat struct {
}

func (c Cat) Bark() string {
	return "にゃあにゃあ"
}

type PetType int

const (
	DogType PetType = 1 << iota
	CatType
)

func NewPet(t PetType) Pet {
	switch t {
	case DogType:
		return Dog{}
	case CatType:
		return Cat{}
	default:
		return Dog{}
	}
}

func main() {
	d := NewPet(DogType)
	pp.Println(d.Bark())

	c := NewPet(CatType)
	pp.Println(c.Bark())
}
