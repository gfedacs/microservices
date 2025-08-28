module github.com/gfedacs/microservices/shipping

go 1.24.4

require (
	go.opentelemetry.io/otel v1.37.0
	go.opentelemetry.io/otel/trace v1.37.0
	gorm.io/driver/mysql v1.6.0
)

require (
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel/metric v1.37.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/gfedacs/microservices-proto/golang/shipping v0.0.0-20250828131226-c2bb762c181e
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/sirupsen/logrus v1.9.3
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	go.opentelemetry.io/otel/sdk v1.37.0
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/grpc v1.75.0
	gorm.io/gorm v1.30.2
)

replace github.com/gfedacs/microservices-proto/golang/shipping => ../../microservices-proto/golang/shipping