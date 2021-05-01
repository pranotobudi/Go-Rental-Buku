package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
)

func ResetPassword(c *gin.Context) {
	email := c.PostForm("email")

	msg := []byte("To: " + email + "\r\n" +
		"Subject: Reset Password - Rental Buku!\r\n" +
		"\r\n" +
		"This is the email body.\r\n" +
		"http://localhost:8080/resetpassword/submitpage?email=" + email)

	toEmail := []string{email}
	SendEmail(toEmail, msg)

}

func ResetPasswordSubmitGet(c *gin.Context) {
	query := c.Request.URL.Query()
	email := query["email"]
	c.JSON(http.StatusOK, gin.H{"email": email[0]})

}

func ResetPasswordSubmitPost(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	passwordConfirm := c.PostForm("passwordConfirm")
	user, _ := database.GetUserByEmail(email)
	if password == passwordConfirm {
		user.Password = password
		database.UpdateUserByObject(user)
	}
}
