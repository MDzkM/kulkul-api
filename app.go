package main

import (
	"database/sql"
	
	"github.com/MDzkM/kulkul-api/handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()
	db := initDB("fridge.db")
	migrate(db)

	e.GET("/fridge", handlers.GetFridges(db))
	e.POST("/fridge", handlers.PutFridge(db))
	e.PUT("/fridge", handlers.EditFridge(db))
	e.DELETE("/fridge/:id", handlers.DeleteFridge(db))

	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS fridges(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		model VARCHAR NOT NULL,
		owner VARCHAR NOT NULL,
		image VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
