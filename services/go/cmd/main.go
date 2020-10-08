package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"service/internal/calculator"
	pb "service/protos/calculator"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx := context.Background()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	endpoint := fmt.Sprintf(":%s", port)
	log.Printf("gRPC endpoint [%s]", endpoint)
	// Register REST API
	mux := http.NewServeMux()
	gwmux := runtime.NewServeMux()
	mux.Handle("/", gwmux)
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	// // dial to gRPC for handling request
	if err := pb.RegisterCalculatorHandlerFromEndpoint(ctx, gwmux, endpoint, opts); err != nil {
		panic(fmt.Errorf("failed registering grpc-gateway: %v", err))
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, calculator.New())
	reflection.Register(grpcServer)
	log.Printf("Starting: gRPC Listener [%s]\n", endpoint)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: dispatcher(ctx, grpcServer, mux),
	}
	// start server
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("error starting server :%v", err)
	}
}

func dispatcher(ctx context.Context, grpcHandler http.Handler, httpHandler http.Handler) http.Handler {
	hf := func(w http.ResponseWriter, r *http.Request) {
		req := r.WithContext(ctx)
		contentTypeHeader := r.Header.Get("Content-Type")
		if r.ProtoMajor == 2 && strings.HasPrefix(contentTypeHeader, "application/grpc") {
			fmt.Println("dispatching to grpc server")
			grpcHandler.ServeHTTP(w, req)
		} else if r.ProtoMajor == 2 {
			// TODO(castlemilk): explore why handler for protocol http/2 doesnt work
			httpHandler.ServeHTTP(w, req)
		} else {
			fmt.Println("dispatching to http server")
			httpHandler.ServeHTTP(w, req)
		}
	}
	return h2c.NewHandler(http.HandlerFunc(hf), &http2.Server{})
}
