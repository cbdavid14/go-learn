package main

import (
	"fmt"
)

func main() {

	sample1 := []string{"a", "b", "c"}
	sample2 := sample1
	fmt.Printf("sample1, len %d, %v\n", len(sample1), sample1)
	fmt.Printf("sample2, len %d, %v\n", len(sample2), sample2)

	sample3 := [5]string{"a", "b", "c", "e", "f"}
	fmt.Printf("sample3, len %d, %v\n", len(sample3), sample3)
	fmt.Printf("sample3, len %d, %v\n", len(sample3), sample3[:])
	fmt.Printf("sample3, len %d, %v\n", len(sample3), sample3[1:])
	fmt.Printf("sample3, len %d, %v\n", len(sample3), sample3[:3])

	sample4 := make([]string, 0)
	sample4 = append(sample4, sample1...)
	fmt.Printf("sample4, len %d, direction %v, %v\n", len(sample4), &sample4, sample4)

	test(sample4)
	fmt.Printf("sample4, len %d, %v\n", len(sample4), sample4)

	sample5 := [][]int{{1, 3, 4}, {10, 20, 30}}
	fmt.Printf("sample5, len %d, %v\n", len(sample5), sample5)

	sample6 := map[string]int{"demo": 5000, "demo2": 6000}
	sample6["tom"] = 1000
	fmt.Printf("sample6, len %d, %v\n", len(sample6), sample6)

	sample61 := map[string]int{}
	sample61["tom"] = 1000
	fmt.Printf("sample6_1, len %d, %v\n", len(sample61), sample61)

	type employed struct {
		name string
		age  int
	}
	sample7 := map[string]employed{}

	sample7["temp"] = employed{
		name: "david",
		age:  123,
	}

	fmt.Printf("sample7, len %d, %v\n", len(sample7), sample7)

	sample8 := map[int]employed{
		0: employed{name: "david", age: 10},
		1: employed{name: "david1", age: 11},
		2: employed{name: "david2", age: 12},
		3: employed{name: "david3", age: 12},
		4: employed{name: "david4", age: 12},
		5: employed{name: "david5", age: 12},
		6: employed{name: "david6", age: 12},
		7: employed{name: "david7", age: 12},
	}

	fmt.Printf("sample8, len %d, %v\n", len(sample8), sample8)

	for key, value := range sample8 {
		fmt.Printf("sample8, key %d, value %v\n", key, value)
	}

}

func test(sample []string) {
	sample[0] = "update test"
	fmt.Printf("sample test, len %d, %v\n", len(sample), sample)

}
