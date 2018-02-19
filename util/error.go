package util

import (
	"fmt"
	"github.com/astaxie/beego"
)

func Error(c *beego.Controller, Error string) {
	fmt.Println("OMG OMG\n\n\n\n")
	fmt.Println(Error, c.Data)
}
