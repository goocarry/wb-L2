package unpacking

import (
	"log"
	"testing"
)

func TestUnpackString(t *testing.T) {
	t.Run("unpack string", func (t *testing.T) {
		unpackString := UnpackString("a4bc2d5e")
		if unpackString != "aaaabccddddde" {
			log.Fatal("unpacking result is not correct")
		}
	})

	t.Run("unpack string with incorrect input", func (t *testing.T) {
		unpackString := UnpackString("45")
		if unpackString != "" {
			log.Fatal("unpacking result is not correct")
		}
	})

	t.Run("unpack string with no compression", func (t *testing.T) {
		unpackString := UnpackString("abcd")
		if unpackString != "abcd" {
			log.Fatal("unpacking result is not correct")
		}
	})

	t.Run("unpack string with empty string", func (t *testing.T) {
		unpackString := UnpackString("")
		if unpackString != "" {
			log.Fatal("unpacking result is not correct")
		}
	})
}