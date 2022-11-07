package main

import "fmt"

func main() {

	type employed struct {
		name string
		age  int
	}

	emp1 := map[int]employed{}

	emp1[0] = employed{
		name: "david",
		age:  10,
	}
	emp1[1] = employed{
		name: "jhon",
		age:  20,
	}
	fmt.Printf("emp1 is %v \n", emp1)

	//dest := map[int]employed{}
	dest := emp1

	emp1[0] = employed{
		name: "david2",
		age:  100,
	}
	fmt.Printf("dest is %v \n", dest)
	fmt.Printf("emp1 is %v \n", emp1)
	fmt.Printf("====================================\n")
	dest2 := make(map[int]employed)
	for key, value := range emp1 {
		dest2[key] = value
	}
	fmt.Printf("dest2 is %v \n", dest2)
	emp1[1] = employed{
		name: "david5555",
	}
	fmt.Printf("dest2 is %v \n", dest2)
	fmt.Printf("emp1 is %v \n", emp1)
}
