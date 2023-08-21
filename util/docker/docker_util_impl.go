package docker

type dockerUtil struct{}

var _ (Util) = (*dockerUtil)(nil)

func (d *dockerUtil) GetJava() string {
	return dockerJava
}

func (d *dockerUtil) GetGolang() string {
	return dockerGolang
}

func (d *dockerUtil) GetJavaScript() string {
	return dockerNode
}
