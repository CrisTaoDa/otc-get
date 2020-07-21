package controller

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"otc-get/app/service"
)

type IndexController struct {
	Ctx iris.Context
}

func (c *IndexController) BeforeActivation(b mvc.BeforeActivation) {
	//首页
	b.Handle("GET", "/", "Index")
	b.Handle("GET", "/okex/otc/buy/usdt", "OkexOtcBuyUsdt")
	b.Handle("GET", "/okex/otc/sell/btc", "OkexOtcSellBtc")
	b.Handle("GET", "/okex/btc-usdt", "OkexBtcUsdt")
	b.Handle("GET", "/test", "Test")

}

func (c *IndexController) Test() {
	service.NewOkexService().LoadOkexTest()
	c.Ctx.JSON("ok")
}
func (c *IndexController) OkexOtcSellBtc() {
	service.NewOkexService().LoadOkexOTCSellBTC()
	c.Ctx.JSON("ok")
}
func (c *IndexController) OkexBtcUsdt() {
	service.NewOkexService().LoadOkexUSDTBTC()
	c.Ctx.JSON("ok")
}
func (c *IndexController) OkexOtcBuyUsdt() {
	service.NewOkexService().LoadOkexOTC()
	c.Ctx.JSON("ok")
}

func (c *IndexController) Index() {
	err := c.Ctx.View("index/main.html")
	if err != nil {
		fmt.Println(err)
	}
}
