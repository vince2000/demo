package main

import (
	"fmt"
	"net/http"
)

func main() {
 //dasdsadsadsadsadsa
	list:=make([]string,0)
	list=append(list,"hahha")
	list=append(list,"qw")
	list=append(list,"as")
	list=append(list,"ds")
	list=append(list,"cxz")
	for _,r:=range list  {
		fmt.Println(r)
	}
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
//非常有意思