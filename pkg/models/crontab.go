package models

// Crontab defines structure for crontab
type Crontab struct {
	Records []Record
}

// Record defines single record at the crontab
type Record struct {
	Schedule string
	Command string
}