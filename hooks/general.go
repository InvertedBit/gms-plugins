package hooks

import "context"

// HookHandler is a function that handles a hook event
type HookHandler func(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error)

type Hook struct {
	Name        string
	Description string
	Priority    int
	Handler     HookHandler
}
