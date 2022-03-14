package db

import "fmt"

type Db struct {
	DbFile string
}

type Urls struct {
	Id              int
	Url             string
	Title           string
	Visit_count     int
	Typed_count     int
	Last_visit_time int
	Hidden          int
}

func (db Db) PrintDbFileName() {
	fmt.Printf("Database file: %s", db.DbFile)
}
