package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.New()
	engine.Use(gin.Logger())
	group := engine.Group("/api")
	{
		g1 := group.Group("nobreaks")
		{
			g1.GET("test")
		}


		g2 := group.Group("hystrix")
		{
			g2.GET("test")
		}


		g3 := group.Group("sentinel")
		{
			g3.GET("test")
		}
	}
}
