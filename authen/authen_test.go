package authen

import (
	"fmt"
	"testing"
)

func TestAuthen(t *testing.T) {

	//t_, d_, s_, k_, i_ := Authen("ftx", "maker", "", "websocket_login", map[string]string{})

	//t_, d_, s_, k_, i_ := Authen("ftx", "test", "GET", "/api/fills?market=SRM/USD", map[string]string{})

	//t_, d_, s_, k_, i_ := Authen("ftx", "maker", "POST", "/api/orders", map[string]string{})

	t_, d_, s_, k_, i_ := Authen("ftx", "sponsor", "WS", "websocket_login", map[string]string{})

	fmt.Println(t_, d_, s_, k_, i_, "test>>>")

	//_ = AuthenPOST("ftx", "maker")
	//fmt.Println(t_, "test>>>")

}

// "/api/fills?market=SRM/USD"
// "/api/orders"
