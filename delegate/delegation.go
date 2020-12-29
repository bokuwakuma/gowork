package main

import (
	"fmt"
)

var ACCELERATE_DIF = 3

// インタフェースと実装すべきメソッド
type Machine interface {
	Run()
	Accelerate()
}

func Accelerate(m Machine) {
	for i := 0; i < ACCELERATE_DIF; i++ {
		m.Run()
	}
}

// Machine <- Vehicle
type Vehicle struct {
}

// Vehicleのimplement
func (v *Vehicle) Run() {
	fmt.Println("走行する")
}

// Vehicleのimplementだがinterfaceで定義したメソッドに委譲
func (v *Vehicle) Accelerate() {
	Accelerate(v)
}

// Machine <- Airplane
type Airplane struct {
}

// Airplaneの個別メソッド
// *小文字は明示的にprivate
func (a *Airplane) fly() {
	fmt.Println("飛行する")
}

// Airplaneのimplement
// AirplaneにとってRunは飛ぶこと
func (a *Airplane) Run() {
	a.fly()
}

func (a *Airplane) Accelerate() {
	Accelerate(a)
}

func main() {
	v := new(Vehicle)
	fmt.Println("The Vehicle runs")
	v.Run()

	a := new(Airplane)
	fmt.Println("The Airplane runs")
	a.Run()

	fmt.Println("The Vehicle accelerates")
	v.Accelerate()

	fmt.Println("The Airplane accelerates")
	a.Accelerate()
}
