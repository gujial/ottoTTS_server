package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gujial/ottoTTS"
	"net/http"
)

type MessageRequest struct {
	Message string `json:"message" binding:"required"`
}

func main() {
	// 初始化 ottoTTS
	ottoTTS.InitializeTTS()

	r := gin.Default()

	r.POST("/speak", func(c *gin.Context) {
		var req MessageRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		wavData, err := ottoTTS.Speech(req.Message)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate speech"})
			return
		}

		c.DataFromReader(http.StatusOK, int64(len(wavData)), "audio/wav",
			http.NoBody, map[string]string{
				"Content-Disposition": `attachment; filename="speech.wav"`,
			})
		c.Data(http.StatusOK, "audio/wav", wavData)
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
