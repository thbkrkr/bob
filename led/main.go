package main

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/s", "./views/")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/s")
	})

	r.GET("/ping", ping)
	r.GET("/led/:mode/:led", led)

	r.Run(":8080")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func led(c *gin.Context) {
	mode := c.Param("mode")
	led := c.Param("led")

	err := bobled(mode, led)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, gin.H{
		"mode": mode,
		"led":  led,
	})
}

func bobled(mode string, led string) error {
	script := "scripts/bob-led.py"
	cmd := exec.Command(script, mode, led)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
