package immu_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func (suite *ImmuTestSuite) TestSet() {

	t := suite.T()

	ud := uuid.New()

	err := suite.m.SetSecret(context.Background(), ud, "key", "bla-bla-bla-bla-bla-bla")
	require.NoError(t, err)
}

func (suite *ImmuTestSuite) TestSet_NoConnection() {

	t := suite.T()

	ud := uuid.New()

	err := suite.failM.SetSecret(context.Background(), ud, "key", "bla-bla-bla-bla-bla-bla")
	require.Error(t, err)
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

func (suite *ImmuTestSuite) TestGet_NoConnection() {

	ud := uuid.New()

	t := suite.T()

	_, err := suite.failM.GetSecret(context.Background(), ud, "key")
	require.Error(t, err)
}
