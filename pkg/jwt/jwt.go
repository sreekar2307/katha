package jwt

import "context"

type JWT interface {
	Token(context.Context, map[string]any) (string, error)
	Validate(context.Context, string) (map[string]any, error)
}
