package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   int    `json:"id"`
}

type pInfo struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	ZipCode     string `json:"zip_code"`
}

var db *sql.DB

func connectToDB() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/golangdb")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Database connected correctly")
}

func getPerson(c *gin.Context) {
	var personsInfo pInfo
	personID := c.Param("person_id")

	query := `
	SELECT p.name, ph.number, a.city, a.state, a.street1, IFNULL(a.street2, ''), a.zip_code
	FROM person p
	JOIN phone ph ON ph.person_id = p.id
	JOIN address_join aj ON aj.person_id = p.id
	JOIN address a ON aj.address_id = a.id
	WHERE p.id = ?`

	err := db.QueryRow(query, personID).Scan(
		&personsInfo.Name,
		&personsInfo.PhoneNumber,
		&personsInfo.City,
		&personsInfo.State,
		&personsInfo.Street1,
		&personsInfo.Street2,
		&personsInfo.ZipCode,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, personsInfo)
}

func createNewPerson(c *gin.Context) {
	var newPerson pInfo
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		return
	}

	result, err := db.Exec("INSERT INTO person(name, age) VALUES(?, ?)", newPerson.Name, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message error": err.Error()})
		return
	}
	personID, _ := result.LastInsertId()

	_, err = db.Exec("INSERT INTO phone(number, person_id) VALUES(?, ?)", newPerson.PhoneNumber, personID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message error": err.Error()})
		return
	}

	result, err = db.Exec("INSERT INTO address(city, state, street1, street2, zip_code) VALUES(?, ?, ?, ?, ?)",
		newPerson.City, newPerson.State, newPerson.Street1, newPerson.Street2, newPerson.ZipCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message error": err.Error()})
		return
	}
	addressID, _ := result.LastInsertId()

	_, err = db.Exec("INSERT INTO address_join(person_id, address_id) VALUES(?, ?)", personID, addressID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newPerson)
}

func main() {
	connectToDB()
	defer db.Close()

	r := gin.Default()

	r.GET("/person/:person_id/info", getPerson)
	r.POST("/person/create", createNewPerson)

	r.Run(":8080")
}
