package proxy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Goalt/hostmonitor/internal/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    apipb "github.com/Goalt/hostmonitor/internal/infrastructure/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Proxy struct {
    cnf config.ProxyServer
    server *http.Server
}

func NewProxy(cnf config.ProxyServer) *Proxy {
    return &Proxy{
        cnf: cnf,
    }
}

func (p *Proxy) Run(grpcServerPort int) error {
    conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%v", grpcServerPort),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
        return err
	}

	mux := runtime.NewServeMux()
    err = apipb.RegisterHostMonitorServiceHandler(context.Background(), mux, conn)
	if err != nil {
		return err
	}
    
    // register the gRPC Service Handler(s) to the mux
    p.server = &http.Server{
        Addr:         fmt.Sprintf(":%v", p.cnf.Port),
        Handler:      cors(mux),
    }

    return p.server.ListenAndServe()
}

func (p *Proxy) Stop() error {
    return p.server.Close()
}

func allowedOrigin(origin string) bool {
    // if viper.GetString("cors") == "*" {
    //     return true
    // }
    // if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
    //     return true
    // }
    return true
}

func cors(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if allowedOrigin(r.Header.Get("Origin")) {
            w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
            w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
        }
        if r.Method == "OPTIONS" {
            return
        }
        h.ServeHTTP(w, r)
    })
}