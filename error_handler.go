package goadmin

import (
	"github.com/CloudyKit/jet"
	"github.com/labstack/echo/v4"
	"gitlab.com/app-builder/backoffice/api.git/admin/api"
	"net/http"
	"strings"
)

func errorHandler(e error, ctx echo.Context) {
	if strings.HasPrefix(ctx.Path(), "/api") {
		JSONError(e, ctx)
		return
	}

	HTMLError(e, ctx)
}

type viewData struct {
	Code    int
	Title   string
	Error   string
	Details string
}

func (data viewData) JetVars() jet.VarMap {
	vars := make(jet.VarMap)
	vars.Set("code", data.Code)
	vars.Set("error", data.Error)
	vars.Set("title", data.Title)
	vars.Set("details", data.Details)

	return vars
}

func (viewData) JetData() map[string]interface{} {
	return nil
}

func HTMLError(e error, ctx echo.Context) {
	code := http.StatusInternalServerError
	title, details := "", ""
	if he, ok := e.(*echo.HTTPError); ok {
		code = he.Code

		if he.Internal != nil {
			title = he.Internal.Error()
		} else {
			switch code {
			case http.StatusBadRequest:
				title = "Bad Request"
			case http.StatusInternalServerError:
				title = "Internal Server Error"
			case http.StatusNotFound:
				title = "Not Found"
			}
		}

		//details = he.Internal.Error()
	}

	data := &viewData{
		Code:    code,
		Title:   title,
		Error:   e.Error(),
		Details: details,
	}

	err := ctx.Render(code, "errors/error.jet", data)
	if err != nil {
		ctx.Logger().Error(err)
	}

	ctx.Logger().Error(e)
}

func JSONError(e error, ctx echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := e.(*echo.HTTPError); ok {
		code = he.Code
	}

	resp := &api.Response{
		Success: false,
		Error:   e.Error(),
	}

	if err := ctx.JSON(code, resp); err != nil {
		ctx.Logger().Error(err)
	}

	ctx.Logger().Error(e)
}
