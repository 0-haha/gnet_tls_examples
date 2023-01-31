module gnet_test

require (
	github.com/panjf2000/gnet/v2 v2.2.5
	github.com/secure-for-ai/goktls v1.20.0-rc3.2
)

require (
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace github.com/panjf2000/gnet/v2 v2.2.5 => github.com/0-haha/gnet/v2 v2.2.5-tls-rc1

// for deveoplement
// replace github.com/panjf2000/gnet/v2 v2.2.5 => ../gnet/

go 1.20
