package immu

import (
	"fmt"
	immudb "github.com/codenotary/immudb/pkg/client"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"kimcha/internal/usecase"
)

type Manager struct {
	client   immudb.ImmuClient
	user     []byte
	password []byte
	db       string
}

func NewManager() (usecase.Manager, error) {
	host := viper.GetString("immu.host")
	port := viper.GetInt("immu.port")

	opts := immudb.DefaultOptions().
		WithAddress(host).
		WithPort(port)
	client := immudb.NewClient().WithOptions(opts)

	user := viper.GetString("immu.user")
	password := viper.GetString("immu.password")
	db := viper.GetString("immu.db")

	return &Manager{
		client:   client,
		user:     []byte(user),
		password: []byte(password),
		db:       db,
	}, nil
}

func (m *Manager) openSession(ctx context.Context) error {
	return m.client.OpenSession(ctx, m.user, m.password, m.db)
}

func (m *Manager) closeSession(ctx context.Context) error {
	return m.client.CloseSession(ctx)
}

func (m *Manager) SetSecret(ctx context.Context, projectId uuid.UUID, key, value string) error {

	err := m.openSession(ctx)
	defer func() {
		_ = m.closeSession(ctx)
	}()
	if err != nil {
		return err
	}

	_, err = m.client.VerifiedSet(
		ctx,
		[]byte(fmt.Sprintf("%s.%s", projectId.String(), key)),
		[]byte(value),
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetSecret(ctx context.Context, projectId uuid.UUID, key string) (string, error) {

	err := m.openSession(ctx)
	defer func() {
		_ = m.closeSession(ctx)
	}()

	if err != nil {
		return "", err
	}

	entry, err := m.client.Get(
		ctx,
		[]byte(fmt.Sprintf("%s.%s", projectId.String(), key)),
	)
	if err != nil {
		return "", err
	}

	return string(entry.Value), nil
}
