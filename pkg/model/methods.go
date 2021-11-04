package model

//func (f Feed) String() (string, error) {
//	data, err := xml.MarshalIndent(f, "", "  ")
//	if err != nil {
//		return "", err
//	}
//	// strip empty line from default xml header
//	s := xml.Header[:len(xml.Header)-1] + string(data)
//	return s, nil
//}

// MakeChannelEvents splits a *MSG into StudentUID to a map[StudentUID]string
//func (msg *MSG) MakeChannelEvents() {
//	for _, feedEntry := range msg.Feed.Entry {
//		if feedEntry.Content.AvbrottEvent != nil {
//			ce := &ChannelEvent{
//				Payload: &EventPayload{
//					StudentUID: feedEntry.Content.AvbrottEvent.StudentUID,
//					EntryID:    feedEntry.ID,
//					EventType:  EventTypeAvbrottEvent,
//				},
//			}
//			msg.Events = append(msg.Events, ce)
//		}
//
//		if feedEntry.Content.AterbudHandelse != nil {
//			ce := &ChannelEvent{
//				Payload: &EventPayload{
//					StudentUID: feedEntry.Content.AterbudHandelse.StudentUID,
//					EntryID:    feedEntry.ID,
//					EventType:  EventTypeAterbudHandelse,
//				},
//			}
//			msg.Events = append(msg.Events, ce)
//		}
//
//		if feedEntry.Content.PaborjatUtbildningstillfalleEvent != nil {
//			ce := &ChannelEvent{
//				Payload: &EventPayload{
//					StudentUID: feedEntry.Content.PaborjatUtbildningstillfalleEvent.StudentUID,
//					EntryID:    feedEntry.ID,
//					EventType:  EventTypePaborjatUtbildningstillfalleEvent,
//				},
//			}
//			msg.Events = append(msg.Events, ce)
//		}
//
//		if feedEntry.Content.ForstagangsregistreringHandelse != nil {
//			ce := &ChannelEvent{
//				Payload: &EventPayload{
//					StudentUID: feedEntry.Content.ForstagangsregistreringHandelse.StudentUID,
//					EntryID:    feedEntry.ID,
//					EventType:  EventTypeForstagangsregistreringHandelse,
//				},
//			}
//			msg.Events = append(msg.Events, ce)
//		}
//
//		if feedEntry.Content.StudentEvent != nil {
//			ce := &ChannelEvent{
//				Payload: &EventPayload{
//					StudentUID: feedEntry.Content.StudentEvent.StudentUID,
//					EntryID:    feedEntry.ID,
//					EventType:  EventTypeStudentEvent,
//				},
//			}
//			msg.Events = append(msg.Events, ce)
//		}
//
//		if feedEntry.Content.StudentrestriktionEvent != nil {
//			ce := &ChannelEvent{
//				Payload: &EventPayload{
//					StudentUID: feedEntry.Content.StudentrestriktionEvent.StudentUID,
//					EntryID:    feedEntry.ID,
//					EventType:  EventTypeStudentrestriktionEvent,
//				},
//			}
//			msg.Events = append(msg.Events, ce)
//		}
//
//		if feedEntry.Content.LokalStudentEvent != nil {
//			ce := &ChannelEvent{
//				Payload: &EventPayload{
//					StudentUID: feedEntry.Content.LokalStudentEvent.StudentUID,
//					EntryID:    feedEntry.ID,
//					EventType:  EventTypeLokalStudentEvent,
//				},
//			}
//			msg.Events = append(msg.Events, ce)
//		}
//
//	}
//}

// AddTimestamp adds a timestamp to a map[string]time.Time
//func (c *ChannelEvent) AddTimestamp(title string) {
//	ts := &EventTimestamp{
//		Title:     title,
//		Timestamp: Now(),
//	}
//	c.Timestamps = append(c.Timestamps, ts)
//
//}
