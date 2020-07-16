package msteams_webhook

// https://docs.microsoft.com/en-us/outlook/actionable-messages/message-card-reference

type InputType string
type ActionType string

const (
	DateInput        InputType = "DateInput"
	MultichoiceInput           = "MultichoiceInput"
	TextInput                  = "TextInput"
)
const (
	HttpPOST           ActionType = "HttpPOST"
	ActionCard                    = "ActionCard"
	OpenUri                       = "OpenUri"
	InvokeAddInCommand            = "InvokeAddInCommand"
)

type MessageCard struct {
	Context          string     `json:"@context"`
	Type             string     `json:"@type"`
	ThemeColor       string     `json:"themeColor"`
	Text             string     `json:"text"`
	Title            string     `json:"title"`
	Sections         []*Section `json:"sections"`
	PotentialAction  []*Actions `json:"potentialAction"`
	HideOriginalBody bool       `json:"hideOriginalBody"`
}

type Section struct {
	Title            string   `json:"title"`
	StartGroup       bool     `json:"startGroup"`
	ActivityImage    string   `json:"activityImage"`
	ActivityTitle    string   `json:"activityTitle"`
	ActivitySubtitle string   `json:"activitySubtitle"`
	ActivityText     string   `json:"activityText"`
	HeroImage        *Image   `json:"heroImage"`
	Images           []*Image `json:"images"`
	Text             string   `json:"text"`
	Facts            []*Fact  `json:"facts"`
}

type Actions struct {
	Type            ActionType `json:"@type"`
	Name            string     `json:"name"`
	Target          string     `json:"target"`
	Headers         []*Header  `json:"headers"`
	Body            string     `json:"body"`
	BodyContentType string     `json:"bodyContentType"`
	Inputs          []*Input   `json:"inputs"`
	Actions         []*Actions `json:"actions"`
}

type Input struct {
	Type       InputType `json:"@type"`
	Id         string    `json:"id"`
	Title      string    `json:"title"`
	IsRequired bool      `json:"is_required"`
	Value      *string   `json:"value"`
}

type Image struct {
	Image string `json:"image"`
	Title string `json:"title"`
}

type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
