package main

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/gujial/ottoTTS"
	"log"
	"net/http"
	"strconv"
)

type MessageRequest struct {
	Message string `json:"message" binding:"required"`
}

type config struct {
	Debug bool `toml:"Debug"`
	Port  int  `toml:"Port"`
}

var cfg config

func loadConfig() {
	_, err := toml.DecodeFile("./config.toml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 初始化 ottoTTS
	ottoTTS.InitializeTTS()
	loadConfig()
	var mode string
	if cfg.Debug {
		mode = gin.DebugMode
	} else {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	r := gin.Default()

	r.POST("/speak", func(c *gin.Context) {
		var req MessageRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		wavData, err := ottoTTS.Speech(req.Message)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate speech"})
			return
		}

		c.DataFromReader(http.StatusOK, int64(len(wavData)), "audio/wav",
			http.NoBody, map[string]string{
				"Content-Disposition": `attachment; filename="speech.wav"`,
			})
		c.Data(http.StatusOK, "audio/wav", wavData)
	})

	err := r.Run(":" + strconv.Itoa(cfg.Port))
	if err != nil {
		return
	}
}
