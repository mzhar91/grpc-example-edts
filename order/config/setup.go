package config

import (
	"time"
	
	"github.com/grpc-example-edts/order/helpers/apierror"
)

func GetTimeoutContext() time.Duration {
	return time.Duration(Env.CtxTimeout) * time.Second
}

func ApiSetup() {
	apierror.Setup()
}
