package authroutes

import (
	"net/mail"

	"github.com/antonybholmes/go-auth"
	"github.com/antonybholmes/go-auth/userdb"
	"github.com/antonybholmes/go-edb-api/consts"
	"github.com/antonybholmes/go-edb-api/routes"
	"github.com/rs/zerolog/log"

	"github.com/labstack/echo/v4"
)

func EmailUpdatedResp(c echo.Context) error {
	return routes.MakeOkPrettyResp(c, "email updated")
}

// Start passwordless login by sending an email
func SendChangeEmailRoute(c echo.Context) error {
	return routes.NewValidator(c).ParseLoginRequestBody().LoadAuthUserFromToken().Success(func(validator *routes.Validator) error {
		authUser := validator.AuthUser
		req := validator.Req

		newEmail, err := mail.ParseAddress(req.Email)

		if err != nil {
			return routes.ErrorReq(err)
		}

		otpJwt, err := auth.ChangeEmailToken(c, authUser, newEmail, consts.JWT_PRIVATE_KEY)

		if err != nil {
			return routes.ErrorReq(err)
		}

		var file string

		if req.CallbackUrl != "" {
			file = "templates/email/email/change/web.html"
		} else {
			file = "templates/email/email/change/api.html"
		}

		go BaseSendEmailWithToken("Update Email",
			authUser,
			newEmail,
			file,
			otpJwt,
			req.CallbackUrl,
			req.Url)

		//if err != nil {
		//	return routes.ErrorReq(err)
		//}

		return routes.MakeOkPrettyResp(c, "check your email for a change email link")
	})
}

func UpdateEmailRoute(c echo.Context) error {
	return routes.NewValidator(c).ParseLoginRequestBody().LoadAuthUserFromToken().Success(func(validator *routes.Validator) error {

		if validator.Claims.Type != auth.TOKEN_TYPE_CHANGE_EMAIL {
			return routes.WrongTokentTypeReq()
		}

		err := auth.CheckOtpValid(validator.AuthUser, validator.Claims.Otp)

		if err != nil {
			return routes.ErrorReq(err)
		}

		authUser := validator.AuthUser
		uuid := authUser.Uuid

		log.Debug().Msgf("change email %s", validator.Req.Email)

		err = userdb.SetEmail(validator.AuthUser.Uuid, validator.Req.Email)

		if err != nil {
			return routes.ErrorReq(err)
		}

		authUser, err = userdb.FindUserByUuid(uuid)

		if err != nil {
			return routes.ErrorReq(err)
		}

		return SendEmailChangedEmail(c, authUser)
	})
}

func SendEmailChangedEmail(c echo.Context, authUser *auth.AuthUser) error {

	file := "templates/email/email/updated.html"

	go SendEmailWithToken("Email Address Changed",
		authUser,
		file,
		"",
		"",
		"")

	//if err != nil {
	//	return routes.ErrorReq(err)
	//}

	return EmailUpdatedResp(c)

}
