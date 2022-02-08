package mail

import (
	"fmt"
	"net/smtp"
	"testing"
	// email "github.com/jordan-wright/email"
)

type FakeEmailSender struct {
	Text []byte
}

func (s *FakeEmailSender) Send(from, subject, text, mailserver string, to, bcc, cc []string, a smtp.Auth) error {
	s.Text = []byte(text)
	return nil
}

func TestSendEmailWithDisclaimer2(t *testing.T) {
	from := "jxs1211@gmail.com"
	to := []string{from}
	subject := "go email test subject"
	text := "go email test body"
	mailserver := "smtp.gmail.com:578"
	a := smtp.PlainAuth("", from, "Jxs93503", "smtp.gmail.com")

	s := &FakeEmailSender{}
	err := SendEmailWithDisclaimer2(s, subject, from, text, mailserver, to, []string{}, []string{}, a)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	want := fmt.Sprintf("%s \n\n %s", text, DISCLAIMER)
	got := string(s.Text)
	if got != want {
		t.Fatalf("want: %s, got: %s\n", want, got)
	}
}
