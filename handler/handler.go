package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rmn-kmr/go-url-shortener/shortner"
	"github.com/rmn-kmr/go-url-shortener/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var createRequest UrlCreationRequest
	// c.GetRawData()

	// bodyAsByteArray, _ := ioutil.ReadAll(c.Request.Body)
	// jsonBody := string(bodyAsByteArray)
	// fmt.Println(jsonBody, "MMM")

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	shortUrl := shortner.GenerateShortLink(createRequest.LongUrl, createRequest.UserId)
	store.SaveUrlMapping(shortUrl, createRequest.LongUrl, createRequest.UserId)

	host := "http://localhost:9988/"
	c.JSON(200, gin.H{
		"meesage":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveUrlMapping(shortUrl)
	c.Redirect(302, initialUrl)
}
