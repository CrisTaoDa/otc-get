package service

import (
	"encoding/json"
	"fmt"
	"net/url"
	"otc-get/app/curl"
	"otc-get/app/util"
	"strconv"
	"time"
)

type OkexService interface {
	LoadOkexOTC() string
	LoadOkexUSDTBTC() string
	LoadOkexOTCSellBTC() string
	LoadOkexTest() string
}

func NewOkexService() OkexService {
	return &okexService{}
}

type okexService struct {
}

func (o okexService) LoadOkexTest() string {

	buyPrice, _ := strconv.ParseFloat(o.LoadOkexOTC(), 64)
	usdtBtc, _ := strconv.ParseFloat(o.LoadOkexUSDTBTC(), 64)
	sellBtc, _ := strconv.ParseFloat(o.LoadOkexOTCSellBTC(), 64)
	fmt.Println(buyPrice,usdtBtc,sellBtc)
	return ""
}

func (o okexService) LoadOkexOTCSellBTC() string {
	baseUrl := "https://www.okex.com/v3/c2c/tradingOrders/book"
	params := url.Values{}
	params.Set("t", strconv.FormatInt(time.Now().Unix()*1000, 10))
	params.Set("quoteCurrency", "CNY")
	params.Set("baseCurrency", "BTC")
	params.Set("side", "buy")
	params.Set("paymentMethod", "all")
	params.Set("userType", "all")
	//params.Set("quoteMinAmountPerOrder", "5000")
	res, err := curl.Get(baseUrl, params)
	if err != nil {
		fmt.Println(err)
	}
	temp := make(map[string]interface{})
	err = json.Unmarshal(res, &temp)
	if err != nil {
		fmt.Println(err)
	}
	a := temp["data"].(map[string]interface{})["buy"].([]interface{})
	if len(a) > 3 {
		// 将前三个个写入本地文件
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		content := "当前时间:" + timeNow + " 卖BTC前三的价格为:" + a[0].(map[string]interface{})["price"].(string) + "/" + a[1].(map[string]interface{})["price"].(string) + "/" + a[2].(map[string]interface{})["price"].(string)
		content += "\n"
		util.WriteFile(content, util.LocalPath()+"data/okex-sell-btc.txt")
		return a[0].(map[string]interface{})["price"].(string)
	}
	return ""
}

func (o okexService) LoadOkexUSDTBTC() string {
	baseUrl := "https://www.okex.com/api/spot/v3/instruments/BTC-USDT/trades"
	params := url.Values{}
	res, err := curl.Get(baseUrl, params)
	if err != nil {
		fmt.Println(err)
	}
	temp := make([]map[string]string, 0)
	err = json.Unmarshal(res, &temp)
	if err != nil {
		fmt.Println(err)
	}
	return temp[0]["price"]
}

func (o okexService) LoadOkexOTC() string {
	baseUrl := "https://www.okex.com/v3/c2c/tradingOrders/book"
	params := url.Values{}
	params.Set("t", strconv.FormatInt(time.Now().Unix()*1000, 10))
	params.Set("quoteCurrency", "CNY")
	params.Set("baseCurrency", "USDT")
	params.Set("side", "sell")
	params.Set("paymentMethod", "aliPay")
	params.Set("userType", "all")
	params.Set("quoteMinAmountPerOrder", "5000")
	res, err := curl.Get(baseUrl, params)
	if err != nil {
		fmt.Println(err)
	}
	temp := make(map[string]interface{})
	err = json.Unmarshal(res, &temp)
	if err != nil {
		fmt.Println(err)
	}
	a := temp["data"].(map[string]interface{})["sell"].([]interface{})
	if len(a) > 3 {
		// 将第三个写入本地文件
		timeNow := time.Now().Format("2006-01-02 15:04:05")

		content := "当前时间:" + timeNow + " USDT兑OKEX前三的价格为:" + a[0].(map[string]interface{})["price"].(string) + "/" + a[1].(map[string]interface{})["price"].(string) + "/" + a[2].(map[string]interface{})["price"].(string)
		content += "\n"
		util.WriteFile(content, util.LocalPath()+"data/okex.txt")
		return a[0].(map[string]interface{})["price"].(string)
	}
	return ""
}
