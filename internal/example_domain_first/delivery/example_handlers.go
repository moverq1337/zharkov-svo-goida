package delivery

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpcLib "github.com/neiasit/grpc-library/core"
	examplegen "github.com/neiasit/service-boilerplate/pkg/api/grpc/golang"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type ExampleHandlers struct {
	logger *slog.Logger
	examplegen.UnimplementedExampleServiceServer
}

func RegisterHandlers(
	grpcSrv *grpc.Server,
	grpcGtw *runtime.ServeMux,
	grpcConfig *grpcLib.Config,
	logger *slog.Logger,
) {

	ctx := context.Background()

	impl := &ExampleHandlers{logger: logger}
	examplegen.RegisterExampleServiceServer(grpcSrv, impl)
	err := examplegen.RegisterExampleServiceHandlerFromEndpoint(ctx, grpcGtw, grpcConfig.Address(), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		logger.Error("failed to register example service handler", "error", err)
		return
	}
	ctx = context.WithValue(ctx, "parameter", "im fucking parametr")
	slog.InfoContext(ctx, "hello world")
}

func (h *ExampleHandlers) ExampleReq(context.Context, *examplegen.Example) (*examplegen.Example, error) {
	return &examplegen.Example{
		Title:       "Kowlad",
		Description: "Гений черт возьми",
	}, nil
}
