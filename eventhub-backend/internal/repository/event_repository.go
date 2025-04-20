package repository

type EventRepository interface {
	GetAll(userID uint) ([]EventResponse, []EventResponse, error)
	Join(userID, eventID uint) error
	Quit(userID, eventID uint) error
}
