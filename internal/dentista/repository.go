package dentista

import (
	"errors"
	"prova/internal/domain"
	"prova/pkg/store/dentistaStore"
)

// declarando os metodos que o repository tem
type Repository interface {
	GetByID(id int) (domain.Dentista, error)
	Create(d domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
	Update(id int, d domain.Dentista) (domain.Dentista, error) 
}

// fazendo a comunicação com o SGBD
type repository struct {
	store dentistaStore.Store
}

// instanciando um repository
func NewRepository(store dentistaStore.Store) Repository {
	return &repository{store}
}

// função getByID
func (r *repository) GetByID(id int) (domain.Dentista, error) {
	pReturn, err := r.store.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}

	return pReturn, nil
}

//funcao para criar um novo dentista
func (r *repository) Create(d domain.Dentista) (domain.Dentista, error) {
	if !r.validateMatricula(d.Matricula) {
		return domain.Dentista{}, errors.New("Matricula já existe")
	}
	return r.store.Create(d)
}

//funcao para validar matriula
func (r *repository) validateMatricula(matricula string) bool {
	for _, dentista := range r.store.GetAll() {
		if dentista.Matricula == matricula {
			return false
		}
	}
	return true
}

//funcao para deletar um dentista
func (r *repository) Delete(id int) error {
	err := r.store.Delete(id)
	if err == nil {
		return nil
	}

	return err
}

//funcao update
func (r *repository) Update(id int, d domain.Dentista) (domain.Dentista, error) {
	for _, dentista := range r.store.GetAll() {
		if dentista.Id == id {
			return r.store.Update(id, d)
		}
	}
	return domain.Dentista{}, errors.New("dentista nao encontrado")
}