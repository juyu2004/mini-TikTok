package consul

import (
	"fmt"
	"os"
	"github.com/hashicorp/consul/api"
)

func Register(service, id string, port int) error {
	addr := os.Getenv("CONSUL_ADDR")
	if addr == "" { addr = "http://localhost:8500" }
	cfg := api.DefaultConfig()
	cfg.Address = addr
	cli, err := api.NewClient(cfg)
	if err != nil { return err }
	reg := &api.AgentServiceRegistration{
		ID: id, Name: service,
		Port: port, Address: "host.docker.internal",
		Check: &api.AgentServiceCheck{
			GRPC: fmt.Sprintf("host.docker.internal:%d", port),
			Interval: "10s",
			Timeout: "3s",
		},
	}
	return cli.Agent().ServiceRegister(reg)
}
