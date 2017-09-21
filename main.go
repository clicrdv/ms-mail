package main

import (
	"github.com/clicrdv/ms-mail/mail"
)

func main() {
	targetMailMap := map[string]string{
		"mikrob - perso": "mikrob@yopmail.com",
		"mikrob - pro":   "mikrob+3@yopmail.com",
	}

	sm := mail.SendgridMail{
		FromName:     "No Reply ClicRDV",
		FromEmail:    "noreply@clicrdv.com",
		ReplyToName:  "No Reply ClicRDV",
		ReplyToEmail: "noreply@clicrdv.com",
		HtmlContent:  "<html><body><b>This is bold html</b></body></html>",
		TextContent:  "This is text content",
		Subject:      "Mail From MS Mail",
		ToMap:        targetMailMap,
	}
	sm.SendMail()
}
