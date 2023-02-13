package internal

type Service struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Active  bool   `json:"active"`
	Proxy   Proxy  `json:"proxy"`
}

type Proxy struct {
	// Enabling this flag instructs Janus that when proxying this API,
	// it should always include the matching URI prefix in the upstream request's URI
	AppendPath bool      `json:"append_path"`
	ListenPath string    `json:"listen_path"`
	Upstreams  Upstreams `json:"upstreams"`
	Methods    []string  `json:"methods"`
}
type Upstreams struct {
	Balancing string   `json:"balancing"`
	Targets   []Target `json:"targets"`
}

type Target struct {
	Target string `json:"target"`
	Weight int    `json:"weight"`
}
