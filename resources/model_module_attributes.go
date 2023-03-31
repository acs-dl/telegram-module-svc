/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ModuleAttributes struct {
	// indicates whether module (gitlab, telegram etc.) or service (unverified, role etc.)
	IsModule bool `json:"is_module"`
	// Module url
	Link string `json:"link"`
	// Module name
	Name string `json:"name"`
	// Module prefix to use in FE
	Prefix string `json:"prefix"`
	// Module name to use in FE
	Title string `json:"title"`
	// Module topic for sender and others
	Topic string `json:"topic"`
}
