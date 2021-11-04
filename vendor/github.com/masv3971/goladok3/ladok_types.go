package goladok3

// Permissions is a simplify permissions object
type Permissions map[int64]string

// Link is a general ladok link structure
type Link struct {
	Method    string `json:"method"`
	URI       string `json:"uri"`
	MediaType string `json:"mediaType"`
	Rel       string `json:"rel"`
}

// Benamning is a general ladok Benamning structure
type Benamning struct {
	Sprakkod string     `json:"Sprakkod"`
	Text     string     `json:"Text"`
	Link     []struct{} `json:"link"`
}

type (
	// FeedID ladok id is an int
	FeedID string
)

type serviceTypes struct {
	client       *Client
	service      string
	acceptHeader string
}

const (
	envIntTestAPI = "Int-test-API"
	envProdAPI    = "Prod-API"
	envTestAPI    = "Test-API"

	// LokalStudentEventName type
	LokalStudentEventName = "LokalStudentEvent"
	// AnvandareAndradEventName type
	AnvandareAndradEventName = "AnvandareAndradEvent"
	// AnvandareSkapadEventName type
	AnvandareSkapadEventName = "AnvandareSkapadEvent"
	// ExternPartEventName type
	ExternPartEventName = "ExternPartEvent"
	// KontaktuppgifterEventName type
	KontaktuppgifterEventName = "KontaktuppgifterEvent"
	// ResultatPaModulAttesteratEventName type
	ResultatPaModulAttesteratEventName = "ResultatPaModulAttesteratEvent"
	// ResultatPaHelKursAttesteratEventName type
	ResultatPaHelKursAttesteratEventName = "ResultatPaHelKursAttesteratEvent"
)

var (
	// ContentTypeStudiedeltagandeJSON server response content type
	ContentTypeStudiedeltagandeJSON = "application/vnd.ladok-studiedeltagande+json;charset=UTF-8"
	// ContentTypeKataloginformationJSON server response content type
	ContentTypeKataloginformationJSON = "application/vnd.ladok-kataloginformation+json;charset=UTF-8"
	// ContentTypeStudentinformationJSON server response content type
	ContentTypeStudentinformationJSON = "application/vnd.ladok-studentinformation+json;charset=UTF-8"
	// ContentTypeAtomXML server response content type
	ContentTypeAtomXML = "application/atom+xml;charset=UTF-8"

	// TypeStudentinformation type
	TypeStudentinformation = "studentinformation"
	//  = "externstudentevent"
)
