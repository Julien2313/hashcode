package main

import (
	"fmt"
	"sort"
	"os"
	"math/rand"
)

type UBERGOOGLE struct {
	R, C, F, N, B, T int
	cars             []*CAR
	rides            []*RIDE
}

type CAR struct {
	nbrTickUsed  int
	History       []RIDE
	x, y          int
}

type RIDES []*RIDE
type RIDE struct {
	ID               int
	x0, y0, x1, y1, s, f int
	Score   float64
	tickStarted int
	averageDistance float64
}

func (thisRide *RIDE) ComputeAverageDistance(rides []*RIDE) {
	thisRide.averageDistance = 0
	nbrMaxRides := 100

	for cptRides := 0; cptRides < nbrMaxRides; cptRides++ {
		randRide := rand.Intn(len(rides))
		if thisRide == rides[randRide] {
			continue
		}
		thisRide.averageDistance += float64(thisRide.Distance(rides[randRide]))
	}

	thisRide.averageDistance /= float64(nbrMaxRides)
}

func (ride *RIDE) markRide(tickStart int, bonus int, distance int, pCar *CAR, p *UBERGOOGLE) {

	ride.ComputeAverageDistance(p.rides)
	ride.Score = 0
	if ride.s >= tickStart {
		ride.Score += float64(bonus)
	}
	ride.Score -= float64(ride.f)
	ride.Score -= float64(max(0, ride.s - tickStart) + distance)
	ride.Score -= ride.averageDistance
	if tickStart + ride.Lenght() > ride.f {
		ride.Score = -9999999
	}
}

func (thisRide *RIDE) Distance(ride *RIDE) int {
	return abs(thisRide.x1 - ride.x0) + abs(thisRide.y1 - ride.y0)
}

func (ride *RIDE) Lenght() int {
	return abs(ride.x1 - ride.x0) + abs(ride.y1 - ride.y0)
}

func (pCar *CAR) Distance(ride *RIDE) int {
	return abs(pCar.x - ride.x0) + abs(pCar.y - ride.y0)
}
//4 34 65 266
//222
func (pCar *CAR) AddRide(ride *RIDE, tick int) {
	if ride.ID ==  {
		fmt.Println(ride)
		os.Exit(1)
	}
	ride.tickStarted = tick+pCar.Distance(ride)
	if ride.tickStarted + ride.Lenght() > ride.f {
		fmt.Println("Took but no points")
		fmt.Println(ride.tickStarted, ride.Lenght(), ride.f)
		os.Exit(1)
	}
	pCar.History = append(pCar.History, *ride)

	pCar.x, pCar.y = ride.x1, ride.y1
	pCar.nbrTickUsed = pCar.Distance(ride) + ride.Lenght() + max(0, ride.s - ride.tickStarted)
	if pCar.nbrTickUsed < 0 {
		fmt.Println("nbrTickUsed < 0")
		os.Exit(1)
	}
}

func (pCar *CAR) ChooseRide(p *UBERGOOGLE, tick int, bonus int) bool{

	for _, ride := range p.rides {
		ride.markRide(tick + pCar.Distance(ride), bonus, pCar.Distance(ride), pCar, p)
	}

	sort.Sort(RIDES(p.rides))

	if len(p.rides) == 0 {
		return false
	}

	if p.rides[0].Score != -9999999 {
		pCar.AddRide(p.rides[0], tick)
		p.rides = append(p.rides[:0], p.rides[0+1:]...)
	}
	return true
}

func (p *UBERGOOGLE) moveAllCarsAtStep(tick int) int {
	for numRide := len(p.rides) - 1; numRide >= 0; numRide--{
		if tick > p.rides[numRide].f {
			p.rides = append(p.rides[:numRide], p.rides[numRide+1:]...)
		}
	}

	for _, car := range p.cars {
		if car.nbrTickUsed > 0 				{ break }
		if !car.ChooseRide(p, tick, p.B) 	{ break }
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
			return p.T
		} else {
			toMinus = 1
		}
	}
	tick += toMinus
	return tick
}


func (p RIDES) Len() int                  { return len(p) }
func (p RIDES) Swap(i, j int)             { p[i], p[j] = p[j], p[i] }
func (p RIDES) Less(i, j int) bool        { return p[i].Score > p[j].Score }

func (p *UBERGOOGLE) Len() int            { return len(p.cars) }
func (p *UBERGOOGLE) Swap(i, j int)       { p.cars[i], p.cars[j] = p.cars[j], p.cars[i] }
func (p *UBERGOOGLE) Less(i, j int) bool  { return p.cars[i].nbrTickUsed < p.cars[j].nbrTickUsed }
