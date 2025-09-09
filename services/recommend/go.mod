module github.com/example/mini-tiktok/services/recommend

go 1.22

require (
	google.golang.org/grpc v1.64.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.10
	github.com/redis/go-redis/v9 v9.7.0
	github.com/hashicorp/consul/api v1.26.1
)
replace github.com/example/mini-tiktok/proto => ../../
replace github.com/example/mini-tiktok/pkg => ../../pkg
