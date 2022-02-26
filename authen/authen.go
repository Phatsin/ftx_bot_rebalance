package authen

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Phatsin/main/ftx_bot/accountdata"
)

func Authen(broker, acc, method, path string, param interface{}) (int64, string, string, string, []string) {

	acc_ := accountdata.GetAccountData(broker, acc)

	secret := acc_.Secret

	time_ := time.Now().UnixNano() / int64(time.Millisecond)

	payload := ""
	switch method {
	case "WS":
		payload = fmt.Sprintf("%d%s", time_, path)

	case "GET":
		payload = fmt.Sprintf("%d%s%s", time_, "GET", path)

	case "DELETE":
		payload = fmt.Sprintf("%d%s%s", time_, "DELETE", path)

	case "POST":
		payload = fmt.Sprintf("%d%s%s%s", time_, "POST", path, param)
	}
	h := hmac.New(sha256.New, []byte(secret))

	h.Write([]byte(payload))

	sha := hex.EncodeToString(h.Sum(nil))

	return time_, acc_.Sub, sha, acc_.Key, acc_.Instument
}
