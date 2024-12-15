package list

import "ipr/modules/daily_activity/repository"

type ResultItemsBuilder struct {
}

type DailyActivitiesCollection struct {
	Date            string              `json:"date"`
	DurationHours   int                 `json:"hours"`
	DurationMinutes int                 `json:"minutes"`
	Items           []DailyActivityItem `json:"items"`
}

type DailyActivityItem struct {
	ID              string  `json:"id"`
	ProjectName     *string `json:"projectName"`
	DurationHours   int     `json:"hours"`
	DurationMinutes int     `json:"minutes"`
	Description     string  `json:"description"`
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
			ID:              item.ID,
			ProjectName:     &item.ProjectName,
			DurationHours:   DurationHours,
			DurationMinutes: DurationMinutes,
			Description:     item.Description,
		})

		groupedDurationInSeconds[date] += durationInSeconds
	}

	var result []DailyActivitiesCollection
	for date, groupedItems := range grouped {
		groupDurationHours, groupDurationMinutes := toHoursWithMinutes(groupedDurationInSeconds[date])

		result = append(result, DailyActivitiesCollection{
			Date:            date,
			DurationHours:   groupDurationHours,
			DurationMinutes: groupDurationMinutes,
			Items:           groupedItems,
		})
	}

	return result
}

func toHoursWithMinutes(seconds int) (hours, minutes int) {
	hours = seconds / 3600
	minutes = (seconds % 3600) / 60

	return hours, minutes
}
