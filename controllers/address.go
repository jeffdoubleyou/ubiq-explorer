package controllers

import (
	"github.com/astaxie/beego"
	"ubiq-explorer/daos"
	"ubiq-explorer/models"
	"ubiq-explorer/services"
)

// Operations about uncles
type AddressController struct {
	beego.Controller
}

// @Title List
// @Description Get all named addresses
// @Success 200 {object} []models.AddressInfo
// @Failure 404 no found addresses
// @router /list [get]
func (c *AddressController) List() {
	addressDAO := daos.NewAddressDAO()
	var address = services.NewAddressService(*addressDAO)
	page, err := address.List()
	if err != nil {
		c.Data["json"] = &models.APIError{err.Error()}
		c.Ctx.ResponseWriter.WriteHeader(404)
	} else {
		c.Data["json"] = page
	}
	c.ServeJSON()
}
