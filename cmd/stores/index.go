package stores

type Meeting struct {
	Title       string
	Description string
	Dates       string
}

type Meetings []Meeting

func NewMeeting(title string, description string, dates string) Meeting {
	return Meeting{
		Title:       title,
		Description: description,
		Dates:       dates,
	}
}

type SessionData struct {
	Data     Data
	FormData Meeting
}

type Data struct {
	Meetings Meetings
}

func newData() Data {
	return Data{
		Meetings: []Meeting{
			NewMeeting("Party Time", "", "2024-07-11,2024-07-12,2024-07-13"),
			NewMeeting("Game night", "", "2024-07-21 17:00:00,2024-07-21 18:00:00,2024-07-21 19:00:00"),
		},
	}
}

func newSessionData() SessionData {
	return SessionData{
		Data: newData(),
	}
}

var SessionDataStore = newSessionData()
