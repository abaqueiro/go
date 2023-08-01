/* parse unsigned integer from CLI first argument */
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len( os.Args ) == 1 {
		// write error to standard error stream
		fmt.Fprintf( os.Stderr, "ERROR: you should provide an integer as first argument.\n" )
		os.Exit(1)
	}

	// parsing integer
	var msg string
	n, err := strconv.ParseInt( os.Args[1], 10, 64 )
	if err != nil {
		msg = fmt.Sprintf( "ERROR: %v\n", err )
		fmt.Fprintf( os.Stderr, msg )
		os.Exit(1)
	}

	// write output to standard output
	msg = fmt.Sprintf( "Parsed integer: %d\n", n )
	fmt.Fprintf( os.Stdout, msg )
}
