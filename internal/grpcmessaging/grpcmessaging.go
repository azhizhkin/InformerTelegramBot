package grpcmessaging

import (
	"context"
	"log"
	"net"

	"github.com/AndreyZhizhkin/informertelegrambot/internal/appconfig"
	grpcMsg "github.com/AndreyZhizhkin/informertelegrambot/internal/proto/com/informertelegrambot"
	"github.com/AndreyZhizhkin/informertelegrambot/internal/telegrambotservice"
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

func NewGRPCMessagingService(appconfig *appconfig.AppConfig, botService *telegrambotservice.TelegramBotService, grpcServer *grpc.Server, lis *net.Listener) *GRPCMessagingService {
	service := GRPCMessagingService{
		config:     appconfig,
		botService: botService,
		server:     grpcServer,
		lis:        lis,
	}

	return &service
}
func (s *GRPCMessagingService) NewMessage(ctx context.Context, msgPB *grpcMsg.Message) (*grpcMsg.MessageResponse, error) {
	err := s.botService.SendMessage(msgPB.Username, msgPB.Text)
	resp := grpcMsg.MessageResponse{}
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
