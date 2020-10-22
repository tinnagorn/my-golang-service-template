package utility

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

const letterAlphaNumeric = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandAlphaNumeric(n int) string {
	b := make([]byte, n)
	max := big.NewInt(int64(len(letterAlphaNumeric)))
	for i := range b {
		randomBigInt, err := rand.Int(rand.Reader, max)
		if err == nil {
			randomNumber, _ := strconv.Atoi(randomBigInt.String())
			b[i] = letterAlphaNumeric[randomNumber]
		}
	}
	return string(b)
}

func RandNumeric(n int) string {
	max := big.NewInt(int64(math.Pow(10, float64(n))))
	randomBigInt, _ := rand.Int(rand.Reader, max)
	randomNum, _ := strconv.Atoi(randomBigInt.String())        // Convert big.Int to int.
	result := fmt.Sprintf("%0"+strconv.Itoa(n)+"d", randomNum) // Format to n len
	return result
}
