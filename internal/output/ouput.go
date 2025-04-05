package output

import (
	"encoding/json"
	"fmt"
	"gh-api/internal/events"
	"time"
)

func PrintEvents(eventsSlice []events.Event, period time.Duration) error {
	currentTime := time.Now()
	for _, event := range eventsSlice {
		if currentTime.Sub(event.CreatedAt) > period {
			continue
		}
		switch event.Type {
		case "PushEvent":
			var payload events.PushEventPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for PushEvent: %w", err)
			}
			fmt.Printf("pushed %d commits to %s at %s (%s)\n", payload.Size, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "IssuesEvent":
			var payload events.IssueEventsPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for IssuesEvent: %w", err)
			}
			fmt.Printf("%s issue \"%s\" at %s (%s)\n", payload.Action, payload.Issue.Title, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "IssueCommentEvent":
			var payload events.IssueEventsPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for IssueCommentEvent: %w", err)
			}
			fmt.Printf("%s comment in issue \"%s\" at %s (%s)\n", payload.Action, payload.Issue.Title, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "PullRequestEvent":
			var payload events.PullRequestEventsPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for PullRequestEvent: %w", err)
			}
			fmt.Printf("%s pull request \"%s\" in %s at %s (%s)\n", payload.Action, payload.PullRequest.Title, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "PullRequestReviewEvent":
			var payload events.PullRequestEventsPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for PullRequestReviewEvent: %w", err)
			}
			fmt.Printf("%s review for pull request \"%s\" in %s at %s (%s)\n", payload.Action, payload.PullRequest.Title, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "PullRequestReviewCommentEvent":
			var payload events.PullRequestEventsPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for PullRequestReviewCommentEvent: %w", err)
			}
			fmt.Printf("%s comment in review for pull request \"%s\" in %s at %s (%s)\n", payload.Action, payload.PullRequest.Title, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "PullRequestReviewThreadEvent":
			var payload events.PullRequestEventsPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for PullRequestReviewThreadEvent: %w", err)
			}
			fmt.Printf("%s thread on review for pull request \"%s\" in %s at %s (%s)\n", payload.Action, payload.PullRequest.Title, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "CreateEvent":
			var payload events.CreateEventPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for CreateEvent: %w", err)
			}
			fmt.Printf("created %s %s %s at %s (%s)\n", payload.RefType, payload.Ref, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "ForkEvent":
			var payload events.ForkEventPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for ForkEvent: %w", err)
			}
			fmt.Printf("Forked from %s (forkee: %s) at %s (%s)\n", event.Repo.Name, payload.Forkee.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "WatchEvent":
			var payload events.WatchEventPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for WatchEvent: %w", err)
			}
			fmt.Printf("%s watching %s at %s (%s)\n", payload.Action, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "ReleaseEvent":
			var payload events.ReleaseEventPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for ReleaseEvent: %w", err)
			}
			fmt.Printf("%s release %s in %s at %s (%s)\n", payload.Action, payload.Release.Name, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "CommitCommentEvent":
			var payload events.CommitCommentEventPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for CommitCommentEvent: %w", err)
			}
			fmt.Printf("%s comment in %s at %s (%s)\n", payload.Action, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "DeleteEvent":
			var payload events.DeleteEventPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for DeleteEvent: %w", err)
			}
			fmt.Printf("deleted %s %s %s at %s (%s)\n", payload.RefType, payload.Ref, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "GollumEvent":
			var payload events.GollumEventPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for GollumEvent: %w", err)
			}
			fmt.Printf("updated %d pages in %s at %s (%s):\n", len(payload.Pages), event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
			for i, page := range payload.Pages {
				fmt.Printf("%d.: %s page \"%s\"\n", i+1, page.Action, page.PageName)
			}
		case "MemberEvent":
			var payload events.MemberEventPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for PullMemberEvent: %w", err)
			}
			fmt.Printf("%s member %s in %s at %s (%s)\n", payload.Action, payload.Member.Login, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		case "SponsorshipEvent":
			var payload events.SponsorshipEventPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				return fmt.Errorf("Error parsing payload for SponsorshipEvent: %w", err)
			}
			fmt.Printf("%s sponsorship (sponsor: %s, sponsee: %s) at %s (%s)\n", payload.Action, payload.Sponsorship.Sponsor.Login, payload.Sponsorship.Sponsee, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		default:
			fmt.Printf("I don't know this event type yet: %s\n", event.Type)
		}
	}
	return nil
}
