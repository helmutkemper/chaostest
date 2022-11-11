package interfaces

import (
	"github.com/helmutkemper/chaostest/support/commonTypes"
	"time"
)

type ParserReceiverInterface interface {
	Receiver(
		data interface{},
	) (
		event commonTypes.ReceiveEvent,
		eventDate time.Time,
		keyToCache,
		dataToCache interface{},
		err error,
	)
}
