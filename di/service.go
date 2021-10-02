package di

type PersonService struct {
	config     *Config
	repository *PersonRepository
}

func (service *PersonService) FindAll() []*Person {
	if service.config.Enabled {
		return service.repository.FindAll()
	}

	return []*Person{}
}

// NewPersonService 返回 persion 的配置, 但是需要 PersonRepository 传入
func NewPersonService(config *Config, repository *PersonRepository) *PersonService {
	return &PersonService{config: config, repository: repository}
}
