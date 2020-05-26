package main

import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
	"flag"
)

// Main function
func main () {
	fmt.Println("LetterCount. v.0.01")
	viewFile(getString());

}

// View the file and count letters
func viewFile(file string, err error) {
	if err != nil {
		fmt.Println("Can't read the file:",err);
	} else {
		x, y := countLetters(file);
		z := x+y
		fmt.Println("Vowels:",x,"Consonants:",y,"Total:",z);
	}
}

func getString() (string, error){
	var file string;
	flag.StringVar(&file, "f", "", "Set filepath.");
	flag.Parse();

	var bytes []byte;
	var err error;

	// If flag is empty, ask for user input
	if file != "" {
		bytes, err = ioutil.ReadFile(file)
	} else {
		fmt.Print("Enter string: ")
		bytes, err = []byte(scanStr()), nil;
	}

	if err != nil {
		return "", err;
	}
	return string(bytes), err;
}

// Return vowel and consonant count
func countLetters(s string) (int,int){
	x, y := 0, 0;
	rs := []rune(s)
	for i:= range rs {
		if rs[i]>='a' && rs[i]<='z' || rs[i]>='A' && rs[i]<='Z' || rs[i] == 'å' || rs[i]=='ä' || rs[i]=='ö' || rs[i] == 'Å' || rs[i] == 'Ä' || rs[i] == 'Ö' {
			if rs[i] == 'a' || rs[i] == 'e' || rs[i] == 'i' || rs[i] == 'o' || rs[i] == 'u' || rs[i] == 'y' || rs[i] == 'å' || rs[i] == 'ä' || rs[i] == 'ö' || rs[i] == 'A' || rs[i] == 'E' || rs[i] == 'I' || rs[i] == 'O' || rs[i] == 'U' || rs[i] == 'Y' || rs[i] == 'Å' || rs[i] == 'Ä' || rs[i] == 'Ö' {
				x++;
			} else {
				y++;
			}
		}
	}
	return x, y;
}

// Scans user input from command prompt
func scanStr() string{
	scanner := bufio.NewScanner(os.Stdin)
	var input string;
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input;
}
