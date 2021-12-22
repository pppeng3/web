package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var ApiKey = "OU7095QHgxCqzztMFAbxYfbVa4kkXRphJoK8C4zOuyxdCUTlfPQBF6n9hh0wLao3"
var ApiSecret = "cE2dr5VeBd5EOXuqpcoPxsZKM9iHLH1qZxnAkedZbHp7pWt7wLUdYYSWJU5BzG4S"
var BaseUri = "https://api.binance.com"
var SubUri = "/api/v3/myTrades"

func main() {
	c := http.DefaultClient
	req, err := http.NewRequest(`GET`, BaseUri+SubUri, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-MBX-APIKEY", ApiKey)
	req.URL.RawQuery = makeQuery(`symbol=BNBBTC&startTime=1636959270028`)
	fmt.Println(req.URL.RawQuery)
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", res)
}

func makeQuery(query string) string {
	if query == `` {
		query = fmt.Sprintf("timestamp=%d&recvWindow=60000", time.Now().UnixNano()/1e6-5000)
	} else {
		query = fmt.Sprintf("%s&timestamp=%d&recvWindow=60000", query, time.Now().UnixNano()/1e6-5000)
	}
	sign := hmacSha256([]byte(ApiSecret), []byte(query))
	query += "&signature=" + sign
	return query
}

func hmacSha256(key, data []byte) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}
