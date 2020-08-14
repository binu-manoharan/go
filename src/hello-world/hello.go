package main

import (
    "fmt"
    "errors"
    "math"
)

func main() {
	// printHello()
	// add10And20()
    // typeOfX()
    // fmt.Println(conditionalTest(32))
	// arrayTest()
	// mapTest()
    // loopTest()
    // testError(16)
    // testError(-16)
    // printStudent(student{id: 1, name: "Binu"})
    // testPointers()
}

func testPointers() {
    x := 10
    incrementX(&x)
    fmt.Println("Incremented value of x is", x)
}

func incrementX(x *int) {
    *x++
}

type student struct {
    id int
    name string
}

func printStudent(student student) {
    fmt.Println(student.name, "has id", student.id)
}

func testError(num float64) {
    result, err := sqrt(num)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(result)
    }
}

func sqrt(num float64) (float64, error) {
    if num < 0 {
        return 0, errors.New("Undefined for negative numbers")
    }

    return math.Sqrt(num), nil
}

func loopTest() {
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	alphabets := []string{"a", "b", "c"}
	for index, value := range alphabets {
		fmt.Println("index:", index, "value:", value)
	}

	fruits := make(map[string]string)
	fruits["a"] = "apple"
    fruits["b"] = "banana"
    
    for key, value := range fruits {
        fmt.Println("key: ", key, "value:", value)
    }
}

func mapTest() {
	wordCount := make(map[string]int)
	wordCount["hello"] = 1
	wordCount["world"] = 1

	fmt.Println(wordCount)
}

func arrayTest() {
	var x [5]int
	x[1] = 10
	y := []int{1, 2, 3, 4, 5}
	y = append(y, 6)

	fmt.Println(x)
	fmt.Println(y)
}

func conditionalTest(x int) string {
    if x > 32 {
        return "Ancient"
    } else if x == 32 {
        return "just right"
    } else {
        return "too young"
    }
}

func typeOfX() {
	const X float64 = 20.0

	fmt.Println(X)
	fmt.Printf("type of x: %T\n", X)
}
func add10And20() {
	x := 10
	y := 20

	sum := x + y

	fmt.Println("10 + 20 = ", sum)
}

func printHello() {
	fmt.Println("Hello!")
}
