package graphql

import "context"

type Resolver struct{}

func (r *Resolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

func (r *Resolver) Reping(ctx context.Context) (string, error) {
	return "repung", nil
}
