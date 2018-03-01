package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func (p *UberGoogle) read() {
	b, err := ioutil.ReadFile("input/" + fileName + ".in")
	if err != nil {
		log.WithFields(logrus.Fields{
			"function": "read",
			"fileName": fileName + ".in",
		}).Panic(err)
	}
	p.decode(b)
}

func (p *UberGoogle) decode(b []byte) {
	tmp := bytes.NewReader(b)
	s := bufio.NewScanner(tmp)
	s.Scan()

	log.WithFields(logrus.Fields{
		"line": s.Text(),
	}).Debug("parsing line")
	args := toSliceOfInt(s.Text())
	p.R, p.C, p.F, p.N, p.B, p.T = args[0], args[1], args[2], args[3], args[4], args[5]

	log.WithFields(logrus.Fields{
		"R": p.R,
		"C": p.C,
		"F": p.F,
		"N": p.N,
		"B": p.B,
		"T": p.T,
	}).Info("Global parameters")
	for i := 0; i < p.N; i++ {
		s.Scan()
		log.WithFields(logrus.Fields{
			"line": s.Text(),
		}).Debug("parsing line")
		args := toSliceOfInt(s.Text())
		p.rides[i] = &ride{i, args[0], args[1], args[2], args[3], args[4], args[5], make([]float64, 2)} //make([]*decisionRide, 0, p.N-1)
		log.WithFields(logrus.Fields{
			"ride number": i,
			"a":           args[0],
			"b":           args[1],
			"x":           args[2],
			"y":           args[3],
			"s":           args[4],
			"f":           args[5],
		}).Debug("ride parameter")

		ride := p.rides[i]
		longueurRide := abs(ride.a-ride.x) + abs(ride.b-ride.y)
		bonus := 0
		if max(ride.a+ride.b, ride.s) == ride.s {
			bonus = p.B
		}
		if longueurRide+max(ride.a+ride.b, ride.s) < ride.f {
			p.rides[i].scoreN[0] = float64(longueurRide+bonus) / float64(1+max(ride.a+ride.b, ride.s))
		}

		//for j := 0; j < p.N; j++ {
		//	if j != i {
		//		p.rides[i].nextRides = append(p.rides[i].nextRides, &decisionRide{
		//			ride: p.rides[j],
		//			//otherRides: make([]*ride, 0, 10),
		//		})
		//	}
		//}
	}

	// create slices & maps
	p.myCars = make(map[int][]*Car)
	p.cars = make([]Car, p.F)
	for i := range p.cars {
		p.cars[i].History = make([]int, 0, 10)
		p.myCars[0] = append(p.myCars[0], &p.cars[i])
	}

	log.WithFields(logrus.Fields{
		"function": "decode",
	}).Info("parsing input")
}

func (p *UberGoogle) encode() string {
	var output string
	first := true
	for _, c := range p.cars {
		if !first {
			output += "\n"
		} else {
			first = false
		}
		output += strconv.Itoa(len(c.History))
		for _, rideID := range c.History {
			output += " " + strconv.Itoa(rideID)
		}
	}
	return output
}

func (p *UberGoogle) write() {
	output := p.encode()
	err := ioutil.WriteFile("output/"+fileName+".out", []byte(output), 0644)
	if err != nil {
		log.WithFields(logrus.Fields{
			"function": "write",
		}).Panic(err)
	}
}

func toSliceOfInt(line string) []int {
	args := strings.Split(line, " ")
	rep := make([]int, len(args))
	var err error
	for i, v := range args {
		rep[i], err = strconv.Atoi(v)
		if err != nil {
			log.WithFields(logrus.Fields{
				"function": "toSliceOfInt",
			}).Panic(err)
		}
	}
	return rep
}
