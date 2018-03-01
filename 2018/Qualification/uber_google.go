package main

import "github.com/sirupsen/logrus"

type UberGoogle struct {
	R, C, F, N, B, T int
	cars             []Car
	rides            map[int]*ride
	myCars           map[int][]*Car
}

type Car struct {
	N       int
	History []int // ride IDs
	x, y    int

	currentScore float64
}

type ride struct {
	ID               int
	a, b, x, y, s, f int

	//timeToRide int

	scoreN []float64
	//nextRides []*decisionRide
}

type decisionRide struct {
	t          int
	score      float64
	otherRides []*ride
	ride       *ride
}

func (p *UberGoogle) moveAllCarsAtStep(step int) {
	for index, car := range p.myCars[step] {
		if len(p.rides) == 0 {
			return
		}
		for i, ride := range p.rides {

			if ride.s > step+abs(ride.a-car.x)+abs(ride.b-car.y)+1000 {
				break
			}
			longueurRide := abs(ride.a-ride.x) + abs(ride.b-ride.y)
			bonus := 0
			if max(ride.a+ride.b, ride.s) == ride.s {
				bonus = p.B
			}
			if longueurRide+max(ride.a+ride.b, ride.s) < ride.f {
				p.rides[i].scoreN[0] = float64(longueurRide+bonus) / float64(1+max(ride.a+ride.b, ride.s))
			} else {
				continue
			}
			p.rides[i].scoreN[1] = 0
			for _, ride1 := range p.rides {
				//ride1.scoreN[1] = ride1.scoreN[0]
				bonus := 0
				timeToReachStart := abs(ride1.a-car.x) + abs(ride1.b-car.y)
				if max(timeToReachStart, ride1.s) == ride1.s {
					bonus = p.B
				}

				longueurRide := abs(ride1.a-ride1.x) + abs(ride1.b-ride1.y)

				if longueurRide+max(timeToReachStart, ride1.s) < ride1.f {
					p.rides[i].scoreN[1] += float64(longueurRide+bonus) / float64(1+max(timeToReachStart, ride1.s))
				}
			}
			p.rides[i].scoreN[1] = p.rides[i].scoreN[1]/float64(len(p.rides)) + p.rides[i].scoreN[0]
		}

		// prendre le ride le plus avantageux
		log.WithFields(logrus.Fields{
			"index":           index,
			"remaining rides": len(p.rides),
		}).Debug("picking the best ride")
		max := 0.0
		indexMax := -1
		for i, ride := range p.rides {
			if indexMax == -1 || ride.scoreN[1] > max {
				max = ride.scoreN[1]
				indexMax = i
			}
		}
		car.currentScore += max
		//for currentRide := p.rides[0];;{
		//
		//}
		rideIndex := indexMax      // l'ID du ride
		ride := p.rides[rideIndex] // ide
		car.History = append(car.History, ride.ID)
		// calculer le nouveau N
		timeToReachStart := abs(ride.a-car.x) + abs(ride.b-car.y)
		timeToRide := abs(ride.a-ride.x) + abs(ride.b-ride.y)

		newN := car.N + timeToReachStart + timeToRide
		log.WithFields(logrus.Fields{
			"timeToReachStart": timeToReachStart,
			"timeToRide":       timeToRide,
		}).Debug("calculating new N")

		// enlever le ride effectué
		delete(p.rides, rideIndex)
		log.WithFields(logrus.Fields{
			"ride": rideIndex,
		}).Debug("deleting ride")

		// permuter la Car vers son nouvel emplacement N
		log.WithFields(logrus.Fields{
			"New N": newN,
		}).Debug("update myCars map")
		car.N = newN
		p.myCars[newN] = append(p.myCars[newN], car)
		p.myCars[step][index] = nil
	}
}

//pickARide := rand.Int() % len(p.rides)
//i := 0
//for rideIndex := range p.rides {
//	if i == pickARide {
//		pickARide = rideIndex
//	}
//	i++
//}
//ride := p.rides[pickARide]
//car.History = append(car.History, ride.ID)
