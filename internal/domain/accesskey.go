package domain

import (
	"github.com/google/uuid"
	"time"
)

type AccessKey struct {
	Id        uuid.UUID
	ProjectId uuid.UUID
	Note      string
	Key       string
	// Secret key mask AB****************************************12
	Mask      string
	Signature []byte
	Expires   *time.Time
}
