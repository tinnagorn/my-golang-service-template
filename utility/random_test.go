package utility

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestRandAlphaNumeric4DigitsShouldCorrect(t *testing.T) {
	rand := RandAlphaNumeric(4)
	if len(rand) != 4 {
		t.Errorf("Expected random string with length 4 but got : %s", rand)
	}
}

func TestRandAlphaNumeric6DigitsShouldCorrect(t *testing.T) {
	rand := RandAlphaNumeric(6)
	if len(rand) != 6 {
		t.Errorf("Expected random string with length 6 but got : %s", rand)
	}
}

func TestRandNumeric4DigitsShouldCorrect(t *testing.T) {
	rand := RandNumeric(4)
	if len(rand) != 4 {
		t.Errorf("Expected random numeric string with length 4 but got : %s", rand)
	}

	if _, err := strconv.Atoi(rand); err != nil {
		t.Errorf("Expected to get random only number but got : %s", rand)
	}
}

func TestRandNumeric6DigitsShouldCorrect(t *testing.T) {
	rand := RandNumeric(6)
	if len(rand) != 6 {
		t.Errorf("Expected random numeric string with length 6 but got : %s", rand)
	}

	if _, err := strconv.Atoi(rand); err != nil {
		t.Errorf("Expected to get random only number but got : %s", rand)
	}
}

func TestRandNumeric6DigitsMillionRoundLessThanTwoSecond(t *testing.T) {
	start := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		RandNumeric(6)
	}
	stop := time.Now().UnixNano()
	assert.Less(t, stop-start, int64(2000000000)) // Must finish less than 2 seconds
}
