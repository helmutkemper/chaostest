package interfaces

// DataToSimulateInterface interface de funções necessárias para a simulação de dados.
type DataToSimulateInterface interface {
	// Populate função encarregada de popular o dado simulado.
	//
	// Imagine como o usuário preencheria a informação e tente simular o que o usuário faria.
	Populate() (err error)

	// Update função encarregada de atualizar dados simulados.
	//
	// Imagine como o usuário atualizaria o dado para deixar a simulação o mais real possível.
	Update() (err error)

	// Get retorna um novo dado simulado populado.
	Get() (data interface{})

	// GetID retorna o id do último dado gerado.
	GetID() (ID interface{}, err error)
}
