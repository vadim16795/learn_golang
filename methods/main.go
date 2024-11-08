package main

import (
	"fmt"
	"math/rand"
)

type House struct {
	garage bool
	size   int
}

// type Data int

// func(reciver_name Type) method_name(parameter_list)(return_type){
// Code
// }

type Player interface {
	Kickball()
	Name()
}

type FootballPlayer struct {
	stamina int
	power   int
	name    string
}

func (f FootballPlayer) Name() {
	fmt.Println("hello")

}

func (f FootballPlayer) Kickball() {
	shot := f.stamina + f.power
	fmt.Println("I'm kicking the ball", shot)
}

func (r *House) HasGarage() bool {
	r.garage = true
	return r.garage
}

func (r *House) SizeGarage() int {
	r.size = 10
	return r.size
}

// func (r Data) multiple(v Data) Data {
// 	return r * v
// }

func main() {

	house := House{}
	house.HasGarage()
	house.SizeGarage()
	team := make([]Player, 11)
	for i := 0; i < len(team); i++ {
		team[i] = FootballPlayer{
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
	for i := 0; i < len(team); i++ {
		// team[i].Name()
		team[i].Kickball()

	}
	// x := Data(15)
	// y := Data(20)
	// res := x.multiple(y)
	// fmt.Println(res)

}
