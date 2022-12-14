package domain

type Paciente struct {
	Id          	int     `json:"id"`
	Nome        	string  `json:"nome" binding:"required"`
	Sobrenome		string  `json:"sobrenome" binding:"required"`
	RG		   		string  `json:"rg" binding:"required"`
	DataCadastro	string  `json:"data_cadastro" binding:"required"`
}
