package main

import (
	"log"
	"net"

	"github.com/google/uuid"
	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	pb "github.com/clicrdv/ms-grpc-stubs/mailservice"
	"github.com/clicrdv/ms-mail/mail"
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
	log.Print("Starting microservice grpc listening on 3008")
	lis, err := net.Listen("tcp", "0.0.0.0:3008")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterClicRdvMailServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
