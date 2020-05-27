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

// Package variables
var layoutFIN = "02.01.2006"

// Input Person info
func query() Person{
	var p Person
	locHelsinki, _ := time.LoadLocation("Europe/Helsinki")
	fmt.Print("Syötä henkilön nimi: ")
	p.name = scan();
	fmt.Print("Syötä henkilön syntymäpäivä (muodossa DD.MM.YYYY): ")
	p.bday, _ = time.ParseInLocation(layoutFIN, scan(), locHelsinki)

	return p;
}

// Print Person profile
func profile(p Person) {
	fmt.Println("Nimi:",p.name)
	fmt.Println("Syntymäpäivä:",p.bday.Format(layoutFIN))
	fmt.Println("Ikä:",calcAge(p.bday))
}

// TODO: Calculate age
// Currently does not consider leap years
func calcAge(bday time.Time) int{
	now := time.Now()

	//fmt.Println(bday)
	//fmt.Println(now)

	diff:=now.Sub(bday)
	years:=(diff.Hours()/24)/365;

	return int(years);
}


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
