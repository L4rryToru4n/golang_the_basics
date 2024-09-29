package main

import "fmt"

func main() {
	count := 1

	for count <= 10 {
		fmt.Println("Loop of", count)
		count++
	}

	fmt.Println("Done..")

	// For with statment
	for counter := 0; counter < 10; counter++ {
		fmt.Println("Loop statement of ", counter)
	}

	fmt.Println("Finish..")

	// For range
	// using array
	names := []string{"Lee", "'Apollo'", "Adama"}
	for index, name := range names {
		fmt.Println("index", index, ":", name)
	}

	// using map
	personel := map[string]string{
		"name": "Lee 'Apollo' Adama",
		"role": "Viper Pilot",
		"rank": "Liutenant",
	}

	for key, value := range personel {
		fmt.Println(key, value)
	}

	// using map
	// without using key

	for _, value := range personel {
		fmt.Println(value)
	}

	/*
		Break and continue
	*/
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("[Break] Loop number ", i)
	}

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}

		fmt.Println("[Continue] Loop number ", i)
	}

}
