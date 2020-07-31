package app

import "time"

type Stats struct {
	Name      string    `json:"name"`
	Port      string    `json:"port"`
	Proxy     string    `json:"proxy"`
	Version   string    `json:"version"`
	StartedAt time.Time `json:"startedat,string"`
}

func GetStats(s *Stats) *Stats {

	return s
}
