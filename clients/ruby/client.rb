require 'sinatra'
require 'json'
require 'grpc'

root_dir = File.expand_path(File.dirname(__FILE__))
helloworld_pb_dir = File.join(root_dir, 'helloworld')
store_pb_dir = File.join(root_dir, 'store')

$LOAD_PATH.unshift(helloworld_pb_dir) unless $LOAD_PATH.include?(helloworld_pb_dir)
$LOAD_PATH.unshift(store_pb_dir) unless $LOAD_PATH.include?(store_pb_dir)

require 'helloworld_services_pb'
require 'store_services_pb'

set :bind, '0.0.0.0'
set :port, 8080

get '/hello/:server/:name' do
  client = Helloworld::Greeter::Stub.new("#{params["server"]}-grpc-server:50051", :this_channel_is_insecure)
  client.say_hello(Helloworld::HelloRequest.new(name: params["name"])).message
end

get '/item/:server/:id' do
  client = Store::Store::Stub.new("#{params["server"]}-grpc-server:50051", :this_channel_is_insecure)
  client.get_line_item(Store::LineItemRequest.new(id: params["id"].to_i)).to_h.merge(
    client: 'ruby'
  ).to_json
end
