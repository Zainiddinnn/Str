package main 
import "fmt"
func main() {
	a := []int{4, 5, 6}
	b := make([]int, 3, 3)
	n := copy(a, b)
	fmt.Printf("a = %v\n", a)                  // a = [1 2 3]
	fmt.Printf("b = %v\n", b)                  // b = [1 2 3]
	fmt.Printf("Скопировано %d элемента\n", n) // Скопировано 3 элемента
}