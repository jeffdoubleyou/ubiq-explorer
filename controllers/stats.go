package controllers

import (
	"github.com/astaxie/beego"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
	"ubiq-explorer/services"
)

// Operations about blocks
type StatsController struct {
	beego.Controller
}

// @Title Get
// @Description	Get stats
// @Param blocks query	int	true	"Number of blocks to get stats for"
// @Success 200 {object} models.Stats
// @Failure 404 not found
// @router /get [get]
func (c *StatsController) Get() {
	blocks, err := c.GetInt("blocks")
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		statsService := services.NewStatsService()
		stats, err := statsService.Get(blocks)
		if err != nil {
			c.Data["json"] = &models.APIError{err.Error()}
			c.Ctx.ResponseWriter.WriteHeader(404)
		} else {
			c.Data["json"] = stats
		}
	}
	c.ServeJSON()
}

// @Title HashRateHistory
// @Description Get Hash Rate History
// @Success 200 {object} []interface{}
// @Failure 404 not found
// @router /hashRateHistory [get]
func (c *StatsController) HashRateHistory() {
	statsService := services.NewStatsService()
	history, err := statsService.HashRateHistory()
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = history
	}
	c.ServeJSON()
}

// @Title DifficultyHistory
// @Description Get Difficulty History
// @Success 200 {object} []interface{}
// @Failure 404 not found
// @router /difficultyHistory [get]
func (c *StatsController) DifficultyHistory() {
	statsService := services.NewStatsService()
	history, err := statsService.DifficultyHistory()
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = history
	}
	c.ServeJSON()
}

// @Title BlockTimeHistory
// @Description Get Block Time History
// @Success 200 {object} []interface{}
// @Failure 404 not found
// @router /blockTimeHistory [get]
func (c *StatsController) BlockTimeHistory() {
	statsService := services.NewStatsService()
	history, err := statsService.BlockTimeHistory()
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = history
	}
	c.ServeJSON()
}

// @Title UncleRateHistory
// @Description Get Uncle Rate History
// @Success 200 {object} []interface{}
// @Failure 404 not found
// @router /uncleRateHistory [get]
func (c *StatsController) UncleRateHistory() {
	statsService := services.NewStatsService()
	history, err := statsService.UncleRateHistory()
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = history
	}
	c.ServeJSON()
}

// @Title Miners
// @Description Get Miner Stat History
// @Success 200 {object} []models.MinerList
// @Failure 404 not found
// @router /miners [get]
func (c *StatsController) Miners() {
	statsService := services.NewStatsService()
	minerWindow, _ := beego.AppConfig.Int("stats::miner_window")
	history, err := statsService.MinerList(minerWindow)
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = history
	}
	c.ServeJSON()
}

// @Title Pools
// @Description Get Network Pool Stats
// @Success 200 {object} []models.Pool
// @Failure 404 Not found
// @router /pools [get]
func (c *StatsController) Pools() {
	poolsDAO := daos.NewPoolsDAO()
	poolsService := services.NewPoolsService(*poolsDAO)
	pools, err := poolsService.List()
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = pools
	}
	c.ServeJSON()
}
