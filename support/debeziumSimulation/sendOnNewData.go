package debeziumSimulation

import (
	"encoding/json"
	"errors"
	"github.com/helmutkemper/util"
	"log"
	"time"
)

// sendOnNewData
//
// Envia um dado criado pelo sistema de mensageria.
//
//	Entrada:
//	  after: dado criado;
//
//	Saída:
//	  err: Objeto de erro padrão do go.
func (e *DebeziumSimulation) sendOnNewData(after interface{}) (err error) {
	if e.messagingTopicOnCreate == "" {
		util.TraceToLog()
		err = errors.New(KErrorMessagingTopicOnCreateIsNotSet)
		return
	}

	if e.messagingSystem == nil {
		util.TraceToLog()
		err = errors.New(KErrorMessagingInterfaceIsNotSet)
		return
	}

	var dataToSend []byte

	e.Before = nil
	e.After = after
	e.Operation = KOperationCreate
	e.EventDate = time.Now().Unix()

	dataToSend, err = json.Marshal(e)
	if err != nil {
		util.TraceToLog()
		log.Printf("json.Marshal(e).error: %v", err.Error())
		return
	}

	err = e.messagingSystem.Publish(e.messagingTopicOnCreate, dataToSend)
	if err != nil {
		util.TraceToLog()
	}

	return
}
