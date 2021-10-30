package db_test

import (
	"database/sql"
	"log"
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
)

var Db *sql.DB

func setUp(){
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProducts(Db)
}

func createTable(db *sql.DB){
	table:= `CREATE TABLE products(
		"id" string, 
		"name" string, 
		"price" float, 
		"status" string
		);`

	stmt, err:= db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}
func createProducts(db *sql.DB){
	insert:= `INSERT INTO products values(
		"abc", 
		"Product Test", 
		0, 
		"disabled"
		);`

	stmt, err:= db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Read(t *testing.T){
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Read("abc")

	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t,0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}
func TestProductDb_Write(t *testing.T){
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product:= application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	result, err := productDb.Write(product)

	require.Nil(t, err)
	require.Equal(t,product.GetName() , result.GetName() )
	require.Equal(t, product.GetPrice(),result.GetPrice())
	require.Equal(t, product.GetStatus(),result.GetStatus())


	product.Status = "enabled"


	result, err = productDb.Write(product)

	require.Nil(t, err)
	require.Equal(t,product.GetName() , result.GetName() )
	require.Equal(t, product.GetPrice(),result.GetPrice())
	require.Equal(t, product.GetStatus(),result.GetStatus())

}