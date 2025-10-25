package main

import (
	"fmt"

	"github.com/nih13/dbeedb/pkg/core"
)

// func main() {
// 	r := gin.Default()

// 	// Define a route for the root path
// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Hello, Gin!")

// 	})

// 	// Start the server
// 	r.Run(":8080")
// }

func main() {
	fruits := core.New()
	fruits.Set("a", "apple")
	fruits.Set("b", "mango")
	fmt.Println(fruits.Get("b"))
	fruits.Print()
	fruits.Delete("b")
	fruits.Print()

}
