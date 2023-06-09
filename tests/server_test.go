package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
}

// ========================

func TestServer(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
