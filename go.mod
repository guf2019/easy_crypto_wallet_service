module github.com/guf2019/easy_crypto_wallet_service

go 1.16

require (
	github.com/ethereum/go-ethereum v1.10.11
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.2.0
	github.com/ilyakaznacheev/cleanenv v1.2.5
	github.com/lib/pq v1.10.1
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.24.0
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	google.golang.org/genproto v0.0.0-20210506142907-4a47615972c2
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/aopoltorzhicky/go_kraken/rest v0.0.3 => github.com/oomag/go_kraken/rest v0.0.4

replace github.com/adshao/go-binance/v2 v2.2.1 => github.com/oomag/go-binance/v2 v2.2.3
