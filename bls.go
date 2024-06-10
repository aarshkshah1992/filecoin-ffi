//go:build cgo && darwin && arm64
// +build cgo,darwin,arm64

package ffi

import "C"
import (
	prebuilt "github.com/aarshkshah1992/prebuilt-ffi-darwin-arm64"
)

// Hash computes the digest of a message
func Hash(message Message) Digest {
	return prebuilt.Hash(message)
}
