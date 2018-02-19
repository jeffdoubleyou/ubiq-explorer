package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
	"ubiq-explorer/services"
)

// Operations about balances
type BalanceController struct {
	beego.Controller
}

// @Title History
// @Description Get a list of balances by address
// @Param	address	query	string	true		"Address to get balances of"
// @Param   limit	query   int     true        "Number of balance history records to retrieve"
// @Param   cursor	query   string 	false       "Cursor string of last history result"
// @Success 200 {object} models.BalancePage
// @Failure 404 not found
// @router /history [get]
func (c *BalanceController) History() {
	address := c.GetString("address")
	limit, _ := c.GetInt("limit")
	cursor := c.GetString("cursor")
	balanceDAO := daos.NewBalanceDAO()
	balanceService := services.NewBalanceService(*balanceDAO)
	page, err := balanceService.History(common.HexToAddress(address), limit, cursor)
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = page
	}
	c.ServeJSON()
}

// @Title Get
// @Description Get balance for an address
// @Param   address	query	string	true	"Address to get balance for"
// @Param	block	query	int	false	"Block number to get balance at or 0 for current block"
// @Success 200 {object} models.Balance
// @Failure 404 not found
// @router /get [get]
func (c *BalanceController) Get() {
	address := c.GetString("address")
	block, _ := c.GetInt64("block")
	balanceDAO := daos.NewBalanceDAO()
	balanceService := services.NewBalanceService(*balanceDAO)
	balance, err := balanceService.Get(common.HexToAddress(address), block)
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = balance
	}
	c.ServeJSON()
}

// @Title RichList
// @Description
// @Success 200 {object} models.CurrentBalancePage
// @Failure 404 not found
// @router /richList [get]
func (c *BalanceController) RichList() {
	balanceDAO := daos.NewBalanceDAO()
	balanceService := services.NewBalanceService(*balanceDAO)
	rich, err := balanceService.RichList(100)
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = rich
	}
	c.ServeJSON()
}
