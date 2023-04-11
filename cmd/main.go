package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-test-vebra-group/internal/jsoncustom"
	"io"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal("Eror init db: ", err)
	}
	defer db.Close()

	router := gin.Default()
	router.GET("/api", func(c *gin.Context) {
		resp, err := http.Get("https://randomuser.me/api/")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch result data"})
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}

		var response jsoncustom.Response
		if err := json.Unmarshal(body, &response); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response"})
			return
		}

		for _, result := range response.Results {
			_, err := db.Exec(`
				INSERT INTO result (gender, title, first_name, last_name, street_number, street_name, city, state, 
				                    country, postcode, latitude, longitude, timezone_offset, timezone_description, email,
				                    username, password, salt, md5, sha1, sha256, dob, age, registered_date, 
				                    registered_age, phone, cell, national_id_name, national_id_value, picture_large_url, 
				                    picture_medium_url, picture_thumbnail_url, nationality)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21,
				        $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33) 
				    `, result.Gender, result.Name.Title, result.Name.First, result.Name.Last, result.Email, result.Login.UUID,
				result.Login.Username, result.Login.Password, result.Login.Salt, result.Login.MD5, result.Login.SHA1,
				result.Login.SHA256, result.DOB.Date, result.DOB.Age, result.Phone, result.Cell, result.ID.Name, result.ID.Value,
				result.Picture.Large, result.Picture.Medium, result.Picture.Thumbnail)

			if err != nil {
				log.Fatal("Failed to insert result data: ", err)
			}
		}

		var info jsoncustom.Info
		if err := json.Unmarshal(body, &response); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode info"})
			return
		}

		_, err = db.Exec(`
				INSERT INTO info (seed, result, page, version)
				VALUES ($1, $2, $3, $4)
				`, info.Seed, info.Result, info.Page, info.Version)
		if err != nil {
			log.Fatal("Failed to insert info data: ", err)
		}

		c.JSON(http.StatusOK, gin.H{"message": "Result data successfully retrieved and stored"})
	})

	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
