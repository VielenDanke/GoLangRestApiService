package sqlstore_test

import "github.com/vielendanke/restful-service/internal/app/sqlstore"

func TestUserRepository_SaveUser(t *testing.T) {
	u := model.TestUser(t)

	store := newTestStore()
}