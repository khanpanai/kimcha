package immu_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func (suite *ImmuTestSuite) TestProject_Create() {
	t := suite.T()

	ud, err := suite.m.CreateProject(context.Background(), "test project")
	require.NoError(t, err)
	require.NotEqual(t, uuid.Nil, ud)
}
