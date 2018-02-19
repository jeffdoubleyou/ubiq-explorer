package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
	"ubiq-explorer/services"
)

// Operations about uncles
type UncleController struct {
	beego.Controller
}

// @Title Miner
// @Description Get a list of uncles mined by address
// @Param	address	query	string	true		"Address of miner"
// @Param   limit	query   int     true        "Number of uncles to retrieve"
// @Param   cursor	query   string 	false       "Cursor string of last uncle result"
// @Param   skip	query	int		false		"Number of uncles to skip after cursor"
// @Success 200 {object} models.BlockPage
// @Failure 404 not found
// @router /miner [get]
func (c *UncleController) Miner() {
	address := c.GetString("address")
	limit, _ := c.GetInt("limit")
	cursor := c.GetString("cursor")
	uncleDAO := daos.NewUncleDAO()
	uncleService := services.NewUncleService(*uncleDAO)
	page, err := uncleService.Miner(common.HexToAddress(address), limit, cursor)
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = page
	}
	c.ServeJSON()
}

// @Title Block
// @Description Get all uncles in the given block
// @Param	block	query	int	true "The block number to retreive uncles for"
// @Success 200 {object} models.UnclePage
// @Failure 404 no uncles found in block
// @router /block [get]
func (c *UncleController) Block() {
	block, err := c.GetInt64("block")
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		blockNumber := big.NewInt(block)
		uncleDAO := daos.NewUncleDAO()
		var uncle = services.NewUncleService(*uncleDAO)
		page, err := uncle.Block(*blockNumber, 1000, "")
		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = page
		}
	}
	c.ServeJSON()
}
