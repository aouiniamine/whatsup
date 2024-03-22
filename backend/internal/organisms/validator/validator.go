package validator

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func randRangeCrypto() (int64, error) {
	max := int64(999999)
	min := int64(100001)

	diff := big.NewInt(max - min)

	n, err := rand.Int(rand.Reader, diff)
	if err != nil {
		return 0, err
	}
	return n.Int64() + min, nil
}

func ValidateEmail(email string) error {
	randomInt, err := randRangeCrypto()
	if err != nil {
		return err
	}

	fmt.Println(randomInt)
	return nil
}
