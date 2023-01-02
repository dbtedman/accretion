package git

type SSH interface {
	ToString() string
}

type ssh struct {
	value string
}

func (s ssh) ToString() string {
	return s.value
}

func NewSSH(value string) SSH {
	return ssh{
		value,
	}
}
