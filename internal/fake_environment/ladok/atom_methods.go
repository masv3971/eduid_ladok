package ladok

import (
	"eduid_ladok/pkg/model"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (s *AtomService) entrySI() []model.FeedEntry {
	var feedEntries []model.FeedEntry
	for x := 1; x <= 10; x++ {
		var nin string
		if x < 10 {
			nin = fmt.Sprintf("190%d-0%d-0%d-xxxx", x, x, x)
		} else {
			nin = fmt.Sprintf("19%d-0%d-0%d-xxxx", x, x, x)
		}

		studentEvent := model.FeedEntry{
			Category: model.Category{
				Label: "Event-typ",
				Term:  "se.ladok.studentinformation.interfaces.events.StudentEvent",
			},
			ID:      uuid.New().String(),
			Updated: time.Now(),
			Content: model.Content{
				Text: "",
				Type: "application/vnd.ladok+xml",
				StudentEvent: &model.SIStudentEvent{
					Efternamn:         fmt.Sprintf("testEfternamn-%d", x),
					ExterntStudentUID: uuid.New().String(),
					Fodelsedata:       "",
					Fornamn:           fmt.Sprintf("testFornamn-%d", x),
					Kon:               1,
					Personnummer:      nin,
					StudentUID:        model.StudentUID(uuid.New().String()),
				},
			},
		}
		feedEntries = append(feedEntries, studentEvent)

		studentrestriktionEvent := model.FeedEntry{
			Category: model.Category{
				Label: "Event-typ",
				Term:  "se.ladok.studentinformation.interfaces.events.StudentrestriktionEvent",
			},
			ID:      uuid.New().String(),
			Updated: time.Now(),
			Content: model.Content{
				Text: "",
				Type: "application/vnd.ladok+xml",
				StudentrestriktionEvent: &model.SIStudentrestriktionEvent{
					Anteckning:            fmt.Sprintf("Anteckning-%d", x),
					Giltighetsperiod:      time.Now(),
					StudentUID:            model.StudentUID(uuid.New().String()),
					StudentrestriktionUID: uuid.New().String(),
				},
			},
		}
		feedEntries = append(feedEntries, studentrestriktionEvent)

		lokalStudentEvent := model.FeedEntry{
			Category: model.Category{
				Label: "Event-typ",
				Term:  "se.ladok.studentinformation.interfaces.events.LokalStudentEvent",
			},
			ID:      uuid.New().String(),
			Updated: time.Now(),
			Content: model.Content{
				Text: "",
				Type: "application/vnd.ladok+xml",
				LokalStudentEvent: &model.SILokalStudentEvent{
					Efternamn:         fmt.Sprintf("testEfternamn-%d", x),
					ExterntStudentUID: uuid.New().String(),
					Fodelsedata:       "",
					Fornamn:           fmt.Sprintf("testFornamn-%d", x),
					Kon:               1,
					Personnummer:      nin,
					StudentUID:        model.StudentUID(uuid.New().String()),
				},
			},
		}
		feedEntries = append(feedEntries, lokalStudentEvent)
	}
	return feedEntries
}

func (s *AtomService) entrySD() []model.FeedEntry {
	var feedEtries []model.FeedEntry
	for x := 1; x <= 10; x++ {

		aterbudEvent := model.FeedEntry{
			Category: model.Category{
				Label: "Event-typ",
				Term:  "se.ladok.studiedeltagande.interfaces.events.ÅterbudEvent",
			},
			ID:      uuid.New().String(),
			Updated: time.Now(),
			Content: model.Content{
				Text: "",
				Type: "application/vnd.ladok+xml",
				AterbudHandelse: &model.SDAterbudHandelse{
					Text:                    "",
					Sd:                      "http://schemas.ladok.se/studiedeltagande",
					Dap:                     "http://schemas.ladok.se/dap",
					Base:                    "http://schemas.ladok.se",
					EventContext:            model.EventContext{Text: "", LarosateID: "42"},
					Handelsetid:             time.Now(),
					HandelseUID:             uuid.New().String(),
					StudentUID:              model.StudentUID(uuid.New().String()),
					UtbildningstillfalleUID: uuid.New().String(),
					Aterbudsdatum:           time.Now(),
				},
			},
		}
		feedEtries = append(feedEtries, aterbudEvent)

		avbrottEvent := model.FeedEntry{
			Category: model.Category{Label: "Event-typ", Term: "se.ladok.studiedeltagande.interfaces.events.AvbrottEvent"},
			ID:       uuid.New().String(),
			Updated:  time.Now(),
			Content: model.Content{
				Text: "",
				Type: "application/vnd.ladok+xml",
				AvbrottEvent: &model.SDAvbrottEvent{
					Text:             "",
					Sd:               "http://schemas.ladok.se/studiedeltagande",
					Dap:              "http://schemas.ladok.se/dap",
					Base:             "http://schemas.ladok.se",
					EventContext:     model.EventContext{Text: "", LarosateID: "42"},
					Handelsetid:      time.Now(),
					HandelseUID:      uuid.New().String(),
					StudentUID:       model.StudentUID(uuid.New().String()),
					Avbrottsdatum:    time.Time{},
					UtbildningUID:    uuid.New().String(),
					UtbildningstypID: 1,
				},
			},
		}

		feedEtries = append(feedEtries, avbrottEvent)

		PaborjatUtbildningstillfalleEvent := model.FeedEntry{
			Category: model.Category{
				Label: "Event-typ",
				Term:  "se.ladok.studiedeltagande.interfaces.events.PåbörjatUtbildningstillfälleEvent",
			},
			ID:      uuid.New().String(),
			Updated: time.Now(),
			Content: model.Content{
				Text: "",
				Type: "application/vnd.ladok+xml",
				PaborjatUtbildningstillfalleEvent: &model.SDPaborjatUtbildningstillfalleEvent{
					Text:                       "",
					Sd:                         "http://schemas.ladok.se/studiedeltagande",
					Dap:                        "http://schemas.ladok.se/dap",
					Base:                       "http://schemas.ladok.se",
					EventContext:               model.EventContext{Text: "", LarosateID: "42"},
					Handelsetid:                time.Now(),
					HandelseUID:                uuid.New().String(),
					Senaredelmarkering:         false,
					StudentUID:                 model.StudentUID(uuid.New().String()),
					TillfallesdeltagandeUID:    uuid.New().String(),
					UtbildningUID:              uuid.New().String(),
					UtbildningstillfalleUID:    uuid.New().String(),
					UtbildningstillfallestypID: 1,
					UtbildningsversionUID:      uuid.New().String(),
				},
			},
		}
		feedEtries = append(feedEtries, PaborjatUtbildningstillfalleEvent)

		forstagangsregistreringHandelse := model.FeedEntry{
			Category: model.Category{},
			ID:       uuid.New().String(),
			Updated:  time.Now(),
			Content: model.Content{
				Text: "",
				Type: "application/vnd.ladok+xml",
				ForstagangsregistreringHandelse: &model.SDForstagangsregistreringHandelse{
					Text: "",
					Sd:   "http://schemas.ladok.se/studiedeltagande",
					Dap:  "http://schemas.ladok.se/dap",
					Base: "http://schemas.ladok.se",
					EventContext: model.EventContext{
						Text:       "",
						LarosateID: "42",
					},
					Handelsetid:             time.Now(),
					HandelseUID:             uuid.New().String(),
					KursUID:                 uuid.New().String(),
					KurstillfalleUID:        uuid.New().String(),
					Period:                  1,
					StudentUID:              model.StudentUID(uuid.New().String()),
					TillfallesdeltagandeUID: uuid.New().String(),
					UtbildningsversionUID:   uuid.New().String(),
					Studieavgiftsbetalning:  "studiedeltagande.domain.studieavgiftsbetalningstyp.ej_aktuell",
				},
			},
		}
		feedEtries = append(feedEtries, forstagangsregistreringHandelse)

	}
	return feedEtries
}

func (s *AtomService) handlerSI(w http.ResponseWriter, r *http.Request) {
	feed := model.Feed{
		XMLName: xml.Name{},
		Xmlns:   "http://www.w3.org/2005/Atom",
		Title: model.Title{
			Text: "Title",
			Type: "text",
		},
		Link: []model.EventHeaderLink{
			{
				Text: "",
				Rel:  "self",
				Type: "application/atom+xml",
				Href: "http://localhost/ladok/feeds/recent",
			},
			{
				Text: "",
				Rel:  "via",
				Type: "application/atom+xml",
				Href: "http://localhost/ladok/feeds/2",
			},
			{
				Text: "",
				Rel:  "prev-archive",
				Type: "application/atom+xml",
				Href: "http://localhost/ladok/feeds/1",
			},
		},
		ID: "urn:id:2",
		Generator: model.Generator{
			Text: "Uppfoljning",
			URI:  "http://ladok.se/studentinformation",
		},
		Updated: time.Now(),
		Entry:   s.entrySI(),
	}

	atom, err := feed.String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, atom)
}

func (s *AtomService) handlerSD(w http.ResponseWriter, r *http.Request) {
	feed := model.Feed{
		XMLName: xml.Name{},
		Xmlns:   "http://www.w3.org/2005/Atom",
		Title: model.Title{
			Text: "Title",
			Type: "text",
		},
		Link: []model.EventHeaderLink{
			{
				Text: "",
				Rel:  "self",
				Type: "application/atom+xml",
				Href: "http://localhost/ladok/feeds/recent",
			},
			{
				Text: "",
				Rel:  "via",
				Type: "application/atom+xml",
				Href: "http://localhost/ladok/feeds/2",
			},
			{
				Text: "",
				Rel:  "prev-archive",
				Type: "application/atom+xml",
				Href: "http://localhost/ladok/feeds/1",
			},
		},
		ID: "urn:id:2",
		Generator: model.Generator{
			Text: "Uppfoljning",
			URI:  "http://ladok.se/studiedeltagande",
		},
		Updated: time.Now(),
		Entry:   s.entrySD(),
	}

	atom, err := feed.String()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, atom)
}
