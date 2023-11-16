package main

import (
	"time"

	"algoliteos/initialization"
	"algoliteos/logger"
)

func main() {
	initialization.InitBase()
	Router := initialization.Routers()
	s := initialization.InitServer(Router)

	time.Sleep(10 * time.Microsecond)

	err := s.ListenAndServe()
	if err != nil {
		logger.Error("An error occurred starting HTTP listener %v", err)
	}

}
