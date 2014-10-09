package utils

import (
	"testing"
	"fmt"
	"time"
)


func TestIdGen(_ *testing.T) {
	fmt.Println(GenerateId("foobar", 6))
	fmt.Println(GenerateId("福州-", 12))
	fmt.Println(GenerateId("福州-北京",9))
}


func TestTimeBasedId(t *testing.T) {
	id, tm := GenerateTimeBasedId()
	fmt.Println(id)

	tm, err := GetUnixTimeFromTimeBasedId(id)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(time.Unix(tm, 0))
}


type signatureTestCase struct {
	timestamp string
	nonce     string
	token     string
	output    string
}

var signatureTestCases = []signatureTestCase{
	{"12321", "adaveq", "foobar", "b60d0aa011c0b2fcd490171f5fe4c347278b728e"},
	{"adaveq", "12321", "foobar", "b60d0aa011c0b2fcd490171f5fe4c347278b728e"},
	{"1232234123451", "adav2341asdf=-eq", "foobar", "720fd5568eabf4cb9f9bd81dea494880d138253a"},
}

func TestSignature(t *testing.T) {
	for _, tc := range signatureTestCases {
		actual := GenerateSignature(tc.timestamp, tc.nonce, tc.token)
		if actual != tc.output {
			t.Errorf("Expected: %s  Actual:%s", tc.output, actual)
		}
	}
}
