package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type products struct {
	Title string
	Name  string
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, getData())
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

func getData() []products {
	urlExample := "postgres://postgres:q1w2e3r4@localhost:5432/chinook?application_name=golang"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, _ := conn.Query(context.Background(), "select album.title, artist.name from album join artist on album.artist_id=artist.artist_id")
	products, _ := pgx.CollectRows(rows, pgx.RowToStructByName[products])

	return products
}
