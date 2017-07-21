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

func connectToSMTPServer(server, port string, istls bool, tlsconfig *tls.Config) (net.Conn, *AppError) {
	var conn net.Conn
	var err error

	if istls {
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

func newSMTPClient(conn net.Conn, username, password, server, port, ctype string, tlsconfig *tls.Config) (*smtp.Client, *AppError) {
	c, err := smtp.NewClient(conn, server+":"+port)
	if err != nil {
		l4g.Error("utils.mail.new_client.open.error", err)
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

func TestMailConnection(username, password, server, port, ctype string, tlsconfig *tls.Config) {
	istls := true
	if ctype == CONN_SECURITY_PLAIN && tlsconfig == nil {
		istls = false
	}
	conn, err1 := connectToSMTPServer(server, port, istls, tlsconfig)
	if err1 != nil {
		l4g.Error("utils.mail.test.configured.error", err1.Message, err1.DetailedError)
		return
	}
	c, err2 := newSMTPClient(conn, username, password, server, port, ctype, tlsconfig)

	if err2 != nil {
		l4g.Error("utils.mail.test.configured.error", err2.Message, err2.DetailedError)
		return
	}
	defer c.Quit()
	defer c.Close()
}

func GetMailClient(username, password, server, port, ctype string, tlsconfig *tls.Config) (*smtp.Client, *AppError) {
	tls := true
	if ctype == CONN_SECURITY_PLAIN && tlsconfig == nil {
		tls = false
	}
	fmt.Printf("dial to server \n")
	conn, err1 := connectToSMTPServer(server, port, tls, tlsconfig)
	fmt.Printf("dial to server completely \n")
	if err1 != nil {
		l4g.Error("utils.mail.test.configured.error", err1.Message, err1.DetailedError)
		return nil, err1
	}

	fmt.Printf("sign on server \n")
	c, err2 := newSMTPClient(conn, username, password, server, port, ctype, tlsconfig)
	fmt.Printf("sign on server completely \n")

	if err2 != nil {
		l4g.Error("utils.mail.test.configured.error", err2.Message, err2.DetailedError)
		return nil, err1
	}
	return c, nil
}

func SendMail(c *smtp.Client, from, to mail.Address, subject, body string) *AppError {
	l4g.Debug("utils.mail.send_mail.sending.debug", to, subject)

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
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

	if err := c.Mail(from.Address); err != nil {
		return NewLocAppError("SendMail", "utils.mail.send_mail.from_address.app_error", nil, err.Error())
	}

	if err := c.Rcpt(to.Address); err != nil {
		return NewLocAppError("SendMail", "utils.mail.send_mail.to_address.app_error", nil, err.Error())
	}

	w, err := c.Data()
	if err != nil {
		return NewLocAppError("SendMail", "utils.mail.send_mail.msg_data.app_error", nil, err.Error())
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return NewLocAppError("SendMail", "utils.mail.send_mail.msg.app_error", nil, err.Error())
	}

	err = w.Close()
	if err != nil {
		return NewLocAppError("SendMail", "utils.mail.send_mail.close.app_error", nil, err.Error())
	}
	return nil
}
