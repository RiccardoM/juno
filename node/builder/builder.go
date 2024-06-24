package builder

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
<<<<<<< HEAD
	"github.com/forbole/juno/v5/node"
	nodeconfig "github.com/forbole/juno/v5/node/config"
	"github.com/forbole/juno/v5/node/local"
	"github.com/forbole/juno/v5/node/remote"
=======

	"github.com/forbole/juno/v6/node"
	nodeconfig "github.com/forbole/juno/v6/node/config"
	"github.com/forbole/juno/v6/node/local"
	"github.com/forbole/juno/v6/node/remote"
>>>>>>> 7327795 (fix: added back authz message handler in worker (#121))
)

func BuildNode(cfg nodeconfig.Config, txConfig client.TxConfig, codec codec.Codec) (node.Node, error) {
	switch cfg.Type {
	case nodeconfig.TypeRemote:
		return remote.NewNode(cfg.Details.(*remote.Details))
	case nodeconfig.TypeLocal:
		return local.NewNode(cfg.Details.(*local.Details), txConfig, codec)
	case nodeconfig.TypeNone:
		return nil, nil

	default:
		return nil, fmt.Errorf("invalid node type: %s", cfg.Type)
	}
}
