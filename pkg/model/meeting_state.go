package model

type MeetingState string

const (
	Booked     MeetingState = "1"
	Canceled   MeetingState = "2"
	InProgress MeetingState = "3"
	Ended      MeetingState = "4"
)
