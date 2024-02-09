package immu

import (
	"context"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"log"
)

func (m *Manager) CreateProject(ctx context.Context, name string) (uuid.UUID, error) {

	id := ulid.Make()

	projectUuid, _ := uuid.FromBytes(id.Bytes())

	err := m.openSession(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	tx, err := m.client.NewTx(context.TODO())
	if err != nil {
		log.Fatal("failed to start transaction", err)
	}

	err = tx.SQLExec(ctx, "INSERT INTO projects (id, name) values (@id, @name)", map[string]interface{}{
		"id":   id.Bytes(),
		"name": name,
	})
	if err != nil {
		_ = tx.Rollback(ctx)
		return uuid.Nil, err
	}

	_, err = tx.Commit(ctx)
	if err != nil {
		log.Fatal("failed to start transaction", err)
	}

	return projectUuid, nil
}
