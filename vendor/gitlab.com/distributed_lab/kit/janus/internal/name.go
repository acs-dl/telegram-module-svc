package internal

import (
	"fmt"
	"regexp"
	"strings"
)

// GetName create URL-friendly name for service
// All parameters will be replaced with "x" in order to avoid adding the same paths differing only in the name of the parameters
func GetName(endpoint, method string) string {
	methodName := strings.ToLower(method)
	//TODO check how root works
	if len(endpoint) == 1 {
		return methodName
	}
	t := endpoint[1:]
	r := regexp.MustCompile(`{([a-z\s-_]+)}`)
	t = r.ReplaceAllString(t, "x")
	t = strings.Replace(t, ":", "-", -1)
	t = strings.Replace(t, "/", "-", -1)
	t = strings.Replace(t, "_", "-", -1)

	return fmt.Sprintf("%s-%s", methodName, t)
}

func GetRouteForGoji(endpoint string) string {
	sp := strings.Split(endpoint[1:], "/")
	var result string
	for _, s := range sp {
		if strings.HasPrefix(s, ":") {
			s = fmt.Sprintf("{%s}", strings.TrimPrefix(s, ":"))
		}
		result = fmt.Sprintf("%s/%s", result, s)
	}
	return result
}
