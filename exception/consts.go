package exception

import "errors"

var (
	AlreadyHookExistsError = errors.New("Hook is already exists ")
	NotFoundHooks          = errors.New("Not Found Hooks")
	NotFoundFile           = errors.New("Not Found File")
	NotFoundEntity         = errors.New("Not Found Entity")
	Forbidden              = errors.New("Forbidden")
	NotAuthorized          = errors.New("Not authorized")
)
