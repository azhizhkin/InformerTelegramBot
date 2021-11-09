package grpcmessaging

import (
	"context"
	"log"
	"net"

	"github.com/azhizhkin/informertelegrambot/internal/appconfig"
	grpcMsg "github.com/azhizhkin/informertelegrambot/internal/proto/com/informertelegrambot"
	"github.com/azhizhkin/informertelegrambot/internal/telegrambotservice"
	"google.golang.org/grpc"
)

type GRPCMessagingService struct {
	config     *appconfig.AppConfig
	botService *telegrambotservice.TelegramBotService
	isRunning  bool
	server     *grpc.Server
	lis        *net.Listener
	grpcMsg.UnimplementedInformerBotMessagingServer
}

func NewGRPCMessagingService(appConfig *appconfig.AppConfig, botService *telegrambotservice.TelegramBotService, grpcServer *grpc.Server, lis *net.Listener) *GRPCMessagingService {
	service := GRPCMessagingService{
		config:     appConfig,
		botService: botService,
		server:     grpcServer,
		lis:        lis,
	}

	return &service
}
func (s *GRPCMessagingService) NewMessage(ctx context.Context, msgPb *grpcMsg.NewMessageRequest) (*grpcMsg.NewMessageResponse, error) {
	err := s.botService.SendMessage(msgPb.Username, msgPb.Text)
	resp := grpcMsg.NewMessageResponse{}
	return &resp, err
}

func (s *GRPCMessagingService) Run() {
	grpcMsg.RegisterInformerBotMessagingServer(s.server, s)
	log.Printf("gRPPC server listening at %v", (*s.lis).Addr())
	s.isRunning = true
	if err := s.server.Serve(*s.lis); err != nil {
		s.isRunning = false
		log.Fatalf("gRPC failed to serve: %v", err)
	}
}

func (s *GRPCMessagingService) Stop() {
	if s.isRunning {
		s.server.GracefulStop()
		s.isRunning = false
		log.Println("gRPC messaging service stopped...")
	}
}
