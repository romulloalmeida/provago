package paciente

import (
	//"errors"
	"prova/internal/domain"
)

// declarando os metodos que o service tem
type Service interface {
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
	Update(id int, p domain.Paciente) (domain.Paciente, error) 
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
func (s *service) GetByID(id int) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

//funcao para criar um novo paciente
func (s *service) Create(p domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

//funcao para deletar um paciente
func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

//funcao update
func (s *service) Update(id int, u domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	if u.Nome != "" {
		p.Nome = u.Nome
	}
	if u.Sobrenome != "" {
		p.Sobrenome = u.Sobrenome
	}
	if u.RG != "" {
		p.RG = u.RG
	}
	if u.DataCadastro != "" {
		p.DataCadastro = u.DataCadastro
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}