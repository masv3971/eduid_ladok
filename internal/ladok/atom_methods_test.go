package ladok

import (
	"context"
	"testing"

	"github.com/masv3971/goladok3/testinginfra"
	"github.com/stretchr/testify/assert"
)

func TestAtomRun(t *testing.T) {
	tts := []struct {
		name string
		keys []string
	}{
		{
			name: "OK",
			keys: []string{
				testinginfra.AnvandareAndradEventID,
				testinginfra.AnvandareSkapadEventID,
				testinginfra.KontaktuppgifterEventID,
				testinginfra.ResultatPaModulAttesteratEventID,
				testinginfra.ExternPartEventID,
				testinginfra.LokalStudentEventID,
				testinginfra.ResultatPaHelKursAttesteratEventID,
			},
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			service, server := mockService(t, 200)
			defer server.Close()

			service.Atom.run(context.TODO())

			for _, key := range tt.keys {
				select {
				case msg := <-service.Atom.Channel:
					assert.Equal(t, key, msg.Event.EntryID)
				}
			}
		})
	}
}

func TestKeyValueStore(t *testing.T) {
	t.SkipNow()
	type want struct {
		value interface{}
		found bool
	}
	tts := []struct {
		name string
		add  string
		get  string
		want want
	}{
		{
			name: "Found",
			add:  "A",
			get:  "A",
			want: want{
				value: "A",
				found: true,
			},
		},
		{
			name: "Not Found",
			add:  "A",
			get:  "B",
			want: want{
				value: nil,
				found: false,
			},
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			//service, server := mockService(t, 200)
			//defer server.Close()

			//	service.Atom.db.Add(tt.add, tt.add, memCache.DefaultExpiration)

			//	got, found := service.Atom.db.HGet()

			//	assert.Equal(t, tt.want.value, got)
			//	assert.Equal(t, tt.want.found, found)
		})
	}
}
