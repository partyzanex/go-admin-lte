package goadmin

import (
	"encoding/hex"
	"github.com/labstack/echo/v4"
	"github.com/partyzanex/go-admin-bootstrap/encrypt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

func Login(ctx *AdminContext) error {
	_, err := ctx.Cookie(AccessCookieName)
	if err == nil {
		return ctx.Redirect(http.StatusFound, ctx.URL("/"))
	}

	u := ctx.User()
	if u != nil {
		return ctx.Redirect(http.StatusFound, ctx.URL("/"))
	}

	data := &Data{
		Title: "Login",
	}
	data.Breadcrumbs.Add("Login", ctx.URL(LoginURL))

	if ctx.Request().Method == http.MethodPost {
		result, err := auth(ctx)
		if err != nil && err != ErrUserNotFound && err != ErrWrongPassword {
			return err
		}
		if err == nil {
			return ctx.Redirect(http.StatusFound, ctx.URL(DashboardURL))
		} else {
			data.Set("err", err.Error())
		}

		data.Set("login", result.Login)
		data.Set("password", result.Password)
	}

	return ctx.Render(http.StatusOK, "auth/login", data)
}

func Logout(ctx *AdminContext) error {
	if user := ctx.User(); user != nil {
		ctx.SetCookie(&http.Cookie{
			Name:    AccessCookieName,
			Expires: time.Now().Add(-24 * time.Hour),
			Domain:  ctx.Request().Host,
			Path:    "/",
		})
	}

	return ctx.Redirect(http.StatusFound, ctx.URL(LoginURL))
}

func auth(ctx *AdminContext) (result User, err error) {

	login := ctx.FormValue("login")
	password := ctx.FormValue("password")

	result.Login = login
	result.Password = password

	user, err := ctx.UserCase().SearchByLogin(ctx.Ctx(), login)
	if err != nil {
		return result, err
	}

	ok, err := ctx.UserCase().ComparePassword(user, password)
	log.Println(ok, err)
	if err != nil {
		return result, err
	}

	if !ok {
		return result, ErrWrongPassword
	}

	token, err := ctx.UserCase().CreateAuthToken(ctx.Ctx(), user)
	if err != nil {
		return result, errors.Wrap(err, "creating auth token failed")
	}

	err = ctx.UserCase().SetLastLogged(ctx.Ctx(), user)
	if err != nil {
		return result, errors.Wrap(err, "updating user failed")
	}

	key, iv := encrypt.KeysFromString(ctx.RealIP() + ctx.Request().UserAgent())

	enc, err := encrypt.New("aes-256-cbc", key, iv)
	if err != nil {
		return result, errors.Wrap(err, "creating encryption failed")
	}

	tokenValue, err := enc.Encrypt([]byte(token.Token))
	if err != nil {
		return result, errors.Wrap(err, "encryption failed")
	}

	http.SetCookie(ctx.Response(), &http.Cookie{
		Name:    AccessCookieName,
		Value:   hex.EncodeToString(tokenValue),
		Expires: token.DTExpired,
		Path:    "/",
	})

	return result, nil
}

func AuthByCookie(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ac, ok := ctx.(*AdminContext)
		if !ok {
			return ErrContextNotConfigured
		}

		u, err := authByCookie(ac)
		if err != nil {
			return err
		}

		u.Current = true
		ctx.Set(UserContextKey, u)

		return withViewData(handlerFunc)(ctx)
	}
}

func authByCookie(ctx *AdminContext) (*User, error) {
	t, err := ctx.Cookie(AccessCookieName)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err)
	}

	value, err := hex.DecodeString(t.Value)
	if err != nil {
		return nil, errors.Wrap(err, "decoding cookie value failed")
	}

	key, iv := encrypt.KeysFromString(ctx.RealIP() + ctx.Request().UserAgent())

	enc, err := encrypt.New("aes-256-cbc", key, iv)
	if err != nil {
		return nil, errors.Wrap(err, "creating encryption failed")
	}

	tokenValue, err := enc.Decrypt(value)
	if err != nil {
		return nil, errors.Wrap(err, "decryption failed")
	}

	c := ctx.Request().Context()
	token, err := ctx.UserCase().SearchToken(c, string(tokenValue))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	if token.Type != AuthToken {
		return nil, echo.NewHTTPError(http.StatusForbidden)
	}

	if token.IsExpired() {
		return nil, echo.NewHTTPError(http.StatusNotFound)
	}

	token.User.Current = true

	return token.User, nil
}
