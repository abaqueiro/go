package main

import (
	"bytes"
	"fmt"
	//"slices"
)

func main(){
	data := []byte{0x40,0x28,0x14,0x17,0x00,0x11,0x12,0x13,0xAA,0xBB}
	search_1 := []byte{0x11,0x12}
	search_2 := []byte{0x99,0x98}
	
	fmt.Println( "Data:" )
	fmt.Printf( "%X\n", data )
	fmt.Println()

	fmt.Println( "Search 1:")
	fmt.Printf( "%X\n", search_1 )
	fmt.Println( "search_1 is contained in data:", bytes.Contains( data, search_1 ) )
	fmt.Println()

	fmt.Println( "Search 2:")
	fmt.Printf( "%X\n", search_2 )
	fmt.Println( "search_2 is contained in data:", bytes.Contains( data, search_2 ) )
	fmt.Println()
}