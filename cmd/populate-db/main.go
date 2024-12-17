// $ source .env.development # or use dotenv
// $ go run ./cmd/populate-db
package main

import (
	"log"
	"math/rand/v2"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/utgwkk/20241217-sketch-mysql-and-goqu/internal"
)

func main() {
	db, err := internal.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	goquDB := goqu.Dialect("mysql").DB(db)
	if _, err := goquDB.Truncate("tbl").Executor().Exec(); err != nil {
		log.Fatal(err)
	}

	numRows := 1000000
	rows := make([]goqu.Record, 0, numRows)
	for i := range numRows {
		id := i + 1
		row := goqu.Record{
			"id":        id,
			"is_active": rand.Float64() < 0.005,
		}
		rows = append(rows, row)
	}

	if _, err := goquDB.Insert("tbl").Rows(rows).Executor().Exec(); err != nil {
		log.Fatal(err)
	}
}
