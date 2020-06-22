package consul

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	consulsd "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
	"os"
	"time"
)

// ConsulRegister consul 注册器
type ConsulRegister struct {
	ConsulAddress                  string        // consul address
	ServiceName                    string        // 服务名称
	ServiceIP                      string        // 实例IP
	Tags                           []string      // 服务标签
	ServicePort                    int           //service port
	DeregisterCriticalServiceAfter time.Duration // 单位分钟
	Interval                       time.Duration // 单位秒
}

func NewConsulRegister(consulAddress, serviceName, ServiceIP string, servicePort int, tags []string) *ConsulRegister {
	return &ConsulRegister{
		ConsulAddress:                  consulAddress,
		ServiceName:                    serviceName,
		ServiceIP:                      ServiceIP,
		Tags:                           tags,
		ServicePort:                    servicePort,
		DeregisterCriticalServiceAfter: time.Duration(1) * time.Minute,
		Interval:                       time.Duration(10) * time.Second,
	}
}

// NewConsulGRPCRegister gRPC服务 注册器
func (r *ConsulRegister) NewConsulGRPCRegister() (*consulsd.Registrar, error) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	consulConfig := api.DefaultConfig()
	consulConfig.Address = r.ConsulAddress
	consulClient, err := api.NewClient(consulConfig)

	if err != nil {
		_ = level.Error(logger).Log("err", err)
		return nil, err
	}
	client := consulsd.NewClient(consulClient)

	reg := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%v-%v-%v", r.ServiceName, r.ServiceIP, r.ServicePort),
		Name:    fmt.Sprintf("grpc.health.v1.%v", r.ServiceName),
		Tags:    r.Tags,
		Port:    r.ServicePort,
		Address: r.ServiceIP,
		Check: &api.AgentServiceCheck{
			Interval:                       r.Interval.String(),
			GRPC:                           fmt.Sprintf("%v:%v/%v", r.ServiceIP, r.ServicePort, r.ServiceName),
			DeregisterCriticalServiceAfter: r.DeregisterCriticalServiceAfter.String(),
		},
	}
	return consulsd.NewRegistrar(client, reg, logger), nil
}
