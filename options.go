package sdk

import (
	"github.com/Ccheers/crawlab-go-sdk/interfaces"
	"github.com/apex/log"
)

type LoggerOption func(l log.Interface)

type ClientOption func(c interfaces.Client)

type ResultServiceOption func(svc interfaces.ResultService)
