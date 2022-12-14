package debeziumSimulation

import (
	"encoding/json"
	"errors"
	"github.com/helmutkemper/util"
	"log"
	"time"
)

// sendOnSimulationEnd
//
// Envia o dado indicativo de fim de simulação.
//
//	Saída:
//	  err: Objeto de erro padrão do go.
func (e *DebeziumSimulation) sendOnSimulationEnd() (err error) {
	if e.messagingTopicOnTerminate == "" {
		util.TraceToLog()
		err = errors.New(KErrorMessagingTopicOnTerminateIsNotSet)
		return
	}

	if e.messagingSystem == nil {
		util.TraceToLog()
		err = errors.New(KErrorMessagingTopicOnStartIsNotSet)
		return
	}

	var dataToSend []byte

	e.Before = nil
	e.After = nil
	e.Operation = KOperationTerminate
	e.EventDate = time.Now().Unix()

	dataToSend, err = json.Marshal(e)
	if err != nil {
		util.TraceToLog()
		log.Printf("json.Marshal(e).error: %v", err.Error())
		return
	}

	err = e.messagingSystem.Publish(e.messagingTopicOnTerminate, dataToSend)
	if err != nil {
		util.TraceToLog()
	}

	return
}
