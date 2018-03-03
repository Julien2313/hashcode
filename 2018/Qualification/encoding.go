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
	p.R, p.C, p.F, p.N, p.B, p.T = args[0], args[1], args[2], args[3], args[4], args[5]
	fmt.Println(p)
	p.rides = make([]*RIDE, p.N)

	for numRide := 0; numRide < p.N; numRide++ {
		s.Scan()
		args := toSliceOfInt(s.Text())
		p.rides[numRide] = &RIDE{numRide, args[0], args[1], args[2], args[3], args[4], args[5], 0.0, 0}
	}

	p.cars = make([]*CAR, p.F)
	for cptCar := 0; cptCar < p.F; cptCar++ {
		p.cars[cptCar] = &CAR{0, make([]RIDE, 0), 0, 0}
	}
}

func (p *UBERGOOGLE) encode() string {
	var output string
	totScore := 0
	first := true
	for _, c := range p.cars {
		if !first {
			output += "\n"
		} else {
			first = false
		}
		output += strconv.Itoa(len(c.History))
		for _, ride := range c.History {
			totScore += ride.Lenght()
			if ride.tickStarted == ride.s {
				totScore += p.B
			}
			output += " " + strconv.Itoa(ride.ID)

		}
	}
	//fmt.Println(output)
	fmt.Println(totScore)
	/*for _, ride := range p.rides {
	  fmt.Println(ride.Used)
	}*/
	return output
}

func (p *UBERGOOGLE) write() {
	output := p.encode()
	err := ioutil.WriteFile(PATH + "output/"+fileName+".out", []byte(output), 0644)
	if err != nil {
		panic(err)
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


