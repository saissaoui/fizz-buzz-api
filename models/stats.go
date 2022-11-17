package models

type StatisticsResponse struct {
	Stats []*RequestStats `json:"stats"`
}

type RequestStats struct {
	Request string `json:"request"`
	Count   int    `json:"count"`
}
