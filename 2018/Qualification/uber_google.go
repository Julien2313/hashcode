package main

import (
	"fmt"
	"sort"
	"os"
	"math/rand"
)

type UBERGOOGLE struct {
	Row, Column, FleetNbr, NbrRides, Bonus, TotalTicks int
	cars             []*CAR
	rides            []*RIDE
	ridesMissed      []*RIDE
	Distances        map[int]map[int]int
	CptRidesMissed  int
	RealTimeScoring int
}

type CAR struct {
	ID int
	nbrTickUsed  int
	History       []RIDE
	x, y          int
	done bool
}

type RIDES []*RIDE
type RIDE struct {
	ID               int
	x0, y0, x1, y1, start, finish int
	Score   int
	tickStarted int
	averageDistances int
	closestRide int
	lenght int
}

func (ride *RIDE) FindClosestRideNearEnd(rides []*RIDE, p *UBERGOOGLE, tickStart int) int{

	if len(rides) == 1 {
		return 0
	}
	nbrMaxRides := 100
	closest := 999999

	for cptRides := 0; cptRides < nbrMaxRides; cptRides++ {
		randRide := rand.Intn(len(rides))
		if ride == rides[randRide] {
			cptRides--
			continue
		}
		//dist := int(p.Distances[ride.ID][rides[randRide].ID])
		dist := ride.DistanceEndToStart(rides[randRide])
		if (tickStart + dist) + rides[randRide].lenght > rides[randRide].finish {
			continue
		}
		if dist < closest {
			closest = dist
			if dist == 0 {
				break
			}
		}
	}

	return closest
}

func (car *CAR) verifyHistory() bool {
	car_ := CAR{}
	car_.x, car_.y = 0, 0
	tick := 0
	for cptRide, ride := range car.History {

		tick += car_.Distance(&ride)
		tick += max(0, ride.start - tick)
		tick += ride.lenght

		if tick > ride.finish {
			fmt.Println(car.History[cptRide-2].ID, car.History[cptRide-1].ID, ride.ID)
			fmt.Println()
			return false
		}
		car_.x, car_.y = ride.x1, ride.y1
	}
	return true
}

func (thisRide *RIDE) ComputeAverageDistance(rides []*RIDE, p *UBERGOOGLE) {

	if len(rides) == 1 {
		thisRide.averageDistances = 0
		return
	}
	averageDistances := 0
	nbrMaxRides := 100

	for cptRides := 0; cptRides < nbrMaxRides; cptRides++ {
		randRide := rand.Intn(len(rides))
		if thisRide == rides[randRide] {
			cptRides--
			continue
		}
		/*if p.Distances[thisRide.ID][rides[randRide].ID] == 0 {
		  p.Distances[thisRide.ID][rides[randRide].ID] = int32(thisRide.DistanceEndToStart(rides[randRide]))
		}
		averageDistances += p.Distances[thisRide.ID][rides[randRide].ID]*/

		averageDistances += thisRide.DistanceEndToStart(rides[randRide])
	}
	thisRide.averageDistances = averageDistances / nbrMaxRides
}

func (ride *RIDE) markRide(currentTick int, distanceFromStartRide int, p *UBERGOOGLE, pCar *CAR) {
	tickStartRide := currentTick + distanceFromStartRide
	ride.Score = 0
	if fileName == "d_metropolis" {
		ride.Score -= ride.FindClosestRideNearEnd(p.rides, p, currentTick + max(0, ride.start - tickStartRide))
	} else {
		ride.ComputeAverageDistance(p.rides, p)
		ride.Score -= ride.averageDistances
	}

	if fileName != "c_no_hurry" {
		if ride.start >= tickStartRide {
			ride.Score += p.Bonus
		}
		ride.Score -= ride.finish*2
	}
	ride.Score -= max(0, ride.start - tickStartRide) + distanceFromStartRide
	if tickStartRide + ride.lenght > ride.finish {
		ride.Score = -9999999
	}
}

func (thisRide *RIDE) DistanceEndToStart(ride *RIDE) int {
	return abs(thisRide.x1 - ride.x0) + abs(thisRide.y1 - ride.y0)
}

func (ride *RIDE) Lenght() int {
	return abs(ride.x1 - ride.x0) + abs(ride.y1 - ride.y0)
}

func (pCar *CAR) Distance(ride *RIDE) int {
	return abs(pCar.x - ride.x0) + abs(pCar.y - ride.y0)
}

func (pCar *CAR) AddRide(rideToAdd *RIDE, tickWhenArrivedAtRide int, p *UBERGOOGLE) {
	rideToAdd.tickStarted = tickWhenArrivedAtRide

	p.RealTimeScoring += rideToAdd.lenght
	if rideToAdd.start >= tickWhenArrivedAtRide {
		p.RealTimeScoring += p.Bonus
	}

	rideToAdd.tickStarted += max(0, rideToAdd.start - rideToAdd.tickStarted)

	if rideToAdd.tickStarted + rideToAdd.lenght > rideToAdd.finish {
		fmt.Println("Took but no points")
		fmt.Println(rideToAdd)
		os.Exit(1)
	}
	pCar.History = append(pCar.History, *rideToAdd)

	pCar.nbrTickUsed = pCar.Distance(rideToAdd) + rideToAdd.lenght + max(0, rideToAdd.start - tickWhenArrivedAtRide)
	pCar.x, pCar.y = rideToAdd.x1, rideToAdd.y1
	if pCar.nbrTickUsed < 0 {
		fmt.Println("nbrTickUsed < 0")
		os.Exit(1)
	}
}

func (pCar *CAR) ChooseRide(p *UBERGOOGLE, currentTick int) bool{
	for _, ride := range p.rides {
		ride.markRide(currentTick, pCar.Distance(ride), p, pCar)
	}

	sort.Sort(RIDES(p.rides))

	if len(p.rides) == 0 {
		return false
	}

	if p.rides[0].Score != -9999999 {
		pCar.AddRide(p.rides[0], currentTick + pCar.Distance(p.rides[0]), p)
		/*
		for i, ride := range p.rides {
		  if ride.ID == p.rides[0].closestRide {
		    pCar.AddRide(p.rides[i], tick + pCar.nbrTickUsed)
		    p.rides = append(p.rides[:i], p.rides[i+1:]...)
		    break
		  }
		}*/

		p.rides = append(p.rides[:0], p.rides[0+1:]...)
		//p.RemoveDistancesOfRide(p.rides[0].ID)
	} else {
		pCar.done = true
	}
	return true
}

func (p *UBERGOOGLE) RemoveDistancesOfRide(numRide int) {
	delete(p.Distances, numRide)
	/*for cptRides := 0; cptRides < p.N; cptRides++ {
	  delete(p.Distances[cptRides], numRide)
	}*/
}

func (p *UBERGOOGLE) moveAllCarsAtTick(currentTick int) int {
	for numRide := len(p.rides) - 1; numRide >= 0; numRide--{
		if currentTick > p.rides[numRide].finish {
			//p.RemoveDistancesOfRide(p.rides[numRide].ID)
			p.ridesMissed =  append(p.ridesMissed, p.rides[numRide])
			p.rides = append(p.rides[:numRide], p.rides[numRide+1:]...)
			p.CptRidesMissed++
		}
	}

	for _, car := range p.cars {
		if car.nbrTickUsed > 0 				      { break }
		if car.done            				      { continue }
		if !car.ChooseRide(p, currentTick) 	{ break }
	}

	sort.Sort(p)
	toMinus := 0

	for _, car := range p.cars {
		if car.nbrTickUsed == 0 { continue}
		if toMinus == 0         { toMinus = car.nbrTickUsed ; car.nbrTickUsed = 0
		} else                  { car.nbrTickUsed -= toMinus }
	}

	if toMinus == 0 {
		if len(p.rides) == 0 {
			return p.TotalTicks
		} else {
			toMinus = 1
		}
	}
	currentTick += toMinus
	return currentTick
}


func (p RIDES) Len() int                  { return len(p) }
func (p RIDES) Swap(i, j int)             { p[i], p[j] = p[j], p[i] }
func (p RIDES) Less(i, j int) bool        { return p[i].Score > p[j].Score }

func (p *UBERGOOGLE) Len() int            { return len(p.cars) }
func (p *UBERGOOGLE) Swap(i, j int)       { p.cars[i], p.cars[j] = p.cars[j], p.cars[i] }
func (p *UBERGOOGLE) Less(i, j int) bool  { return p.cars[i].nbrTickUsed < p.cars[j].nbrTickUsed }
