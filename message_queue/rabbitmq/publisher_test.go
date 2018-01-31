package rabbitmq

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"go_rtb/internal/message_queue/payload"
	"testing"
)

func TestPublisher_Publish(t *testing.T) {
	assertor := assert.New(t)

	conn := &Connector{}
	err := conn.Connect("amqp://rtb:rtb@127.0.0.1:5672/")
	assertor.Nil(err)

	pub := &Publisher{}
	pub.Init(conn)

	token := uuid.NewV4()

	callbackPayload := payload.SSPWinConfirmPayload{Token: token.String()}
	callbackPayloadJson, _ := json.Marshal(callbackPayload)
	err = pub.Publish(callbackPayloadJson, "win_confirm_queue")

	assertor.Nil(err)
}
