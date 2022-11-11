package dataTest

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
)

// DataTest estrutura de dado simulado.
type DataTest struct {
	Id   string
	Name string
}

// GetID retorna o id do último dado gerado.
func (e *DataTest) GetID() (ID interface{}, err error) {
	ID = e.Id
	return
}

// getNextId retorna um novo ID a ser usado no próximo dado gerado.
func (e *DataTest) getNextId() (id string) {
	return gofakeit.UUID()
}

// Get retorna um novo dado simulado populado.
func (e *DataTest) Get() (data interface{}) {
	return *e
}

// Populate função encarregada de popular o dado simulado.
func (e *DataTest) Populate() (err error) {
	e.Id = e.getNextId()
	e.Name = gofakeit.Name()
	return
}

// Update função encarregada de atualizar dados simulados.
func (e *DataTest) Update() (err error) {
	e.Name = gofakeit.Name()
	return
}

// UnmarshalJSON converte o json novamente em dado.
//
//	Nota:
//	  * Alguns bancos de dados isam booleano como sendo inteiro e o dado usado no Golang usa
//	    booleano. Quando isto acontece, esta conversão é necessária.
func (e *DataTest) UnmarshalJSON(data []byte) (err error) {
	err = json.Unmarshal(data, e)
	if err != nil {
		return
	}

	return
}

// MarshalJSON converte o dado em json
//
//	Nota:
//	  * Alguns bancos de dados isam booleano como sendo inteiro e o dado usado no Golang usa
//	    booleano. Quando isto acontece, esta conversão é necessária.
func (e *DataTest) MarshalJSON() (data []byte, err error) {
	return json.Marshal(&struct {
		Id   string
		Name string
	}{
		Id:   e.Id,
		Name: e.Name,
	})
}
