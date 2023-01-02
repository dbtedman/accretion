package git

type Repository interface {
	CloneHTTPS() string
	CloneSSH() string
	URL() string
}

type repository struct {
	cloneHTTPS string
	cloneSSH   string
	url        string
}

func (my repository) CloneHTTPS() string {
	return my.cloneHTTPS
}

func (my repository) CloneSSH() string {
	return my.cloneSSH
}

func (my repository) URL() string {
	return my.url
}

func NewRepository(cloneHTTPS string, cloneSSH string, url string) Repository {
	return repository{
		cloneHTTPS,
		cloneSSH,
		url,
	}
}
