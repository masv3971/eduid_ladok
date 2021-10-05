package ladok

import (
	"context"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.step.sm/crypto/x509util"
)

func TestNewCertificateService(t *testing.T) {
	s := mockNewLadokService(t)

	type have struct {
		notAfter, notBefore int
	}

	tts := []struct {
		name string
		have have
		want error
	}{
		{
			name: "OK - 100 days left",
			have: have{
				notBefore: 0,
				notAfter:  100,
			},
			want: errors.New(""),
		},
		{
			name: "OK - 89 day left",
			have: have{
				notBefore: 0,
				notAfter:  89,
			},
			want: errors.New(""),
		},
		{
			name: "OK - 29 day left",
			have: have{
				notBefore: 0,
				notAfter:  29,
			},
			want: errors.New(""),
		},
		{
			name: "OK - 1 day left",
			have: have{
				notBefore: 0,
				notAfter:  1,
			},
			want: errors.New(""),
		},
		{
			name: "Old certificate",
			have: have{
				notBefore: -2,
				notAfter:  -1,
			},
			want: model.ErrCertificateNotValid,
		},
	}

	for _, tt := range tts {
		mockNewCertificateBundle(t, tt.have.notBefore, tt.have.notAfter, s.config.LadokCertificatePath, "pfx", s.config.LadokCertificatePassword)

		cs, err := NewCertificateService(context.TODO(), s, logger.New("test"))
		if err != nil {
			assert.EqualError(t, tt.want, err.Error())
			return
		}

		assert.NotEmpty(t, cs.CRT.NotAfter, "should not be empty")

		assert.Equal(t, x509util.Fingerprint(cs.CRT), cs.SHA256Fingerprint)
	}
}
