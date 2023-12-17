package docker

import (
	"io"
	"net/http"

	"github.com/tae2089/bob-logging/logger"
)

var _ (Util) = (*dockerUtil)(nil)

type dockerUtil struct{}

func (d *dockerUtil) GetFileByLang(lang string) string {
	switch lang {
	case "java":
		return d.getDockerfile("https://raw.githubusercontent.com/tae2089/docker/main/gradle-spring/Dockerfile")
	case "nodejs":
		return d.getDockerfile("https://raw.githubusercontent.com/tae2089/docker/main/nodejs/Dockerfile")
	case "go":
		return d.getDockerfile("https://raw.githubusercontent.com/tae2089/docker/main/go/Dockerfile")
	case "python":
		return d.getDockerfile("https://raw.githubusercontent.com/tae2089/docker/main/python/Dockerfile")
	default:
		return "invalid language"
	}
}

func (d *dockerUtil) getDockerfile(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err)
		return ""
	}
	defer resp.Body.Close()
	// 결과 출력
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		return ""
	}
	return string(data)
}
