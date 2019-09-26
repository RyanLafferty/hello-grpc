# Source Reference: https://github.com/PathMotion/firestore-emulator-docker/blob/master/Dockerfile

# Use google cloud sdk
FROM google/cloud-sdk:alpine

# Install Java 8 for Datastore emulator
RUN apk --update --no-cache  add openjdk8-jre
RUN gcloud components install cloud-firestore-emulator beta

# Volume to persist Datastore data
VOLUME /opt/data

# Copy firestore start script
COPY ./firestore/start-firestore.sh .

EXPOSE 8080

# Start firestore
CMD ["sh", "start-firestore.sh"]
