package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func productsShow(c echo.Context) error {
	return c.String(http.StatusOK, "lista de produtos")
}
func getProduct(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func showProduct(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+"menber:"+member)
}
func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+"email:"+email)
}
func saveProduct(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name: "+name+"email:"+email)
}
func deleteProduct(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "Product deleted:"+id)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world!")
	})
	//route
	e.GET("/products", productsShow)
	e.GET("/product/:id", getProduct)
	e.GET("/showProduct", showProduct)
	e.POST("/saveProduct", saveProduct)
	e.PUT("/product/:id", deleteProduct)
	e.Logger.Fatal(e.Start(":1234"))
}
