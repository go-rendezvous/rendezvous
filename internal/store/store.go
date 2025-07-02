package store

// Factory to make stores for database
type DBFactory interface {
	MeetingStore() MeetingStore
}
