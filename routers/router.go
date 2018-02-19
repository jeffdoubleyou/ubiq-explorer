// @APIVersion 1.0.0
// @Title ubiq-explorer API
// @Description Ubiq explorer created for ubiq.cc
// @Contact jeffdoubleyou@gmail.com
// @TermsOfServiceUrl http://www.ubiq.cc
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ubiq-explorer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/transaction",
			beego.NSInclude(
				&controllers.TransactionController{},
			),
		),
		beego.NSNamespace("/block",
			beego.NSInclude(
				&controllers.BlockController{},
			),
		),
		beego.NSNamespace("/uncle",
			beego.NSInclude(
				&controllers.UncleController{},
			),
		),
		beego.NSNamespace("/balance",
			beego.NSInclude(
				&controllers.BalanceController{},
			),
		),
		beego.NSNamespace("/token",
			beego.NSInclude(
				&controllers.TokenController{},
			),
		),
		beego.NSNamespace("/stats",
			beego.NSInclude(
				&controllers.StatsController{},
			),
		),
		beego.NSNamespace("/address",
			beego.NSInclude(
				&controllers.AddressController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
