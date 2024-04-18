package main

import (
	"fmt"
	"os"
	"strconv"
)

func reducefunction(map2 map[string]string) {
	map1 := make(map[string]string)
	result := 0
	for key1, value1 := range map2 {
		for key2, value2 := range map2 {
			if key1 == key2 {
				_, err1 := strconv.Atoi(value1)
				if err1 != nil {
					panic(err1)
				}

				num2, err2 := strconv.Atoi(value2)
				if err2 != nil {
					panic(err2)
				}

				result = result + num2
			}
		}
		str := strconv.Itoa(result)
		map1[key1] = str
		result = 0
	}
	fmt.Println(map1)
}

func main() {
	file, err := os.Open("document.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var word string
	map1 := make(map[string]string)

	for {
		_, err := fmt.Fscanf(file, "%s", &word)
		if err != nil {
			break
		}
		map1[word] = "1"

	}
	fmt.Println(map1)
	reducefunction(map1)
}
