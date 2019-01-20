package main

import "fmt"

func main() {
	ids := []int{33, 34, 62, 68, 30}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
	// let's add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("sum", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: ", k)
	}
}
