package main

import (
	"context"
	"encoding/json"
	"etl/internal/domain"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

func main() {

	db := &domain.Database{}
	routes := &domain.Routes{}

	dbDescr, err := os.ReadFile("jsons/db.json")
	if err != nil {
		log.Fatal(err)
	}
	routesDescr, err := os.ReadFile("jsons/routes.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(dbDescr, db)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(routesDescr, routes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db.ToSql())
	fmt.Println(routes.ToSql())
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:1234@localhost:5432/dst")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(context.Background(), db.ToSql())
	if err != nil {
		log.Fatal(err)
	}
	//_, err = conn.Exec(context.Background(), routes.ToSql())
	//if err != nil {
	//	log.Fatal(err)
	//}
}
