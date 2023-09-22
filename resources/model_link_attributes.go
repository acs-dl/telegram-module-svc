/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type LinkAttributes struct {
	// indicates whether link exists
	IsExists bool `json:"is_exists"`
	// link to repository or group
	Link       string `json:"link"`
	Submodules []Chat `json:"submodules"`
}
