package debeziumSimulation

import "errors"

// GetUpdateList
//
// Retorna a lista de dados a serem atualizados
//
//	Saída:
//	  list: map[string]struct{Id string `json:"id"`, Name string `json:"name"`}
func (e *DebeziumSimulation) GetUpdateList() (list interface{}, err error) {
	if e.realDataPointer == nil {
		err = errors.New(KErrorUseSetDataFunctionFirst)
		return
	}

	return e.update, err
}
