package main

import (
	"encoding/json"
	"fmt"
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

}
