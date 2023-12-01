package main

import (
	"bytes"
	"fmt"
	"time"
)

// Event служит для представления событий в календаре
type Event struct {
	ID      string `json:"id"`
	Date    Date   `json:"date"`
	Content string `json:"content"`
}

// Date изпользуется для работы с датами и реалирует json.Marshaler
type Date struct {
	time.Time
}

const dateLayout = "2006-01-02"

// UnmarshalJSON имплементирует десериализацию для json.Marshaler
func (d *Date) UnmarshalJSON(data []byte) error {
	data = bytes.Trim(data, "\"")
	if len(data) == 0 {
		return fmt.Errorf("indalid date format: '%s'", string(data))
	}

	date, err := time.Parse(dateLayout, string(data))
	if err != nil {
		return fmt.Errorf("indalid date format: '%s'", string(data))
	}
	d.Time = date
	return nil
}

// MarshalJSON имплементирует сериализацию для json.Marshaler
func (d *Date) MarshalJSON() ([]byte, error) {
	dateStr := d.Format(dateLayout)
	stamp := fmt.Sprintf(`%q`, dateStr)
	return []byte(stamp), nil
}
