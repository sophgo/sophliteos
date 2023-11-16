package router

import (
	"algoliteos/router/algorithm"
)

type RouterGroup struct {
	Algorithm algorithm.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
