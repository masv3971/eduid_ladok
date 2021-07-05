package model

import (
	"encoding/xml"
	"time"
)

type (
	// StudentUID is ladok uniq identifyer for student
	StudentUID string
)

// EventTimestamp collects timestamps
type EventTimestamp struct {
	Title     string
	Timestamp time.Time
}

// MSG holds feed and added metadata
type MSG struct {
	Feed   *Feed
	Events []*ChannelEvent
}

// EventPayload consists of the actual payload for each message
type EventPayload struct {
	StudentUID StudentUID
	EntryUID   string
	EventType  string
}

// ChannelEvent is the message on the channel, from ladok-atom to aggregate
type ChannelEvent struct {
	Payload     *EventPayload
	Timestamps  []*EventTimestamp
	ChannelName string
}

// SD -> Studiedeltanage
// SI -> Studentinformation

// EventHeaderLink holds Link
type EventHeaderLink struct {
	Text string `xml:",chardata"`
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
	Href string `xml:"href,attr"`
}

// Feed holds the header for each feed
type Feed struct {
	XMLName   xml.Name          `xml:"feed"`
	Xmlns     string            `xml:"xmlns,attr"`
	Title     Title             `xml:"title"`
	Link      []EventHeaderLink `xml:"link"`
	ID        string            `xml:"id"`
	Generator Generator         `xml:"generator"`
	Updated   time.Time         `xml:"updated"`
	Entry     []FeedEntry       `xml:"entry"`
}

// Title ladok
type Title struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

// Generator ladok
type Generator struct {
	Text string `xml:",chardata"`
	URI  string `xml:"uri,attr"`
}

// Category ladok
type Category struct {
	Label string `xml:"label"`
	Term  string `xml:"term"`
}

// EventContext ladok
type EventContext struct {
	Text       string `xml:",chardata"`
	LarosateID string `xml:"LarosateID"`
}

// SDAterbudHandelse ladok
type SDAterbudHandelse struct {
	Text                    string       `xml:",chardata"`
	Sd                      string       `xml:"sd,attr"`
	Dap                     string       `xml:"dap,attr"`
	Base                    string       `xml:"base,attr"`
	EventContext            EventContext `xml:"EventContext"`
	Handelsetid             time.Time    `xml:"Handelsetid"`
	HandelseUID             string       `xml:"HandelseUID"`
	StudentUID              StudentUID   `xml:"StudentUID"`
	UtbildningstillfalleUID string       `xml:"UtbildningstillfalleUID"`
	Aterbudsdatum           time.Time    `xml:"Aterbudsdatum"`
}

// SDAvbrottEvent ladok
type SDAvbrottEvent struct {
	Text             string       `xml:",chardata"`
	Sd               string       `xml:"sd,attr"`
	Dap              string       `xml:"dap,attr"`
	Base             string       `xml:"base,attr"`
	EventContext     EventContext `xml:"EventContext"`
	Handelsetid      time.Time    `xml:"Handelsetid"`
	HandelseUID      string       `xml:"HandelseUID"`
	StudentUID       StudentUID   `xml:"StudentUID"`
	Avbrottsdatum    time.Time    `xml:"Avbrottsdatum"`
	UtbildningUID    string       `xml:"UtbildningUID"`
	UtbildningstypID int          `xml:"UtbildningstypID"`
}

// SDPaborjatUtbildningstillfalleEvent ladok
type SDPaborjatUtbildningstillfalleEvent struct {
	Text                       string       `xml:",chardata"`
	Sd                         string       `xml:"sd,attr"`
	Dap                        string       `xml:"dap,attr"`
	Base                       string       `xml:"base,attr"`
	EventContext               EventContext `xml:"EventContext"`
	Handelsetid                time.Time    `xml:"Handelsetid"`
	HandelseUID                string       `xml:"HandelseUID"`
	Senaredelmarkering         bool         `xml:"Senaredelmarkering"`
	StudentUID                 StudentUID   `xml:"StudentUID"`
	TillfallesdeltagandeUID    string       `xml:"TillfallesdeltagandeUID"`
	UtbildningUID              string       `xml:"UtbildningUID"`
	UtbildningstillfalleUID    string       `xml:"UtbildningstillfalleUID"`
	UtbildningstillfallestypID int          `xml:"UtbildningstillfallestypID"`
	UtbildningsversionUID      string       `xml:"UtbildningsversionUID"`
}

//SDForstagangsregistreringHandelse ladok
type SDForstagangsregistreringHandelse struct {
	Text                    string       `xml:",chardata"`
	Sd                      string       `xml:"sd,attr"`
	Dap                     string       `xml:"dap,attr"`
	Base                    string       `xml:"base,attr"`
	EventContext            EventContext `xml:"EventContext"`
	Handelsetid             time.Time    `xml:"Handelsetid"`
	HandelseUID             string       `xml:"HandelseUID"`
	KursUID                 string       `xml:"KursUID"`
	KurstillfalleUID        string       `xml:"KurstillfalleUID"`
	Period                  int          `xml:"Period"`
	StudentUID              StudentUID   `xml:"StudentUID"`
	TillfallesdeltagandeUID string       `xml:"TillfallesdeltagandeUID"`
	UtbildningsversionUID   string       `xml:"UtbildningsversionUID"`
	Studieavgiftsbetalning  string       `xml:"Studieavgiftsbetalning"`
}

// SIStudentEvent ladok
type SIStudentEvent struct {
	Efternamn         string     `xml:"Efternamn"`
	ExterntStudentUID string     `xml:"ExterntStudentUID"`
	Fodelsedata       string     `xml:"Fodelsedata"`
	Fornamn           string     `xml:"Fornamn"`
	Kon               int        `xml:"Kon"`
	Personnummer      string     `xml:"Personnummer"`
	StudentUID        StudentUID `xml:"StudentUID"`
}

// SIStudentrestriktionEvent ladok
type SIStudentrestriktionEvent struct {
	Anteckning            string     `xml:"Anteckning"`
	Giltighetsperiod      time.Time  `xml:"Giltighetsperiod"`
	StudentUID            StudentUID `xml:"StudentUID"`
	StudentrestriktionUID string     `xml:"StudentrestriktionUID"`
	//RestriktionEventPart
	//Restriktionstyp"
}

// SILokalStudentEvent ladok
type SILokalStudentEvent struct {
	Efternamn         string     `xml:"Efternamn"`
	ExterntStudentUID string     `xml:"ExterntStudentUID"`
	Fodelsedata       string     `xml:"Fodelsedata"`
	Fornamn           string     `xml:"Fornamn"`
	Kon               int        `xml:"Kon"`
	Personnummer      string     `xml:"Personnummer"`
	StudentUID        StudentUID `xml:"StudentUID"`
}

// Content ladok
type Content struct {
	Text                              string                               `xml:",chardata"`
	Type                              string                               `xml:"type,attr"`
	AterbudHandelse                   *SDAterbudHandelse                   `xml:"AterbudHandelse,omitempty"`
	AvbrottEvent                      *SDAvbrottEvent                      `xml:"AvbrottEvent,omitempty"`
	PaborjatUtbildningstillfalleEvent *SDPaborjatUtbildningstillfalleEvent `xml:"PaborjatUtbildningstillfalleHandelse,omitempty"`
	ForstagangsregistreringHandelse   *SDForstagangsregistreringHandelse   `xml:"ForstagangsregistreringHandelse,omitempty"`
	StudentEvent                      *SIStudentEvent                      `xml:"StudentEvent"`
	StudentrestriktionEvent           *SIStudentrestriktionEvent           `xml:"StudentrestriktionEvent"`
	LokalStudentEvent                 *SILokalStudentEvent                 `xml:"LokalStudentEvent"`
}

// FeedEntry holds types for an entry
type FeedEntry struct {
	Category Category  `xml:"category"`
	ID       string    `xml:"ID"`
	Updated  time.Time `xml:"updated"`
	Content  Content   `xml:"content"`
}

// EduIDIAMBody dummy
type EduIDIAMBody struct{}

// SIStudentRest student in ladok
type SIStudentRest struct {
	XMLName                           xml.Name  `xml:"Student"`
	Avliden                           bool      `xml:"Avliden" json:"avliden"`
	Efternamn                         string    `xml:"Efternamn" json:"efternamn"`
	ExterntUID                        string    `xml:"ExterntUID" json:"external_uid"`
	FelVidEtableringExternt           bool      `xml:"FelVidEtableringExternt" json:"fel_vid_etablering_externt"`
	FolkbokforingsbevakningTillOchMed time.Time `xml:"FolkbokforingsbevakningTillOchMed" json:"folkbokforingsbevakning_till_och_med"`
	Fodelsedata                       string    `xml:"Fodelsedata" json:"fodelsedata"`
	Fornamn                           string    `xml:"Fornamn" json:"fornamn"`
	KonID                             int       `xml:"KonID" json:"kod_id"`
	Personnummer                      string    `xml:"Personnummer" json:"personnummer"`
	//UnikaIdentifierare
}
