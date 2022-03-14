package main

import (
	"database/sql"
	"dbBrowser/db"
	"flag"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {
	var tables bool
	var table string
	var dbfile string

	flag.BoolVar(&tables, "tables", false, "Prints current version and exits")
	flag.StringVar(&dbfile, "dbfile", "", "Database file path to investigate")
	flag.StringVar(&table, "table", "", "Table name to investigate")
	flag.Parse()

	db := db.Db{
		DbFile: dbfile,
	}

	sqliteDatabase, _ := sql.Open("sqlite3", db.DbFile) // Open the created SQLite File
	defer sqliteDatabase.Close()                        // Defer Closing the database

	if dbfile != "" {
		if tables {
			fmt.Println("Getting table names...")
			printTables(sqliteDatabase)
			return
		}

		if table != "" {
			printTable(sqliteDatabase, table)
		}
	} else {
		fmt.Println("[!] Please specify the database file path with --dbfile.")
	}

}

func printTable(sqliteDatabase *sql.DB, tableName string) {
	fmt.Println("-----------------------------------------")
	fmt.Printf("Investigating the table: %s |", tableName)
	fmt.Println()
	fmt.Println("-----------------------------------------")
	fmt.Println()
	sqlQuery := "SELECT * FROM " + tableName
	rows, err := sqliteDatabase.Query(sqlQuery)

	var urlsList []db.Urls
	var urls db.Urls

	var (
		id              int
		url             string
		title           string
		visit_count     int
		typed_count     int
		last_visit_time int
		hidden          int
	)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&id, &url, &title, &visit_count, &typed_count, &last_visit_time, &hidden)

		urls = db.Urls{
			Id:              id,
			Url:             url,
			Title:           title,
			Visit_count:     visit_count,
			Typed_count:     typed_count,
			Last_visit_time: last_visit_time,
			Hidden:          hidden}

		urlsList = append(urlsList, urls)
	}

	fmt.Print(" ID  |            TITLE              |                            URL                            | VISIT COUNT ")
	fmt.Println()
	for _, row := range urlsList {
		fmt.Printf(" %d | %s | %s | %d ", row.Id, row.Title, row.Url, row.Visit_count)
		fmt.Println()
		fmt.Println("------------------------")
	}
}

func printTables(sqliteDatabase *sql.DB) {
	tables, err := sqliteDatabase.Query("SELECT name FROM sqlite_schema WHERE type ='table' AND name NOT LIKE 'sqlite_%';")
	if err != nil {
		log.Fatal(err)
	}
	defer tables.Close()

	for tables.Next() { // Iterate and fetch the records from result cursor
		var name string
		tables.Scan(&name)
		log.Printf("|[+] Table Name: %s", name)

	}

}
