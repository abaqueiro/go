/* implementation of echo unix command */
package main

import "fmt"
import "os"

func main(){
	var arg_n int = len(os.Args)
	var s string

	if arg_n > 1 {
		s += os.Args[1]
	}
	for i:=2; i < arg_n; i++ {
		s += " " + os.Args[i]
	}
	fmt.Println( s )
}
