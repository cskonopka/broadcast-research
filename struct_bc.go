package main

type BC struct {
	Source          string
	Destination     string
	OutputDirectory string
	Operation       BroadcastResearch
}

type BroadcastResearch interface {
	Create(bc *BC, message string) error
}
