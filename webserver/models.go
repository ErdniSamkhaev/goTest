package main

type Status struct {
	Service string `json:"service"`
	Version string `json:"version"`
	Healthy bool   `json:"healthy"`
}

type Result struct {
	Address string `json:"address"`
	Alive   bool   `json:"alive"`
}
