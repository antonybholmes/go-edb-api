package auth

import (
	"bytes"
	"fmt"
	"html/template"
	"net/url"
	"strings"

	"github.com/antonybholmes/go-auth"
	"github.com/antonybholmes/go-edb-api/consts"
	"github.com/antonybholmes/go-edb-api/routes"
	"github.com/antonybholmes/go-mailer/email"
)

const DO_NOT_REPLY = "Please do not reply to this message. It was sent from a notification-only email address that we don't monitor."

type EmailBody struct {
	Name       string
	From       string
	Time       string
	Link       string
	DoNotReply string
}

// Generic method for sending an email with a token in it. For APIS this is a token to use in the request, for websites
// it can craft a callback url with the token added as a parameter so that the web app can deal with the response.
func SendEmailWithToken(subject string,
	authUser *auth.AuthUser,
	file string,
	token string,
	callbackUrl string,
	vistUrl string) error {

	var body bytes.Buffer

	t, err := template.ParseFiles(file)

	if err != nil {
		return routes.ErrorReq(err)
	}

	var firstName string = ""

	if len(authUser.Name) > 0 {
		firstName = authUser.Name
	} else {
		firstName = authUser.Email.Address
	}

	firstName = strings.Split(firstName, " ")[0]

	time := fmt.Sprintf("%d minutes", int(auth.TOKEN_TYPE_OTP_TTL_MINS.Minutes()))

	if callbackUrl != "" {
		callbackUrl, err := url.Parse(callbackUrl)

		if err != nil {
			return routes.ErrorReq(err)
		}

		params, err := url.ParseQuery(callbackUrl.RawQuery)

		if err != nil {
			return routes.ErrorReq(err)
		}

		if vistUrl != "" {
			params.Set("url", vistUrl)
		}

		params.Set("otp", token)

		callbackUrl.RawQuery = params.Encode()

		link := callbackUrl.String()

		err = t.Execute(&body, EmailBody{
			Name:       firstName,
			Link:       link,
			From:       consts.NAME,
			Time:       time,
			DoNotReply: DO_NOT_REPLY,
		})

		if err != nil {
			return routes.ErrorReq(err)
		}
	} else {
		err = t.Execute(&body, EmailBody{
			Name:       firstName,
			Link:       token,
			From:       consts.NAME,
			Time:       time,
			DoNotReply: DO_NOT_REPLY,
		})

		if err != nil {
			return routes.ErrorReq(err)
		}
	}

	err = email.SendHtmlEmail(authUser.Email, subject, body.String())

	if err != nil {
		return err
	}

	return nil
}
