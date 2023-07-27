package tests

import (
	"testing"

	"cyberpull.com/gosrv"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	ServerHost string = "localhost"
	ServerPort string = "1988"
)

type GoSRVTestSuite struct {
	suite.Suite

	server gosrv.Server
	client gosrv.Client
}

func (s *GoSRVTestSuite) SetupSuite() {
	// Start GoSRV Server
	require.NoError(s.T(), startServer(s))

	// Start GoSRV Client
	require.NoError(s.T(), startClient(s))
}

func (s *GoSRVTestSuite) TearDownSuite() {
	if s.client != nil {
		// Stop GoSRV Client
		require.NoError(s.T(), s.client.Stop())
	}

	if s.server != nil {
		// Stop GoSRV Server
		require.NoError(s.T(), s.server.Stop())
	}
}

func (s *GoSRVTestSuite) TestRequest() {
	value, err := gosrv.SendRequest[string](s.client, "GET", "/demo/name", "Howdy")
	require.NoError(s.T(), err)

	assert.Equal(s.T(), "Fine, thanks", value)
}

func (s *GoSRVTestSuite) TestUpdate() {
	updateChan := make(chan string, 1)

	s.client.On("GET", "/demo/update", func(data gosrv.Update) {
		var err error
		var update string

		err = data.GetError()
		require.NoError(s.T(), err)

		err = data.ParseContent(&update)
		require.NoError(s.T(), err)

		updateChan <- update
	})

	resp, err := gosrv.SendRequest[string](s.client, "GET", "/demo/update", "Hello")
	require.NoError(s.T(), err)

	assert.Equal(s.T(), "Demo Update", <-updateChan)
	assert.Equal(s.T(), "Hi", resp)
}

// ========================

func TestGoSRV(t *testing.T) {
	suite.Run(t, new(GoSRVTestSuite))
}
