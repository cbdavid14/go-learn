package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page1"`
	Fruits []string `json:"fruits1"`
}

func main() {

	jsonB, _ := json.Marshal(true)

	fmt.Println(string(jsonB))

	jsonArr, _ := json.Marshal([]string{"demo1", "demo2", "demo3"})
	fmt.Println(string(jsonArr))

	response1 := &response1{
		Page:   1,
		Fruits: []string{"apple", "orange", "banana"},
	}
	jsonResp, _ := json.Marshal(response1)
	fmt.Println(string(jsonResp))

	response2 := &response2{
		Page:   1,
		Fruits: []string{"apple", "orange", "banana1"},
	}
	jsonResp2, _ := json.Marshal(response2)
	fmt.Println(string(jsonResp2))

	jsonSstring := `{"page":1, "fruits":[1,2,3,4,5,6], "tv":{"new":"lg"}}`
	fmt.Println(jsonSstring)

	byt := []byte(jsonSstring)
	var data map[string]interface{}

	err := json.Unmarshal(byt, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)

	listAlumns := map[string]int{"daviud": 30, "jhon": 50}
	fmt.Println(listAlumns)

}
