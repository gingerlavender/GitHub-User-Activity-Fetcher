package events

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	PushEvent = iota
	IssuesEvent
	IssueCommentEvent
	PullRequestEvent
	PullRequestReviewEvent
	PullRequestReviewCommentEvent
	PullRequestReviewThreadEvent
	CreateEvent
	ForkEvent
	WatchEvent
	ReleaseEvent
	CommitCommentEvent
	DeleteEvent
	GollumEvent
	MemberEvent
	SponsorshipEvent
	EventsAmount = 16
)

type Repo struct {
	Name string `json:"name"`
}

type PullRequest struct {
	Title string `json:"title"`
}

type Comment struct {
	Body string `json:"body"`
}

type Issue struct {
	Title string `json:"title"`
}

type Event struct {
	CreatedAt time.Time       `json:"created_at"`
	Type      string          `json:"type"`
	Repo      Repo            `json:"repo"`
	Payload   json.RawMessage `json:"payload"`
}

type PushEventPayload struct {
	Size int `json:"size"`
}

type IssueEventsPayload struct {
	Action string `json:"action"`
	Issue  Issue  `json:"issue"`
}

type PullRequestEventsPayload struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
}

type GollumEventPayload struct {
	Pages []struct {
		Action   string `json:"action"`
		PageName string `json:"page_name"`
	} `json:"pages"`
}

type MemberEventPayload struct {
	Action string `json:"action"`
	Member struct {
		Login string `json:"login"`
	} `json:"member"`
}

type SponsorshipEventPayload struct {
	Action      string `json:"action"`
	Sponsorship struct {
		Sponsor struct {
			Login string `json:"login"`
		} `json:"sponsor"`
		Sponsee struct {
			Login string `json:"login"`
		} `json:"sponsee"`
	} `json:"sponsorship"`
}

type CreateEventPayload struct {
	Ref     string `json:"ref"`
	RefType string `json:"ref_type"`
}

type ForkEventPayload struct {
	Forkee Repo `json:"forkee"`
}

type WatchEventPayload struct {
	Action string `json:"action"`
}

type ReleaseEventPayload struct {
	Action  string `json:"action"`
	Release struct {
		Name string `json:"name"`
	} `json:"release"`
}

type CommitCommentEventPayload struct {
	Action string `json:"action"`
}

type DeleteEventPayload struct {
	Ref     string `json:"ref"`
	RefType string `json:"ref_type"`
}

func FetchEvents(username string) ([]Event, error) {
	resp, err := http.Get("https://api.github.com/users/" + username + "/events")
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch: %w", err)
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status of response is not OK, but following: %s", resp.Status)
	}
	defer resp.Body.Close()
	var events []Event
	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&events); err != nil {
		return nil, fmt.Errorf("Unable to parse: %w", err)
	}
	return events, nil
}

func GetEventName(eventConstant int) string {
	eventNames := [EventsAmount]string{"PushEvent", "IssuesEvent", "IssueCommentEvent", "PullRequestEvent", "PullRequestReviewEvent", "PullRequestReviewCommentEvent",
		"PullRequestReviewThreadEvent", "CreateEvent", "ForkEvent", "WatchEvent", "ReleaseEvent", "CommitCommentEvent", "DeleteEvent", "GollumEvent", "MemberEvent", "SponsorshipEvent"}
	if eventConstant < 0 || eventConstant >= EventsAmount {
		panic("Unknown event constant")
	}
	return eventNames[eventConstant]
}

func GetIndex(eventType string) int {
	switch eventType {
	case "PushEvent":
		return PushEvent
	case "IssuesEvent":
		return IssuesEvent
	case "IssueCommentEvent":
		return IssueCommentEvent
	case "PullRequestEvent":
		return PullRequestEvent
	case "PullRequestReviewEvent":
		return PullRequestReviewEvent
	case "PullRequestReviewCommentEvent":
		return PullRequestReviewCommentEvent
	case "PullRequestReviewThreadEvent":
		return PullRequestReviewThreadEvent
	case "CreateEvent":
		return CreateEvent
	case "ForkEvent":
		return ForkEvent
	case "WatchEvent":
		return WatchEvent
	case "ReleaseEvent":
		return ReleaseEvent
	case "CommitCommentEvent":
		return CommitCommentEvent
	case "DeleteEvent":
		return DeleteEvent
	case "GollumEvent":
		return GollumEvent
	case "MemberEvent":
		return MemberEvent
	case "SponsorshipEvent":
		return SponsorshipEvent
	default:
		panic("Unknown event type")
	}
}
