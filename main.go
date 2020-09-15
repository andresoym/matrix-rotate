package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"POST", "OPTION"}
	config.AllowHeaders = []string{"*"}

	r.Use(cors.New(config))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.POST("/", PostAction)

	//session := r.Group("/session")

	err := r.Run(":8080")
	if err != nil {
		log.Println("Fail starting service http: ", err)
	}

}

// PostAction Controller
func PostAction(c *gin.Context) {
	var err error
	var m Matriz

	err = c.Bind(&m)
	if err != nil {
		c.JSON(400, gin.H{"error": "Parámetro array debe ser enviado como una matriz", "ejemplo": m.Random()})
		return
	}

	n := len(m.Array)
	if n < 1 {
		c.JSON(400, gin.H{"error": "No es una matriz NxN", "ejemplo": m.Random()})
		return
	}

	for i := 0; i < n; i++ {
		if len(m.Array[i]) != n {
			c.JSON(400, gin.H{"error": "No es una matriz NxN", "ejemplo": m.Random()})
			return
		}
	}

	m.Rotate()

	c.JSON(200, m)
}

//Matriz struct
type Matriz struct {
	Array [][]int `json:"array" binding:"required"`
}

//Rotate function to rotate matriz
func (m *Matriz) Rotate() {
	n := len(m.Array)

	for x := 0; x < n/2; x++ {
		for y := x; y < n-x-1; y++ {
			temp := m.Array[x][y]
			m.Array[x][y] = m.Array[y][n-1-x]
			m.Array[y][n-1-x] = m.Array[n-1-x][n-1-y]
			m.Array[n-1-x][n-1-y] = m.Array[n-1-y][x]
			m.Array[n-1-y][x] = temp
		}
	}

}

//Random function to generate random matrix
func (m Matriz) Random() [][]int {
	var array [][]int

	rand.Seed(time.Now().Unix())

	x := rand.Intn(6)
	if x <= 1 {
		x = 3
	}
	for i := 0; i < x; i++ {
		array = append(array, rand.Perm(50)[:x])
	}
	return array
}
