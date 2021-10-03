# frozen_string_literal: true

require 'pinger_services_pb.rb'
require 'pinger_pb.rb'
class ApplicationController < ActionController::API
  def ping
    pinger_stub = Pinger::Pinger::Stub.new('localhost:5300', :this_channel_is_insecure)

    pong = pinger_stub.ping(Pinger::Empty.new)
    render json: { pong: pong.text }
  end
end
