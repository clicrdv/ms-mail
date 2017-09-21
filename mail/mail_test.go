package mail

import (
	"testing"
)

var (
	targetMailMap = map[string]string{
		"mail perso": "mikrob@yopmail.com",
		"mail pro":   "mikrob+3@yopmail.com",
	}
	sm = SendgridMail{
		FromName:     "No Reply ClicRDV",
		FromEmail:    "noreply@clicrdv.com",
		ReplyToName:  "No Reply ClicRDV",
		ReplyToEmail: "noreply@clicrdv.com",
		HtmlContent:  "<html><body><b>This is bold html</b></body></html>",
		TextContent:  "This is text content",
		Subject:      "Mail From MS Mail",
		ToMap:        targetMailMap,
	}
)

func TestBuildMail(t *testing.T) {

	mail := sm.BuildMail()
	if mail.From.Address != sm.FromEmail {
		t.Errorf("Expected from email to be %s but found %s", sm.FromEmail, mail.From.Address)
	}

	for idx, p := range mail.Personalizations {
		if _, ok := targetMailMap[p.To[idx].Name]; !ok {
			t.Errorf("Expected to find [%s] in To (recipient) of mail personnalization but didn't found", p.To[idx].Name)
		}
		email := targetMailMap[p.To[idx].Name]
		if email != p.To[idx].Address {
			t.Errorf("Expected to find [%s] in To (recipient) of mail personnalization but didn't found", p.To[idx].Address)
		}

	}
}
