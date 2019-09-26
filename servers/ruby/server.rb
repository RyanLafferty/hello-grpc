require 'logging'
require 'grpc'
require "google/cloud/firestore"

root_dir = File.expand_path(File.dirname(__FILE__))
helloworld_pb_dir = File.join(root_dir, 'helloworld')
store_pb_dir = File.join(root_dir, 'store')

$LOAD_PATH.unshift(helloworld_pb_dir) unless $LOAD_PATH.include?(helloworld_pb_dir)
$LOAD_PATH.unshift(store_pb_dir) unless $LOAD_PATH.include?(store_pb_dir)

require 'helloworld_services_pb'
require 'store_services_pb'

# Add logging to GRPC module
module GRPC
  extend Logging.globally
end

# Append standard out to logging and set level to info
Logging.logger.root.appenders = Logging.appenders.stdout
Logging.logger.root.level = :info

class GreeterServer < Helloworld::Greeter::Service
  def say_hello(hello_req, _unused_call)
    puts "Greeting: #{hello_req.name}"
    Helloworld::HelloReply.new(message: "Hello #{hello_req.name}")
  end
end

class StoreServer < Store::Store::Service
  def get_line_item(line_item_req, _unused_call)
    puts "Received request for line item #{line_item_req.id}"
    puts "Connecting to project '#{ENV["FIRESTORE_PROJECT_ID"]}' on firestore emulator: #{ENV["FIRESTORE_EMULATOR_HOST"]}"
    firestore = Google::Cloud::Firestore.new(
      project_id: ENV["FIRESTORE_PROJECT_ID"],
      emulator_host: ENV["FIRESTORE_EMULATOR_HOST"]
    )

    doc_ref = firestore.doc("line_items/#{line_item_req.id}")
    document = doc_ref.get()

    if document.exists?
      puts "Fetched line item: #{document.data}"
      Store::LineItem.new(document.data.merge(source: 'ruby'))
    else
      puts "Error: Could not fetch line item #{line_item_req.id}"
    end
  end
end

def main
  s = GRPC::RpcServer.new
  s.add_http2_port('0.0.0.0:50051', :this_port_is_insecure)

  s.handle(GreeterServer)
  s.handle(StoreServer)

  s.run_till_terminated_or_interrupted([1, 'int', 'SIGQUIT'])
end

main
