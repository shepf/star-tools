package impl

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/shepf/star-tools/node/impl/full"

	"github.com/shepf/star-tools/api"
	"github.com/shepf/star-tools/node/impl/common"
)

var log = logging.Logger("node")

type FullNodeAPI struct {
	common.CommonAPI

	//自定义 star监控 相关api
	full.MonitorAPI
}

var _ api.FullNode = &FullNodeAPI{}
