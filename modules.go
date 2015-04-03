package main

import (
	_ "github.com/kanicz/logspout/adapters/raw"
	_ "github.com/kanicz/logspout/adapters/syslog"
	_ "github.com/kanicz/logspout/httpstream"
	_ "github.com/kanicz/logspout/routesapi"
	_ "github.com/kanicz/logspout/transports/tcp"
	_ "github.com/kanicz/logspout/transports/udp"
	_ "github.com/kanicz/logspout/transports/logstash"
)
