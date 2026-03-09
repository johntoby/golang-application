package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var movies = []string{
	"The Shawshank Redemption",
	"The Godfather",
	"The Dark Knight",
	"Pulp Fiction",
	"Forrest Gump",
	"Inception",
	"The Matrix",
	"Interstellar",
	"Gladiator",
	"Titanic",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		movie := movies[rand.Intn(len(movies))]
		c.HTML(http.StatusOK, "index.html", gin.H{
			"movie": movie,
		})
	})
	
	r.LoadHTMLGlob("templates/*")
	r.Run(":8080")
}
