package consulta

import (
	//"errors"
	"prova/internal/domain"
)

// declarando os metodos que o service tem
type Service interface {
	GetByID(id int) (domain.Consulta, error)
	Create(c domain.Consulta) (domain.Consulta, error)
	Delete(id int) error
	Update(id int, c domain.Consulta) (domain.Consulta, error) 
}

// fazendo a comunicação com o repository
type service struct {
	r Repository
}

// instanciando o service
func NewService(r Repository) Service {
	return &service{r}
}

// função getByID
func (s *service) GetByID(id int) (domain.Consulta, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Consulta{}, err
	}
	return p, nil
}

//funcao para criar um novo consulta
func (s *service) Create(c domain.Consulta) (domain.Consulta, error) {
	c, err := s.r.Create(c)
	if err != nil {
		return domain.Consulta{}, err
	}
	return c, nil
}

//funcao para deletar um consulta
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

//funcao update
func (s *service) Update(id int, u domain.Consulta) (domain.Consulta, error) {
	c, err := s.r.GetByID(id)
	if err != nil {
		return domain.Consulta{}, err
	}
	if u.Paciente != "" {
		c.Paciente = u.Paciente
	}
	if u.Dentista != "" {
		c.Dentista = u.Dentista
	}
	if u.DataHora != "" {
		c.DataHora = u.DataHora
	}
	c, err = s.r.Update(id, c)
	if err != nil {
		return domain.Consulta{}, err
	}
	return c, nil
}