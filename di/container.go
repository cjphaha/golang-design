package di

import (
	"go.uber.org/dig"
)

func main2() {
	container := dig.New()
	container.Provide(NewConfig)
	container.Provide(ConnectDatabase)
	container.Provide(NewPersonRepository)
	container.Provide(NewPersonService)
	container.Provide(NewServer)
	// 这里可以自动整理依赖，就不用手动整理或者使用过多的 init 了
	container.Invoke(func(server *Server) {
		server.Run()
	})
}
