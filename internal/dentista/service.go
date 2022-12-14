package dentista

import (
	//"errors"
	"prova/internal/domain"
)

// declarando os metodos que o service tem
type Service interface {
	GetByID(id int) (domain.Dentista, error)
	Create(d domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
	Update(id int, d domain.Dentista) (domain.Dentista, error) 
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
func (s *service) GetByID(id int) (domain.Dentista, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

//funcao para criar um novo dentista
func (s *service) Create(d domain.Dentista) (domain.Dentista, error) {
	d, err := s.r.Create(d)
	if err != nil {
		return domain.Dentista{}, err
	}
	return d, nil
}

//funcao para deletar um dentista
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

//funcao update
func (s *service) Update(id int, u domain.Dentista) (domain.Dentista, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	if u.Nome != "" {
		d.Nome = u.Nome
	}
	if u.Sobrenome != "" {
		d.Sobrenome = u.Sobrenome
	}
	if u.Matricula != "" {
		d.Matricula = u.Matricula
	}
	d, err = s.r.Update(id, d)
	if err != nil {
		return domain.Dentista{}, err
	}
	return d, nil
}