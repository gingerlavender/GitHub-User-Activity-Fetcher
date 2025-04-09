package output

import (
	"encoding/json"
	"fmt"
	"github.com/gingerlavender/GitHub-User-Activity-Fetcher/internal/events"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
	"slices"
	"strings"
	"time"
)

var outputs = make(map[string]func(event events.Event) error)

func InitOutputs() {
	outputs["PushEvent"] = func(event events.Event) error {
		var payload events.PushEventPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for PushEvent: %w", err)
		}
		fmt.Printf("pushed %d commits to %s at %s (%s)\n", payload.Size, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["IssuesEvent"] = func(event events.Event) error {
		var payload events.IssueEventsPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for IssuesEvent: %w", err)
		}
		fmt.Printf("%s issue \"%s\" at %s (%s)\n", payload.Action, payload.Issue.Title, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["IssueCommentEvent"] = func(event events.Event) error {
		var payload events.IssueEventsPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for IssueCommentEvent: %w", err)
		}
		fmt.Printf("%s comment in issue \"%s\" at %s (%s)\n", payload.Action, payload.Issue.Title, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["PullRequestEvent"] = func(event events.Event) error {
		var payload events.PullRequestEventsPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for PullRequestEvent: %w", err)
		}
		fmt.Printf("%s pull request \"%s\" in %s at %s (%s)\n", payload.Action, payload.PullRequest.Title, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["PullRequestReviewEvent"] = func(event events.Event) error {
		var payload events.PullRequestEventsPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for PullRequestReviewEvent: %w", err)
		}
		fmt.Printf("%s review for pull request \"%s\" in %s at %s (%s)\n", payload.Action, payload.PullRequest.Title, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["PullRequestReviewCommentEvent"] = func(event events.Event) error {
		var payload events.PullRequestEventsPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for PullRequestReviewCommentEvent: %w", err)
		}
		fmt.Printf("%s comment in review for pull request \"%s\" in %s at %s (%s)\n", payload.Action, payload.PullRequest.Title, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["PullRequestReviewThreadEvent"] = func(event events.Event) error {
		var payload events.PullRequestEventsPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for PullRequestReviewThreadEvent: %w", err)
		}
		fmt.Printf("%s thread on review for pull request \"%s\" in %s at %s (%s)\n", payload.Action, payload.PullRequest.Title, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["CreateEvent"] = func(event events.Event) error {
		var payload events.CreateEventPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for CreateEvent: %w", err)
		}
		fmt.Printf("created %s %s %s at %s (%s)\n", payload.RefType, payload.Ref, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["ForkEvent"] = func(event events.Event) error {
		var payload events.ForkEventPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for ForkEvent: %w", err)
		}
		fmt.Printf("Forked from %s (forkee: %s) at %s (%s)\n", event.Repo.Name, payload.Forkee.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["WatchEvent"] = func(event events.Event) error {
		var payload events.WatchEventPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for WatchEvent: %w", err)
		}
		fmt.Printf("%s watching %s at %s (%s)\n", payload.Action, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["ReleaseEvent"] = func(event events.Event) error {
		var payload events.ReleaseEventPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for ReleaseEvent: %w", err)
		}
		fmt.Printf("%s release %s in %s at %s (%s)\n", payload.Action, payload.Release.Name, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["CommitCommentEvent"] = func(event events.Event) error {
		var payload events.CommitCommentEventPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for CommitCommentEvent: %w", err)
		}
		fmt.Printf("%s comment in %s at %s (%s)\n", payload.Action, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["DeleteEvent"] = func(event events.Event) error {
		var payload events.DeleteEventPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for DeleteEvent: %w", err)
		}
		fmt.Printf("deleted %s %s %s at %s (%s)\n", payload.RefType, payload.Ref, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["GollumEvent"] = func(event events.Event) error {
		var payload events.GollumEventPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for GollumEvent: %w", err)
		}
		fmt.Printf("updated %d pages in %s at %s (%s):\n", len(payload.Pages), event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		for i, page := range payload.Pages {
			fmt.Printf("%d.: %s page \"%s\"\n", i+1, page.Action, page.PageName)
		}
		return nil
	}
	outputs["MemberEvent"] = func(event events.Event) error {
		var payload events.MemberEventPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for PullMemberEvent: %w", err)
		}
		fmt.Printf("%s member %s in %s at %s (%s)\n", payload.Action, payload.Member.Login, event.Repo.Name, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
	outputs["SponsorshipEvent"] = func(event events.Event) error {
		var payload events.SponsorshipEventPayload
		if err := json.Unmarshal(event.Payload, &payload); err != nil {
			return fmt.Errorf("Error parsing payload for SponsorshipEvent: %w", err)
		}
		fmt.Printf("%s sponsorship (sponsor: %s, sponsee: %s) at %s (%s)\n", payload.Action, payload.Sponsorship.Sponsor.Login, payload.Sponsorship.Sponsee, event.CreatedAt.Format(time.DateTime), event.CreatedAt.Location().String())
		return nil
	}
}

func PrintEvents(eventsSlice *[]events.Event, period time.Duration, eventType string) error {
	InitOutputs()
	currentTime := time.Now()
	for _, event := range *eventsSlice {
		if currentTime.Sub(event.CreatedAt) > period || eventType != "" && eventType != event.Type {
			continue
		}
		if output, ok := outputs[event.Type]; ok {
			if err := output(event); err != nil {
				return fmt.Errorf("Error in printing event: %w", err)
			}
		} else {
			return fmt.Errorf("I don't know such event type yet: %s", event.Type)
		}
	}
	return nil
}

func GetEventsMap(eventsSlice *[]events.Event, period time.Duration, eventType string) map[string]map[string]int {
	eventsCount := make(map[string]map[string]int)
	currentTime := time.Now()
	for _, event := range *eventsSlice {
		if currentTime.Sub(event.CreatedAt) > period || eventType != "" && eventType != event.Type {
			continue
		}
		if _, ok := eventsCount[event.CreatedAt.Format(time.DateOnly)]; !ok {
			eventsCount[event.CreatedAt.Format(time.DateOnly)] = make(map[string]int)
		}
		eventsCount[event.CreatedAt.Format(time.DateOnly)][event.Type]++
	}
	return eventsCount
}

func DrawEventsPlot(eventsSlice *[]events.Event, period time.Duration, eventType string) error {
	var (
		makeVisible [events.EventsAmount]bool
		counts      [events.EventsAmount][]opts.BarData
		eventsCount = GetEventsMap(eventsSlice, period, eventType)
		dates       = make([]string, 0, len(eventsCount))
	)
	for i := range events.EventsAmount {
		counts[i] = make([]opts.BarData, len(eventsCount))
	}
	for date, _ := range eventsCount {
		dates = append(dates, date)
	}
	slices.SortFunc(dates, func(a, b string) int {
		return strings.Compare(a, b)
	})
	for i := range dates {
		for eventType, count := range eventsCount[dates[i]] {
			counts[events.GetIndex(eventType)][i] = opts.BarData{Value: count}
			makeVisible[events.GetIndex(eventType)] = true
		}
	}
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "User Activity"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Date"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "Events"}),
	)
	for i := range events.EventsAmount {
		if makeVisible[i] {
			bar.SetXAxis(dates).AddSeries(events.GetEventName(i), counts[i])
		}
	}
	file, err := os.Create("activity_" + time.Now().Format("2006-01-02_15-04-05") + ".html")
	if err != nil {
		return fmt.Errorf("Error creating file for plot: %w", err)
	}
	defer file.Close()
	bar.Render(file)
	return nil
}
