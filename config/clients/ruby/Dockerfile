FROM ruby:2.6-buster

RUN gem install json sinatra grpc

COPY ./clients/ruby /app/
COPY ./config/protocol_buffers/compiled/ruby /app
WORKDIR /app

EXPOSE 8080
CMD ["ruby", "client.rb"]
