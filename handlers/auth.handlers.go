package handlers

import (
	"chopitto-task/lang"
	"chopitto-task/services"
	"chopitto-task/views/authviews"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/language"
)

const (
	AUTH_SESSIONS_KEY string = "authenticate-sessions"
	AUTH_KEY          string = "authenticated"
	USER_ID_KEY       string = "user_id"
	USERNAME_KEY      string = "username"
	TIME_ZONE_KEY     string = "time_zone"
	LANG_KEY          string = "lang"
)

var DEFAULT_LANG language.Tag = language.Japanese

type AuthService interface {
	CreatePerson(person services.Person) error
	CheckEmail(email string) (services.Person, error)
}

type AuthHandler struct {
	PersonServices AuthService
}

// Auth Handler constructor
func NewAuthHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{
		PersonServices: authService,
	}
}

func (authHandler *AuthHandler) detectLang(c echo.Context) (language.Tag, error) {
	sess, _ := session.Get(AUTH_SESSIONS_KEY, c)

	queryLang := c.QueryParam("lang")
	if queryLang != "" {
		print("Using query lang")
		println(queryLang)

		var newLang language.Tag

		switch queryLang {
		case "ja":
			fallthrough
		case "ja-JP":
			newLang = language.Japanese
		case "en":
			fallthrough
		default:
			newLang = language.English
		}

		sess.Values[LANG_KEY] = newLang.String()
		sess.Save(c.Request(), c.Response())

		println(newLang.String())
		return newLang, nil
	}

	cookieStr, ok := sess.Values[LANG_KEY].(string)
	cookieLang := language.Make(cookieStr)

	print("Cookie lang: ")
	println(cookieStr)
	println(cookieLang.String())
	print("ok?")
	println(ok)
	if ok && cookieLang != language.Und {
		println("Returning cookie lang!")
		return cookieLang, nil
	}

	headerLangs, _, err := language.ParseAcceptLanguage(c.Request().Header.Values("Accept-Language")[0])
	if err != nil {
		println("Parse Accept Language Error:: Default Lang Selected")
		println(err)
		println(DEFAULT_LANG.String())
		sess.Values[LANG_KEY] = DEFAULT_LANG.String()
		sess.Save(c.Request(), c.Response())
		c.Set(LANG_KEY, DEFAULT_LANG.String())
		return DEFAULT_LANG, err
	}

	var newLang language.Tag
	for _, headerLang := range headerLangs {
		switch headerLang {
		case language.Japanese:
			newLang = language.Japanese
			sess.Values[LANG_KEY] = newLang.String()
			sess.Save(c.Request(), c.Response())
			println("Accept Language Headers selected")
			println(newLang.String())
			return newLang, nil
		case language.English:
			newLang = language.English
			sess.Values[LANG_KEY] = newLang.String()
			sess.Save(c.Request(), c.Response())
			println("Accept Language Headers selected")
			println(newLang.String())
			return newLang, nil
		}
	}

	println("Default Lang Selected")
	println(DEFAULT_LANG.String())
	sess.Values[LANG_KEY] = DEFAULT_LANG.String()
	sess.Save(c.Request(), c.Response())
	return DEFAULT_LANG, nil

}

// Sets FROMPROTECTED if the person is authenticated
func (authHandler *AuthHandler) flagsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Detect Language and set value in cookie
		lang, _ := authHandler.detectLang(c)
		c.Set(LANG_KEY, lang.String())

		sess, _ := session.Get(AUTH_SESSIONS_KEY, c)
		auth, ok := sess.Values[AUTH_KEY].(bool)
		// If not authenticated
		if !ok || !auth {
			// Set from protected to false
			c.Set("FROMPROTECTED", false)
			// Proceed with next handler
			return next(c)
		}

		// Set from protected to true
		c.Set("FROMPROTECTED", true)
		// Proceed with next handler
		return next(c)
	}
}

// Middleware for logged in users
// Returns 401 error if user is not logged in
func (authHandler *AuthHandler) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Detect Language and set value in cookie
		lang, _ := authHandler.detectLang(c)
		c.Set(LANG_KEY, lang.String())

		// Get current session from session store ??
		sess, _ := session.Get(AUTH_SESSIONS_KEY, c)
		// Check AUTH key in session store
		auth, ok := sess.Values[AUTH_KEY].(bool)
		// If key is not found or auth is false
		if !ok || !auth {
			// Not authenticated, set protected to fasle
			c.Set("FROMPROTECTED", false)

			// return 401 unauthorised error
			return echo.NewHTTPError(echo.ErrUnauthorized.Code, "Please provide valid credentials")
		}

		// Get userId from session store
		userId, ok := sess.Values[USER_ID_KEY].(int)
		// if not 0
		if ok && userId != 0 {
			// Set context userId
			c.Set(USER_ID_KEY, userId)
		}

		// Get username from session store
		username, ok := sess.Values[USERNAME_KEY].(string)
		if ok && len(username) != 0 {
			c.Set(USERNAME_KEY, username)
		}

		// Get Timezone from session store
		timezone, ok := sess.Values[TIME_ZONE_KEY].(string)
		if ok && len(timezone) != 0 {
			c.Set(TIME_ZONE_KEY, timezone)
		}

		// Set Protected to true
		c.Set("FROMPROTECTED", true)

		// No error
		return next(c)
	}
}

// Handle '/' endpoint
func (authHandler *AuthHandler) homeHandler(c echo.Context) error {
	fromProtected, ok := c.Get("FROMPROTECTED").(bool)
	if !ok {
		return errors.New("invalid type for key FROMPROTECTED")
	}

	l := c.Get(LANG_KEY).(string)
	messages := lang.Messages[l]

	homeView := authviews.Home(fromProtected, messages.HomePageStrings)

	c.Set("ISERROR", false)

	return renderView(c, authviews.HomeIndex(
		"| Home",
		"",
		time.Now().Format("2006-01-02_15-04-05"),
		fromProtected,
		c.Get("ISERROR").(bool),
		getFlashMessages(c, "error"),
		getFlashMessages(c, "success"),
		homeView,
		messages.BaseLayoutStrings,
	))

}

// Handle /register endpoint
func (authHandler *AuthHandler) registerHandler(c echo.Context) error {
	fromProtected, ok := c.Get("FROMPROTECTED").(bool)
	if !ok {
		return errors.New("invalid type for key FROMPROTECTED")
	}

	l := c.Get(LANG_KEY).(string)
	messages := lang.Messages[l]

	registerView := authviews.Register(fromProtected, messages.RegisterPageStrings)

	c.Set("ISERROR", false)

	if c.Request().Method == "POST" {
		// Get values from <form>
		person := services.Person{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Name:     c.FormValue("name"),
		}

		// TODO: Error if email empty
		// TODO: Error if password empty
		// TODO: Error if name empty

		// Create person in db
		err := authHandler.PersonServices.CreatePerson(person)
		if err != nil {
			// Check error type
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				err = errors.New("the email is already in use")
				setFlashMessages(c, "error", fmt.Sprintf(
					"something went wrong: %s",
					err,
				))
				return c.Redirect(http.StatusSeeOther, "/register")
			}

			return echo.NewHTTPError(
				echo.ErrInternalServerError.Code,
				fmt.Sprintf(
					"something went wrong: %s",
					err,
				))
		}

		setFlashMessages(c, "success", "You have successfully registered!!")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	return renderView(c, authviews.RegisterIndex(
		"| Sign Up",
		"",
		time.Now().Format("2006-01-02_15-04-05"),
		fromProtected,
		c.Get("ISERROR").(bool),
		getFlashMessages(c, "error"),
		getFlashMessages(c, "success"),
		registerView,
		messages.BaseLayoutStrings,
	))

}

func (authHandler *AuthHandler) loginHandler(c echo.Context) error {
	fromProtected, ok := c.Get("FROMPROTECTED").(bool)
	if !ok {
		return errors.New("invalid type for key FROMPROTECTED")
	}

	l := c.Get(LANG_KEY).(string)
	messages := lang.Messages[l]
	loginView := authviews.Login(fromProtected, messages.LoginPageStrings)

	c.Set("ISERROR", false)

	if c.Request().Method == "POST" {
		// TODO Check for empty form values

		// obtaining the time zone from the POST request of the login request
		tzone := ""
		if len(c.Request().Header["X-Timezone"]) != 0 {
			tzone = c.Request().Header["X-Timezone"][0]
		}

		// Check Authentication credentials
		person, err := authHandler.PersonServices.CheckEmail(c.FormValue("email"))
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				// FIX: THis is a security flaw, fix
				setFlashMessages(c, "error", "There is no user with that email")
				// TODO: Send form back, rather than redirect whole page?
				return c.Redirect(http.StatusSeeOther, "/login")
			}

			return echo.NewHTTPError(
				echo.ErrInternalServerError.Code,
				fmt.Sprintf(
					"something went wrong: %s",
					err,
				))
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(person.Password),
			[]byte(c.FormValue("password")),
		)
		if err != nil {
			setFlashMessages(c, "error", "Incorrect password")

			return c.Redirect(http.StatusSeeOther, "/login")
		}

		// Get session and setting cookies
		sess, _ := session.Get(AUTH_SESSIONS_KEY, c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   3600 * 100, // TODO: Lower age
			HttpOnly: true,
		}

		// Set user as authenticated, their username,
		// their ID and the client's time zone
		sess.Values = map[interface{}]interface{}{
			AUTH_KEY:      true,
			USER_ID_KEY:   person.Id,
			USERNAME_KEY:  person.Name,
			TIME_ZONE_KEY: tzone,
		}

		sess.Save(c.Request(), c.Response())

		setFlashMessages(c, "success", "You have successfully logged in!!")

		return c.Redirect(http.StatusSeeOther, "/task")
	}

	return renderView(c, authviews.LoginIndex(
		"| Login",
		"",
		time.Now().Format("2006-01-02_15-04-05"),
		fromProtected,
		c.Get("ISERROR").(bool),
		getFlashMessages(c, "error"),
		getFlashMessages(c, "success"),
		loginView,
		messages.BaseLayoutStrings,
	))
}

func (authHandler *AuthHandler) logoutHandler(c echo.Context) error {
	sess, _ := session.Get(AUTH_SESSIONS_KEY, c)
	// Revoke users authentication
	sess.Values = map[interface{}]interface{}{
		AUTH_KEY:      false,
		USER_ID_KEY:   "",
		USERNAME_KEY:  "",
		TIME_ZONE_KEY: "",
	}

	// Save revoked cookie
	sess.Save(c.Request(), c.Response())

	// Send logout toast
	setFlashMessages(c, "success", "You have successfully logged out!!")

	// Set context auth status to false
	c.Set("FROMPROTECTED", false)

	// Redirect to login page
	return c.Redirect(http.StatusSeeOther, "/login")
}
