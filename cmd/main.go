package main

import (
	"fmt"
	"github.com/Mamvriyskiy/TeamDev/pkg/repository"
	"github.com/Mamvriyskiy/TeamDev/pkg/service"
	"github.com/Mamvriyskiy/TeamDev/pkg/handler"
	"github.com/Mamvriyskiy/TeamDev/pkg"
	"log"
	// "github.com/joho/godotenv"
	// "github.com/spf13/viper"
)

func main() {
	fmt.Println("Server start...")
	// if err := initConfig(); err != nil {
	// 	logger.Log("Error", "initCongig", "Error config DB:", err, "")
	// 	fmt.Println(err)
	// 	return
	// }
	// logger.Log("Info", "", "InitConfig", nil)

	// if err := godotenv.Load(); err != nil {
	// 	logger.Log("Error", "Load", "Load env file:", err, "")
	// 	fmt.Println(err)
	// 	return
	// }
	// logger.Log("Info", "", "Load env", nil)

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Println("Error", "initCongig", "Error config DB:", err, "")
		return
	}

	repos := repository.NewRepository(db)
	services := service.NewServicesPsql(repos)
	handlers := handler.NewHandler(services)

	// logger.Log("Info", "", "The connection to the database is established", nil)

	srv := new(pkg.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("Error", "Run", "Error occurred while running http server:", err, "")
		return
	}
}

// func initConfig() error {
// 	viper.AddConfigPath("configs")
// 	viper.SetConfigName("config")
// 	return viper.ReadInConfig()
// }
