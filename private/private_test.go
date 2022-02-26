package private

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPrivate(t *testing.T) {

	// for _, v := range j {
	// 	//fmt.Println(v)
	// 	fmt.Printf("%+v\n", v)
	// }

	// fmt.Println(j.Result)

	//j := GetOrders("ftx", "sponsor")

	//j := GetFills("ftx", "sponsor", 30)

	j := GetFillsToCSV("ftx", "mltoken", 100)

	//j := GetFills("ftx", "sponsor")

	//j := GetBalances("ftx", "sponsor")

	//j := PlaceOrder("ftx", "sponsor", "sol/USD", "buy", 0.2, 98)

	//j := CancelOrder("ftx", "sponsor", "0")

	//fmt.Println(j)
	//fmt.Printf("%+v\n", j)

	//b, _ := json.MarshalIndent(j.Result[0], "", "  ")
	//fmt.Println(string(b))

	// fmt.Println(j.Result[0].Price)

	//j := GetMarket("SRM/USD")

	b, _ := json.MarshalIndent(j, "", "  ")
	fmt.Println(string(b))

	// for _, s := range j.Result {
	// 	fmt.Println(s.Coin)
	// }

}
