
 # use echo framework to create restfull  in golang
1 - go version 
2 - mkdir myapp
3 - go mod init myapp
4 - create server.go
5 - go run server.go
#  rotas -

	e.GET("/products", productsShow)
	e.GET("/product/:id", getUser)
	e.GET("/showProduct", showUser)
	e.POST("/saveProduct", saveUser)
	e.PUT("/product/:id", deleteUser)