package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
)

func (p *pizza) read() {
	b, err := ioutil.ReadFile("input/" + fileName + ".in")
	if err != nil {
		panic(err)
	}
	p.decode(b)
}

func (p *pizza) decode(b []byte) {
	tmp := bytes.NewReader(b)
	s := bufio.NewScanner(tmp)
	s.Scan()

	args := toSliceOfInt(s.Text())
	p.R = args[0]
	p.C = args[1]
	p.L = args[2]
	p.H = args[3]

	for y := 0; y < p.R; y++ {
		s.Scan()
		row := s.Bytes()
		for x, c := range row {
			p.Cells[point{x, y}] = &Cell{c, nil, make([]*slice, 0, 10)}
		}
	}

}

func (p *pizza) encode() string {
	var output string
	count := 0
	for coord, slice := range p.Slices {
		if slice.used {
			count++
			output += "\n" + strconv.Itoa(coord.y0) + " " + strconv.Itoa(coord.x0) + " " + strconv.Itoa(coord.y1) + " " + strconv.Itoa(coord.x1)
		}
	}
	return strconv.Itoa(count) + output
}

func (p *pizza) write() {
	output := p.encode()
	err := ioutil.WriteFile("output/"+fileName+".out", []byte(output), 0644)
	if err != nil {
		return
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
