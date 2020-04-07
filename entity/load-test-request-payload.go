package entity

type LoadTestRequestPayload struct {
	NumberTests	int `json:"numberTests"`
	IOIncluded bool `json:"ioIncluded"`
}