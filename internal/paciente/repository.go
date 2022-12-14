package paciente

import (
	"errors"
	"prova/internal/domain"
	"prova/pkg/store/pacienteStore"
)

// declarando os metodos que o repository tem
type Repository interface {
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
	Update(id int, p domain.Paciente) (domain.Paciente, error) 
}

// fazendo a comunicação com o SGBD
type repository struct {
	store pacienteStore.Store
}

// instanciando um repository
func NewRepository(store pacienteStore.Store) Repository {
	return &repository{store}
}

// função getByID
func (r *repository) GetByID(id int) (domain.Paciente, error) {
	pReturn, err := r.store.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}

	return pReturn, nil
}

//funcao para criar um novo paciente
func (r *repository) Create(p domain.Paciente) (domain.Paciente, error) {
	if !r.validateRG(p.RG) {
		return domain.Paciente{}, errors.New("RG já existe")
	}
	return r.store.Create(p)
}

//funcao para validar RG
func (r *repository) validateRG(rg string) bool {
	for _, paciente := range r.store.GetAll() {
		if paciente.RG == rg {
			return false
		}
	}
	return true
}

//funcao para deletar um paciente
func (r *repository) Delete(id int) error {
	err := r.store.Delete(id)
	if err == nil {
		return nil
	}

	return err
}

//funcao update
func (r *repository) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	for _, paciente := range r.store.GetAll() {
		if paciente.Id == id {
			return r.store.Update(id, p)
		}
	}
	return domain.Paciente{}, errors.New("paciente nao encontrado")
}