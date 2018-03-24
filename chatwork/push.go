package chatwork

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ntrv/webhooks"
	"github.com/ntrv/webhooks/github"
)

func (c client) HandlePush(payload interface{}, header webhooks.Header, w http.ResponseWriter) {
	pl := payload.(github.PushPayload)
	j, _ := json.Marshal(pl)
	fmt.Printf("%v\n", string(j))
	http.Error(w, "Status OK", http.StatusOK)
}
