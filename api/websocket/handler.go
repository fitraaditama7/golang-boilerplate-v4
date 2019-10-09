package websocket

import (
	"context"
	"golang-websocket/api/helper"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/gin-gonic/gin"
)

type WebsocketHandler struct {
	App *firebase.App
}

func (w WebsocketHandler) SendAndroidMessage(c *gin.Context) {

	var res = c.Writer
	ctx := context.Background()
	client, err := w.App.Messaging(ctx)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
	}

	registrationToken := c.Request.FormValue("token")

	message := &messaging.Message{
		Data: map[string]string{
			"score": "850",
			"time":  "2:45",
		},
		Token: registrationToken,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		helper.ErrorCustomStatus(res, http.StatusInternalServerError, err.Error())
	}
	helper.Responses(res, http.StatusOK, "Successfully Send Message", response)
}
