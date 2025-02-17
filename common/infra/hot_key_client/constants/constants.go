package constants

var (
	ClientChannel       string
	ClientServiceName   string
	WorkerChannelIdList string
)

func init() {
	ClientChannel = "client-channel"
	ClientServiceName = "test-service"
	WorkerChannelIdList = "worker-channel-id-list"
}
