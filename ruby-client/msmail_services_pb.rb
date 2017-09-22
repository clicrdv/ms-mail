# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: msmail.proto for package 'mailservice'
# Original file comments:
# GRPC description for ms mail
#

require 'grpc'
require 'msmail_pb'

module Mailservice
  module ClicRdvMailService
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'mailservice.ClicRdvMailService'

      rpc :SendMail, ClicRdvMail, SendMailStatus
    end

    Stub = Service.rpc_stub_class
  end
end