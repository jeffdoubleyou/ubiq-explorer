package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
	"ubiq-explorer/services"
)

// Operations about blocks
type BlockController struct {
	beego.Controller
}

// @Title Get
// @Description	Get block
// @Param block query	string	true	"Block number to retrieve"
// @Success 200 {object} types.Block
// @Failure 404 block not found
// @router /get [get]
func (c *BlockController) Get() {
	block, err := c.GetInt64("block")
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		blockDAO := daos.NewBlockDAO()
		blockService := services.NewBlockService(*blockDAO)
		b, err := blockService.Get(block)
		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = b //.Header()
		}
	}
	c.ServeJSON()
}

// @Title List
// @Description Get a list of blocks
// @Param   limit	query   int     true        "Number of blocks to retrieve"
// @Param   cursor	query   string 	false       "Cursor string of last block result"
// @Param   skip	query	int		false		"Number of blocks to skip after cursor"
// @Success 200 {object} models.BlockPage
// @Failure 404 block not found
// @router /list [get]
func (c *BlockController) List() {
	limit, _ := c.GetInt("limit")
	cursor := c.GetString("cursor")
	blockDAO := daos.NewBlockDAO()
	blockService := services.NewBlockService(*blockDAO)
	page, err := blockService.List(limit, cursor)
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = page
	}
	c.ServeJSON()
}

// @Title Miner
// @Description Get a list of blocks mined by address
// @Param	address	query	string	true		"Address of miner"
// @Param   limit	query   int     true        "Number of blocks to retrieve"
// @Param   cursor	query   string 	false       "Cursor string of last block result"
// @Param   skip	query	int		false		"Number of blocks to skip after cursor"
// @Success 200 {object} models.BlockPage
// @Failure 404 not found
// @router /miner [get]
func (c *BlockController) Miner() {
	address := c.GetString("address")
	limit, _ := c.GetInt("limit")
	cursor := c.GetString("cursor")
	blockDAO := daos.NewBlockDAO()
	blockService := services.NewBlockService(*blockDAO)
	page, err := blockService.Miner(common.HexToAddress(address), limit, cursor)
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = page
	}
	c.ServeJSON()
}
