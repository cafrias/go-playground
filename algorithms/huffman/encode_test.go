package huffman

import (
	"reflect"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	input := strings.NewReader("abbcccdddd英英英英英")
	expectedCode := []byte{0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	code, _, err := Encode(input)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		t.FailNow()
	}

	if !reflect.DeepEqual(code, expectedCode) {
		t.Errorf("Invalid code\nExpected:%v\nGot:%v\n", expectedCode, code)
		t.FailNow()
	}
}
