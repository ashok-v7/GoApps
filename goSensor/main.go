package main

import "fmt"

type Thermostat struct {
	ID    string
	value float64
}

//Go supports methods defined on struct types.

func (t *Thermostat) Value() float64 {
	return t.value
}

func (t *Thermostat) Set(value float64) {

	t.value = value
}

func (*Thermostat) kind() string {
	return "thermostat"
}

func main() {
	t := Thermostat{ID: "Living room", value: 16.20}
	fmt.Printf("%s before: %.2f\n", t.ID, t.Value())

	t.Set(18)
	fmt.Printf("%s after : %.2f\n", t.ID, t.Value())

	fmt.Printf("%s\n", t.kind())
}
