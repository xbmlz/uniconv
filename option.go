package uniconv

const (
	defaultHost = "localhost"
	defaultPort = 2002
)

type Office struct {
	path string
	host string
	port int
}

type Option func(*Office)

func WithPath(path string) Option {
	return func(c *Office) {
		c.path = path
	}
}

func WithHost(host string) Option {
	return func(c *Office) {
		c.host = host
	}
}

func WithPort(port int) Option {
	return func(c *Office) {
		c.port = port
	}
}
