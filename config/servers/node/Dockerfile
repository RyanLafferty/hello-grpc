FROM node:12-alpine

COPY ./servers/node /app/
COPY ./config/protocol_buffers /app/protocol_buffers
WORKDIR /app
RUN npm install

CMD ["npm", "start"]
