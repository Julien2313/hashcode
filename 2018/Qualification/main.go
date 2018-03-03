package main

import (
	"fmt"
)

const (
	fileName = "a_example"
	//fileName = "b_should_be_easy"
	//fileName = "c_no_hurry"
	//fileName = "d_metropolis"
	//fileName = "e_high_bonus"

)

func main() {

	fmt.Println("go")
	myS := &UBERGOOGLE{}
	myS.read()

	fmt.Println()
	for step := 0; step < myS.T; {
		fmt.Println(step, "/", myS.T, " : ", len(myS.rides))
		step = myS.moveAllCarsAtStep(step)
	}
	myS.write()
}
