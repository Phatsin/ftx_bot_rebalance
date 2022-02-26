package accountdata

import (
	"fmt"
	"testing"
)

func TestAccountData(t *testing.T) {

	r := GetAccountData("ftx", "sponsor")

	fmt.Println(r.Key, "test>>>")
	fmt.Println(r.Secret, "test>>>")
	fmt.Println(r.Sub, "test>>>")

}
