package git

func ParseCloneSSH(value string) (CloneSSH, error) {
	return cloneSSH{
		value,
	}, nil
}

type CloneSSH interface {
	String() string
}

type cloneSSH struct {
	value string
}

// ensure cloneSSH implements CloneSSH interface.
var _ CloneSSH = cloneSSH{}

func (my cloneSSH) String() string {
	return my.value
}
