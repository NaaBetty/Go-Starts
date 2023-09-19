package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// movie shop
type movie struct{
	ID string	`json:"id"`
	Title string `json:"title"`
	Maker string  `json:"maker"`
	Category string `json:"category"`
	Copies  int	`json:"compies"`
}

var movies = []movie{
	{ID: "1", Title: "Aki and Pawpaw", Maker: "Marcel Proust", Category: "Comedy", Copies: 7},
	{ID: "2", Title: "Mr. Ibu", Maker: "F. Scott Fitzgerald", Category: "Drama" ,Copies: 5},
	{ID: "3", Title: "Osofia", Maker: "Leo Tolstoy", Category: "Horror" ,Copies: 6},
	{ID: "4", Title: "Kyeiwaaa", Maker: "Leo Tolstoy", Category: "Horror" ,Copies: 4},
}


//api request for movie shop
func getMovies(c *gin.Context){
	c.IndentedJSON(http.StatusOK, movies)
}

func movieById(c *gin.Context){
	id := c.Param("id")
	movie, err := getMovieById(id)

	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not in collection"})
		return
	}
	c.IndentedJSON(http.StatusOK, movie)
}

func checkoutMovie(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	movie, err := getMovieById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie not found."})
		return
	}

	if movie.Copies <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Movie not available."})
		return
	}

	movie.Copies -= 1
	c.IndentedJSON(http.StatusOK, movie)
}

func returnMovie(c *gin.Context){
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	movie, err := getMovieById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Movie not found."})
		return
	}
	movie.Copies += 1
	c.IndentedJSON(http.StatusOK, movie)

}

func getMovieById(id string) (*movie, error){
	for i, m := range movies{
		if m.ID == id {
			return &movies[i], nil
		}
	}
	return nil, errors.New("movie not in collection")
}

func addMovie(c *gin.Context){
	var newMovie movie
	if err := c.BindJSON(&newMovie); err != nil{
		return
	}
	movies = append(movies, newMovie)
	
	c.IndentedJSON(http.StatusAccepted, newMovie)

}
//end of movie shop



func main()  {
	router := gin.Default()
	router.GET("/movies", getMovies)
	router.GET("/movies/:id", movieById)
	router.POST("/movies", addMovie)
	router.PATCH("/checkout/:id", checkoutMovie)
	router.PUT("/return", returnMovie)
	router.Run("localhost:8080")
	fmt.Println("naa")
	
}

