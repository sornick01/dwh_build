package main

import (
	"context"
	smp "dwh/internal/domain"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"io/ioutil"
	"log"
)

func main() {

	//db := smp.Database{
	//	Name: "db",
	//	Tables: map[string]*smp.Table{
	//		"table1": &smp.Table{
	//			Attributes: map[string]*smp.Attribute{
	//				"column1": &smp.Attribute{
	//					Type: "serial",
	//				},
	//				"column2": {
	//					Type: "int",
	//				},
	//			},
	//		},
	//		"table2": {
	//			Attributes: map[string]*smp.Attribute{
	//				"column1": &smp.Attribute{
	//					Type: "varchar",
	//				},
	//				"column2": {
	//					Type: "text",
	//				},
	//			},
	//		},
	//	},
	//	Pool: nil,
	//}
	db := smp.Database{}

	b, err := ioutil.ReadFile("jsons/db.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, &db)
	if err != nil {
		log.Fatal(err)
	}

	db.Pool, err = pgxpool.New(context.Background(), "postgres://postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db.CreateDatabase())
	//fmt.Println(db.Tables["table1"].CreateTable(nil))
}
