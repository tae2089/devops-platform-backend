package docker

import "sync"

type Util interface {
	GetFileByLang(lang string) string
}

var (
	util Util
	once sync.Once
)

func NewDockerUtil() Util {
	once.Do(func() {
		util = &dockerUtil{}
	})
	return util
}
