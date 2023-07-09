package tests

import (
	"financialManagement/pkg/hashing"
	"testing"
)

func TestHashing(t *testing.T) {
	testData := []string{
		"a",
		"ab",
		"abc",
		"abcd",
		"abcde",
		"abcdef",
		"abcdefg",
		"abcdefgh",
		"abcdefghi",
		"abcdefghij",
		"abcdefghijk",
		"abcdefghijkl",
		"abcdefghijklm",
		"abcdefghijklmn",
		"abcdefghijklmno",
		"abcdefghijklmnop",
		"abcdefghijklmnopq",
		"abcdefghijklmnopqr",
		"abcdefghijklmnopqrs",
		"abcdefghijklmnopqrst",
		"abcdefghijklmnopqrstu",
		"abcdefghijklmnopqrstuv",
		"abcdefghijklmnopqrstuvw",
		"abcdefghijklmnopqrstuvwx",
		"abcdefghijklmnopqrstuvwxy",
		"abcdefghijklmnopqrstuvwxyz",
	}
	for _, data := range testData {
		hash := hashing.GetHash(data)
		if hash == data {
			t.Errorf("Hash equals string: %s", data)
		}
	}
}
