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
	b.Handle("GET", "/test2", "Test2")
	b.Handle("GET", "/test3", "Test3")
	b.Handle("GET", "/test4", "Test4")
	b.Handle("GET", "/test5", "Test5")

}

func (c *IndexController) Test() {
	service.NewOkexService().LoadOkexTest()
	c.Ctx.JSON("ok")
}
func (c *IndexController) Test2() {
	service.NewOkexService().LoadOkexTest2()
	c.Ctx.JSON("ok")
}
func (c *IndexController) Test3() {
	res := service.NewOkexService().LoadOkexTest3()
	c.Ctx.JSON(res)
}
func (c *IndexController) Test4() {
	res := service.NewOkexService().LoadOkexTest4()
	c.Ctx.JSON(res)
}
func (c *IndexController) Test5() {
	name := c.Ctx.URLParam("name")
	mode := c.Ctx.URLParam("mode")
	res := service.NewOkexService().LoadOkexTest5(name, mode)
	c.Ctx.JSON(res)
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
