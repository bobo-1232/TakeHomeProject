package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
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
	Age         int    `json:"age"`
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
	SELECT p.name, ph.number, a.city, a.state, a.street1, IFNULL(a.street2, ''), a.zip_code, p.age
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
		&personsInfo.Age,
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

	result, err := db.Exec("INSERT INTO person(name, age) VALUES(?, ?)", newPerson.Name, newPerson.Age)
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

func getPersonsPaginated(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")

	pageInt, err1 := strconv.Atoi(page)
	limitInt, err2 := strconv.Atoi(limit)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	offset := (pageInt - 1) * limitInt

	rows, err := db.Query("SELECT id, name, age FROM person LIMIT ? OFFSET ?", limitInt, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var persons []Person
	for rows.Next() {
		var p Person
		if err := rows.Scan(&p.ID, &p.Name, &p.Age); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		persons = append(persons, p)
	}

	var total int
	err = db.QueryRow("SELECT COUNT(*) FROM person").Scan(&total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	totalPages := (total + limitInt - 1) / limitInt

	c.JSON(http.StatusOK, gin.H{
		"data":     persons,
		"has_next": pageInt < totalPages,
	})
}

func main() {
	connectToDB()
	defer db.Close()

	r := gin.Default()
	r.Use(cors.Default()) // Allow React to access Go APIs

	r.GET("/person/:person_id/info", getPerson)
	r.POST("/person/create", createNewPerson)
	r.GET("/persons", getPersonsPaginated)

	r.Run(":8080")
}
