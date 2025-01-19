package httpexec

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/cresplanex/bloader/internal/logger"
	"github.com/cresplanex/bloader/internal/utils"
)

// RequestCountLimit represents the request count limit
type RequestCountLimit struct {
	Enabled bool
	Count   int
}

// MassRequestContent represents the request content
type MassRequestContent[Req ExecReq] struct {
	Req          Req
	Interval     time.Duration
	ResponseWait bool
	ResChan      chan<- ResponseContent
	CountLimit   RequestCountLimit
	ResponseType ResponseType
}

// MassRequestExecute executes the request
func (q MassRequestContent[Req]) MassRequestExecute(
	ctx context.Context,
	log logger.Logger,
) error {
	go func() {
		// defer close(q.ResChan) // TODO: close channel
		waitForResponse := q.ResponseWait
		var count int
		chanForWait := make(chan struct{})
		// defer close(chanForWait) // TODO: close channel

		client := &http.Client{
			Timeout: 10 * time.Minute,
			Transport: &utils.DelayedTransport{
				Transport: &http.Transport{
					MaxIdleConns:        200,
					MaxIdleConnsPerHost: 180,
					IdleConnTimeout:     5 * time.Minute,
				},
				// Delay:     2 * time.Second,
			},
		}
		ticker := time.NewTicker(q.Interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				log.Info(ctx, "request processing is interrupted due to context termination",
					logger.Value("on", "RequestContent.QueryExecute"))
				return
			case <-ticker.C:
				if count > 0 && waitForResponse {
					select {
					case <-ctx.Done():
						log.Info(ctx, "request processing is interrupted due to context termination",
							logger.Value("on", "RequestContent.QueryExecute"))
						return
					case <-chanForWait:
					}
				}

				if q.CountLimit.Enabled && count >= q.CountLimit.Count {
					log.Info(ctx, "request processing is interrupted due to count limit",
						logger.Value("on", "RequestContent.QueryExecute"))
					select {
					case <-ctx.Done():
						log.Info(ctx, "request processing is interrupted due to context termination",
							logger.Value("on", "RequestContent.QueryExecute"))
						return
					case q.ResChan <- ResponseContent{
						WithCountLimit: true,
					}: // do nothing
					}

					return
				}

				go func(countInternal int) {
					defer func() {
						if waitForResponse {
							chanForWait <- struct{}{}
						}
					}()

					req, err := q.Req.CreateRequest(ctx, log, countInternal)
					if err != nil {
						log.Error(ctx, "failed to create request",
							logger.Value("error", err), logger.Value("on", "RequestContent.QueryExecute"))
						select {
						case <-ctx.Done():
							log.Info(ctx, "request processing is interrupted due to context termination",
								logger.Value("on", "RequestContent.QueryExecute"))
							return
						case q.ResChan <- ResponseContent{
							Success:      false,
							HasSystemErr: true,
							Count:        countInternal,
						}: // do nothing
						}

						return
					}

					log.Debug(ctx, "sending request",

						logger.Value("url", req.URL),
						logger.Value("count", countInternal),
					)
					startTime := time.Now()
					resp, err := client.Do(req)
					endTime := time.Now()
					log.Debug(ctx, "received response",

						logger.Value("url", req.URL),
						logger.Value("count", countInternal),
					)
					if err != nil {
						log.Error(ctx, "response error",
							logger.Value("error", err), logger.Value("url", req.URL))
						select {
						case <-ctx.Done():
							log.Info(ctx, "request processing is interrupted due to context termination",
								logger.Value("url", req.URL))
							return
						case q.ResChan <- ResponseContent{
							Success:      false,
							StartTime:    startTime,
							EndTime:      endTime,
							Count:        countInternal,
							ResponseTime: endTime.Sub(startTime).Milliseconds(),
							HasSystemErr: true,
						}: // do nothing
							log.Error(ctx, "response error",
								logger.Value("startTime", startTime),
								logger.Value("endTime", endTime),
								logger.Value("count", countInternal),
								logger.Value("responseTime", endTime.Sub(startTime).Milliseconds()),
								logger.Value("err", err),
							)
						}

						return
					}
					defer resp.Body.Close()

					statusCode := resp.StatusCode
					var response any
					responseByte, err := io.ReadAll(resp.Body)
					if err != nil {
						log.Error(ctx, "failed to read response",
							logger.Value("error", err), logger.Value("url", req.URL))
						select {
						case <-ctx.Done():
							log.Info(ctx, "request processing is interrupted due to context termination",
								logger.Value("url", req.URL))
							return
						case q.ResChan <- ResponseContent{
							Success:        false,
							Res:            response,
							StartTime:      startTime,
							EndTime:        endTime,
							Count:          countInternal,
							ResponseTime:   endTime.Sub(startTime).Milliseconds(),
							StatusCode:     statusCode,
							ParseResHasErr: true,
						}: // do nothing
							log.Error(ctx, "failed to read response",
								logger.Value("responseByte", string(responseByte)),
								logger.Value("response", response),
								logger.Value("startTime", startTime),
								logger.Value("endTime", endTime),
								logger.Value("count", countInternal),
								logger.Value("statusCode", statusCode),
								logger.Value("err", err),
							)
						}
						return
					}
					switch ResponseType(q.ResponseType) {
					case ResponseTypeJSON:
						err = json.Unmarshal(responseByte, &response)
					case ResponseTypeXML:
						err = xml.Unmarshal(responseByte, &response)
					case ResponseTypeYAML:
						err = yaml.Unmarshal(responseByte, &response)
					case ResponseTypeText:
						response = string(responseByte)
					case ResponseTypeHTML:
						response = string(responseByte)
					default:
						err = fmt.Errorf("invalid response type: %s", q.ResponseType)
					}
					if err != nil {
						log.Error(ctx, "failed to parse response",
							logger.Value("error", err), logger.Value("url", req.URL))
						select {
						case <-ctx.Done():
							log.Info(ctx, "request processing is interrupted due to context termination",
								logger.Value("url", req.URL))
							return
						case q.ResChan <- ResponseContent{
							Success:        false,
							Res:            response,
							ByteResponse:   responseByte,
							StartTime:      startTime,
							EndTime:        endTime,
							Count:          countInternal,
							ResponseTime:   endTime.Sub(startTime).Milliseconds(),
							StatusCode:     statusCode,
							ParseResHasErr: true,
						}: // do nothing
							log.Error(ctx, "failed to parse response",
								logger.Value("responseByte", string(responseByte)),
								logger.Value("response", response),
								logger.Value("startTime", startTime),
								logger.Value("endTime", endTime),
								logger.Value("count", countInternal),
								logger.Value("statusCode", statusCode),
								logger.Value("err", err),
							)
						}
						return
					}

					log.Debug(ctx, "response OK",
						logger.Value("url", req.URL))
					responseContent := ResponseContent{
						Success:      true,
						ByteResponse: responseByte,
						Res:          response,
						StartTime:    startTime,
						EndTime:      endTime,
						Count:        countInternal,
						ResponseTime: endTime.Sub(startTime).Milliseconds(),
						StatusCode:   statusCode,
					}
					select {
					case q.ResChan <- responseContent:
					case <-ctx.Done():
						log.Info(ctx, "request processing is interrupted due to context termination",
							logger.Value("url", req.URL))
						return
					}
				}(count)

				count++
			}
		}
	}()

	return nil
}

var _ MassRequestExecutor = MassRequestContent[ExecReq]{}
