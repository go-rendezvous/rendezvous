package store

// Factory to make stores for database
type DBFactory interface {
	Migrate() error
	MeetingStore() MeetingStore
	UserStore() UserStore
}
