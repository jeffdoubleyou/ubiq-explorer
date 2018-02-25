package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ubiq-explorer/controllers:AddressController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:AddressController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:BalanceController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:BalanceController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:BalanceController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:BalanceController"],
		beego.ControllerComments{
			Method: "History",
			Router: `/history`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:BalanceController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:BalanceController"],
		beego.ControllerComments{
			Method: "RichList",
			Router: `/richList`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:BlockController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:BlockController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:BlockController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:BlockController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:BlockController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:BlockController"],
		beego.ControllerComments{
			Method: "Miner",
			Router: `/miner`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"],
		beego.ControllerComments{
			Method: "BlockTimeHistory",
			Router: `/blockTimeHistory`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"],
		beego.ControllerComments{
			Method: "DifficultyHistory",
			Router: `/difficultyHistory`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"],
		beego.ControllerComments{
			Method: "HashRateHistory",
			Router: `/hashRateHistory`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"],
		beego.ControllerComments{
			Method: "Miners",
			Router: `/miners`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"],
		beego.ControllerComments{
			Method: "Pools",
			Router: `/pools`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:StatsController"],
		beego.ControllerComments{
			Method: "UncleRateHistory",
			Router: `/uncleRateHistory`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"],
		beego.ControllerComments{
			Method: "Address",
			Router: `/address`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"],
		beego.ControllerComments{
			Method: "Balance",
			Router: `/balance`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"],
		beego.ControllerComments{
			Method: "From",
			Router: `/from`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"],
		beego.ControllerComments{
			Method: "ListTokens",
			Router: `/listTokens`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"],
		beego.ControllerComments{
			Method: "ListTransactions",
			Router: `/listTransactions`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"],
		beego.ControllerComments{
			Method: "Symbol",
			Router: `/symbol`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TokenController"],
		beego.ControllerComments{
			Method: "To",
			Router: `/to`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TransactionController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TransactionController"],
		beego.ControllerComments{
			Method: "Block",
			Router: `/block`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TransactionController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TransactionController"],
		beego.ControllerComments{
			Method: "From",
			Router: `/from`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TransactionController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TransactionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TransactionController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TransactionController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TransactionController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TransactionController"],
		beego.ControllerComments{
			Method: "Receipt",
			Router: `/receipt`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:TransactionController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:TransactionController"],
		beego.ControllerComments{
			Method: "To",
			Router: `/to`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:UncleController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:UncleController"],
		beego.ControllerComments{
			Method: "Block",
			Router: `/block`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ubiq-explorer/controllers:UncleController"] = append(beego.GlobalControllerRouter["ubiq-explorer/controllers:UncleController"],
		beego.ControllerComments{
			Method: "Miner",
			Router: `/miner`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
