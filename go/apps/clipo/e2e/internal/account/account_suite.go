package account

import "github.com/stretchr/testify/suite"

var httpEndpoint string
var mailEndpoint string
var db string

type TestSuite struct {
	suite.Suite
	apiUrl string
}
