package plugin_pprof //nolint: stylecheck

import (
	"net/http"
	"net/http/pprof"

	"github.com/v4fly/v4ray-core/v0/main/commands/base"
	"github.com/v4fly/v4ray-core/v0/main/plugins"
)

var pprofPlugin plugins.Plugin = func(cmd *base.Command) func() error {
	addr := cmd.Flag.String("pprof", "", "")
	return func() error {
		if *addr != "" {
			h := http.NewServeMux()
			h.HandleFunc("/debug/pprof/", pprof.Index)
			h.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
			h.HandleFunc("/debug/pprof/profile", pprof.Profile)
			h.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
			h.HandleFunc("/debug/pprof/trace", pprof.Trace)
			return (&http.Server{Addr: *addr, Handler: h}).ListenAndServe()
		}
		return nil
	}
}

func init() {
	plugins.RegisterPlugin(pprofPlugin)
}
