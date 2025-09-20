package events

import (
	"calendarapi/internal/api/v1/events/dto"
	"sort"
	"time"

	"github.com/teambition/rrule-go"
)

type eventsService struct {
	repo Repository
}

func NewEventsService(repo Repository) Service {
	return &eventsService{repo: repo}
}

func (s *eventsService) GetAll() ([]*Events, error) {
	return s.repo.GetAll()
}

func (s *eventsService) GetByRange(startDate, endDate time.Time) ([]dto.GroupedEvent, error) {
	events, err := s.repo.GetByRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	var allExpanded []dto.CalendarEventResponse
	for _, e := range events {
		expanded := expandRecurringEvent(*e, startDate, endDate)
		allExpanded = append(allExpanded, expanded...)
	}

	grouped := groupEventsByDate(allExpanded)

	return grouped, nil
}

func groupEventsByDate(events []dto.CalendarEventResponse) []dto.GroupedEvent {
	grouped := make(map[string][]dto.CalendarEventResponse)

	/*for _, ev := range events {
		// Ambil hanya tanggal (yyyy-mm-dd) dari Start
		dateStr := ev.Start[:10]
		grouped[dateStr] = append(grouped[dateStr], ev)
	}*/

	for _, ev := range events {
		startDate := ev.Start[:10]
		endDate := ev.End[:10]

		if startDate == endDate {
			grouped[startDate] = append(grouped[startDate], ev)
		} else {
			firstPart := ev
			firstPart.End = startDate + "T23:59:59+07:00"
			grouped[startDate] = append(grouped[startDate], firstPart)

			secondPart := ev
			secondPart.Start = endDate + "T00:00:00+07:00"
			grouped[endDate] = append(grouped[endDate], secondPart)
		}
	}

	var result []dto.GroupedEvent
	for date, evs := range grouped {
		sort.Slice(evs, func(i, j int) bool {
			return evs[i].Start < evs[j].Start
		})

		result = append(result, dto.GroupedEvent{
			Date:   date,
			Events: evs,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Date < result[j].Date
	})

	return result
}

func expandRecurringEvent(e Events, startRange, endRange time.Time) []dto.CalendarEventResponse {
	var expanded []dto.CalendarEventResponse

	if e.RecurrenceRule == nil || *e.RecurrenceRule == "" {
		expanded = append(expanded, dto.CalendarEventResponse{
			ID:          int(e.ID),
			Title:       e.Title,
			Description: e.Description,
			Start:       e.StartTime.Format(time.RFC3339),
			End:         e.EndTime.Format(time.RFC3339),
			AllDay:      e.IsAllDay,
			Location:    derefString(e.Location),
		})
		return expanded
	}

	r, err := rrule.StrToRRule("DTSTART:" + e.StartTime.UTC().Format("20060102T150405Z") + "\nRRULE:" + *e.RecurrenceRule)
	if err != nil {
		expanded = append(expanded, dto.CalendarEventResponse{
			ID:          int(e.ID),
			Title:       e.Title,
			Description: e.Description,
			Start:       e.StartTime.Format(time.RFC3339),
			End:         e.EndTime.Format(time.RFC3339),
			AllDay:      e.IsAllDay,
			Location:    derefString(e.Location),
		})
		return expanded
	}

	occurances := r.Between(startRange, endRange, true)
	for _, occ := range occurances {
		end := occ.Add(e.EndTime.Sub(e.StartTime))
		startStr := occ.Format(time.RFC3339)
		endStr := end.Format(time.RFC3339)

		if e.IsAllDay {
			startStr = occ.Format("2006-01-02")
			endStr = occ.Format("2006-01-02")
		}

		expanded = append(expanded, dto.CalendarEventResponse{
			ID:          int(e.ID),
			Title:       e.Title,
			Description: e.Description,
			Start:       startStr,
			End:         endStr,
			AllDay:      e.IsAllDay,
			Location:    derefString(e.Location),
		})
	}

	return expanded
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
