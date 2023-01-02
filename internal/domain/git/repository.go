package git

type Repository interface {
	CloneHTTPS() string
	CloneSSH() SSH
	URL() string
}

type repository struct {
	cloneHTTPS string
	cloneSSH   SSH
	url        string
}

func (my repository) CloneHTTPS() string {
	return my.cloneHTTPS
}

func (my repository) CloneSSH() SSH {
	return my.cloneSSH
}

func (my repository) URL() string {
	return my.url
}

func NewRepository(cloneHTTPS string, cloneSSH SSH, url string) Repository {
	return repository{
		cloneHTTPS,
		cloneSSH,
		url,
	}
}
