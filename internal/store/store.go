package store

type DBFactory interface {
	MeetingStore() MeetingStore
}
