package syncutil

import (
	"bytes"
	"io"
	"time"
)

type secret struct {
	ID         string
	CreateTine time.Time
	token      string
}

func (s *secret) Read(p []byte) (int, error) {
	return bytes.NewBuffer(p).WriteString(s.token)
}

func NewSecret() io.Reader {
	return &secret{
		ID:         "dummy id",
		CreateTine: time.Now(),
		token:      "dymmy token",
	}
}
