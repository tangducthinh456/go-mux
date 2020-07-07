package main

import (
	"os"
	"testing"
	"log"
)

var a App

func TestMain(m *testing.M){
	a.Initialize(
		//os.Getenv("APP_DB_USERNAME"),
		//os.Getenv("APP_DB_PASSWORD"),
		"postgres",
		"",
		"postgres",
		//os.Getenv("APP_DB_NAME"),
		//"postgres", "1", "postgres",
	)
	ensureTableExist()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExist(){
	if _, err := a.DB.Exec(tableCreationQuery); err != nil{
		log.Fatal(err)
	}
}

func clearTable(){
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXIST products
(
	id SERIAL,
	name TEXT NOT NULL,
	price NUMERIC(10, 2) NOT NULL DEFAULT 0.00,
	CONSTRAINT products_pkey PRIMARY KEY (id)
)`
