package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	
	"github.com/MDzkM/kulkul-api/models"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetFridges(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetFridges(db))
	}
}

func PutFridge(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var fridge models.Fridge

		c.Bind(&fridge)

		id, err := models.PutFridge(db, fridge.Model, fridge.Owner, fridge.Image)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}

	}
}

func EditFridge(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var fridge models.Fridge
		c.Bind(&fridge)

		_, err := models.EditFridge(db, fridge.ID, fridge.Model, fridge.Owner, fridge.Image)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"updated": fridge,
			})
		} else {
			return err
		}
	}
}

func DeleteFridge(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		_, err := models.DeleteFridge(db, id)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}

	}
}
