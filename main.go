package main

import (
	"log"
	"net"

	"github.com/google/uuid"
	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/clicrdv/ms-mail/mail"
	pb "github.com/clicrdv/ms-mail/mailservice"
)

type server struct{}

func (s *server) SendMail(ctxt context.Context, mailToSend *pb.ClicRdvMail) (*pb.SendMailStatus, error) {
	log.Print("Received GRPC Call")
	log.Printf("Received arguments : %s, %s", mailToSend.GetFromEmail(), mailToSend.GetHtmlContent())
	sm := mail.SendgridMail{
		FromName:     mailToSend.GetFromName(),
		FromEmail:    mailToSend.GetFromEmail(),
		ReplyToName:  mailToSend.GetReplyToName(),
		ReplyToEmail: mailToSend.GetReplyToEmail(),
		HtmlContent:  mailToSend.GetHtmlContent(),
		TextContent:  mailToSend.GetTextContent(),
		Subject:      mailToSend.GetSubject(),
		ToMap:        mailToSend.GetToMap(),
		UUID:         uuid.New().String(),
	}
	status, uuid := sm.SendMail()
	log.Print("Finished grpc processing of mail with uuid:", sm.UUID)
	return &pb.SendMailStatus{Status: status, UniqueId: uuid}, nil
}

func main() {
	log.Print("Starting microservice grpc listening on 50052")
	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterClicRdvMailServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
