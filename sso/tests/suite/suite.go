package suite

import (
	"testing"

	ssov1 "github.com/Vladislav747/protos/gen/go/sso"
	"sso/internal/config"
)

type Suite struct {
	*testing.T                  // потребуется для вызова метода *testing.T внутри Suite
	Cfg        *config.Config   // конфигурация приложения
	AuthClient ssov1.AuthClient // клиент для взаимодействия с сервисом Auth
}

func New(t *testing.T) (context.Context, *Suite) {
	// помогает отлавливать ошибки в тестах
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadByPath("config/local_tests.yaml")

	ctx, cancelCtx := context.WithCancel(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	cc, err := grpc.DialContext(
		ctx,
		grpcAddress(cfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		AuthClient: ssov1.NewAuthClient(conn),
	}

}

/* grpcAddress соединяет хост и порт для соединения с gRPC сервером */
func grpcAddress(cfg *config.Config) string {
	return net.JoinHostPort(grpcHost, strconv.Itoa(cfg.GRPC.Port))
}
