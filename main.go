package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("db/ping", handler)
	r.Run()
}

func handler(c *gin.Context) {
	res, err := pingDB2()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": res})
}

func pingDB2() (string, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        "localhost", 5432, "postgres", "secret", "postgres")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		println("Could not open the connection")
		return "", err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return "", err
	}

	var str string
	err = db.QueryRow("SELECT NOW()").Scan(&str)
	if err != nil {
		return "", err
	}

	return str, nil
}
