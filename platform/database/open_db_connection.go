package database

import "github.com/kazimovzaman2/online_book/app/queries"

type Queries struct {
	*queries.BookQueries
}

func OpenDBConnection() (*Queries, error) {
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		BookQueries: &queries.BookQueries{DB: db},
	}, nil
}
