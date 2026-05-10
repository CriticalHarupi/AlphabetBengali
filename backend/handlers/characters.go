package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"alphabetbengali/data"
)

func GetCharacters(c *gin.Context) {
	c.JSON(http.StatusOK, data.Characters)
	// test
}

func GetCharacterByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	for _, ch := range data.Characters {
		if ch.ID == id {
			c.JSON(http.StatusOK, ch)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "character not found"})
}
