package pool

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/pool/obj"

type ICreateResponse interface {
	ToPoolObject() *obj.Pool
}