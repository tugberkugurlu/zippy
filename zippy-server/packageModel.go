package main

import "time"

type packageModel struct {
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	CreatedOn time.Time `json:"createdOn"`
}

type packages []packageModel
