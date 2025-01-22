package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// represents pets data
type pet struct{
    ID  string   `json:"id"`
    Name string  `json:"name"`
    Image string `json:"image"`
    Votes int    `json:"votes"`
}

var pets = []pet{
    {ID: "1", Name: "Mr. Cute", Image: "https://thumbs.gfycat.com/EquatorialIckyCat-max-1mb.gif", Votes: 1},
    {ID: "2", Name: "Mx. Monkey", Image: "https://thumbs.gfycat.com/FatalInnocentAmericanshorthair-max-1mb.gif", Votes: 3},
    {ID: "3", Name: "Ms. Zebra", Image: "https://media2.giphy.com/media/20G9uNqE3K4dRjCppA/source.gif", Votes:  9},
    {ID: "4", Name: "Dr.Lion", Image: "http://bestanimations.com/Animals/Mammals/Cats/Lions/animated-lion-gif-11.gif", Votes: 7},
    {ID: "5", Name: "Mr. Panda", Image: "https://media.giphy.com/media/ALalVMOVR8Qw/giphy.gif", Votes: 20},
}

func main(){
    router := gin.Default()
    router.GET("/pets", getPets)

    router.Run("localhost:7070")
}

// creating a get funtion to obtain pets
func getPets(c *gin.Context){
    c.IndentedJSON(http.StatusOK, pets)
}



