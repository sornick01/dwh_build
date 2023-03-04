package main

import (
	"dwh/internal/domain"
	"encoding/json"
	"fmt"
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

	fmt.Println(db.ToSql())

	err = os.WriteFile("output/creation.sql", []byte(db.ToSql()), 0777)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db.ToSql())
	//conn, err := pgx.Connect(context.Background(), "postgresql://postgres:1234@localhost:5432/postgres")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//_, err = conn.Exec(context.Background(), db.ToSql())
	//if err != nil {
	//	log.Fatal(err)
	//}
}
