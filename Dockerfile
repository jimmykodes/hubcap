FROM node:14 as dist
WORKDIR /app
COPY web .
RUN npm i
RUN npm run build
RUN npm run generate

FROM golang:1 as binary
WORKDIR /app
COPY . .
ENV GOBIN=/
RUN go install .

FROM debian:buster
COPY --from=dist /app/dist /dist
COPY --from=binary /vehicle_maintenance /vehicle_maintenance
RUN apt-get update --fix-missing && \
    apt-get install -yqq \
    ca-certificates
CMD "./vehicle_maintenance"
