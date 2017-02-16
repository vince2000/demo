package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"

	"io/ioutil"
	"encoding/json"
)

// com cc

type User struct {
	Name string
	Age int `json:"age"`
}

// b dd

func main() {
 //dasdsadsadsadsadsa
	/*list:=make([]string,0)
	list=append(list,"hahha")
	list=append(list,"qw")
	list=append(list,"as")
	list=append(list,"ds")
	list=append(list,"cxz")
	for _,r:=range list  {
		fmt.Println(r)
	}*/
	/*c:=make(chan string,0)

	go func() {c<-"haha"}()
	a:=<-c
	fmt.Println(a)*/
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		/*	name1:=c.FormValue("name")
		      fmt.Println(name1)*/
		body, _ := ioutil.ReadAll(c.Request().Body)
		//vm := map[string]interface{}{}
		user := &User{}
		json.Unmarshal(body, &user)
		fmt.Println("name: ", user)
		//name:= c.QueryParam("name")
		//m:=make(map[string]string)
		fmt.Println("hahah")
		return c.JSON(http.StatusOK, user)
		//return c.String(http.StatusOK,name)
	})
          e.Logger.Fatal(e.Start(":1323"))
}
//非常有意思
