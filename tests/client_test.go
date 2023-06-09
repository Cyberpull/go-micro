package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
}

// ========================

func TestClient(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
