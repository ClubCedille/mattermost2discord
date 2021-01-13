package main

import (
	"fmt"
	"os"

	"github.com/ClubCedille/mattermost2discord/api"
)

// PORT - constant for env port variable.
const PORT string = "PORT"

func main() {
	api.SetupServer().Run(fmt.Sprintf(":%s", os.Getenv(PORT)))
}
