package constants

var (
	ClientChannel       string
	ClientServiceName   string
	WorkerChannelIdList string
)

func Init(serviceName string) {
	ClientChannel = "client-channel"
	ClientServiceName = serviceName
	WorkerChannelIdList = "worker-channel-id-list"
}
