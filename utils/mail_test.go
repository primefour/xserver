package utils

import (
	"net/mail"
	"testing"
)

func TestSendMail(t *testing.T) {

	var emailSubject string = "Testing this email"
	var emailBody string = "This is a test from autobot"

	client, cerr := GetMailClient("primefour", "four.c..", "123.125.50.138", "25", CONN_SECURITY_PLAIN, nil)
	if cerr != nil {
		t.Error("get client fail ")
		return
	}

	defer func() {
		client.Quit()
		client.Close()
	}()

	fromMail := mail.Address{Name: "lihaihui", Address: "primefour@163.com"}
	toMail := mail.Address{Name: "lihaihui", Address: "lihaihui@zenmen.com"}

	if err := SendMail(client, fromMail, toMail, emailSubject, emailBody); err != nil {
		t.Log(err)
		t.Fatal("Should connect to the STMP Server")
	}

	/*
		else {
			//Check if the email was send to the rigth email address
			var resultsMailbox JSONMessageHeaderInbucket
			err := RetryInbucket(5, func() error {
				var err error
				resultsMailbox, err = GetMailBox(emailTo)
				return err
			})
			if err != nil {
				t.Log(err)
				t.Log("No email was received, maybe due load on the server. Disabling this verification")
			}
			if err == nil && len(resultsMailbox) > 0 {
				if !strings.ContainsAny(resultsMailbox[0].To[0], emailTo) {
					t.Fatal("Wrong To recipient")
				} else {
					if resultsEmail, err := GetMessageFromMailbox(emailTo, resultsMailbox[0].ID); err == nil {
						if !strings.Contains(resultsEmail.Body.Text, emailBody) {
							t.Log(resultsEmail.Body.Text)
							t.Fatal("Received message")
						}
					}
				}
			}
		}
	*/
}
