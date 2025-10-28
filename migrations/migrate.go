package main

import (
	"github.com/shemigam1/hng13-currency-exchange/db"
	"github.com/shemigam1/hng13-currency-exchange/models"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(&models.CountryInfo{})
}
