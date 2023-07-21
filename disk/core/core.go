package main

import (
	"disk/sql"
	"flag"
	"fmt"

	"disk/core/internal/config"
	"disk/core/internal/handler"
	"disk/core/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "./core/etc/core-api.yaml", "the config file")

func main() {
	flag.Parse()
	sql.Viper_init()
	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
