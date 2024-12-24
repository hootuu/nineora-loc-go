package restx

import (
	"github.com/go-resty/resty/v2"
	"github.com/hootuu/gelato/configure"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/logger"
	"github.com/hootuu/gelato/sys"
	"github.com/hootuu/nineorai/io"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var gLogger = logger.GetLogger("nineora")

var gCli = NewClient()

func Rest[REQ io.RequestData, RESP any](path string, req *io.Request[REQ]) *io.Response[RESP] {
	if req == nil {
		return io.FailResponse[RESP]("UNKNOWN_ID", errors.Verify("require req"))
	}
	if len(req.Signatures) == 0 {
		return io.FailResponse[RESP](req.ID, errors.Verify("require signatures"))
	}
	bodyDataBytes, err := req.Marshal()
	if err != nil {
		return io.FailResponse[RESP](req.ID, err)
	}
	var restResp io.Response[RESP]
	restReq := gCli.R().SetBody(bodyDataBytes)
	if gLogger.Level() < zapcore.InfoLevel {
		gLogger.Info("call rest req", zap.String("path", path), zap.Any("req", req))
		s := time.Now().UnixMilli()
		defer func() {
			gLogger.Info("call rest req back", zap.String("path", path), zap.Any("resp", &restResp),
				zap.Int64("elapse", time.Now().UnixMilli()-s))
		}()
	}
	_, nErr := restReq.SetResult(&restResp).Post(path)
	if nErr != nil {
		return io.FailResponse[RESP](req.ID, errors.System("rest failed:"+nErr.Error(), nErr))
	}
	return &restResp
}

func NewClient() *resty.Client {
	cfgBaseUrl := configure.GetString("nineora.loc.gw", "http://localhost:8080")
	cfgRetryWaitTime := configure.GetDuration("nineora.loc.wait.retry", 2)
	cfgRetryMaxWaitTime := configure.GetDuration("nineora.loc.wait.retry.max", 10)
	cfgTimeout := configure.GetDuration("nineora.loc.timeout", 60)
	cli := resty.New().
		SetBaseURL(cfgBaseUrl).
		SetRetryWaitTime(cfgRetryWaitTime * time.Second).
		SetRetryMaxWaitTime(cfgRetryMaxWaitTime * time.Second).
		SetTimeout(cfgTimeout * time.Second).
		//SetPreRequestHook(doBeforeRequest).
		OnAfterResponse(doAfterResponse)
	if sys.RunMode.IsRd() {
		cli.EnableTrace()
	}
	return cli
}

//func doBeforeRequest(_ *resty.Client, request *http.Request) error {
//	return nil
//}

func doAfterResponse(_ *resty.Client, r *resty.Response) error {
	if r == nil {
		return errors.System("response is nil")
	}
	if !r.IsSuccess() {
		return errors.System("response is not success")
	}

	return nil
}
