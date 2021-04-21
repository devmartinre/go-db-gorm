package main

import (
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

	// myProduct := model.Product{}
	// storage.DB().First(&myProduct, 2)
	// fmt.Println(myProduct)

	// myProduct := model.Product{}
	// myProduct.ID = 3

	// storage.DB().Model(&myProduct).Updates(
	// 	model.Product{Name: "Curso de Java", Price: 120},
	// )

	// myProduct := model.Product{}
	// myProduct.ID = 3

	// storage.DB().Delete(&myProduct)

	// myProduct := model.Product{}
	// myProduct.ID = 2

	// storage.DB().Unscoped().Delete(&myProduct)

	// storage.DB().Model(&model.InvoiceItem{}).AddForeignKey("product_id",
	// 	"products(id)", "RESTRICT", "RESTRICT")
	// storage.DB().Model(&model.InvoiceHeader{}).AddForeignKey("invoice_header_id",
	// 	"invoice_headers(id)", "RESTRICT", "RESTRICT")

	invoice := model.InvoiceHeader{
		Client: "Martin Rangel",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 1},
		},
	}

	storage.DB().Create(&invoice)
}
