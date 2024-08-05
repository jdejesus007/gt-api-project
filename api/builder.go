package api

import "github.com/jdejesus007/gt-api-project/api/provider"

// Describe builder / creator
type Builder interface {
	WithRepositoryProvider(provider.RepositoryProvider) Builder
	// Finalize creates an instance of an API. If the dependency provider is not
	// set, Finalize will panic
	Finalize() API
}

type builder struct {
	provider provider.RepositoryProvider
}

func NewBuilder() Builder {
	return &builder{}
}

func (b *builder) WithRepositoryProvider(p provider.RepositoryProvider) Builder {
	b.provider = p
	return b
}

func (b *builder) Finalize() API {
	if b.provider == nil {
		panic("no dependency provider set")
	}
	return &Implementation{
		repositoryProvider: b.provider,
	}
}
