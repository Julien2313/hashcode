package main

import (
	"fmt"
	"time"
	"math/rand"
)

const (
	//fileName = "a_example"
	//fileName = "b_should_be_easy"
	fileName = "c_no_hurry"
	//fileName = "d_metropolis"
	//fileName = "e_high_bonus"

)

func main() {

	start := time.Now()
	rand.Seed(time.Now().Unix())
	for {
		fmt.Println("go")
		myS := &UBERGOOGLE{}
		myS.read()

		fmt.Println("Init finished !")
		cptStep := 0
		for step := 0; step < myS.TotalTicks; {
			step = myS.moveAllCarsAtTick(step)
			cptStep++
			if cptStep%1 == 0 {
				fmt.Println(step, "/", myS.TotalTicks, " : ", len(myS.rides), myS.CptRidesMissed, myS.RealTimeScoring)
			}
			if len(myS.rides) == 0{
				break
			}
		}

		for _,car := range myS.cars {
			break
			for cptRide := 0; cptRide < len(car.History) - 1; cptRide++ {
				ride1 := car.History[cptRide]
				ride2 := car.History[cptRide+1]
				tickWhenEndRide1 := ride1.tickStarted + ride1.lenght
				numberTicksFree := ride2.tickStarted - (tickWhenEndRide1 + ride1.DistanceEndToStart(&ride2))

				fmt.Println(numberTicksFree, tickWhenEndRide1)

				for cptRideMissed := 0; cptRideMissed < len(myS.ridesMissed); cptRideMissed++ {
					rideMissed := myS.ridesMissed[cptRideMissed]

					tickWhenStartRideMissed := tickWhenEndRide1 + ride1.DistanceEndToStart(rideMissed) + max(0, rideMissed.start - (tickWhenEndRide1 + ride1.DistanceEndToStart(rideMissed)))

					if rideMissed.start > tickWhenStartRideMissed {
						continue
					}
					tickWhenEnded := tickWhenStartRideMissed + rideMissed.lenght
					if rideMissed.finish > tickWhenEnded {
						continue
					}
					timeUsed := ride1.DistanceEndToStart(rideMissed) + rideMissed.lenght + rideMissed.DistanceEndToStart(&ride2)

					fmt.Println("-->", timeUsed)
				}
			}
		}

		myS.write()
	}

	elapsed := time.Since(start)
	fmt.Println("Took", elapsed)
}
