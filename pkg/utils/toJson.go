package utils

import (
	"encoding/json"
	"log"
)

// ToJSON safely marshals an interface into a JSON byte slice
func ToJSON(v interface{}) []byte {
	bytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return []byte("{}")
	}
	return bytes
}
