package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"flag"
)

// Main function
func main () {
	fmt.Println("Rot13-kääntäjä");		// Transl. "Rot13-converter"
	fmt.Println(rot13(setInput()));
}

// Scans user input from command prompt
func scan() string{
	scanner := bufio.NewScanner(os.Stdin)
	var input string;
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input;
}

// Checks if flag -s (string) is set, if not asks user input
func setInput() string {
	var input string;
	flag.StringVar(&input, "s", "", "Set string to be encoded/decoded.")
	flag.Parse();
	if len(input) <= 0 {
		fmt.Print("Syötä tekstiä käännettäväksi: ");	// Transl. "Input text to be converted: "
		input = scan();
		return input;
	} else {
		fmt.Println("Teksti käännettäväksi:",input)	// Trans. "Text to be converted"
		return input;
	}
}

// strings.Map(x,y) changes letters in a string 
func rot13(input string) string{
	return strings.Map(rot13map, input);
}

// Increases int32 decimal by 13 if characters A-N (upper & lower). Example: A > M, a > m
// Decreases int32 decimal by 13 if characters M-Z (upper & lower). Example: M > A, m > a
func rot13map(r rune) rune {
	if r>='a' && r<='m' || r>='A' && r<='M' {
		return r+13;
	} else if r>='n' && r<='z' || r>='N' &&  r<='Z' {
		return r-13;
	}
	return r;
}
