package models

// Workshop represents data about workshops stored in database
type Workshop struct {
	WorkshopID   int
	WorkshopName string
	ChairmanID   int64
}
