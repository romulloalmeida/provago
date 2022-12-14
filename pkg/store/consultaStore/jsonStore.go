package consultaStore

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
	GetAll() []domain.Consulta
	GetByID(id int) (domain.Consulta, error)
	Create(c domain.Consulta) (domain.Consulta, error)
	Delete(id int) error
	Update(id int, c domain.Consulta) (domain.Consulta, error)
}

// relacionando com a entidade consulta
type store struct {
	list []domain.Consulta
}

// fazendo a comunicacao com o banco de dados, no caso o arquivo json
func NewJsonStore() Store {
	data, err := loadConsultas()
	if err != nil {
		panic("Ocorreu um erro ao consultar os dados.")
	}
	log.Printf("consultas carregados na memória: %d", len(data))
	return &store{list: data}
}

//fucao getAll
func (st *store) GetAll() []domain.Consulta {
	return st.list
}

//funcao getByID
func (st *store) GetByID(id int) (domain.Consulta, error) {
	log.Printf("consultas carregados na memória: %d", len(st.list))
	for _, consulta := range st.list {
		if consulta.Id == id {
			return consulta, nil
		}
	}

	return domain.Consulta{}, errors.New("consulta nao encontrado")
}

//funcao para carregar os pentistas para memoria
func loadConsultas() ([]domain.Consulta, error) {
	file, err := os.ReadFile("./consultas.json")
	if err != nil {
		return nil, errors.New("ocorreu um erro ao ler o arquivo: " + err.Error())
	}

	var list []domain.Consulta
	err = json.Unmarshal(file, &list)
	if err != nil {
		return nil, errors.New("ocorreu um erro ao converter os dados do arquivo: " + err.Error())
	}

	return list, nil
}

//funcao create
func (st *store) Create(c domain.Consulta) (domain.Consulta, error) {
	c.Id = st.list[len(st.list)-1].Id + 1
	st.list = append(st.list, c)
	st.saveFile()
	return c, nil
}

//funcao delete
func (st *store) Delete(id int) error {
	var consultasUpdated []domain.Consulta
	for _, consulta := range st.list {
		if consulta.Id != id {
			consultasUpdated = append(consultasUpdated, consulta)
		}
	}
	if len(consultasUpdated) > 0 && len(consultasUpdated) < len(st.list) {
		st.list = consultasUpdated
		st.saveFile()
		return nil
	}

	return errors.New("consulta nao encontrado")
}

//funcao saveFile
func (st *store) saveFile() {
	productsToSave, err := json.Marshal(st.list)
	if err != nil {
		panic("erro ao converter os dados para salvar no arquivo")
	}

	err = os.Remove("consultas.json")
	if err != nil {
		panic("erro ao deletar os dados para salvar no arquivo")
	}

	file, err := os.Create("consultas.json")
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
func (st *store) Update(id int, c domain.Consulta) (domain.Consulta, error) {
	var consultasUpdated []domain.Consulta
	for _, consulta := range st.list {
		if consulta.Id == id {
			consulta = c
			fmt.Println(consulta)
		}
		consultasUpdated = append(consultasUpdated, consulta)
	}
	st.list = consultasUpdated
	st.saveFile()
	return c, nil
}