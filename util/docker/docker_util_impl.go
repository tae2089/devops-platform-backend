package docker

var _ (Util) = (*dockerUtil)(nil)

type dockerUtil struct{}

func (d *dockerUtil) GetFileByLang(lang string) string {

	switch lang {
	case "java":
		return dockerJava
	case "node":
		return dockerNode
	case "go":
		return dockerGolang
	case "fastapi":
		return dockerFastAPI
	default:
		return ""
	}
}
