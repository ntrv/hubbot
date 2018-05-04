package github

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

func (hook Webhook) runProcess(payload []byte, fn ProcessPayloadFunc, c echo.Context) error {
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
		return fn(de, c)
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
		json.Unmarshal(payload, &prr)
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
	return echo.NewHTTPError(
		http.StatusInternalServerError,
		"Ending GitHub Event",
	)
}
