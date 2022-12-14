package consulta

import (
	"prova/internal/domain"
	"errors"
	"prova/pkg/store/consultaStore"
)

// declarando os metodos que o repository tem
type Repository interface {
	GetByID(id int) (domain.Consulta, error)
	Create(c domain.Consulta) (domain.Consulta, error)
	Delete(id int) error
	Update(id int, c domain.Consulta) (domain.Consulta, error) 
}

// fazendo a comunicação com o SGBD
type repository struct {
	store consultaStore.Store
}

// instanciando um repository
func NewRepository(store consultaStore.Store) Repository {
	return &repository{store}
}

// função getByID
func (r *repository) GetByID(id int) (domain.Consulta, error) {
	pReturn, err := r.store.GetByID(id)
	if err != nil {
		return domain.Consulta{}, err
	}

	return pReturn, nil
}

//funcao para criar um novo consulta
func (r *repository) Create(c domain.Consulta) (domain.Consulta, error) {
	return r.store.Create(c)
}

//funcao para deletar um consulta
func (r *repository) Delete(id int) error {
	err := r.store.Delete(id)
	if err == nil {
		return nil
	}

	return err
}

//funcao update
func (r *repository) Update(id int, c domain.Consulta) (domain.Consulta, error) {
	for _, consulta := range r.store.GetAll() {
		if consulta.Id == id {
			return r.store.Update(id, c)
		}
	}
	return domain.Consulta{}, errors.New("consulta nao encontrado")
}