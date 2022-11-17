package stats

import "fizz-buzz-api/services/statistics"

type Response struct {
	Stats []*RequestStats `json:"stats"`
}

type RequestStats struct {
	Request string `json:"request"`
	Count   int    `json:"count"`
}

func NewResponse(stats []statistics.RequestStats) Response {
	responseStats := make([]*RequestStats, 0)
	for _, s := range stats {
		responseStats = append(responseStats, &RequestStats{
			Request: s.Request,
			Count:   s.Count,
		})
	}
	return Response{
		responseStats,
	}
}
