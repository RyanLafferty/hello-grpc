FROM node:12-alpine

COPY ./clients/node /app/
COPY ./config/protocol_buffers /app/protocol_buffers
WORKDIR /app
RUN npm install

EXPOSE 8080
CMD ["npm", "start"]
