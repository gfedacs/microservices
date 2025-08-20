module github.com/gfedacs/microservices/order

go 1.24.4

require (
	github.com/gfedacs/microservices-proto/golang/order v0.0.0
	github.com/gfedacs/microservices-proto/golang/payment v0.0.0
	google.golang.org/grpc v1.74.1
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.30.0
)

require (
	filippo.io/edwards25519 v1.1.0 
	github.com/go-sql-driver/mysql v1.9.3 
	github.com/golang/protobuf v1.5.4 
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/jinzhu/inflection v1.0.0 
	github.com/jinzhu/now v1.1.5
	golang.org/x/net v0.42.0 
	golang.org/x/sys v0.34.0 
	golang.org/x/text v0.27.0 
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b
	google.golang.org/protobuf v1.36.6
)

replace github.com/gfedacs/microservices-proto/golang/order => ../../microservices-proto/golang/order

replace github.com/gfedacs/microservices-proto/golang/payment => ../../microservices-proto/golang/payment