package config

import (
	"os"
	"secskill/pkg/bootstrap"
	"secskill/pkg/config"
	"secskill/skapp/model"
	"sync"

	"github.com/go-kit/kit/log"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/spf13/viper"

	"github.com/openzipkin/zipkin-go"
)

const (
	kConfigType = "CONFIG_TYPE"
)

var ZipkinTracer *zipkin.Tracer
var Logger log.Logger

func init() {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	Logger = log.With(Logger, "caller", log.DefaultCaller)
	viper.AutomaticEnv()
	initDefault()
	if err := config.LoadRemoteConfig(); err != nil {
		Logger.Log("Fail to load remote config", err)
	}
	if err := config.Sub("mysql", &config.MysqlConfig); err != nil {
		Logger.Log("Fail to parse mysql", err)
	}
	if err := config.Sub("service", &config.SecKill); err != nil {
		Logger.Log("Fail to parse trace ", err)
	}

	if err := config.Sub("service", &config.Redis); err != nil {
		Logger.Log("Fail to parse redis ", err)
	}
}

func initTracer(zipkinUrl string) {
	var (
		err            error
		userNoopTracer = zipkinUrl == ""
		reporter       = zipkinhttp.NewReporter(zipkinUrl)
	)

	zEP, _ := zipkin.NewEndpoint(bootstrap.DiscoverConfig.ServiceName, bootstrap.HttpConfig.Port)
	ZipkinTracer, err = zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(zEP), zipkin.WithNoopTracer(userNoopTracer))
	if err != nil {
		Logger.Log("err", err)
		os.Exit(1)
	}
	if !userNoopTracer {
		Logger.Log("tracer", "Zipkin", "type", "Native", "URL", zipkinUrl)
	}
}
func initDefault() {
	viper.SetDefault(kConfigType, "yaml")
}

var SkAppContext = &SkAppCtx{
	UserConnMap: make(map[string]chan *model.SecResult, 1024),
	SecReqChan:  make(chan *model.SecRequest, 1024),
}

type SkAppCtx struct {
	SecReqChan       chan *model.SecRequest
	SecReqChanSize   int
	RWSecProductLock sync.RWMutex

	UserConnMap     map[string]chan *model.SecResult
	UserConnMapLock sync.Mutex
}

const (
	ProductStatusNormal       = 0 //商品状态正常
	ProductStatusSaleOut      = 1 //商品售罄
	ProductStatusForceSaleOut = 2 //商品强制售罄
)
