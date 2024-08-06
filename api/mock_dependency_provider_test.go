package api_test

import "github.com/jdejesus007/gt-api-project/internal/db"

type mockDependencyProvider struct {
	database db.DBExecutor
}

func newMockProvider() *mockDependencyProvider {
	return &mockDependencyProvider{
		database: &mockDBExecutor{},
	}
}

func (d *mockDependencyProvider) Database() db.DBExecutor {
	return d.database
}
