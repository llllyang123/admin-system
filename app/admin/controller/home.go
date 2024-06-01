package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) (one string, two string) {
	one = c.Query("one")
	two = c.Query("two")
	fmt.Println("index: ", one, two)
	return one, two
}
