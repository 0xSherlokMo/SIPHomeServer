package main

import (
	"io"
	"log"
	"net/http"

	"github.com/0xSherlokMo/SIPHomeServer/model"
	"github.com/gin-gonic/gin"
	"github.com/twilio/twilio-go/twiml"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	r.POST("/incoming", func(c *gin.Context) {
		var request model.IncomingCall
		err := c.Bind(&request)
		if err != nil {
			log.Printf("Error binding request: %s", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		body, _ := io.ReadAll(c.Request.Body)
		log.Printf("Received request: %s", body)

		doc, el := twiml.CreateDocument()
		el.CreateElement("Reject").CreateAttr("reason", "busy")
		s, _ := doc.WriteToString()
		c.Header("Content-Type", "application/xml")
		c.String(http.StatusOK, s)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}
}
