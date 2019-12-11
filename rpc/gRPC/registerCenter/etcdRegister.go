package registerCenter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
)

type etcdRegisterImpl struct {
	//Address     string
	Endpoints   []string
	DialTimeout time.Duration
}

func NewRegisterImpl(endpoints []string, timeout time.Duration) *etcdRegisterImpl {
	return &etcdRegisterImpl{
		Endpoints:   endpoints,
		DialTimeout: timeout,
	}
}

// etcd 实现注册接口
func (etcd *etcdRegisterImpl) Register(info ServiceDescInfo) error {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcd.Endpoints,
		DialTimeout: etcd.DialTimeout,
	})
	if err != nil {
		log.Printf("etcd connect error: %v", err)
		return err
	}

	resp, err := client.Grant(context.TODO(), int64(info.IntervalTime))
	if err != nil {
		log.Printf("etcd grant error: %v", err)
		return err
	}

	_, err = client.Get(context.Background(), info.ServiceName)
	serviceValue := fmt.Sprintf("%s:%d", info.Host, info.Port)
	if err != nil {
		if err == rpctypes.ErrKeyNotFound {
			if _, err := client.Put(context.TODO(), info.ServiceName, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
				log.Printf("etcd put service (`%s`:`%s`) error: %v", info.ServiceName, serviceValue, err)
				return err
			}
		} else {
			log.Printf("etcd service `%s` connects to etcd3 failed", info.ServiceName)
			return err
		}
	} else {
		if _, err := client.Put(context.Background(), info.ServiceName, serviceValue, clientv3.WithLease(resp.ID)); err != nil {
			log.Printf("etcd refresh service `%s` with ttl to etcd3 failed: %v", info.ServiceName, err)
			return err
		}
	}
	log.Printf("etcd service `%s` register success", info.ServiceName)
	return nil
}

func (etcd *etcdRegisterImpl) Unregister(info ServiceDescInfo) error {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   etcd.Endpoints,
		DialTimeout: etcd.DialTimeout,
	})
	if err != nil {
		log.Printf("etcd unregister service `%s` error when connect to etcd server: %s", info.ServiceName, err.Error())
		return err
	}

	if _, err := client.Delete(context.Background(), info.ServiceName); err != nil {
		log.Printf("etcd: unregister `%s` failed: %s", info.ServiceName, err.Error())
	} else {
		log.Printf("etcd: unregister `%s` ok.", info.ServiceName)
	}
	return err
}
