package auth

import (
	"github.com/antonybholmes/go-auth"
	"github.com/antonybholmes/go-auth/userdb"
	"github.com/antonybholmes/go-edb-api/consts"
	"github.com/antonybholmes/go-edb-api/routes"

	"github.com/labstack/echo/v4"
)

// Start passwordless login by sending an email
func ResetPasswordFromUsernameRoute(c echo.Context) error {
	return routes.NewValidator(c).AuthUserFromUsername().VerifiedEmail().Success(func(validator *routes.Validator) error {
		authUser := validator.AuthUser
		req := validator.Req

		otpJwt, err := auth.ResetPasswordToken(c, authUser, consts.JWT_SECRET)

		if err != nil {
			return routes.ErrorReq(err)
		}

		var file string

		if req.CallbackUrl != "" {
			file = "templates/email/password/reset/web.html"
		} else {
			file = "templates/email/password/reset/api.html"
		}

		err = SendEmailWithToken("Password Reset",
			authUser,
			file,
			otpJwt,
			req.CallbackUrl,
			req.Url)

		if err != nil {
			return routes.ErrorReq(err)
		}

		return routes.MakeOkResp(c, "password reset email sent")
	})
}

func UpdatePasswordRoute(c echo.Context) error {
	return routes.NewValidator(c).ReqBind().AuthUserFromUuid().Success(func(validator *routes.Validator) error {

		if validator.Claims.Type != auth.TOKEN_TYPE_RESET_PASSWORD {
			return routes.WrongTokentTypeReq()
		}

		err := userdb.SetPassword(validator.AuthUser.Uuid, validator.Req.Password)

		if err != nil {
			return routes.ErrorReq(err)
		}

		return SendPasswordEmail(c, validator.AuthUser, validator.Req.Password)
	})
}

func SendPasswordEmail(c echo.Context, authUser *auth.AuthUser, password string) error {

	var file string

	if password == "" {
		file = "templates/email/password/switch-to-passwordless.html"
	} else {
		file = "templates/email/password/updated.html"
	}

	err := SendEmailWithToken("Password Updated",
		authUser,
		file,
		"",
		"",
		"")

	if err != nil {
		return routes.ErrorReq(err)
	}

	return routes.PasswordUpdatedResp(c)

}
