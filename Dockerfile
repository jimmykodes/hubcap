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
RUN go install ./cmd/...

FROM debian:buster
COPY --from=dist /app/dist /dist
COPY --from=binary /server /server
CMD "./server"
