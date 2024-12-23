package list

import (
	"ipr/modules/daily_activity/repository"
	"sort"
	"time"
)

type ResultItemsBuilder struct {
}

type DailyActivitiesCollection struct {
	Date            string              `json:"date"`
	DurationHours   int                 `json:"hours"`
	DurationMinutes int                 `json:"minutes"`
	Activity        []DailyActivityItem `json:"activities"`
}

type DailyActivityItem struct {
	Id              string `json:"id"`
	ProjectName     string `json:"projectName"`
	DurationHours   int    `json:"hours"`
	DurationMinutes int    `json:"minutes"`
	Description     string `json:"description"`
	StartAt         string `json:"startAt"`
	EndAt           string `json:"endAt"`
}

func NewResultItemsBuilder() *ResultItemsBuilder {
	return &ResultItemsBuilder{}
}

func (b *ResultItemsBuilder) build(items []repository.QueryItem) []DailyActivitiesCollection {
	grouped := make(map[string][]DailyActivityItem)
	groupedDurationInSeconds := make(map[string]int)

	for _, item := range items {
		date := item.StartAt.Format("02.01.2006")

		if _, exists := grouped[date]; !exists {
			grouped[date] = []DailyActivityItem{}
		}

		durationInSeconds := int(item.EndAt.Sub(item.StartAt).Seconds())
		DurationHours, DurationMinutes := toHoursWithMinutes(durationInSeconds)

		grouped[date] = append(grouped[date], DailyActivityItem{
			Id:              item.ID,
			ProjectName:     item.ProjectName,
			DurationHours:   DurationHours,
			DurationMinutes: DurationMinutes,
			Description:     item.Description,
			StartAt:         item.StartAt.Format("15:04"),
			EndAt:           item.EndAt.Format("15:04"),
		})

		groupedDurationInSeconds[date] += durationInSeconds
	}

	var result []DailyActivitiesCollection
	for date, activities := range grouped {
		sort.Slice(activities, func(i, j int) bool {
			t1, _ := time.Parse("15:04", activities[i].StartAt)
			t2, _ := time.Parse("15:04", activities[j].StartAt)
			return t1.Before(t2)
		})

		groupDurationHours, groupDurationMinutes := toHoursWithMinutes(groupedDurationInSeconds[date])

		result = append(result, DailyActivitiesCollection{
			Date:            date,
			DurationHours:   groupDurationHours,
			DurationMinutes: groupDurationMinutes,
			Activity:        activities,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		t1, _ := time.Parse("02.01.2006", result[i].Date)
		t2, _ := time.Parse("02.01.2006", result[j].Date)
		return t2.Before(t1)
	})

	return result
}

func toHoursWithMinutes(seconds int) (hours, minutes int) {
	hours = seconds / 3600
	minutes = (seconds % 3600) / 60

	return hours, minutes
}
