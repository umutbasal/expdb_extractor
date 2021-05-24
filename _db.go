package extract

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mithrandie/csvq-driver"
	"github.com/mithrandie/csvq/lib/query"
)

type row struct {
	id          int
	file        string
	description string
	date        string
	author      string
	_type       string
	platform    string
	port        int
}

type Database interface {
	GetRowByFileName(ctx context.Context, filename string) (row, error)
}

type database struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) Database {
	return database{db: db}
}

func (d database) GetRowByFileName(ctx context.Context, fileName string) (row row, err error) {
	db := d.db
	queryString := fmt.Sprintf("SELECT id, file, description, date, author, type, platform, coalesce(port, 0)  FROM `files_exploits.csv` WHERE file = '%s'", fileName)
	r := db.QueryRowContext(ctx, queryString)
	if err := r.Scan(&row.id, &row.file, &row.description, &row.date, &row.author, &row._type, &row.platform, &row.port); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Rows.")
		} else if csvqerr, ok := err.(query.Error); ok {
			panic(fmt.Sprintf("Unexpected error: Number: %d  Message: %s", csvqerr.Number(), csvqerr.Message()))
		} else {
			panic("Unexpected error: " + err.Error())
		}
	} else {
		// fmt.Printf("Result: [id] %3d ", row.id)
	}
	return row, nil
}
