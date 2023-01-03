package git

func ParseCloneSSH(value string) (CloneSSH, error) {
	return cloneSSH{
		value, // TODO: Validate that value matches expected pattern
	}, nil
}

type CloneSSH interface {
	String() string
}

type cloneSSH struct {
	value string
}

// ensure cloneSSH implements CloneSSH interface
var _ CloneSSH = cloneSSH{}

func (my cloneSSH) String() string {
	return my.value
}
