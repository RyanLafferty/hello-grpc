FROM ruby:2.6-buster

RUN gem install logging grpc google-cloud-firestore

COPY ./servers/ruby /app/
COPY ./config/protocol_buffers/compiled/ruby /app
WORKDIR /app

CMD ["ruby", "server.rb"]
