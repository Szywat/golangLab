package structures

import "time"

type Data struct {
	Date   time.Time
	Last   float32
	Volume int32
	Open   float32
	High   float32
	Low    float32
}
