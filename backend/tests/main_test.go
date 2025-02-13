package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestUserAndPostSuite(t *testing.T) {
	suite.Run(t, new(PostTestSuite))
	suite.Run(t, new(UserTestSuite))
}
