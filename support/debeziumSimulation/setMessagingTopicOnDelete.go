package debeziumSimulation

// SetMessagingTopicOnDelete
//
// Define o tópico do sistema de mensageria quando o dado é apagado.
//
//	Entrada:
//	  topic: texto identificador do tópico do sistema de mensageria.
//
//	Nota:
//	  * Use as funções SetMessagingTopicOnCreate(), SetMessagingTopicOnDelete(),
//	    SetMessagingTopicOnStart(), SetMessagingTopicOnUpdate() e SetMessagingTopicOnTerminate() para
//	    definir tópicos específicos.
//	  * Use a função SetMessagingTopic() para definir todos os tópicos simultaneamente, e em seguida
//	    use as demais funções para definir um tópico específico.
func (e *DebeziumSimulation) SetMessagingTopicOnDelete(topic string) {
	e.messagingTopicOnDelete = topic
}
