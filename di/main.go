package di

func main() {
	config := NewConfig()

	db, err := ConnectDatabase(config)

	if err != nil {
		panic(err)
	}

	personRepository := NewPersonRepository(db)

	personService := NewPersonService(config, personRepository)

	server := NewServer(config, personService)

	server.Run()
}
