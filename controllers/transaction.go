package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
	"ubiq-explorer/services"
)

// Operations about transactions
type TransactionController struct {
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
func (c *TransactionController) From() {
	address := c.GetString("address")
	limit, _ := c.GetInt("limit")
	cursor := c.GetString("cursor")

	if address != "" {
		transactionDAO := daos.NewTransactionDAO()
		var txn = services.NewTransactionService(*transactionDAO)
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
func (c *TransactionController) To() {
	address := c.GetString("address")
	limit, _ := c.GetInt("limit")
	cursor := c.GetString("cursor")
	if address != "" {
		transactionDAO := daos.NewTransactionDAO()
		var txn = services.NewTransactionService(*transactionDAO)
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
func (c *TransactionController) List() {
	limit, _ := c.GetInt("limit")
	cursor := c.GetString("cursor")
	transactionDAO := daos.NewTransactionDAO()
	var txn = services.NewTransactionService(*transactionDAO)
	page, err := txn.List(limit, cursor)
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = page
	}
	c.ServeJSON()
}

// @Title Block
// @Description Get all transactions in the given block
// @Param	block	query	int	true "The block number to retreive transactions for"
// @Success 200 {object} models.TrasactionPage
// @Failure 404 no transactions found in block
// @router /block [get]
func (c *TransactionController) Block() {
	block, err := c.GetInt64("block")
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		blockNumber := big.NewInt(block)
		transactionDAO := daos.NewTransactionDAO()
		var txn = services.NewTransactionService(*transactionDAO)
		page, err := txn.Block(*blockNumber, 1000, "")
		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = page
		}
	}
	c.ServeJSON()
}

// @Title Get
// @Description	Get full transaction data
// @Param hash	query	string	true	"Block hash to retrieve"
// @Success 200 {object} models.TxExtraInfo
// @Failure 404 transaction not found
// @router /get [get]
func (c *TransactionController) Get() {
	hash := c.GetString("hash")
	if hash != "" {
		transactionDAO := daos.NewTransactionDAO()
		transactionService := services.NewTransactionService(*transactionDAO)
		txn, err := transactionService.Get(hash)

		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = txn
		}
	} else {
		c.Data["json"] = &models.APIError{"Invalid hash"}
		c.Ctx.ResponseWriter.WriteHeader(400)
	}
	c.ServeJSON()
}

// @Title Receipt
// @Description Get transaction receipt
// @Param hash	query	string	true	"Block hash to retrieve receipt for"
// @Success 200 {object} models.TxReceipt
// @Failure 404 transaction not found
// @router /receipt [get]
func (c *TransactionController) Receipt() {
	hash := c.GetString("hash")
	if hash != "" {
		transactionDAO := daos.NewTransactionDAO()
		transactionService := services.NewTransactionService(*transactionDAO)
		txn, err := transactionService.Receipt(hash)

		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = txn
		}
	} else {
		c.Data["json"] = &models.APIError{"Invalid hash"}
		c.Ctx.ResponseWriter.WriteHeader(400)
	}
	c.ServeJSON()
}

// @Title Debug
// @Description Get internal transactions
// @Param hash	query	string	true	"Block hash to retrieve receipt for"
// @Success 200 {object} models.TxReceipt
// @Failure 404 transaction not found
// @router /debug [get]
func (c *TransactionController) Debug() {
	hash := c.GetString("hash")
	if hash != "" {
		transactionDAO := daos.NewTransactionDAO()
		transactionService := services.NewTransactionService(*transactionDAO)
		txn, err := transactionService.Debug(hash)

		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = txn
		}
	} else {
		c.Data["json"] = &models.APIError{"Invalid hash"}
		c.Ctx.ResponseWriter.WriteHeader(400)
	}
	c.ServeJSON()
}

// @Title Pending
// @Description Get pending transactions
// @Success 200 {object} []models.RpcTransaction
// @Failure 404 transaction not found
// @router /pending [get]
func (c *TransactionController) Pending() {
	transactionDAO := daos.NewTransactionDAO()
	transactionService := services.NewTransactionService(*transactionDAO)
	txn, err := transactionService.Pending()

	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = txn
	}
	c.ServeJSON()
}
