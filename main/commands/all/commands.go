package all

import (
	"github.com/v4fly/v4ray-core/v0/main/commands/all/api"
	"github.com/v4fly/v4ray-core/v0/main/commands/all/tls"
	"github.com/v4fly/v4ray-core/v0/main/commands/base"
)

//go:generate go run v2ray.com/core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		cmdLove,
		tls.CmdTLS,
		cmdUUID,
		cmdVerify,

		// documents
		docFormat,
		docMerge,
	)
}
