package immu_test

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
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
	m     usecase.Manager
	failM usecase.Manager
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

	t.Setenv("immu.port", "123124")
	viper.AutomaticEnv()

	suite.failM = manager.NewManager()
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
