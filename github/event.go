package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/webhooks.v3"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

type ProcessPayloadFunc func(payload interface{}, c echo.Context) error

type Webhook struct {
	eventName gh.Event
	eventFuncs map[gh.Event]ProcessPayloadFunc
}

func NewHook Webhook {
	return Webhook{
		eventName: nil,
		eventFuncs: map[gh.Event]ProcessPayloadFunc{},
	}
}

func (hook Webhook) RegisterEvents(fn ProcessPayloadFunc, events ...github.Event){
	for _, event := range events {
		hook.eventFuncs[event] = fn
	}
}

func (hook Webhook) ParsePayloadHandler(c echo.Context) error {
	c.Request().Method != "POST" {
		return echo.NewHTTPError(
			http.StatusMethodNotAllowed,
			fmt.Sprintf(
				"405 Method not allowed, attempt made using Method: %s",
				c.Request().Method,
			),
		)
	}

	event := c.Request().Header.Get("X-GitHub-Event")
	if len(event) == 0 {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			"Missing X-GitHub-Event Header",
		)
	}

	hook.eventName := gh.Event(event)
	fn, ok := hook.eventFuncs[hook.eventName]
	if !ok {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			fmt.Sprintf(
				"Webhook Event %s not registered, it is recommended to setup only events in github that will be registered in the webhook to avoid unnecessary traffic and reduce potential attack vectors.",
				event,
			),
		)
	}

	payload, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"Cannot read payload",
		)
	}

	return runProcess(payload, c)
}

func (hook Webhook) runProcess(payload []byte, c echo.Context) error {
	switch hook.eventName {
		case gh.CommitCommentEvent:
			var cc gh.CommitCommentPayload
			json.Unmarshal(payload, &cc)
			return fn(cc, c)
		case gh.CreateEvent:
			var cr gh.CreatePayload
			json.Unmarshal(payload, &cr)
			return fn(cr, c)
		case gh.DeleteEvent:
			var de gh.DeletePayload
			json.Unmarshal(payload, &de)
			return fn(d, c)
		case gh.DeploymentEvent:
			var dp gh.DeploymentPayload
			json.Unmarshal(payload, &dp)
			return fn(dp, c)
		case gh.DeploymentStatusEvent:
			var ds gh.DeploymentStatusPayload
			json.Unmarshal(payload, &ds)
			return fn(ds, c)
		case gh.ForkEvent:
			var fk gh.ForkPayload
			json.Unmarshal(payload, &fk)
			return fn(fk, c)
		case gh.GollumEvent:
			var gl gh.GollumPayload
			json.Unmarshal(payload, &gl)
			return fn(gl, c)
		case gh.InstallationEvent, gh.IntegrationInstallationEvent:
			var in gh.InstallationPayload
			json.Unmarshal(payload, &in)
			return fn(in, c)
		case gh.IssueCommentEvent:
			var ic gh.IssueCommentPayload
			json.Unmarshal(payload, &ic)
			return fn(ic, c)
		case gh.IssuesEvent:
			var is gh.IssuesPayload
			json.Unmarshal(payload, &is)
			return fn(is, c)
		case gh.LabelEvent:
			var lb gh.LabelPayload
			json.Unmarshal(payload, &lb)
			return fn(lb, c)
		case gh.MemberEvent:
			var me gh.MemberPayload
			json.Unmarshal(payload, &me)
			return fn(me, c)
		case gh.MembershipEvent:
			var ms gh.MembershipPayload
			json.Unmarshal(payload, &ms)
			return fn(ms, c)
		case gh.MilestoneEvent:
			var mi gh.MilestonePayload
			json.Unmarshal(payload, &mi)
			return fn(mi, c)
		case gh.OrganizationEvent:
			var or gh.OrganizationPayload
			json.Unmarshal(payload, &or)
			return fn(or, c)
		case gh.OrgBlockEvent:
			var ob gh.OrgBlockPayload
			json.Unmarshal(payload, &ob)
			return fn(ob, c)
		case gh.PageBuildEvent:
			var pa gh.PageBuildPayload
			json.Unmarshal(payload, &pa)
			return fn(pa, c)
		case gh.PingEvent:
			var pi gh.PingPayload
			json.Unmarshal(payload, &pi)
			return fn(pi, c)
		case gh.ProjectCardEvent:
			var pc gh.ProjectCardPayload
			json.Unmarshal(payload, &pc)
			return fn(pc, c)
		case gh.ProjectColumnEvent:
			var po gh.ProjectColumnPayload
			json.Unmarshal(payload, &po)
			return fn(po, c)
		case gh.ProjectEvent:
			var pe gh.ProjectPayload
			json.Unmarshal(payload, &pe)
			return fn(pe, c)
		case gh.PublicEvent:
			var pu gh.PublicPayload
			json.Unmarshal(payload, &pu)
			return fn(pu, c)
		case gh.PullRequestEvent:
			var pr gh.PullRequestPayload
			json.Unmarshal(payload, &pr)
			return fn(pr, c)
		case gh.PullRequestReviewEvent:
			var prr gh.PullRequestReviewPayload
			json.Unmarchal(payload, &prr)
			return fn(prr, c)
		case gh.PullRequestReviewCommentEvent:
			var prc gh.PullRequestReviewCommentPayload
			json.Unmarshal(payload, &prc)
			return fn(prc, c)
		case gh.PushEvent:
			var pu gh.PushPayload
			json.Unmarshal(payload, &pu)
			return fn(pu, c)
		case gh.ReleaseEvent:
			var re gh.ReleasePayload
			json.Unmarshal(payload, &re)
			return fn(re, c)
		case gh.RepositoryEvent:
			var rp gh.RepositoryPayload
			json.Unmarshal(payload, &rp)
			return fn(rp, c)
		case gh.StatusEvent:
			var st gh.StatusPayload
			json.Unmarshal(payload, &st)
			return fn(st, c)
		case gh.TeamEvent:
			var te gh.TeamPayload
			json.Unmarshal(payload, &te)
			return fn(te, c)
		case gh.TeamAddEvent:
			var ta gh.TeamAddPayload
			json.Unmarshal(payload, &ta)
			return fn(ta, c)
		case gh.WatchEvent:
			var wa gh.WatchPayload
			json.Unmarshal(payload, &wa)
			return fn(wa, c)
	}
	
}
