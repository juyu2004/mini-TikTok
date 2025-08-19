package consul

import (
        "context"
        "fmt"
        "github.com/hashicorp/consul/api"
        "mini-douyin/src/config"
        "time"
)

// RegisterService 注册服务到Consul
func RegisterService(serviceName, addr string) error {
        client, err := api.NewClient(&api.Config{
                Address: config.ConsulAddr,
        })
        if err != nil {
                return err
        }

        registration := &api.AgentServiceRegistration{
                Name:    serviceName,
                ID:      fmt.Sprintf("%s-%s", serviceName, addr),
                Address: addr,
                Port:    config.RPCPort,
                Check: &api.AgentServiceCheck{
                        HTTP:                           fmt.Sprintf("http://%s:%d/health", addr, config.RPCPort),
                        Timeout:                        "5s",
                        Interval:                       "10s",
                        DeregisterCriticalServiceAfter: "30s",
                },
        }

        return client.Agent().ServiceRegister(registration)
}