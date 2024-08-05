package provider

import "github.com/jdejesus007/gt-api-project/internal/db"

// RepositoryProvider DependencyProvider is used to inject external resources into the request and
// can be retrieved from any request through its context using the api.ProviderKey
type RepositoryProvider interface {
	Database() db.DBExecutor
}
