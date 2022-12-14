package pacienteStore

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
	GetAll() []domain.Paciente
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
	Update(id int, p domain.Paciente) (domain.Paciente, error)
}

// relacionando com a entidade paciente
type store struct {
	list []domain.Paciente
}

// fazendo a comunicacao com o banco de dados, no caso o arquivo json
func NewJsonStore() Store {
	data, err := loadPacientes()
	if err != nil {
		panic("Ocorreu um erro ao consultar os dados.")
	}
	log.Printf("pacientes carregados na memória: %d", len(data))
	return &store{list: data}
}

//fucao getAll
func (st *store) GetAll() []domain.Paciente {
	return st.list
}

//funcao getByID
func (st *store) GetByID(id int) (domain.Paciente, error) {
	log.Printf("produtos carregados na memória: %d", len(st.list))
	for _, paciente := range st.list {
		if paciente.Id == id {
			return paciente, nil
		}
	}

	return domain.Paciente{}, errors.New("paciente not found")
}

//funcao para carregar os pacientes para memoria
func loadPacientes() ([]domain.Paciente, error) {
	file, err := os.ReadFile("./pacientes.json")
	if err != nil {
		return nil, errors.New("ocorreu um erro ao ler o arquivo: " + err.Error())
	}

	var list []domain.Paciente
	err = json.Unmarshal(file, &list)
	if err != nil {
		return nil, errors.New("ocorreu um erro ao converter os dados do arquivo: " + err.Error())
	}

	return list, nil
}

//funcao create
func (st *store) Create(p domain.Paciente) (domain.Paciente, error) {
	p.Id = st.list[len(st.list)-1].Id + 1
	st.list = append(st.list, p)
	st.saveFile()
	return p, nil
}

//funcao delete
func (st *store) Delete(id int) error {
	var pacientesUpdated []domain.Paciente
	for _, paciente := range st.list {
		if paciente.Id != id {
			pacientesUpdated = append(pacientesUpdated, paciente)
		}
	}
	if len(pacientesUpdated) > 0 && len(pacientesUpdated) < len(st.list) {
		st.list = pacientesUpdated
		st.saveFile()
		return nil
	}

	return errors.New("paciente nao encontrado")
}

//funcao saveFile
func (st *store) saveFile() {
	productsToSave, err := json.Marshal(st.list)
	if err != nil {
		panic("erro ao converter os dados para salvar no arquivo")
	}

	err = os.Remove("pacientes.json")
	if err != nil {
		panic("erro ao deletar os dados para salvar no arquivo")
	}

	file, err := os.Create("pacientes.json")
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
func (st *store) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	var pacientesUpdated []domain.Paciente
	for _, paciente := range st.list {
		if paciente.Id == id {
			paciente = p
			fmt.Println(paciente)
		}
		pacientesUpdated = append(pacientesUpdated, paciente)
	}
	st.list = pacientesUpdated
	st.saveFile()
	return p, nil
}