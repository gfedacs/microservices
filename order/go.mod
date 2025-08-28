module github.com/gfedacs/microservices/order

go 1.24.4

require (
	github.com/gfedacs/microservices-proto/golang/order v0.0.1
	github.com/gfedacs/microservices-proto/golang/payment v0.0.1
	github.com/gfedacs/microservices-proto/golang/shipping v0.0.1
	google.golang.org/grpc v1.75.0
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.30.1
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/net v0.42.0 // indirect
	golang.org/x/sys v0.34.0 // indirect
	golang.org/x/text v0.27.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250818200422-3122310a409c // indirect
	google.golang.org/protobuf v1.36.8 // indirect
)

replace github.com/gfedacs/microservices-proto/golang/order => ../../microservices-proto/golang/order

replace github.com/gfedacs/microservices-proto/golang/payment => ../../microservices-proto/golang/payment

replace github.com/gfedacs/microservices-proto/golang/shipping => ../../microservices-proto/golang/shipping