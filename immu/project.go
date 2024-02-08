package immu

import (
	"context"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"kimcha/types"
	"log"
)

func (m *manager) CreateProject(ctx context.Context, name string) (types.ULID, error) {

	id := ulid.Make()

	bytes, _ := uuid.FromBytes(id.Bytes())

	err := m.openSession(ctx)
	if err != nil {
		return "", err
	}

	tx, err := m.client.NewTx(context.TODO())
	if err != nil {
		log.Fatal("failed to start transaction", err)
	}

	err = tx.SQLExec(ctx, "INSERT INTO project (id, name) values (?, ?)", map[string]interface{}{
		"id":   bytes.String(),
		"name": name,
	})
	if err != nil {
		_ = tx.Rollback(ctx)
		return "", err
	}

	_, err = tx.Commit(ctx)
	if err != nil {
		log.Fatal("failed to start transaction", err)
	}

	return types.ULID(id.String()), nil
}
