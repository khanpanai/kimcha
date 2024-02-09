package domain

import "github.com/google/uuid"

type Project struct {
	Id   uuid.UUID
	Name string
}
