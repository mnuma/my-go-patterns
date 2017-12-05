package main

import (
    "github.com/k0kubun/pp"
    "strconv"
)

type Car interface {
    Drive() string
    Stop() string
}

type CarBuilder interface {
    TopSpeed(int) CarBuilder
    Paint(string) CarBuilder
    Build() Car
}

type car struct {
    topSpeed int
    color    string
}

type carBuilder struct {
    speedOption int
    color string
}

func (cb *carBuilder) TopSpeed(speed int) CarBuilder {
    cb.speedOption = speed
    return cb
}

func (cb *carBuilder) Paint(color string) CarBuilder {
    cb.color = color
    return cb
}

func (cb *carBuilder) Build() Car {
    return &car{
        topSpeed: cb.speedOption,
        color:    cb.color,
    }
}

func New() CarBuilder {
    return &carBuilder{}
}

func (c *car) Drive() string {
    return "Driving at speed: " + strconv.Itoa(c.topSpeed)
}

func (c *car) Stop() string {
    return "Stopping a " + string(c.color) + " car"
}

func main() {
    builder := New()
    car := builder.TopSpeed(50).Paint("BLUE").Build()
    pp.Println(car)
    pp.Println(car.Drive())
    pp.Println(car.Stop())
}
