package private

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Phatsin/main/ftx_bot/authen"
	"github.com/Phatsin/main/ftx_bot/models"
)

func CancelOrder(broker, acc, id string) *models.GetData {

	url_ := "https://ftx.com"

	endpoint_ := "/api/orders"

	if id != "0" {
		endpoint_ += "/" + id
	}

	t_, d_, s_, k_, i_ := authen.Authen(broker, acc, "DELETE", endpoint_, map[string]string{})
	_ = i_
	client := http.Client{}

	//fmt.Println(url_ + endpoint_)
	fmt.Println("Cancel orders!")

	req, _ := http.NewRequest("DELETE", url_+endpoint_, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTX-KEY", k_)
	req.Header.Set("FTX-SIGN", s_)
	req.Header.Set("FTX-TS", fmt.Sprintf("%d", t_))
	req.Header.Set("FTX-SUBACCOUNT", d_)

	res, _ := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	r := &models.GetData{}
	_ = json.Unmarshal(body, r)

	return r
}

func GetOrders(broker, acc string) *models.GetData {

	t_, d_, s_, k_, i_ := authen.Authen(broker, acc, "GET", "/api/orders", map[string]string{})
	_ = i_
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://ftx.com/api/orders", nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTX-KEY", k_)
	req.Header.Set("FTX-SIGN", s_)
	req.Header.Set("FTX-TS", fmt.Sprintf("%d", t_))
	req.Header.Set("FTX-SUBACCOUNT", d_)

	res, _ := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	r := &models.GetData{}
	_ = json.Unmarshal(body, r)

	fmt.Println("Get orders!")
	return r
}

func PlaceOrder(broker, acc, market, side string, size, price float32) *models.GetData {

	params := map[string]interface{}{

		"market": market,
		"side":   side,
		"price":  price,
		"type":   "limit",
		"size":   size,
	}
	json_params, _ := json.Marshal(params)

	t_, d_, s_, k_, i_ := authen.Authen(broker, acc, "POST", "/api/orders", json_params)
	_ = i_

	client := http.Client{}
	req, _ := http.NewRequest("POST", "https://ftx.com/api/orders", bytes.NewBuffer(json_params))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTX-KEY", k_)
	req.Header.Set("FTX-SIGN", s_)
	req.Header.Set("FTX-TS", fmt.Sprintf("%d", t_))
	req.Header.Set("FTX-SUBACCOUNT", d_)

	res, _ := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	r := &models.GetData{}
	_ = json.Unmarshal(body, r)

	fmt.Println("Place orders!")
	return r
}

func GetFills(broker, acc string, amount int) *models.GetData {

	a_ := strconv.Itoa(amount)
	t_, d_, s_, k_, i_ := authen.Authen(broker, acc, "GET", "/api/fills?limit="+a_, map[string]string{})
	_ = i_
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://ftx.com/api/fills?limit="+a_, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTX-KEY", k_)
	req.Header.Set("FTX-SIGN", s_)
	req.Header.Set("FTX-TS", fmt.Sprintf("%d", t_))
	req.Header.Set("FTX-SUBACCOUNT", d_)

	res, _ := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	r := &models.GetData{}
	_ = json.Unmarshal(body, r)

	fmt.Println("Get fill!")
	return r
}

func GetFillsToCSV(broker, acc string, amount int) *models.GetData {

	a_ := strconv.Itoa(amount)
	t_, d_, s_, k_, i_ := authen.Authen(broker, acc, "GET", "/api/fills?limit="+a_, map[string]string{})
	_ = i_
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://ftx.com/api/fills?limit="+a_, nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTX-KEY", k_)
	req.Header.Set("FTX-SIGN", s_)
	req.Header.Set("FTX-TS", fmt.Sprintf("%d", t_))
	req.Header.Set("FTX-SUBACCOUNT", d_)

	res, _ := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	r := &models.GetData{}
	_ = json.Unmarshal(body, r)

	fmt.Println("Get fill!")

	file, err := os.Create("fills.csv")
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	//{
	//	"orderId": 116217708349,
	//	"tradeId": 3180784518,
	//	"id": 6411671397,
	//	"size": 1.12,
	//	"price": 89,
	//	"baseCurrency": "SOL",
	//	"quoteCurrency": "USD",
	//	"type": "order",
	//	"market": "SOL/USD",
	//	"side": "buy",
	//	"time": "2022-01-24T06:57:51.719247+00:00"
	//}

	//{
	//	"orderId": 115115224484,
	//	"tradeId": 3139503864,
	//	"id": 6328446598,
	//	"size": 85,
	//	"price": 2.6,
	//	"fee": -0.00085,
	//	"feeRate": -0.00001,
	//	"baseCurrency": "SRM",
	//	"quoteCurrency": "USD",
	//	"type": "order",
	//	"market": "SRM/USD",
	//	"side": "buy",
	//	"time": "2022-01-21T12:08:20.922965+00:00"
	//}

	// Using WriteAll
	var data [][]string
	data = append(data, []string{"orderId", "tradeId", "id", "size", "price", "fee", "feeRate", "baseCurrency", "quoteCurrency", "type", "market", "side", "time"})

	for _, record := range r.Result {
		row := []string{strconv.Itoa(record.OrderId), strconv.Itoa(record.TradeId), strconv.Itoa(record.Id), fmt.Sprintf("%f", record.Size), fmt.Sprintf("%f", record.Price), fmt.Sprintf("%f", record.Fee), fmt.Sprintf("%f", record.FeeRate), record.BaseCurrency, record.QuoteCurrency, record.Type, record.Market, record.Side, fmt.Sprintf("%v", record.Time)}
		data = append(data, row)
	}
	w.WriteAll(data)
	return r
}

func GetBalances(broker, acc string) *models.GetData {

	t_, d_, s_, k_, i_ := authen.Authen(broker, acc, "GET", "/api/wallet/balances", map[string]string{})
	_ = i_
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://ftx.com/api/wallet/balances", nil)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTX-KEY", k_)
	req.Header.Set("FTX-SIGN", s_)
	req.Header.Set("FTX-TS", fmt.Sprintf("%d", t_))
	req.Header.Set("FTX-SUBACCOUNT", d_)

	res, _ := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	r := &models.GetData{}
	_ = json.Unmarshal(body, r)

	fmt.Println("Get balance!")
	return r
}

func GetMarket(s string) *models.GetData_ {

	resp, err := http.Get("https://ftx.com/api/markets/" + s)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	//log.Println(string(body))

	if err != nil {
		log.Println(err)
	}

	r := &models.GetData_{}
	_ = json.Unmarshal(body, r)

	fmt.Println("Get market data!")
	return r

}
