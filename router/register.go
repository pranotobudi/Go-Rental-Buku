package router

import (
	"fmt"
	"net/http"
	"net/smtp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pranotobudi/Go-Rental-Buku/database"
	"github.com/thanhpk/randstr"
)

func UserRegister(c *gin.Context) {
	name := c.PostForm("name")
	address := c.PostForm("address")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role, _ := strconv.Atoi(c.PostForm("Role"))

	var user = database.User{
		Name:              name,
		Address:           address,
		Photo:             "",
		Email:             email,
		Email_verified_at: time.Now(),
		Password:          password,
		Role:              role,
	}

	// fmt.Println("USER: ", user)
	c.JSON(http.StatusOK, user)

	// db := c.Value("db")
	if database.IsValidUserData(user) {
		//Send Confirmation Email
		random := randstr.Hex(16) // generate 128-bit hex string
		err := database.AddRegistration(random)

		msg := []byte("To: oceankingdigital@gmail.com\r\n" +
			"Subject: Confirmation Email from Rental Buku!\r\n" +
			"\r\n" +
			"This is the email body.\r\n" +
			"http://localhost:8080/register/confirmation?email=" + user.Email + "&token=" + random)

		toEmail := []string{"oceankingdigital@gmail.com"}
		SendEmail(toEmail, msg)

		err = database.UpdateUserByObject(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, user)
		}
	}
}

func UserRegisterJSON(c *gin.Context) {
	var user database.User
	c.BindJSON(&user)

	// fmt.Println("USER: ", user)
	c.JSON(http.StatusOK, user)
}

const (
	fromAddress       = "pranotobudi.app@gmail.com"
	fromEmailPassword = "CamelCasePasswordBismillah"
	smtpServer        = "smtp.gmail.com"
	smtpPort          = "587"
)

func SendEmail(toAddress []string, message []byte) {
	// Message.
	//   message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", fromAddress, fromEmailPassword, smtpServer)

	// Sending email.
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, fromAddress, toAddress, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
