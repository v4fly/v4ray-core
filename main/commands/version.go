package commands

import (
	"fmt"

	core "github.com/v4fly/v4ray-core/v0"
	"github.com/v4fly/v4ray-core/v0/main/commands/base"
)

// CmdVersion prints V2Ray Versions
var CmdVersion = &base.Command{
	UsageLine: "{{.Exec}} version",
	Short:     "print V2Ray version",
	Long: `Prints the build information for V2Ray.
`,
	Run: executeVersion,
}

func executeVersion(cmd *base.Command, args []string) {
	printVersion()
}

func printVersion() {
	version := core.VersionStatement()
	for _, s := range version {
		fmt.Println(s)
	}
}
