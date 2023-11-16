package v1

import (
	"algoliteos/api/v1/algorithm"
)

type ApiGroup struct {
	AlgoGroup algorithm.AlgorithmGroup
}

var ApiGroupApp = new(ApiGroup)
