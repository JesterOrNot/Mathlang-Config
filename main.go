package main

func sum(i ...int) int {
	var sum int = 0
	for _, i := range i {
		sum += i
	}
	return sum
}
func main() {
	var x int = sum(1,2,3,4,5,6)
	println(x)
}