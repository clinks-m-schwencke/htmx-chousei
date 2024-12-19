package handlers

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

const (
	SESSION_NAME               string = "fmessages"
	SESSION_FLASH_MESSAGES_KEY string = "flashmessages-key"
)

// Gets the cookie store at the key
func getCookieStore() *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(SESSION_FLASH_MESSAGES_KEY))
}

// Set adds a new message to the cookie store
func setFlashMessages(c echo.Context, kind, value string) {
	// Gets the cookie store
	session, _ := getCookieStore().Get(c.Request(), SESSION_NAME)

	// Adds a flash message TODO: What's a flash message?
	session.AddFlash(value, kind)

	// Saves the message to the request??
	session.Save(c.Request(), c.Response())
}

func getFlashMessages(c echo.Context, kind string) []string {
	// Get session from cookie
	session, _ := getCookieStore().Get(c.Request(), SESSION_NAME)

	// get flash messages
	flashMessages := session.Flashes(kind)

	// If there are any messages...
	if len(flashMessages) > 0 {
		session.Save(c.Request(), c.Response())

		// Declare string slice to append to later
		var flashes []string
		for _, fl := range flashMessages {
			// We add the messages to the slice
			flashes = append(flashes, fl.(string))
		}

		// Return string slice
		return flashes
	}

	// No messages, return nil
	return nil

}
