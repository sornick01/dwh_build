package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"

	"dwh/internal/domain"
)

func main() {

	db := &domain.Database{}

	b, err := os.ReadFile("jsons/db.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, db)

	fmt.Println(db.ToSql())

	err = os.WriteFile("output/creation.sql", []byte(db.ToSql()), 0777)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:postgrespw@localhost:32768/postgres")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec(context.Background(), db.ToSql())
	if err != nil {
		log.Fatal(err)
	}
}
