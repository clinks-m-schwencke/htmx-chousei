package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"chopitto-task/lang"
	"chopitto-task/views/errorviews"
)

func CustomHttpErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	httpError, ok := err.(*echo.HTTPError)
	if ok {
		code = httpError.Code
	}
	c.Logger().Error(err)

	l := c.Get(LANG_KEY).(string)
	messages := lang.Messages[l]

	var errorPage func(fp bool) templ.Component

	switch code {
	case 401:
		errorPage = errorviews.Error401
	case 404:
		errorPage = errorviews.Error404
	case 500:
		errorPage = errorviews.Error500

	}

	c.Set("ISERROR", true)
	fromProtected, ok := c.Get("FROMPROTECTED").(bool)
	if !ok {
		log.Println("WARN: Error handler from protected has invalid type")
		log.Println(fromProtected)
		fromProtected = false
	}

	renderView(c, errorviews.ErrorIndex(
		fmt.Sprintf("| Error (%d)", code),
		"",
		time.Now().Format("2006-01-02_15-04-05"),
		fromProtected,
		true,
		errorPage(fromProtected),
		messages.BaseLayoutStrings,
	))
}

func RouteNotFoundHandler(c echo.Context) error {

	l := c.Get(LANG_KEY).(string)
	messages := lang.Messages[l]

	// Hard code parameters for not found page
	return renderView(c, errorviews.ErrorIndex(
		fmt.Sprintf("| Error (%d)", 404),
		"",
		time.Now().Format("2006-01-02_15-04-05"),
		false,
		true,
		errorviews.Error404(false),
		messages.BaseLayoutStrings,
	))
}
