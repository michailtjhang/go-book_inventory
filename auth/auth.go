package auth

import (
	"book_inventory/models"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

func HomeHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/login")
}

func LoginGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"content": "",
	})
}

func LoginPostHandler(c *gin.Context) {
	var credentials models.Login
	err := c.Bind(&credentials)
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": "USERNAME OR PASSWORD INVALID REQUEST",
		})
	}

	if credentials.Username != models.USER || credentials.Password != models.PASS {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"content": "USERNAME OR PASSWORD INVALID",
		})
	} else {
		// Token

		claim := jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
			Issuer:    "Book Inventory",
			IssuedAt:  time.Now().Unix(),
		}

		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		token, err := sign.SignedString([]byte(models.SECRET))
		if err != nil {
			c.HTML(http.StatusInternalServerError, "login.html", gin.H{
				"content": "USERNAME OR PASSWORD INVALID REQUEST",
			})
			c.Abort()
		}

		q := url.Values{}
		q.Add("token", token)
		location := url.URL{Path: "/books", RawQuery: q.Encode()}
		c.Redirect(http.StatusMovedPermanently, location.String())
	}

}
