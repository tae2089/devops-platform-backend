package docker

import "sync"

type Util interface {
	GetJava() string
	GetGolang() string
	GetJavaScript() string
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
