package main

import (
	"encoding/json"
	"flag"
	"log"
	"math"
	"net/url"
	"strings"
	"time"

	"github.com/Phatsin/main/ftx_bot/authen"
	"github.com/Phatsin/main/ftx_bot/instrumentdata"
	"github.com/Phatsin/main/ftx_bot/models"
	"github.com/Phatsin/main/ftx_bot/private"
	"github.com/recws-org/recws"
)

func main() {

	assign := &models.Assign{}

	flag.StringVar(&assign.Account, "acc", "", "Account")
	flag.StringVar(&assign.Instrument, "inst", "", "Instrument")
	flag.Int64Var(&assign.Fvalue, "f_val", 0, "Fixed of value")
	flag.Float64Var(&assign.Bdist, "b_dist", 0, "distance of the buy order")
	flag.Float64Var(&assign.Sdist, "s_dist", 0, "distance of the sell order")
	flag.Float64Var(&assign.Prest, "p_rest", 0, "Sell price of the rest order")

	flag.Parse()

	if assign.Fvalue != 0 && assign.Bdist != 0 && assign.Sdist != 0 && assign.Prest != 0 {
		inst := strings.ToUpper(assign.Instrument)
		inst_ := strings.Split(inst, "/")
		coin_ := inst_[0]

		ws := recws.RecConn{
			//KeepAliveTimeout: 10 * time.Second,
			KeepAliveTimeout: 0,
		}

		u_broker := url.URL{Scheme: "wss", Host: "ftx.com", Path: "/ws"}

		ws.Dial(u_broker.String(), nil)

		go func() {
			for {
				time.Sleep(58 * time.Second)
				push_ := map[string]interface{}{
					"op": "ping",
				}
				err := ws.WriteJSON(push_)
				if err != nil {
					log.Println("Write: ", err)
					return
				}
			}
		}()

		brokerReq := &models.Req{}

		authen_ := false
		for {
			if !authen_ {
				////////////////////////////////
				t_, d_, s_, k_, i_ := authen.Authen("ftx", assign.Account, "WS", "websocket_login", map[string]string{})
				_ = i_

				push_ := map[string]interface{}{
					"op": "login",
					"args": map[string]interface{}{
						"key":        k_,
						"sign":       s_,
						"time":       t_,
						"subaccount": d_,
					},
				}

				err := ws.WriteJSON(push_)
				if err == nil {
					log.Println("Authen by " + assign.Account)
				}

				push_ = map[string]interface{}{
					"op":      "subscribe",
					"channel": "orders",
				}

				err = ws.WriteJSON(push_)
				if err == nil {
					log.Println("Order subscribe...")
				}

				push_ = map[string]interface{}{
					"op":      "subscribe",
					"channel": "fills",
				}

				err = ws.WriteJSON(push_)
				if err == nil {
					log.Println("Fills subscribe...")
				}

				//////////////////////////////
			}
			err := ws.ReadJSON(brokerReq)
			if err != nil {
				//log.Println("Read: ", err)
				authen_ = false
			}

			if brokerReq.Type != "" && brokerReq.Type != "pong" {
				log.Println("Exchange:" + brokerReq.Type)
			}

			if brokerReq.Type == "error" {
				b, _ := json.MarshalIndent(brokerReq, "", "  ")
				log.Println("Broker recv:", string(b))
			}

			if brokerReq.Channel == "orders" || brokerReq.Channel == "fills" {
				authen_ = true
			}

			if brokerReq.Type == "subscribed" && brokerReq.Channel == "fills" {
				sendOrders(assign.Account, inst, coin_, assign.Fvalue, assign.Bdist, assign.Sdist, assign.Prest)
			}

			if brokerReq.Type == "update" && brokerReq.Channel == "fills" {
				if brokerReq.Data.RemainingSize != 0 {
					time.Sleep(1 * time.Minute)
				}
				sendOrders(assign.Account, inst, coin_, assign.Fvalue, assign.Bdist, assign.Sdist, assign.Prest)
			}

			brokerReq = &models.Req{}

		}
	} else {
		log.Println("error: Require any flag")
	}
}

func sendOrders(account, instrument, coin string, fvalue int64, bdist, sdist, prest float64) {

	instrument_ := instrumentdata.GetInstrumentData("ftx", instrument)

	size_digit := 1 / instrument_.Minsize

	j := private.GetBalances("ftx", account)

	var coin_ float32
	for _, s := range j.Result {
		if s.Coin == coin {
			coin_ = s.Total
		}
	}

	k := private.GetFills("ftx", account, 1)

	mark_price := k.Result[0].Price

	_ = private.CancelOrder("ftx", account, "0")

	l := private.GetMarket(instrument)

	bid := l.Result.Bid
	ask := l.Result.Ask

	buy_price := mark_price - float32(bdist)
	buy_size_ := (float32(fvalue) - (coin_ * buy_price)) / buy_price
	buy_size := float32(math.Floor(float64(buy_size_)*float64(size_digit)) / float64(size_digit))
	sell_price := mark_price + float32(sdist)
	sell_size_ := ((coin_ * sell_price) - float32(fvalue)) / sell_price
	sell_size := float32(math.Floor(float64(sell_size_)*float64(size_digit)) / float64(size_digit))

	if buy_size < instrument_.Minsize {
		//buy_price = float32(fvalue) / (coin_ + instrument_.Minsize)
		buy_price := mark_price - float32(bdist)
		buy_size = instrument_.Minsize
		log.Println(buy_price)
	}
	if sell_size < instrument_.Minsize {
		//sell_price = float32(fvalue) / (coin_ - instrument_.Minsize)
		sell_price = mark_price + float32(sdist)
		sell_size = instrument_.Minsize
	}
	if buy_price > ask {
		buy_price = ask - 0.005
	}
	if sell_price < bid {
		sell_price = bid + 0.005
	}

	rest_size := coin_ - sell_size

	log.Println(buy_price)
	log.Println((buy_size))

	log.Println(sell_price)
	log.Println((sell_size))
	log.Println((rest_size))

	_ = private.PlaceOrder("ftx", account, instrument, "buy", buy_size, buy_price)
	_ = private.PlaceOrder("ftx", account, instrument, "sell", sell_size, sell_price)
	_ = private.PlaceOrder("ftx", account, instrument, "sell", rest_size, float32(prest))
}

//go run main.go -acc=maker -inst=srm/usd -f_val=1500 -b_dist=0.1 -s_dist=0.1 -p_rest=222
//go run main.go -acc=mltoken -inst=srm/usd -f_val=6000 -b_dist=0.1 -s_dist=0.1 -p_rest=222
//go run main.go -acc=sponsor -inst=sol/usd -f_val=300 -b_dist=10 -s_dist=10 -p_rest=888
