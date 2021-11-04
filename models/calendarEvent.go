package models

type CalendarEvent struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Class string `json:"class"`
	Start int64  `json:"start"`
	End   int64  `json:"end"`
}
