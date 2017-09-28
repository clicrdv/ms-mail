package mail

import (
	"fmt"
	"os"
	"strconv"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendgridMail struct {
	FromName     string
	FromEmail    string
	ReplyToName  string
	ReplyToEmail string
	ToMap        map[string]string
	Subject      string
	HtmlContent  string
	TextContent  string
	UUID         string
	GroupID      string
}

func (sm *SendgridMail) NewMail(subject string, fromName string, fromEmail string, toMap map[string]string, htmlContent string, textContent string, groupId string) {
	sm.FromName = fromName
	sm.FromEmail = fromEmail
	sm.ToMap = toMap
	sm.HtmlContent = htmlContent
	sm.TextContent = textContent
	sm.Subject = subject
	sm.GroupID = groupId
}

func (sm *SendgridMail) SendMail() (string, string) {
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = mail.GetRequestBody(sm.BuildMail())
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return strconv.Itoa(response.StatusCode), sm.UUID
}

func (sm *SendgridMail) BuildMail() *mail.SGMailV3 {
	m := mail.NewV3Mail()
	e := mail.NewEmail(sm.FromName, sm.FromEmail)
	m.SetFrom(e)

	m.Subject = sm.Subject

	p := mail.NewPersonalization()
	tos := make([]*mail.Email, len(sm.ToMap))
	idx := 0
	for name, email := range sm.ToMap {
		tos[idx] = mail.NewEmail(name, email)
		idx++
	}
	p.AddTos(tos...)

	// p.Subject = "Hello World from the Personalized SendGrid Go Library"

	p.SetCustomArg("clicrdvid", sm.UUID)
	m.AddPersonalizations(p)
	m.AddCategories("MS-MAIL")

	c := mail.NewContent("text/plain", sm.TextContent)
	m.AddContent(c)

	c = mail.NewContent("text/html", sm.HtmlContent)
	m.AddContent(c)

	replyToEmail := mail.NewEmail(sm.ReplyToName, sm.ReplyToEmail)
	m.SetReplyTo(replyToEmail)

	return m
}
