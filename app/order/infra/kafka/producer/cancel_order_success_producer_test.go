package producer

import (
	"context"
	"testing"
)

func TestCancelOrderSuccessProducer_Produce(t *testing.T) {
	InitCancelOrderSuccessProducer()
	SendCancelOrderSuccessMessage(context.Background(), "111")
}
