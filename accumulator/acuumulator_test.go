package accumulator


import (
	"fmt"
	"testing"
)

func TestHashToPrime(* testing.T) {
	p := hashToPrime([]byte("Test 1"), 2048)
	fmt.Println(p)
}

