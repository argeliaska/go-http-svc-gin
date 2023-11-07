package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

// Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
func main(){
	route := gin.Default()
	route.GET("/testing", getPerson)
	route.POST("/testingpost", postPerson)
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.Run(":8085") // listen and serve on 0.0.0.0:8080
}


func getPerson(c *gin.Context) {
	var person Person
	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	err := c.ShouldBind(&person)
	if  err == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
		c.String(200, "Success")
	} 
}

func postPerson(c *gin.Context){
	var person Person

	err := c.ShouldBindWith(&person, binding.Form)
	if err != nil {
		log.Printf("failed %v", err)
	}

	log.Println(person.Name)
	log.Println(person.Address)
	log.Println(person.Birthday)
	c.String(200, "Success")
}