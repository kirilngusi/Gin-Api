package app

import (
	"github.com/gin-gonic/gin"
)

func SetupAndRunApp() error {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//// load env
	//err := config.LoadENV()
	//if err != nil {
	//	return err
	//}

	//// start database
	//err = database.StartMongoDB()
	//if err != nil {
	//	return err
	//}
	//
	//// defer closing database
	//defer database.CloseMongoDB()

	//// attach middleware
	//app.Use(recover.New())
	//app.Use(logger.New(logger.Config{
	//	Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	//}))
	//
	//// setup routes
	//router.SetupRoutes(app)
	//
	//// attach swagger
	//config.AddSwaggerRoutes(app)
	//
	//// get the port and start
	//port := os.Getenv("PORT")
	//app.Listen(":" + port)

	r.Run() // listen and serve on 0.0.0.0:8080

	return nil
}
