package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
	"fmt"
)

const PATH = "E:/Julien/Documents/GitHub/hashcode/2018/Qualification/"

func (p *UBERGOOGLE) read() {
	b, err := ioutil.ReadFile(PATH + "input/" + fileName + ".in")
	if err != nil {
		panic(err)
	}
	p.decode(b)
}

func (p *UBERGOOGLE) decode(b []byte) {
	tmp := bytes.NewReader(b)
	s := bufio.NewScanner(tmp)
	s.Scan()

	args := toSliceOfInt(s.Text())
	p.Row, p.Column, p.FleetNbr, p.NbrRides, p.Bonus, p.TotalTicks = args[0], args[1], args[2], args[3], args[4], args[5]
	p.CptRidesMissed = 0
	p.RealTimeScoring = 0
	//fmt.Println(p)
	p.rides = make([]*RIDE, p.NbrRides)
	p.ridesMissed = make([]*RIDE, 0)

	for numRide := 0; numRide < p.NbrRides; numRide++ {
		s.Scan()
		args := toSliceOfInt(s.Text())
		p.rides[numRide] = &RIDE{numRide, args[0], args[1], args[2], args[3], args[4], args[5], 0.0, 0, 0, 0, 0}
		p.rides[numRide].lenght = p.rides[numRide].Lenght()
	}
	/*
		p.Distances = make(map[int]map[int]int, p.NbrRides)
		for cptDistances := 0; cptDistances < p.NbrRides; cptDistances++ {
		  p.Distances[cptDistances] = make(map[int]int, p.NbrRides)
		  for cptDistance := 0; cptDistance < p.NbrRides; cptDistance++ {
			p.Distances[cptDistances][cptDistance] = p.rides[cptDistances].DistanceEndToStart(p.rides[cptDistance])
		  }
		  if cptDistances % 100 == 0 {
			fmt.Println(cptDistances, "/", p.NbrRides)
		  }
		}*/

	p.cars = make([]*CAR, p.FleetNbr)
	for cptCar := 0; cptCar < p.FleetNbr; cptCar++ {
		p.cars[cptCar] = &CAR{cptCar, 0, make([]RIDE, 0), 0, 0, false}
	}
}
func (p *UBERGOOGLE) encode(bestC, bestD, bestE *int) string {
	var output string
	totScore := 0
	cptTookButNotGood := 0
	cptTookTot := 0
	first := true
	realScore := 0
	for _, c := range p.cars {
		if !first {
			output += "\n"
		} else {
			first = false
		}
		output += strconv.Itoa(len(c.History))

		tick := 0
		c.x, c.y = 0, 0
		for _, ride := range c.History {

			bonus := false
			if ride.ID == 8737 {
				//fmt.Println("->", tick, ", ID du rid précédent", c.History[cptRide-1].ID, ", quand est-ce que l'on démare:", c.History[cptRide-1].start, ride.start)
			}
			if ride.ID == 8737 {
				//fmt.Println("->distance", c.Distance(&ride))
			}
			tick += c.Distance(&ride)
			if ride.ID == 8737 {
				//fmt.Println("->", tick)
			}
			if tick <= ride.start {
				bonus = true
			}
			tick += max(0, ride.start - tick)
			if ride.ID == 8737 {
				//fmt.Println("->", tick)
			}
			tick += ride.lenght

			if ride.ID == 8737 {
				//fmt.Println("->", tick)
			}
			if tick <= ride.finish {
				realScore += ride.lenght
				if bonus {
					realScore += p.Bonus
				}
			} else {
				//fmt.Println("!", ride.ID)
				//os.Exit(1)
			}
			c.x, c.y = ride.x1, ride.y1

			if ride.tickStarted + ride.lenght > ride.finish {
				cptTookButNotGood++
			} else {
				cptTookTot++
				totScore += ride.lenght
				if ride.tickStarted <= ride.start {
					totScore += p.Bonus
				}

				if totScore != realScore {
					//fmt.Println("!", ride.ID)
					//fmt.Println("lààààààà")
					//os.Exit(1)
				}
			}
			output += " " + strconv.Itoa(ride.ID)

		}
	}
	if  fileName == "a_example" || fileName == "b_should_be_easy" || fileName == "c_no_hurry" && realScore > *bestC || fileName == "d_metropolis" && realScore > *bestD  || fileName == "e_high_bonus" && realScore > *bestE {
		//fmt.Println(output)
		if fileName == "c_no_hurry" {
			*bestC = realScore
		} else if fileName == "d_metropolis"{
			*bestD = realScore
		}else if fileName == "e_high_bonus"{
			*bestE = realScore
		}
		fmt.Println(*bestC, *bestD, *bestE)
		//fmt.Println("Total real Score     : ", realScore)
		//fmt.Println("Total Score          : ", totScore)
		//fmt.Println("Took but no points   : ", cptTookButNotGood)
		//fmt.Println("Tot tooked 	     : ", cptTookTot)
		return output
	}/*
	fmt.Println("Total real Score     : ", realScore)
	fmt.Println("Total Score          : ", totScore)
	fmt.Println("Took but no points   : ", cptTookButNotGood)
	fmt.Println("Tot tooked 	     : ", cptTookTot)*/
	return ""
}

func (p *UBERGOOGLE) write(bestC, bestD, bestE *int) {
	output := p.encode(bestC, bestD, bestE)
	if output != "" {
		err := ioutil.WriteFile(PATH + "output/"+fileName+".out", []byte(output), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func toSliceOfInt(line string) []int {
	args := strings.Split(line, " ")
	rep := make([]int, len(args))
	var err error
	for i, v := range args {
		rep[i], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}
	return rep
}


