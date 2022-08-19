package models

import (
	"database/sql"
)

type Fridge struct {
	ID int `json:"id"`
	Model string `json:"model"`
	Owner string `json:"owner"`
	Image string `json:"image"`
}

type FridgeCollection struct {
	Fridges []Fridge `json:"fridges"`
}

func GetFridges(db *sql.DB) FridgeCollection {
	sql := "SELECT * FROM fridges"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := FridgeCollection{}
	for rows.Next() {
		fridge := Fridge{}
		err2 := rows.Scan(&fridge.ID, &fridge.Model, &fridge.Owner, &fridge.Image)

		if err2 != nil {
			panic(err2)
		}

		result.Fridges = append(result.Fridges, fridge)
	}
	return result
}

func PutFridge(db *sql.DB, model string, owner string, image string) (int64, error) {
	sql := "INSERT INTO fridges(model, owner, image) VALUES(?,?,?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(model, owner, image)

	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func EditFridge(db *sql.DB, fridgeId int, model string, owner string, image string) (int64, error) {
	sql := "UPDATE fridges set model = ?, owner = ?, image = ? WHERE id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(model, owner, image, fridgeId)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func DeleteFridge(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM fridges WHERE id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
