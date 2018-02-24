package main

type pizza struct {
	Cells      map[point]*Cell
	H, L, R, C int

	Slices map[slice]*sliceInfo
}

type slice struct {
	x0, y0, x1, y1 int
	//score int
}

type sliceInfo struct {
	nbChamp, nbTomate int
	score             int
	used              bool
}
type point struct {
	x, y int
}

type Cell struct {
	Ingredient byte

	IsInSlice      *slice
	AvailableSlice []*slice
}

func (p *pizza) score() int {
	nbCell := 0
	for _, c := range p.Slices {
		if c.used {
			nbCell += c.score
		}
	}
	return nbCell
}

func (p *pizza) cutInBigSlice() {

	//fmt.Println(p.H)
	primaryN := PrimeFactors(p.H)
	p.Slices = make(map[slice]*sliceInfo)
	for coordCell := range p.Cells {
		for _, n := range primaryN {
			if coordCell.x+n > p.C || coordCell.y+p.H/n > p.R {
				continue
			}
			coordSlice := slice{coordCell.x, coordCell.y, coordCell.x + n - 1, coordCell.y + p.H/n - 1}
			p.Slices[coordSlice] = &sliceInfo{}
			for j := 0; j < n; j++ {
				for k := 0; k < p.H/n; k++ {
					_, exist := p.Cells[point{coordCell.x + j, coordCell.y + k}]
					if !exist {
						break
					}
					if p.Cells[point{coordCell.x + j, coordCell.y + k}].Ingredient == 'T' {
						p.Slices[coordSlice].nbTomate++
					} else {
						p.Slices[coordSlice].nbChamp++
					}
				}
			}
			if p.Slices[coordSlice].nbTomate < p.L || p.Slices[coordSlice].nbChamp < p.L {
				delete(p.Slices, coordSlice)
				continue
			}
			p.Slices[coordSlice].score = p.Slices[coordSlice].nbChamp + p.Slices[coordSlice].nbTomate
			p.Cells[coordCell].AvailableSlice = append(p.Cells[coordCell].AvailableSlice, &coordSlice)
		}
	}
}

func (p pizza) PlacerSlice() {

	//fmt.Println("R ",p.R," C ",p.C)
	for y := 0; y < p.R; y++ {
		for x := 0; x < p.C; x++ {

			//fmt.Println("X ", x, " Y ", y)
		dance:
			for _, trySlice := range p.Cells[point{x, y}].AvailableSlice {
				for i := trySlice.x0; i <= trySlice.x1; i++ {
					for j := trySlice.y0; j <= trySlice.y1; j++ {
						//fmt.Println("looking for ", trySlice, i, j)
						//_,exist := p.Cells[point{i,j}]
						if p.Cells[point{i, j}].IsInSlice != nil {
							//fmt.Println("break dance")
							continue dance
						}
					}
				}
				for i := trySlice.x0; i <= trySlice.x1; i++ {
					for j := trySlice.y0; j <= trySlice.y1; j++ {
						//fmt.Println("locking ",trySlice,i,j)
						p.Cells[point{i, j}].IsInSlice = trySlice
					}
				}
				p.Slices[*trySlice].used = true
				x += trySlice.x1 - trySlice.x0 - 1
				break
			}
		}
	}
}

//func (p *pizza) FindSubSlice() {
//	for slice,sliceInfo := range p.Slices{
//		for {
//
//			if slice
//		}
//	}
//}

//
//func (s *sliceInfo) String()  string{
//	return fmt.Sprint(s)
//}
