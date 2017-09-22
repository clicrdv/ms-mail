package main

import (
	"log"
	"net"

	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/clicrdv/ms-mail/mail"
	pb "github.com/clicrdv/ms-mail/mailservice"
)

type server struct{}

func (s *server) SendMail(ctxt context.Context, mailToSend *pb.ClicRdvMail) (*pb.SendMailStatus, error) {
	log.Print("Received GRPC Call")
	log.Printf("Received arguments : %s, %s", mailToSend.GetFromEmail(), mailToSend.GetHtmlContent())
	targetMailMap := map[string]string{
		"mikrob - perso": "mikrob@yopmail.com",
		"mikrob - pro":   "mikrob+3@yopmail.com",
	}

	sm := mail.SendgridMail{
		FromName:     "No Reply ClicRDV",
		FromEmail:    "noreply@clicrdv.com",
		ReplyToName:  "No Reply ClicRDV",
		ReplyToEmail: "noreply@clicrdv.com",
		HtmlContent:  "<html><body><b>This is bold html</b></body></html>",
		TextContent:  "This is text content",
		Subject:      "Mail From MS Mail",
		ToMap:        targetMailMap,
	}
	sm.SendMail()
	log.Print("Finished grpc processing")
	return nil, nil
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
