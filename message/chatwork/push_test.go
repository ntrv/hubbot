package chatwork

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func TestPushMsgPush(t *testing.T) {
	var pl github.PushPayload

	raw, err := ioutil.ReadFile("../../example/push.json")
	assert.NoError(t, err)

	assert.NoError(t, json.Unmarshal(raw, &pl))

	msg, err := PushMsg(pl)
	assert.NoError(t, err)

	fmt.Println(msg)
}
