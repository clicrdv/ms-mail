package main

import (
	"log"
	"net"
	"os"
	"strings"

	"github.com/google/uuid"
	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	followpb "github.com/clicrdv/ms-grpc-stubs/followservice"
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

	conn, err := grpc.Dial(os.Getenv("FOLLOW_MAIL_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	emails := []string{}
	for _, value := range sm.ToMap {
		emails = append(emails, value)
	}
	emailsStr := strings.Join(emails, ",")
	clicRdvFollowMail := followpb.ClicRdvFollowMail{
		Email:   emailsStr,
		Event:   "SENT",
		GroupId: "12345",
		Uuid:    sm.UUID,
	}
	followStatusClient := followpb.NewClicRdvFollowMailServiceClient(conn)
	sendMailStatus, err := followStatusClient.NotifySentMail(context.TODO(), &clicRdvFollowMail)
	if err != nil {
		log.Fatal("Error while sending mail status to follow mail service : ", err.Error())
		log.Fatal("Status : ", sendMailStatus)
	}
	return &pb.SendMailStatus{Status: status, UniqueId: uuid}, nil
}

func main() {
	log.Print("Starting microservice grpc listening on 3008")
	listenAddress := os.Getenv("LISTEN_ADDRESS")
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterClicRdvMailServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
