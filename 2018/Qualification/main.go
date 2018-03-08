package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
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

	bestC, bestD, bestE, cpt := 0,0,0,0

	for {
		cpt++
		fmt.Println("cpt:", cpt)
		myS := &UBERGOOGLE{}
		myS.read()

		//fmt.Println("Init finished !")
		cptStep := 0
		for step := 0; step < myS.TotalTicks; {
			step = myS.moveAllCarsAtTick(step)
			cptStep++
			//if cptStep%100 == 0 {
			//	fmt.Println(step, "/", myS.TotalTicks, " : ", len(myS.rides), myS.CptRidesMissed, myS.RealTimeScoring)
			//}
			if len(myS.rides) == 0{
				break
			}
		}

		for _,car := range myS.cars {
			for cptRide := 0; cptRide < len(car.History) - 1; cptRide++ {
				ride1 := car.History[cptRide]
				ride2 := car.History[cptRide+1]
				tickWhenEndRide1 := ride1.tickStarted + ride1.lenght

				numberTicksFree := ride2.tickStarted - (tickWhenEndRide1 + ride1.DistanceEndToStart(&ride2))

				for cptRideMissed := 0; cptRideMissed < len(myS.ridesMissed); cptRideMissed++ {
					rideMissed := myS.ridesMissed[cptRideMissed]
					tickWhenStartRideMissed := tickWhenEndRide1 + ride1.DistanceEndToStart(rideMissed)
					tickWhenStartRideMissed += max(0, rideMissed.start - tickWhenStartRideMissed)
					if rideMissed.start > tickWhenStartRideMissed { continue }

					tickWhenEnded := tickWhenStartRideMissed + rideMissed.lenght
					if rideMissed.finish < tickWhenEnded { continue }
					if tickWhenEnded + rideMissed.DistanceEndToStart(&ride2) > ride2.start { continue }

					timeUsed := ride1.DistanceEndToStart(rideMissed) + rideMissed.lenght + rideMissed.DistanceEndToStart(&ride2)
					if timeUsed > numberTicksFree { continue }
					rideMissed.tickStarted = tickWhenStartRideMissed
					car.History = append(car.History, RIDE{})
					copy(car.History[cptRide+2:], car.History[cptRide+1:])
					car.History[cptRide+1] = *rideMissed
					if !car.verifyHistory() {
						fmt.Println("Ids of the rides:", ride1.ID, rideMissed.ID, ride2.ID)
						fmt.Println("End previous ride:", tickWhenEndRide1)
						fmt.Println("Distance of the two rides:", ride1.DistanceEndToStart(rideMissed))
						fmt.Println("When the missed ride can start:", rideMissed.start)
						fmt.Println("Lenght of the missed ride:", rideMissed.lenght)
						fmt.Println("End of the missed ride:", rideMissed.finish)
						fmt.Println("distance from missed ride to ride2:", rideMissed.DistanceEndToStart(&ride2))
						fmt.Println("When the ride2 can start:", ride2.start)
						fmt.Println("When the ride2 is suppose to start:", ride2.tickStarted)
						fmt.Println("When the ride2 is suppose to end:", ride2.finish)
						fmt.Println("The free time:", numberTicksFree)
						os.Exit(1)
					}
					ride2 = *rideMissed
					myS.ridesMissed = append(myS.ridesMissed[:cptRideMissed], myS.ridesMissed[cptRideMissed+1:]...)
					cptRideMissed--

				}
			}
		}
		myS.write(&bestC, &bestD, &bestE)
	}

	elapsed := time.Since(start)
	fmt.Println("Took", elapsed)
}
