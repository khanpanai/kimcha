package usecase

import (
	"context"
	"github.com/google/uuid"
)

type SecretManager interface {
	SetSecret(ctx context.Context, projectId uuid.UUID, key, value string) error
	GetSecret(ctx context.Context, projectId uuid.UUID, key string) (string, error)
}

type DataManager interface {
	CreateProject(ctx context.Context, name string) (uuid.UUID, error)
}

type Manager interface {
	SecretManager
	DataManager
}
