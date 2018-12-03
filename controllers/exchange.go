package controllers

import (
	"github.com/astaxie/beego"
	"ubiq-explorer/models"
	"ubiq-explorer/services"
)

// Operations about exchange rates
type ExchangeController struct {
	beego.Controller
}

// @Title Symbol
// @Description Get the exchange rate of a symbol
// @Param	symbol	query	string	true		"Symbol of asset to get exchange rates for"
// @Success 200 {object} models.ExchangeRate
// @Failure 404 not found
// @router /symbol [get]
func (c *ExchangeController) Symbol() {
	symbol := c.GetString("symbol")
	if symbol == "" {
		c.Data["json"] = &models.APIError{"Symbol is required"}
		c.Ctx.ResponseWriter.WriteHeader(400)
	} else {
		exchangeService := services.NewExchangeService()
		rate, err := exchangeService.Get(symbol)
		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = rate
		}
	}
	c.ServeJSON()
}

// @Title List
// @Description Get a list of all available exchange rates
// @Success 200 {object} []models.ExchangeRate
// @Failure 404 not found
// @router /list [get]
func (c *ExchangeController) List() {
	exchangeService := services.NewExchangeService()
	rate, err := exchangeService.List()
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(500)
	} else {
		c.Data["json"] = rate
	}
	c.ServeJSON()
}

// @Title History
// @Description Get the exchange rate history of a symbol
// @Param	symbol	query	string	true		"Symbol of asset to get exchange rate history for"
// @Success 200 {object} []models.ExchangeRate
// @Failure 404 not found
// @router /history [get]
func (c *ExchangeController) History() {
	symbol := c.GetString("symbol")
	if symbol == "" {
		c.Data["json"] = &models.APIError{"Symbol is required"}
		c.Ctx.ResponseWriter.WriteHeader(400)
	} else {
		exchangeService := services.NewExchangeService()
		rate, err := exchangeService.History(symbol)
		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = rate
		}
	}
	c.ServeJSON()
}
