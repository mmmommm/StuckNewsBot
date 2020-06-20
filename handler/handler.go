package handler

import (
	"github.com/gin-gonic/gin"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go-send-contact/domain"
	"os"
)

func SendMail(c *gin.Context) {
	var m domain.mail
	err := c.BindJSON(&m)
	if err != nil {
		c.JSON(http.InternalServerError, err.Error())
	}
	from := mail.NewEmail(m.Name, m.Email)
	subject := m.Subject
	to := mail.NewEmail("ryota", "omotenashikyoto2020@gmail.com")
	plainTextContent := m.Text
	htmlContent := m.Text

	message := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		c.JSON(response.StatusCode, err.Error())
	}
	c.JSON(response.StatusCode, "success to send mail !!")
}
