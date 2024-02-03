package main

import "fmt"

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)

	fmt.Printf("Original: %v\n", input)
	fmt.Printf("Reverse: %v\n", rev)
	fmt.Printf("Reversed again: %v\n", doubleRev)
}

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
