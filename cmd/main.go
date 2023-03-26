package main

import (
	"ecommerce/handler"
	"ecommerce/logs"
	"ecommerce/repository"
	"ecommerce/service"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

const jwtSecret = "secret"

func main() {
	fmt.Println("Staring Go")
	initTimeZone()
	initConfig()
	db := initDatabase()
	productRepoDB := repository.NewProductRepositoryDB(db)
	productService := service.NewProductService(productRepoDB)
	productHandler := handler.NewProductHandler(productService)

	orderRepoDB := repository.NewOrderRepositoryDB(db)
	orderService := service.NewOrderService(orderRepoDB, productRepoDB)
	orderHandler := handler.NewOrderHandler(orderService)

	userRepoDB := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepoDB)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Post("/signup", userHandler.SignUp)
	app.Post("/login", userHandler.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSecret),
	}))

	api := app.Group("/api")
	api.Get("/orders", orderHandler.GetOrders)
	api.Get("/order/{id}", orderHandler.GetOrder)
	api.Post("/order/{id}", orderHandler.CreateOrder)
	api.Delete("/order/{id}", orderHandler.CancelOrder)

	api.Get("/products", productHandler.GetAllProduct)
	api.Get("/product/{id}", productHandler.GetProduct)
	api.Post("/product", productHandler.NewProduct)

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	app.Listen(viper.GetString("app.port"))
	fmt.Println("Staring Go")

}
func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("..")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logs.Error(err)
		panic(err)
	}

}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		logs.Error(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		logs.Error(err)
	}
	return db
}
