package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
)

// Main function
func main() {
	profile(query());
}

// Struct for Person info
type Person struct {
	name string
	bday time.Time
}

// Input Person info
func query() Person{
	var p Person
	layoutFIN := "02.01.2006"

	fmt.Print("Syötä henkilön nimi: ")
	p.name = scan();
	fmt.Print("Syötä henkilön syntymäpäivä (muodossa DD.MM.YYYY): ")
	p.bday, _ = time.Parse(layoutFIN, scan())

	return p;
}

// Print Person profile
func profile(p Person) {
	fmt.Println("Nimi:",p.name)
	fmt.Print("Syntymäpäivä: ")
	fmt.Println(p.bday.Date()) //TODO: Format date
	//TODO: Print age
}

/* TODO: Calculate age
func calcAge() {

}*/


// Scanner function
func scan() string{
	scanner := bufio.NewScanner(os.Stdin)
	var input string;

	if scanner.Scan() {
		input = scanner.Text()
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err());
		os.Exit(1)
	}
	return input;
}
