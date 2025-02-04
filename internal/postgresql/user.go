package postgresql

type User struct {
	q *Queries
}

func NewUser(d DBTX) *User {
	return &User{
		q: NewQueries(d),
	}
}
