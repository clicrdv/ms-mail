this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'ruby-client')

$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'msmail_services_pb'

def main
  stub = Mailservice::ClicRdvMailService::Stub.new('localhost:50052', :this_channel_is_insecure)
  mailstatus = stub.send_mail(
    Mailservice::ClicRdvMail.new(
      fromName: "Mikrob From",
      fromEmail: "mikrob@clicrdv.com",
      replyToName: "noreply@clicrdv.com",
      replyToEmail: "noreply@clicrdv.com",
      htmlContent: "<html><body><h1>This is a mail title </h1><b>This is bold content </b></i> This is italitc content</i></body></html>",
      textContent: "this is a text content",
      subject: "MS Mail subject",
      toMap: {
        "pipo" => "pipo@yopmail.com",
        "molo" => "molo@yopmail.com"
      }
    )
  )
  puts "Mail sent : #{mailstatus.status}, id : #{mailstatus.uniqueId}"
end

main
