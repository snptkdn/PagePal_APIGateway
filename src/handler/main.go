package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"page-pal/controller"

	_ "github.com/go-sql-driver/mysql"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
)

var ginLambda *ginadapter.GinLambda

func getGin() *gin.Engine {
	engine := gin.Default()
  
  engine.Use(cors.New(cors.Config{
    // アクセスを許可したいアクセス元
    AllowOrigins: []string{
      os.Getenv("FRONT_HOST"),
    },
    AllowMethods: []string{
        "POST",
        "GET",
        "OPTIONS",
    },
    AllowHeaders: []string{
        "Access-Control-Allow-Credentials",
        "Access-Control-Allow-Headers",
        "Content-Type",
        "Content-Length",
        "Accept-Encoding",
        "Authorization",
    },
    AllowCredentials: true,
  }))
  
	engine.GET("/", controller.IndexHandler)
	engine.GET("/hello", controller.HelloHandler)
	engine.POST("/signup", controller.SignUpHandler)
  engine.POST("/signin", controller.SignInHandler)
  engine.POST("/books", controller.BookHandler)
  engine.POST("/migrate", controller.MigrateHandler)

  return engine
}

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	r := getGin()
	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	environment := os.Getenv("ENVIRONMENT")
	fmt.Printf("environment:%s\n", environment)
	if environment == "develop" {
		getGin().Run(":3000")
	} else {
		lambda.Start(Handler)
	}
}
