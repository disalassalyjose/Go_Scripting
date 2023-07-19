package main

/* https://go.dev/doc/tutorial/web-service-gin -> With the Gin package you can build a web server in go */
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type student struct {
	Name      string `json:"name"`
	Rolnumber int    `json:"rolnumber"`
	Class     int    `json:"class"`
}

var students = []student{
	{Name: "Arun", Rolnumber: 10, Class: 8},
	{Name: "Sanjsy", Rolnumber: 11, Class: 7},
}

func main() {
	router := gin.Default()
	router.GET("/students", getStudents)
	router.GET("/students/:name", getIndividualdetails)
	router.POST("/students", postStudents)

	router.Run("localhost:9999")
}

// getstudent
func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

// poststudent
func postStudents(c *gin.Context) {
	var newStudents student
	if err := c.BindJSON(&newStudents); err != nil {
		return
	}

	students = append(students, newStudents)
	c.IndentedJSON(http.StatusCreated, newStudents)
}

func getIndividualdetails(c *gin.Context) {
	name := c.Param("name")

	for _, a := range students {
		if a.Name == name {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student data not found"})
}
