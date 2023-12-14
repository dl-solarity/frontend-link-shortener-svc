/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

import (
	"encoding/json"
	"time"
)

type ShortLinkAttributes struct {
	CreatedAt time.Time `json:"created_at"`
	// tool path
	Path string `json:"path"`
	// abi or hash data
	Value json.RawMessage `json:"value"`
}
