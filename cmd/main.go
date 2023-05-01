package main

import (
	"context"
	"dwh/internal/domain"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

func main() {

	db := &domain.Database{}

	b, err := os.ReadFile("jsons/db.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, db)

	//fmt.Println(db.ToSql())

	fmt.Println(db.ToSql())
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:1234@localhost:5432/dst")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(context.Background(), db.ToSql())
	if err != nil {
		log.Fatal(err)
	}
}
