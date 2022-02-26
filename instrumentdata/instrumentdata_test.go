package instrumentdata

import (
	"fmt"
	"testing"
)

func TestInstrumentdata(t *testing.T) {

	instrument_ := GetInstrumentData("ftx", "SRM/USD")

	fmt.Println(instrument_.Minsize, "test>>>")

}
