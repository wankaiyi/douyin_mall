package constant

import (
	"github.com/google/uuid"
)

var (
	WorkerChannelId     uuid.UUID
	ClientChannel       string
	WorkerChannelIdList string
)

func init() {

	WorkerChannelId = uuid.New()
	ClientChannel = "client-channel"
	WorkerChannelIdList = "worker-channel-id-list"
}
