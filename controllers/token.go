package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
	"ubiq-explorer/services"
)

// Operations about transactions
type TokenController struct {
	beego.Controller
}

// @Title From
// @Description Get transactions sent from provided address
// @Param   address	query   string  true        "The address to retrieve transactions from"
// @Param   limit	query   int     true        "Number of records to retrieve"
// @Param   cursor	query   string 	false       "Cursor string of last record result"
// @Param   skip	query	int		false		"Number of records to skip after cursor"
// @Success 200 {object} models.Transaction
// @Failure 404 no transactions found
// @router /from [get]
func (c *TokenController) From() {
	address := c.GetString("address")
	limit, _ := c.GetInt("limit")
	cursor := c.GetString("cursor")

	if address != "" {
		tokenDAO := daos.NewTokenDAO()
		var txn = services.NewTokenService(*tokenDAO)
		page, err := txn.From(common.HexToAddress(address), limit, cursor)
		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = page
		}
	}
	c.ServeJSON()
}

// @Title To
// @Description Get transactions sent to provided address
// @Param   address	query   string  true        "The address to retrieve transactions from"
// @Param   limit	query   int     true        "Number of records to retrieve"
// @Param   cursor	query   string 	false       "Cursor string of last record result"
// @Param   skip	query	int		false		"Number of records to skip after cursor"
// @Success 200 {object} models.TransactionPage
// @Failure 404 no transactions found
// @router /to [get]
func (c *TokenController) To() {
	address := c.GetString("address")
	limit, _ := c.GetInt("limit")
	cursor := c.GetString("cursor")
	if address != "" {
		tokenDAO := daos.NewTokenDAO()
		var txn = services.NewTokenService(*tokenDAO)
		page, err := txn.To(common.HexToAddress(address), limit, cursor)
		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = page
		}
	}
	c.ServeJSON()
}

// @Title List
// @Description Get a list of transactions
// @Param   limit	query   int     true        "Number of records to retrieve"
// @Param   cursor	query   string 	false       "Cursor string of last record result"
// @Param   skip	query	int		false		"Number of records to skip after cursor"
// @Success 200 {object} models.TransactionPage
// @Failure 404 no transactions found
// @router /list [get]
func (c *TokenController) List() {
	limit, _ := c.GetInt("limit")
	cursor := c.GetString("cursor")
	tokenDAO := daos.NewTokenDAO()
	var txn = services.NewTokenService(*tokenDAO)
	page, err := txn.List(limit, cursor)
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = page
	}
	c.ServeJSON()
}

// @Title Address
// @Description Get token info at address
// @Param address	query	string	true	"Address to retrieve token at"
// @Success 200 {object} models.TokenInfo
// @Failure 404 no token found at address
// @router /address [get]
func (c *TokenController) Address() {
	address := c.GetString("address")
	if address != "" {
		tokenDAO := daos.NewTokenDAO()
		var tokenService = services.NewTokenService(*tokenDAO)
		token, err := tokenService.GetTokenByAddress(common.HexToAddress(address))
		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = token
		}
	}
	c.ServeJSON()
}

// @Title Symbol
// @Description Get token info by symbol
// @Param symbol	query	string	true	"Token Symbol to retrievei info about"
// @Success 200 {object} models.TokenInfo
// @Failure 404 no token found with given symbol
// @router /symbol [get]
func (c *TokenController) Symbol() {
	symbol := c.GetString("symbol")
	if symbol != "" {
		tokenDAO := daos.NewTokenDAO()
		var tokenService = services.NewTokenService(*tokenDAO)
		token, err := tokenService.GetTokenBySymbol(symbol)
		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = token
		}
	}
	c.ServeJSON()
}

// @Title Balance
// @Description Get token balances by address
// @Param address query string true "Address to get token balances for"
// @Success 200 {object} []models.TokenBalance
// @Failure 404 no balances found for address
// @router /balance [get]
func (c *TokenController) Balance() {
	address := c.GetString("address")
	if address != "" {
		tokenBalanceDAO := daos.NewTokenBalanceDAO()
		var tokenBalanceService = services.NewTokenBalanceService(*tokenBalanceDAO)
		balances, err := tokenBalanceService.Address(common.HexToAddress(address))
		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = balances
		}
	}
	c.ServeJSON()
}
