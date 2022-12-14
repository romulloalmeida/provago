package dentistaStore

import (
	"prova/internal/domain"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

// definindo os metodos do SGBD
type Store interface {
	GetAll() []domain.Dentista
	GetByID(id int) (domain.Dentista, error)
	Create(d domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
	Update(id int, d domain.Dentista) (domain.Dentista, error)
}

// relacionando com a entidade dentista
type store struct {
	list []domain.Dentista
}

// fazendo a comunicacao com o banco de dados, no caso o arquivo json
func NewJsonStore() Store {
	data, err := loadDentistas()
	if err != nil {
		panic("Ocorreu um erro ao consultar os dados.")
	}
	log.Printf("dentistas carregados na memória: %d", len(data))
	return &store{list: data}
}

//fucao getAll
func (st *store) GetAll() []domain.Dentista {
	return st.list
}

//funcao getByID
func (st *store) GetByID(id int) (domain.Dentista, error) {
	log.Printf("dentistas carregados na memória: %d", len(st.list))
	for _, dentista := range st.list {
		if dentista.Id == id {
			return dentista, nil
		}
	}

	return domain.Dentista{}, errors.New("dentista nao encontrado")
}

//funcao para carregar os pentistas para memoria
func loadDentistas() ([]domain.Dentista, error) {
	file, err := os.ReadFile("./dentistas.json")
	if err != nil {
		return nil, errors.New("ocorreu um erro ao ler o arquivo: " + err.Error())
	}

	var list []domain.Dentista
	err = json.Unmarshal(file, &list)
	if err != nil {
		return nil, errors.New("ocorreu um erro ao converter os dados do arquivo: " + err.Error())
	}

	return list, nil
}

//funcao create
func (st *store) Create(d domain.Dentista) (domain.Dentista, error) {
	d.Id = st.list[len(st.list)-1].Id + 1
	st.list = append(st.list, d)
	st.saveFile()
	return d, nil
}

//funcao delete
func (st *store) Delete(id int) error {
	var dentistasUpdated []domain.Dentista
	for _, dentista := range st.list {
		if dentista.Id != id {
			dentistasUpdated = append(dentistasUpdated, dentista)
		}
	}
	if len(dentistasUpdated) > 0 && len(dentistasUpdated) < len(st.list) {
		st.list = dentistasUpdated
		st.saveFile()
		return nil
	}

	return errors.New("dentista nao encontrado")
}

//funcao saveFile
func (st *store) saveFile() {
	productsToSave, err := json.Marshal(st.list)
	if err != nil {
		panic("erro ao converter os dados para salvar no arquivo")
	}

	err = os.Remove("dentistas.json")
	if err != nil {
		panic("erro ao deletar os dados para salvar no arquivo")
	}

	file, err := os.Create("dentistas.json")
	if err != nil {
		panic(fmt.Errorf("erro ao abrir o arquivo: %s", err))
	}
	_, err = file.Write(productsToSave)
	if err != nil {
		panic("erro ao escrever os dados no arquivo")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("erro ao fechar o arquivo")
		}
	}(file)
}

//funcao update
func (st *store) Update(id int, d domain.Dentista) (domain.Dentista, error) {
	var dentistasUpdated []domain.Dentista
	for _, dentista := range st.list {
		if dentista.Id == id {
			dentista = d
			fmt.Println(dentista)
		}
		dentistasUpdated = append(dentistasUpdated, dentista)
	}
	st.list = dentistasUpdated
	st.saveFile()
	return d, nil
}