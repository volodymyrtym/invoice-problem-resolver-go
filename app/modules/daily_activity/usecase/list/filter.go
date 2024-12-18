package list

import (
	"ipr/modules/daily_activity/repository"
	"time"
)

var limitDays = 31

func filterFabric(q Query) *repository.GetListFilter {
	var filter repository.GetListFilter
	filter.UserID = q.UserID

	now := time.Now()

	if q.Range != nil {
		switch *q.Range {
		case "currentMonth":
			startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			endOfMonth := startOfMonth.AddDate(0, 1, -1)
			filter.DateRange = &repository.DateRange{
				StartDate: startOfMonth,
				EndDate:   endOfMonth,
			}

		case "prevMonth":
			startOfPrevMonth := time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, now.Location())
			endOfPrevMonth := startOfPrevMonth.AddDate(0, 1, -1)
			filter.DateRange = &repository.DateRange{
				StartDate: startOfPrevMonth,
				EndDate:   endOfPrevMonth,
			}
		}
	} else {
		filter.Limit = &limitDays
		filter.Page = q.Page
	}

	return &filter
}
