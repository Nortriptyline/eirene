package http

type HTTPError interface {
	error
	StatusCode() int
}
