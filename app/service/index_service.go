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
	LoadOkexOTCBuyBTC() string
	LoadOkexOTCSellBTC() string
	LoadOkexOTCSellUSDT() string
	LoadOkexOTCBuyEOS() string
	LoadOkexOTCSellEOS() string
	LoadOkexUSDTEOS() string

	LoadOkexOTCBuyXXX(name string) string
	LoadOkexOTCSellXXX(name string) string
	LoadOkexUSDTXXX(name string) string
	LoadOkexTest() string
	LoadOkexTest2() string
	LoadOkexTest3() interface{}
	LoadOkexTest4() interface{}
	LoadOkexTest5(name, mode string) interface{}
}

func NewOkexService() OkexService {
	return &okexService{}
}

type okexService struct {
}

func (o okexService) LoadOkexTest5(name, mode string) interface{} {
	if mode == "1" {
		buyPrice, _ := strconv.ParseFloat(o.LoadOkexOTCBuyXXX("USDT"), 64)
		usdtBtc, _ := strconv.ParseFloat(o.LoadOkexUSDTXXX(name), 64)
		sellBtc, _ := strconv.ParseFloat(o.LoadOkexOTCSellXXX(name), 64)
		fmt.Println(buyPrice, usdtBtc, sellBtc)
		return (10000.0 / buyPrice / usdtBtc * 0.9985 * sellBtc)
	} else {
		buyPrice, _ := strconv.ParseFloat(o.LoadOkexOTCBuyXXX(name), 64)
		usdtBtc, _ := strconv.ParseFloat(o.LoadOkexUSDTXXX(name), 64)
		sellBtc, _ := strconv.ParseFloat(o.LoadOkexOTCSellXXX("USDT"), 64)
		fmt.Println(buyPrice, usdtBtc, sellBtc)
		return (10000.0 / buyPrice * usdtBtc * 0.9985 * sellBtc)
	}
}

func (o okexService) LoadOkexOTCSellXXX(name string) string {
	baseUrl := "https://www.okex.com/v3/c2c/tradingOrders/book"
	params := url.Values{}
	params.Set("t", strconv.FormatInt(time.Now().Unix()*1000, 10))
	params.Set("quoteCurrency", "CNY")
	params.Set("baseCurrency", name)
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
		return a[0].(map[string]interface{})["price"].(string)
	}
	return ""
}

func (o okexService) LoadOkexOTCBuyXXX(name string) string {
	baseUrl := "https://www.okex.com/v3/c2c/tradingOrders/book"
	params := url.Values{}
	params.Set("t", strconv.FormatInt(time.Now().Unix()*1000, 10))
	params.Set("quoteCurrency", "CNY")
	params.Set("baseCurrency", name)
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
		return a[0].(map[string]interface{})["price"].(string)
	}
	return ""
}

func (o okexService) LoadOkexUSDTXXX(name string) string {
	baseUrl := "https://www.okex.com/api/spot/v3/instruments/" + name + "-USDT/trades"
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

func (o okexService) LoadOkexTest4() interface{} {

	buyPrice, _ := strconv.ParseFloat(o.LoadOkexOTCBuyEOS(), 64)
	usdtBtc, _ := strconv.ParseFloat(o.LoadOkexUSDTEOS(), 64)
	sellBtc, _ := strconv.ParseFloat(o.LoadOkexOTCSellUSDT(), 64)
	fmt.Println(buyPrice, usdtBtc, sellBtc)
	return 10000.0 / buyPrice * usdtBtc * 0.9985 * sellBtc
}

func (o okexService) LoadOkexTest3() interface{} {

	buyPrice, _ := strconv.ParseFloat(o.LoadOkexOTC(), 64)
	usdtBtc, _ := strconv.ParseFloat(o.LoadOkexUSDTEOS(), 64)
	sellBtc, _ := strconv.ParseFloat(o.LoadOkexOTCSellEOS(), 64)
	fmt.Println(buyPrice, usdtBtc, sellBtc)
	return 10000.0 / buyPrice / usdtBtc * 0.9985 * sellBtc
}

func (o okexService) LoadOkexOTCBuyEOS() string {
	baseUrl := "https://www.okex.com/v3/c2c/tradingOrders/book"
	params := url.Values{}
	params.Set("t", strconv.FormatInt(time.Now().Unix()*1000, 10))
	params.Set("quoteCurrency", "CNY")
	params.Set("baseCurrency", "EOS")
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
		return a[0].(map[string]interface{})["price"].(string)
	}
	return ""
}

func (o okexService) LoadOkexOTCSellEOS() string {
	baseUrl := "https://www.okex.com/v3/c2c/tradingOrders/book"
	params := url.Values{}
	params.Set("t", strconv.FormatInt(time.Now().Unix()*1000, 10))
	params.Set("quoteCurrency", "CNY")
	params.Set("baseCurrency", "EOS")
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
		return a[0].(map[string]interface{})["price"].(string)
	}
	return ""
}

func (o okexService) LoadOkexUSDTEOS() string {
	baseUrl := "https://www.okex.com/api/spot/v3/instruments/EOS-USDT/trades"
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

func (o okexService) LoadOkexTest2() string {

	buyPrice, _ := strconv.ParseFloat(o.LoadOkexOTCBuyBTC(), 64)
	usdtBtc, _ := strconv.ParseFloat(o.LoadOkexUSDTBTC(), 64)
	sellBtc, _ := strconv.ParseFloat(o.LoadOkexOTCSellUSDT(), 64)
	fmt.Println(buyPrice, usdtBtc, sellBtc)
	fmt.Println(10000.0 / buyPrice * usdtBtc * 0.9985 * sellBtc)
	return ""
}

func (o okexService) LoadOkexOTCBuyBTC() string {
	baseUrl := "https://www.okex.com/v3/c2c/tradingOrders/book"
	params := url.Values{}
	params.Set("t", strconv.FormatInt(time.Now().Unix()*1000, 10))
	params.Set("quoteCurrency", "CNY")
	params.Set("baseCurrency", "BTC")
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

func (o okexService) LoadOkexOTCSellUSDT() string {
	baseUrl := "https://www.okex.com/v3/c2c/tradingOrders/book"
	params := url.Values{}
	params.Set("t", strconv.FormatInt(time.Now().Unix()*1000, 10))
	params.Set("quoteCurrency", "CNY")
	params.Set("baseCurrency", "USDT")
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
		return a[0].(map[string]interface{})["price"].(string)
	}
	return ""
}

func (o okexService) LoadOkexTest() string {

	buyPrice, _ := strconv.ParseFloat(o.LoadOkexOTC(), 64)
	usdtBtc, _ := strconv.ParseFloat(o.LoadOkexUSDTBTC(), 64)
	sellBtc, _ := strconv.ParseFloat(o.LoadOkexOTCSellBTC(), 64)
	fmt.Println(buyPrice, usdtBtc, sellBtc)
	fmt.Println(10000.0 / buyPrice / usdtBtc * 0.9985 * sellBtc)
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
