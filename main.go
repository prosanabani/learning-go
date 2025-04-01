package main

type shape interface {
	area() int
	circumference() int
}

type circle struct {
	width, height int
}
type rectangle struct {
	depth, height int
}

func (currentCircle circle) area() int {
	return currentCircle.width * currentCircle.height
}

func (currentCircle rectangle) area() int {
	return currentCircle.height * currentCircle.height
}
func (currentCircle circle) circumference() int {
	return currentCircle.height * currentCircle.width
}
func (currentCircle rectangle) circumference() int {
	return currentCircle.height * currentCircle.depth
}

func PrintArea(currentShape shape) {
	println(currentShape.area())
}

func printCirum(currentShape shape) {
	println(currentShape.circumference())
}

func main() {

	myCircle := circle{width: 10, height: 20}
	myRectangle := rectangle{depth: 10, height: 20}

	printCirum(myCircle)
	printCirum(myRectangle)

}