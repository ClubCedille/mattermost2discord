package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type message struct {
	Text     string `json:"text"`
	Username string `json:"user_name"`
	UserID   string `json:"user_id"`
}

// ProcessMattermost splits the payload
func ProcessMattermost(ctx *gin.Context) {

	var mmContent message

	ctx.BindJSON(&mmContent)

	//trim 2disc eventually
	msgfromMM := strings.Split(mmContent.Text, TriggerWordMattermost)[1]

	userfromMM := mmContent.Username

	discordBot(userfromMM, msgfromMM)
}
