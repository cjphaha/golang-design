package di

import "database/sql"

type PersonRepository struct {
	database *sql.DB
}

// NewPersonRepository 返回一个待创建的数据库连接
func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{database: database}
}

func (repository *PersonRepository) FindAll() []*Person {
	rows, _ := repository.database.Query(
		`SELECT id, name, age FROM people;`,
	)
	defer rows.Close()
	people := []*Person{}
	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)
		rows.Scan(&id, &name, &age)
		people = append(people, &Person{
			Id:   id,
			Name: name,
			Age:  age,
		})
	}
	return people
}
