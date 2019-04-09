package accumulator


import (
	"fmt"
	"testing"
)

func TestHashToPrime1(* testing.T) {
	p := hashToPrime([]byte("Test 1"), 2048)
	fmt.Println(p)
}

func TestHashToPrime2(* testing.T) {
	p := hashToPrime([]byte("Test 2"), 2048)
	fmt.Println(p)
}

func TestHashToPrime3(* testing.T) {
	p := hashToPrime([]byte("Test 3"), 2048)
	fmt.Println(p)
}

