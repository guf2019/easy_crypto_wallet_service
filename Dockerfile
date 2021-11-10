FROM golang:latest

RUN mkdir /easy_crypto_wallet_service


RUN apt -y update && apt -y install git

COPY . /easy_crypto_wallet_service
WORKDIR /easy_crypto_wallet_service

RUN cd /easy_crypto_wallet_service && make install && buf mod update && make generate
RUN cd cmd/standalone/ && go build -o main


RUN chmod +x ./scripts/migrate.sh

ENTRYPOINT ["/easy_crypto_wallet_service/cmd/standalone/main"]