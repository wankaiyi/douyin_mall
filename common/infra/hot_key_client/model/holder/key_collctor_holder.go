package holder

import "douyin_mall/common/infra/hot_key_client/model/utils"

var (
	TurnKeyCollector *utils.TurnKeyCollector
)

func init() {
	TurnKeyCollector = utils.NewTurnKeyCollector()
}
