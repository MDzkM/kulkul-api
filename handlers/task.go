package handlers

import (
	"bytes"
	"encoding/base64"
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"fmt"
	
	"github.com/MDzkM/kulkul-api/models"

	"github.com/labstack/echo"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type H map[string]interface{}

func initAWSConnection() *s3.S3 {
	awsAccessKey := "AKIAW74DYNZLAVCSJQMC"
	awsSecretKey := "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"

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, "")

	_, err := creds.Get()

	if err != nil {
		fmt.Println(err)
	}

	s3Region := "us-east-1"

	cfg := aws.NewConfig().WithRegion(s3Region).WithCredentials(creds)

	s3Connection := s3.New(session.New(), cfg)
	return s3Connection
}

func S3UploadBase64(base64File string, objectKey string) error {
	decode, err := base64.StdEncoding.DecodeString(base64File)

	if err != nil {
		return err
	}

	awsSession := initAWSConnection()

	uploadParams := &s3.PutObjectInput{
		Bucket: aws.String("kulkul"),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(decode),
	}

	_, err = awsSession.PutObject(uploadParams)

	return err
}

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

		keyName := "fridge-" + strconv.Itoa(int(id)) + ".jpg"
		
		b64data := fridge.Image[strings.IndexByte(fridge.Image, ',')+1:]

		err = S3UploadBase64(b64data, keyName)
		
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
