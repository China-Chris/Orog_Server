package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/simance-ai/smdx/pkg/kqx"
	"github.com/simance-ai/smdx/rpcx/ws/internal/consumers"
	"net/http"

	"github.com/zeromicro/go-zero/rest"

	"github.com/simance-ai/smdx/rpcx/ws/internal/config"
	"github.com/simance-ai/smdx/rpcx/ws/internal/server"
	"github.com/simance-ai/smdx/rpcx/ws/internal/svc"
	"github.com/simance-ai/smdx/rpcx/ws/ws"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/ws.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	websocketCtx := svc.NewWebsocketContext(c)
	ctx := svc.NewServiceContext(c, websocketCtx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		ws.RegisterWsServer(grpcServer, server.NewWsServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	//kqx
	queue := kqx.MustNewQueue(c.MarketKlineTopicKafkaConf, consumers.NewMarketKlineConsumer(ctx))
	go queue.Start()
	defer queue.Stop()

	queue2 := kqx.MustNewQueue(c.MarketSwapTopicKafkaConf, consumers.NewMarketTxActivityConsumer(ctx))
	go queue2.Start()
	defer queue2.Stop()

	restEngine, err := getRestEngine(ctx)
	if err != nil {
		panic(err)
	}
	defer restEngine.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func getRestEngine(svcCtx *svc.ServiceContext) (*rest.Server, error) {
	restEngine := rest.MustNewServer(rest.RestConf{
		Host:         svcCtx.Config.Rest.Host,
		Port:         svcCtx.Config.Rest.Port,
		Timeout:      svcCtx.Config.Rest.Timeout,
		CpuThreshold: svcCtx.Config.Rest.CpuThreshold,
	})
	if restEngine == nil {
		return nil, errors.New("restEngine is nil")
	}

	restEngine.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/" {
				http.Error(w, "Not found", http.StatusNotFound)
				return
			}
			if r.Method != "GET" {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			http.ServeFile(w, r, "home.html")
		},
	})

	restEngine.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/stream",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			handler := server.NewWebSocketHandler(svcCtx.Config, svcCtx.WebsocketContext)
			handler.ServeHTTP(w, r)
		},
	})

	go func() {
		fmt.Printf("Starting http socket server at %v:%v...\n", svcCtx.Config.Rest.Host, svcCtx.Config.Rest.Port)
		restEngine.Start()
	}()

	return restEngine, nil
}
