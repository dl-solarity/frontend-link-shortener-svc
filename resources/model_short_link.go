/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ShortLink struct {
	Key
	Attributes ShortLinkAttributes `json:"attributes"`
}
type ShortLinkResponse struct {
	Data     ShortLink `json:"data"`
	Included Included  `json:"included"`
}

type ShortLinkListResponse struct {
	Data     []ShortLink `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustShortLink - returns ShortLink from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustShortLink(key Key) *ShortLink {
	var shortLink ShortLink
	if c.tryFindEntry(key, &shortLink) {
		return &shortLink
	}
	return nil
}
