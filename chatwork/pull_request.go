package chatwork

import (
	"encoding/json"
	"fmt"

	"gopkg.in/go-playground/webhooks.v3"
	"gopkg.in/go-playground/webhooks.v3/github"
)

func (c client) HandlePullRequest(payload interface{}, header webhooks.Header) {
	pl := payload.(github.PullRequestPayload)
	j, _ := json.Marshal(pl)
	fmt.Printf("%v\n", string(j))
}
