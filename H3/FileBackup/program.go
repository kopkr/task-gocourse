// Backup a file (currently single file only). Tested and functional on Linux
package main

import (
	"fmt"
	"os"
	"flag"
	"log"
	"io"
	"path/filepath"
)

// Main function
func main() {
	var from string;
	var to string;
	flag.StringVar(&from, "from", "", "Select file to backup.");
	flag.StringVar(&to, "to", "", "Select backup location (and specify backup filename). Example: 'dir/dir', '../dir/dir/', 'file' 'dir/file'");
	flag.Parse();

	copyFile(to, from)
}

// Make a backup copy file
func copyFile(to, from string) {
	//Read input file
	infile, err := os.Open(from);
	if err != nil {
		log.Fatal(err)
	}
	defer infile.Close();

	//Create a dummy backup file to user specified path
	//TODO: Create parent folders if they do not exist
	path := makePath(from, to);
	outfile, err := os.Create(path);
	if err != nil {
		log.Fatal(err)
	}

	//Copy input file contents to newly created backup-file.
	//TODO: If-else: Check if backup file already exists. If so, ask if you want to overwrite.
	_, err = io.Copy(outfile, infile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Saved backup file to:",path)
	defer outfile.Close();
}

func makePath (from, to string) string{
	// If "to"-flag is empty, save source file's path and name with .bak extension
	if to == "" {
		filename := from+".bak"
		return filename
	}

	// Get information on *os.File pointed at in "to"-flag
	fileInfo, err := os.Stat(to)
	fileExists := true
	if os.IsNotExist(err) {		// Need to check if *os.File exists or fatal error happens
		fileExists = false;
	}

	// If "to"-flag points to a directory, return the original filename with .bak extension
	if fileExists {
		if fileInfo.IsDir() {
			filename := filepath.Clean(to)+"/"+filepath.Base(from)+".bak"
			return filename
		}
	}

	pathname := to+".bak"
	return pathname;
}
