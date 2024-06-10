package ffi

import (
	"fmt"
	"testing"
)

func TestBLSHash(t *testing.T) {
	fooMessage := Message("hello foo")
	fooDigest := Hash(fooMessage)
	fmt.Println("fooDigest: ", fooDigest)
}
