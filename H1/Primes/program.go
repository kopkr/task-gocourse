package main

import (
	"fmt"
	"flag"
)

// Main function. Used to call the two functions and pass int between them.
func main() {
	fmt.Println("Tulosta alkuluvut väliltä 2 ja syötetty arvo.") //Transl. "Print prime numbers between 2 and user input"
	input:=setValue();
	listPrimes(input);
}

// Checks -list flag. If flag is less than 2 (default, also the smallest possible prime), then ask user for input.
func setValue() int {
	var limit int;
	flag.IntVar(&limit, "max", 0, "List all the prime numbers from two to given number.");
	flag.Parse();
	if limit >= 2 {
		fmt.Println("Annettu yläraja", limit,"(flag)"); //Transl. "Provided limit (flag)"
	} else {
		fmt.Print("Määritä yläraja: ") //Transl. "Set limit: "
		fmt.Scanln(&limit)
	}
	return limit;
}

// Check and print all prime numbers between 2 and user input
func listPrimes(input int) {
	for i:=2;i<=input; i++ {
		isPrime:=true;
		for x:=2; x<i; x++{
			if i!=x && i % x == 0 {
				isPrime=false;
			}
		}
		if isPrime {
			fmt.Println(i)
		}
	}
}
