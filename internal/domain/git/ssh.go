package git

func NewSSH(value string) (SSH, error) {
	return ssh{
		value, // TODO: Validate that value matches expected pattern
	}, nil
}

type SSH interface {
	ToString() string
}

type ssh struct {
	value string
}

func (s ssh) ToString() string {
	return s.value
}
