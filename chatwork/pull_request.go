package chatwork

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ntrv/webhooks"
	"github.com/ntrv/webhooks/github"
)

func (c client) HandlePullRequest(payload interface{}, header webhooks.Header, w *http.ResponseWriter) {
	pl := payload.(github.PullRequestPayload)
	j, _ := json.Marshal(pl)
	fmt.Printf("%v\n", string(j))
	http.Error(*w, "Status OK", http.StatusOK)
}
