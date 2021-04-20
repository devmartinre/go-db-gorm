package main

import (
	"fmt"

	"github.com/whiteagleinc-meli/go-db-gorm/model"
	"github.com/whiteagleinc-meli/go-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	// Create tables
	// storage.DB().AutoMigrate(
	// 	&model.Product{},
	// 	&model.InvoiceHeader{},
	// 	&model.InvoiceItem{},
	// )

	// product1 := model.Product{
	// 	Name:  "Curso de Go",
	// 	Price: 120,
	// }

	// obs := "testing with go"

	// product2 := model.Product{
	// 	Name:         "Curso de Testing",
	// 	Price:        150,
	// 	Observations: &obs,
	// }

	// product3 := model.Product{
	// 	Name:  "Curso de Python",
	// 	Price: 200,
	// }

	// storage.DB().Create(&product1)
	// storage.DB().Create(&product2)
	// storage.DB().Create(&product3)

	// products := make([]model.Product, 0)
	// storage.DB().Find(&products)

	// for _, product := range products {
	// 	fmt.Printf("%d - %s\n", product.ID, product.Name)
	// }

	myProduct := model.Product{}
	storage.DB().First(&myProduct, 3)
	fmt.Println(myProduct)
}