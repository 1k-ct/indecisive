package api

import (
	"net/http"

	"github.com/1k-ct/indecisive/pkg/random"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "ping"})
}
func Introduction(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "you can serect some choicies."})
}

// Alteratives 選択肢
type Alteratives struct {
	Items []string `json:"items"`
}

func SelectOne(c *gin.Context) {
	a := Alteratives{}
	if err := c.BindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	r := random.NewCrypto()
	res := r.CryptoRandChooseList(a.Items)
	c.JSON(http.StatusOK, gin.H{"you have to do ": res})
}
