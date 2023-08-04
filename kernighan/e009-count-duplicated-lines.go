/* Count duplicated lines from standard input */
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	countH := make( map[string]int )
	files := os.Args[1:]
	// if no file names in CLI arguments use os.Stdin
	if len(files) == 0 {
		count_for_file( os.Stdin, countH )
	} else {
		for _, fname := range files {
			// opening file
			fh, err := os.Open( fname )
			if err != nil {
				fmt.Fprintf( os.Stderr, "ERROR: %v\n", err )
				continue
			}
			count_for_file( fh, countH )
			fh.Close()
		}
	}
	// print report of repeated lines
	msg := "------------------------------------------------------------"
	fmt.Println(msg)
	fmt.Println("TIMES\tLINE")
	fmt.Println(msg)
	for line, count := range countH {
		if count > 1 {
			fmt.Printf( "%d\t%s\n", count, line )
		}
	}
}

func count_for_file( fh *os.File, countH map[string]int ) {
	scanner := bufio.NewScanner( fh )
	for scanner.Scan() {
		// ignoring scanner.Err()
		countH[ scanner.Text() ]++
	}
}
