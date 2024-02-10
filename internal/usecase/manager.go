package usecase

import (
	"context"
	"github.com/google/uuid"
)

//go:generate mockery --name SecretManager
type SecretManager interface {
	SetSecret(ctx context.Context, projectId uuid.UUID, key, value string) error
	GetSecret(ctx context.Context, projectId uuid.UUID, key string) (string, error)
}

//go:generate mockery --name SecretManager
type DataManager interface {
	CreateProject(ctx context.Context, name string) (uuid.UUID, error)
}

type Manager interface {
	SecretManager
	DataManager
}
