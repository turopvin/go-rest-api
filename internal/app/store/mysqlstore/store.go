package mysqlstore

import "github.com/turopvin/go-rest-api/internal/app/auth"

//implementation for mysql store...
type Store struct {
}

func (s Store) UserRepository() auth.UserRepository {
	panic("implement me")
}
