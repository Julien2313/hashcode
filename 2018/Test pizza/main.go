package main

const fileName = "example"

func main() {
	myPizza := &pizza{
		Cells: make(map[point]*Cell),
	}
	myPizza.read()
	//fmt.Println(myPizza.Cells)
	myPizza.cutInBigSlice()

	//fmt.Println(myPizza.Slices)
	myPizza.PlacerSlice()

	//

	myPizza.write()
}
