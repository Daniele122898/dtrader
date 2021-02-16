package models

type Volume struct {
	Timestamp         string `json:"timestamp"`
	Volume            string `json:"volume"`
	TransparentVolume string `json:"transparent_volume"`
}