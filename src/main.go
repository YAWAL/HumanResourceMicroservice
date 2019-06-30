package main

import (
	"github.com/YAWAL/HumanResourceMicroservice/src/app"
	"github.com/YAWAL/HumanResourceMicroservice/src/logging"
)

func main() {

	err := app.LoadApp()
	if err != nil {
		logging.Log.Errorf("Cannot load config: ", err)
	}

}
