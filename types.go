package caddyapi

type Config struct {
	Apps Apps `json:"apps"`
}

type Apps struct {
	HTTP HTTP `json:"http"`
}

type HTTP struct {
	Servers map[string]Server `json:"servers"`
}

type Server struct {
	Listen []string `json:"listen"`
	Routes []Route  `json:"routes"`
}

type Route struct {
	Handle   []Handle `json:"handle"`
	Match    []Match  `json:"match"`
	Terminal bool     `json:"terminal"`
}

type Match struct {
	Host []string `json:"host,omitempty"`
	Path []string `json:"path,omitempty"`
}

type Handle struct {
	Handler       string         `json:"handler"`
	Routes        []Route        `json:"routes,omitempty"`
	Headers       *Headers       `json:"headers,omitempty"`
	Upstreams     []Upstream     `json:"upstreams,omitempty"`
	LoadBalancing *LoadBalancing `json:"load_balancing,omitempty"`
	Transport     *Transport     `json:"transport,omitempty"`
	StatusCode    int            `json:"status_code,omitempty"`
}

type Headers struct {
	Request *Request `json:"request,omitempty"`
}

type Request struct {
	Set map[string][]string `json:"set,omitempty"`
}

type Upstream struct {
	Dial string `json:"dial"`
}

type LoadBalancing struct {
	SelectionPolicy SelectionPolicy `json:"selection_policy"`
}

type SelectionPolicy struct {
	Policy string `json:"policy"`
}

type Transport struct {
	Protocol string `json:"protocol"`
	TLS      *TLS   `json:"tls"`
}

type TLS struct{}
