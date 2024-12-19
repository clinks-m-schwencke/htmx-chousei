package lang

var Messages = map[string]I18nStrings{
	"en": EnMessages,
	"ja": JaMessages,
}

type I18nStrings struct {
	BaseLayoutStrings   BaseLayoutStrings
	HomePageStrings     HomePageStrings
	LoginPageStrings    LoginPageStrings
	RegisterPageStrings RegisterPageStrings
	TaskPageStrings     TaskPageStrings
}

type BaseLayoutStrings struct {
	Title                string
	Tasks                string
	Logout               string
	Register             string
	Login                string
	LogoutConfirmTitle   string
	LogoutConfirmDetails string
	LogoutConfirmOk      string
	LogoutConfirmCancel  string
}

type HomePageStrings struct {
	Title       string
	Description string
	HaveAccount string
	Login       string
	Register    string
}

type LoginPageStrings struct {
	Title        string
	Email        string
	Password     string
	ViewPassword string
	Login        string
	Disabled     string
}

type RegisterPageStrings struct {
	Title        string
	Email        string
	Name         string
	Password     string
	ViewPassword string
	Register     string
	Disabled     string
}

type TaskPageStrings struct {
	Id                     string
	Complete               string
	Reviewed               string
	Tasks                  string
	CreatedBy              string
	Assigned               string
	Reviewer               string
	DueDate                string
	Options                string
	SelectUserPlaceholder  string
	NoTasks                string
	CreateTask             string
	EditTask               string
	UpdateTask             string
	DeleteTask             string
	Cancel                 string
	DeleteTaskConfirmTitle string
	// DeleteTaskConfirmDetails string
	DeleteTaskConfirmOk string
}
