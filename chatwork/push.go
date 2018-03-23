package chatwork

import (
	"encoding/json"
	"fmt"

	"gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func HandlePush(payload interface{}, header webhooks.Header) {
	pl := payload.(github.PushPayload)
	j, _ := json.Marshal(pl)
	fmt.Printf("%v\n", string(j))
}
