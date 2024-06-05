package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase(databaseFilePath string) error {
	fmt.Println("in create database")
	err := CreateDatabaseFileIfNotExists(databaseFilePath)
	if err != nil {
		return err
	}
	db, err := GetDatabase(databaseFilePath)
	if err != nil {
		return err
	}
	err = createDomainTable(db)
	if err != nil {
		return err
	}
	err = createCompetitionTable(db)
	if err != nil {
		return err
	}
	err = createCompetitorTable(db)
	if err != nil {
		return err
	}
	err = createEventTable(db)
	if err != nil {
		return err
	}
	return nil
}

func createDomainTable(db *sql.DB) error {
	createDomainQuery := `CREATE TABLE IF NOT EXISTS domain(
		source_id INTEGER NOT NULL,
		domain_id TEXT NOT NULL,
		domain_name TEXT NOT NULL,
		matched_at TIMESTAMP NOT NULL);`

	return createTable(db, createDomainQuery)
}

func createCompetitionTable(db *sql.DB) error {
	createCompetitionQuery := `CREATE TABLE IF NOT EXISTS competition(
		competition_id INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT NOT NULL,
		name TEXT NOT NULL);`
	return createTable(db, createCompetitionQuery)
}

func createCompetitorTable(db *sql.DB) error {
	createCompetitorQuery := `CREATE TABLE IF NOT EXISTS competitor(
		competitor_id INTEGER PRIMARY KEY AUTOINCREMENT,
		id TEXT NOT NULL,
		name TEXT NOT NULL);`
	return createTable(db, createCompetitorQuery)
}

func createEventTable(db *sql.DB) error {
	createEventQuery := `CREATE TABLE IF NOT EXISTS event(
		id TEXT NOT NULL,
		source_id INTEGER NOT NULL,
		domain_id TEXT NOT NULL,
		name TEXT NOT NULL,
		full_data INTEGER NOT NULL,
		type TEXT NOT NULL,
		start_date TIMESTAMP,
		competition_id INTEGER NOT NULL,
		competitor1_id INTEGER NOT NULL,
		competitor2_id INTEGER NOT NULL,
		matched_at TIMESTAMP NOT NULL);`
	return createTable(db, createEventQuery)
}

func createTable(db *sql.DB, query string) error {
	if _, err := db.Exec(query); err != nil {
		return err
	}
	return nil
}
