package main

import (
	"github.com/Oluwatise-Ajayi/hng13-currency-exchange/db"
	"github.com/Oluwatise-Ajayi/hng13-currency-exchange/models"
)

func init() {
	db.LoadEnvVariables()
	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(&models.CountryInfo{})
}
