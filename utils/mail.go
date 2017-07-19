package utils

import (
	"crypto/tls"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"mime"
	"net"
	"net/mail"
	"net/smtp"
	"time"
)

const (
	CONN_SECURITY_NONE     = ""
	CONN_SECURITY_PLAIN    = "PLAIN"
	CONN_SECURITY_TLS      = "TLS"
	CONN_SECURITY_STARTTLS = "STARTTLS"
)

func encodeRFC2047Word(s string) string {
	return mime.BEncoding.Encode("utf-8", s)
}

func connectToSMTPServer(server, port string, tls bool, tlsconfig *tls.Config) (net.Conn, *AppError) {
	var conn net.Conn
	var err error

	if tls {
		conn, err = tls.Dial("tcp", server+":"+port, tlsconfig)
		if err != nil {
			return nil, NewLocAppError("SendMail", "utils.mail.connect_smtp.open_tls.app_error", nil, err.Error())
		}
	} else {
		conn, err = net.Dial("tcp", server+":"+port)
		if err != nil {
			return nil, NewLocAppError("SendMail", "utils.mail.connect_smtp.open.app_error", nil, err.Error())
		}
	}

	return conn, nil
}

func newSMTPClient(conn net.Conn, username, password, server, port string, ctype bool, tlsconfig *tls.Config) (*smtp.Client, *AppError) {
	c, err := smtp.NewClient(conn, server+":"+port)
	if err != nil {
		l4g.Error(T("utils.mail.new_client.open.error"), err)
		return nil, NewLocAppError("SendMail", "utils.mail.connect_smtp.open_tls.app_error", nil, err.Error())
	}
	auth := smtp.PlainAuth("", username, password, server+":"+port)

	if ctype == CONN_SECURITY_TLS {
		if err = c.Auth(auth); err != nil {
			return nil, NewLocAppError("SendMail", "utils.mail.new_client.auth.app_error", nil, err.Error())
		}
	} else if ctype == CONN_SECURITY_STARTTLS {
		c.StartTLS(tlsconfig)
		if err = c.Auth(auth); err != nil {
			return nil, NewLocAppError("SendMail", "utils.mail.new_client.auth.app_error", nil, err.Error())
		}
	} else if ctype == CONN_SECURITY_PLAIN {
		// note: go library only supports PLAIN auth over non-tls connections
		if err = c.Auth(auth); err != nil {
			return nil, NewLocAppError("SendMail", "utils.mail.new_client.auth.app_error", nil, err.Error())
		}
	}
	return c, nil
}

func TestMailConnection(config *model.Config) {
	if !config.EmailSettings.SendEmailNotifications {
		return
	}

	conn, err1 := connectToSMTPServer(config)
	if err1 != nil {
		l4g.Error(T("utils.mail.test.configured.error"), T(err1.Message), err1.DetailedError)
		return
	}
	defer conn.Close()

	c, err2 := newSMTPClient(conn, config)
	if err2 != nil {
		l4g.Error(T("utils.mail.test.configured.error"), T(err2.Message), err2.DetailedError)
		return
	}
	defer c.Quit()
	defer c.Close()
}

func SendMail(to, subject, body string) *model.AppError {
	return SendMailUsingConfig(to, subject, body, Cfg)
}

func SendMailUsingConfig(to, subject, body string, config *model.Config) *model.AppError {
	if !config.EmailSettings.SendEmailNotifications || len(config.EmailSettings.SMTPServer) == 0 {
		return nil
	}

	l4g.Debug(T("utils.mail.send_mail.sending.debug"), to, subject)

	fromMail := mail.Address{Name: config.EmailSettings.FeedbackName, Address: config.EmailSettings.FeedbackEmail}
	toMail := mail.Address{Name: "", Address: to}

	headers := make(map[string]string)
	headers["From"] = fromMail.String()
	headers["To"] = toMail.String()
	headers["Subject"] = encodeRFC2047Word(subject)
	headers["MIME-version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"utf-8\""
	headers["Content-Transfer-Encoding"] = "8bit"
	headers["Date"] = time.Now().Format(time.RFC1123Z)

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n<html><body>" + body + "</body></html>"

	conn, err1 := connectToSMTPServer(config)
	if err1 != nil {
		return err1
	}
	defer conn.Close()

	c, err2 := newSMTPClient(conn, config)
	if err2 != nil {
		return err2
	}
	defer c.Quit()
	defer c.Close()

	if err := c.Mail(fromMail.Address); err != nil {
		return model.NewLocAppError("SendMail", "utils.mail.send_mail.from_address.app_error", nil, err.Error())
	}

	if err := c.Rcpt(toMail.Address); err != nil {
		return model.NewLocAppError("SendMail", "utils.mail.send_mail.to_address.app_error", nil, err.Error())
	}

	w, err := c.Data()
	if err != nil {
		return model.NewLocAppError("SendMail", "utils.mail.send_mail.msg_data.app_error", nil, err.Error())
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return model.NewLocAppError("SendMail", "utils.mail.send_mail.msg.app_error", nil, err.Error())
	}

	err = w.Close()
	if err != nil {
		return model.NewLocAppError("SendMail", "utils.mail.send_mail.close.app_error", nil, err.Error())
	}

	return nil
}
