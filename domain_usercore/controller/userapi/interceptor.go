package userapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"zakariawahyu.com/go-gogen-mytodo/shared/model/payload"

	"github.com/gin-gonic/gin"
)

func (r *ginController) authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.GetHeader("Authorization")
		strArr := strings.Split(bearToken, " ")

		tokenInBytes, err := r.jwtToken.VerifyToken(strArr[1])
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		var dataToken payload.DataToken
		err = json.Unmarshal(tokenInBytes, &dataToken)
		log.Println(tokenInBytes)
		log.Println(err)
		log.Println(dataToken)
		if err != nil {
			log.Printf("error decoding sakura response: %v", err)
			if e, ok := err.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("sakura response: %q", tokenInBytes)
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Set("data", dataToken)

		c.AbortWithStatus(http.StatusForbidden)
		return
	}
}

func (r *ginController) authorization() gin.HandlerFunc {

	return func(c *gin.Context) {

		authorized := true

		if !authorized {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
	}
}
