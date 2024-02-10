package immu_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
	"io/fs"
	"kimcha/internal/manager"
	"kimcha/internal/usecase"
	immu2 "kimcha/pkg/immu"
	"os"
	"path/filepath"
	"testing"
)

type ImmuTestSuite struct {
	suite.Suite
	m usecase.Manager
}

func (suite *ImmuTestSuite) SetupSuite() {
	ti := immu2.SetupTestImmu()
	fmt.Println(ti.Port)
	t := suite.T()
	t.Setenv("immu.user", "immudb")
	t.Setenv("immu.password", "immudb")
	t.Setenv("immu.db", "defaultdb")
	t.Setenv("immu.port", ti.Port.Port())

	viper.AutomaticEnv()
	suite.m = manager.NewManager()
}

func (suite *ImmuTestSuite) TestSet() {

	t := suite.T()

	ud := uuid.New()

	err := suite.m.SetSecret(context.Background(), ud, "key", "bla-bla-bla-bla-bla-bla")
	require.NoError(t, err)
}

func (suite *ImmuTestSuite) TestGet() {

	ud := uuid.New()

	val := "bla-bla-bla-bla-bla-bla"

	t := suite.T()

	err := suite.m.SetSecret(context.Background(), ud, "key", val)
	require.NoError(t, err)

	value, err := suite.m.GetSecret(context.Background(), ud, "key")
	require.NoError(t, err)
	assert.EqualValues(t, val, value)
}

func (suite *ImmuTestSuite) TearDownSuite() {
	_ = filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if path == "." || path[0] != '.' {
			return nil
		}

		err = os.Remove(path)
		if err != nil {
			return err
		}

		return nil
	})
}

func TestImmuTestSuite(t *testing.T) {
	suite.Run(t, new(ImmuTestSuite))
}
