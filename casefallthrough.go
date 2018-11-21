package main

import "fmt"

func main() {

	fmt.Println("on the 1 of christmas my true love sent to me:")
	fmt.Println("a partridge in a pear tree")
	fmt.Println("")

	for day := 1; day <= 12; day++ {
		fmt.Println("on the", day, "of christmas my true love sent to me")

		switch day {

		case 12:
			fmt.Println("twelve drummers drumming")
			fallthrough
		case 11:
			fmt.Println("Eleven pipers piping")
			fallthrough
		case 10:
			fmt.Println("ten lords a leaping")
			fallthrough
		case 9:
			fmt.Println("nine ladies dancing")
			fallthrough
		case 8:
			fmt.Println("eight maids a milking")
			fallthrough
		case 7:
			fmt.Println("seven swans swimming")
			fallthrough
		case 6:
			fmt.Println("six geese a laying")
			fallthrough
		case 5:
			fmt.Println("five golden rings")
			fallthrough
		case 4:
			fmt.Println("four calling birds")
			fallthrough
		case 3:
			fmt.Println("three french hens")
			fallthrough
		case 2:
			fmt.Println("two turtle doves")
			fallthrough
		case 1:
			fmt.Println("a partridge in fear tree.")

		}

	}
}
