package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"alphabetbengali/data"
)

func GetConjuncts(c *gin.Context) {
	c.JSON(http.StatusOK, data.Conjuncts)
}

func GetConjunctByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	for _, cj := range data.Conjuncts {
		if cj.ID == id {
			c.JSON(http.StatusOK, cj)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "conjunct not found"})
}
