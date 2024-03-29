package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/mustafakemalgordesli/go-commerce/controllers"
	"github.com/mustafakemalgordesli/go-commerce/pkg/middlewares"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.GET("/", middlewares.VerifyJWT(), func(c *gin.Context) {
		responseData := gin.H{
			"success": true,
			"message": "Hello World!",
		}
		c.JSON(200, responseData)
	})

	authController := controllers.NewAuthController(db)
	authRouter := r.Group("/auth")
	{
		authRouter.POST("/signup", authController.SignUp)
		authRouter.POST("/signin", authController.SignIn)
	}

	categoryController := controllers.NewCategoryController(db)
	categoryRouter := r.Group("/category")
	{
		categoryRouter.GET("/", categoryController.GetAllCategories)
		categoryRouter.POST("/", categoryController.AddCategory)
	}

	return r
}
