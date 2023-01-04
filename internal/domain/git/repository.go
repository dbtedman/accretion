package git

import "net/url"

func NewRepository(cloneHTTPS string, cloneSSH CloneSSH, url url.URL) Repository {
	return repository{
		cloneHTTPS,
		cloneSSH,
		url,
	}
}

type Repository interface {
	CloneHTTPS() string
	CloneSSH() CloneSSH
	URL() url.URL
}

type repository struct {
	cloneHTTPS string
	cloneSSH   CloneSSH
	url        url.URL
}

// ensure repository implements Repository interface.
var _ Repository = repository{}

func (my repository) CloneHTTPS() string {
	return my.cloneHTTPS
}

func (my repository) CloneSSH() CloneSSH {
	return my.cloneSSH
}

func (my repository) URL() url.URL {
	return my.url
}
