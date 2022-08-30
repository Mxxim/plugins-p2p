package plugincom

import "time"

type PluginConfig struct {
	ListenPort     string
	MaxRecvMsgSize int64
	MaxSendMsgSize int64

	SelfHostname string
	SelfDomain   string

	RetryInterval time.Duration
}
