package runner

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/cresplanex/bloader/internal/executor/httpexec"
	"github.com/cresplanex/bloader/internal/logger"
	"github.com/cresplanex/bloader/internal/runner/matcher"
)

// WriteData represents the write data
type WriteData struct {
	Success          bool
	SendDatetime     string
	ReceivedDatetime string
	Count            int
	ResponseTime     int
	StatusCode       string
	RawData          any
}

// ToSlice converts WriteData to slice
func (d WriteData) ToSlice() []string {
	return []string{
		strconv.FormatBool(d.Success),
		d.SendDatetime,
		d.ReceivedDatetime,
		strconv.Itoa(d.Count),
		strconv.Itoa(d.ResponseTime),
		d.StatusCode,
	}
}

type writeSendData struct {
	uid       uuid.UUID
	writeData WriteData
}

// ResponseDataConsumer represents the response data consumer
type ResponseDataConsumer func(
	ctx context.Context,
	log logger.Logger,
	id int,
	data WriteData,
) error

func runResponseHandler(
	ctx context.Context,
	reqTermChan <-chan struct{},
	log logger.Logger,
	id int,
	request ValidMassExecRequest,
	termChan chan<- TermChanType,
	writeErrChan <-chan struct{},
	uidChan <-chan uuid.UUID,
	resChan <-chan httpexec.ResponseContent,
	writeChan chan<- writeSendData,
) {
	defer close(termChan)
	var timeout <-chan time.Time
	if request.Break.Time.Enabled && request.Break.Time.Time > 0 {
		timeout = time.After(request.Break.Time.Time)
	}
	sentUID := make(map[uuid.UUID]struct{})
	for {
		select {
		case uid := <-uidChan:
			delete(sentUID, uid)
		case <-writeErrChan:
			log.Warn(ctx, "Term Condition: Write Error",
				logger.Value("id", id), logger.Value("on", "runResponseHandler"))
			for len(sentUID) > 0 {
				select {
				case uid := <-uidChan:
					delete(sentUID, uid)
				case <-reqTermChan:
					return
				}
			}
			select {
			case termChan <- NewTermChanType(matcher.TerminateTypeByWriteError, ""):
			case <-reqTermChan:
				return
			}
			return
		case <-timeout:
			sentLen := len(sentUID)
			writeErr := false
			for sentLen > 0 {
				select {
				case <-reqTermChan:
					return
				case uid := <-uidChan:
					delete(sentUID, uid)
					sentLen--
				case <-writeErrChan:
					log.Warn(ctx, "write error occurred",
						logger.Value("id", id), logger.Value("on", "runResponseHandler"))
					writeErr = true
				}
			}
			if writeErr {
				log.Warn(ctx, "Term Condition: Write Error",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"))
				select {
				case termChan <- NewTermChanType(matcher.TerminateTypeByWriteError, ""):
				case <-reqTermChan:
					return
				}
				return
			}
			log.Info(ctx, "Term Condition: Time",
				logger.Value("id", id), logger.Value("on", "runResponseHandler"))
			select {
			case termChan <- NewTermChanType(matcher.TerminateTypeByTimeout, ""):
			case <-reqTermChan:
				return
			}
			return
		case <-ctx.Done():
			sentLen := len(sentUID)
			writeErr := false
			for sentLen > 0 {
				select {
				case <-reqTermChan:
					return
				case uid := <-uidChan:
					delete(sentUID, uid)
					sentLen--
				case <-writeErrChan:
					log.Warn(ctx, "write error occurred",
						logger.Value("id", id), logger.Value("on", "runResponseHandler"))
					writeErr = true
				}
			}
			if writeErr {
				log.Warn(ctx, "Term Condition: Write Error",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"))
				select {
				case termChan <- NewTermChanType(matcher.TerminateTypeByWriteError, ""):
				case <-reqTermChan:
					return
				}
				return
			}
			log.Info(ctx, "Term Condition: Context Done",
				logger.Value("id", id), logger.Value("on", "runResponseHandler"))
			select {
			case termChan <- NewTermChanType(matcher.TerminateTypeByContext, ""):
			case <-reqTermChan:
				return
			}
			return
		case v := <-resChan:
			mustWrite := true
			var response any
			err := json.Unmarshal(v.ByteResponse, &response)
			if err != nil {
				log.Error(ctx, "The response is not a valid JSON",
					logger.Value("error", err), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
			}
			_, isMatch := request.RecordExcludeFilter.CountFilter(v.Count)
			if isMatch {
				log.Debug(ctx, "Count output filter found",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
				mustWrite = false
			}
			_, isMatch = request.RecordExcludeFilter.StatusCodeFilter(v.StatusCode)
			if isMatch {
				log.Debug(ctx, "Status output filter found",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
				mustWrite = false
			}
			var matchID string
			matchID, isMatch, err = request.RecordExcludeFilter.ResponseBodyFilter(response)
			if err != nil {
				log.Error(ctx, "failed to search jmespath",
					logger.Value("error", err), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
				sentLen := len(sentUID)
				writeErr := false
				for sentLen > 0 {
					select {
					case <-reqTermChan:
						return
					case uid := <-uidChan:
						delete(sentUID, uid)
						sentLen--
					case <-writeErrChan:
						log.Warn(ctx, "write error occurred",
							logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
						writeErr = true
					}
				}
				if writeErr {
					log.Warn(ctx, "Term Condition: Write Error",
						logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
					select {
					case termChan <- NewTermChanType(matcher.TerminateTypeByWriteError, ""):
					case <-reqTermChan:
						return
					}
					return
				}
				log.Info(ctx, "Term Condition: Response Body Write Filter Error",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
				select {
				case termChan <- NewTermChanType(matcher.TerminateTypeByResponseBodyWriteFilterError, matchID):
				case <-reqTermChan:
					return
				}
				return
			}
			if isMatch {
				log.Debug(ctx, "Response output filter found",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
				mustWrite = false
			}

			if mustWrite {
				uid := uuid.New()
				writeData := WriteData{
					Success:          v.Success,
					SendDatetime:     v.StartTime.Format(time.RFC3339Nano),
					ReceivedDatetime: v.EndTime.Format(time.RFC3339Nano),
					Count:            v.Count,
					ResponseTime:     int(v.ResponseTime),
					StatusCode:       strconv.Itoa(v.StatusCode),
					RawData:          response,
				}
				sentUID[uid] = struct{}{}
				go func() {
					select {
					case <-reqTermChan:
						return
					case writeChan <- writeSendData{
						uid:       uid,
						writeData: writeData,
					}:
					}
				}()
			}

			if v.ReqCreateHasErr {
				sentLen := len(sentUID)
				for sentLen > 0 {
					select {
					case <-reqTermChan:
						return
					case uid := <-uidChan:
						delete(sentUID, uid)
						sentLen--
					case <-writeErrChan:
						log.Warn(ctx, "write error occurred",
							logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
					}
				}
				log.Warn(ctx, "Term Condition: Request Creation Error",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
				select {
				case termChan <- NewTermChanType(matcher.TerminateTypeByCreateRequestError, ""):
				case <-reqTermChan:
					return
				}
				return
			}
			if v.HasSystemErr {
				if request.Break.SysError {
					sentLen := len(sentUID)
					for sentLen > 0 {
						select {
						case <-reqTermChan:
							return
						case uid := <-uidChan:
							delete(sentUID, uid)
							sentLen--
						case <-writeErrChan:
							log.Warn(ctx, "write error occurred",
								logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
						}
					}
					log.Warn(ctx, "Term Condition: System Error",
						logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
					select {
					case termChan <- NewTermChanType(matcher.TerminateTypeBySystemError, ""):
					case <-reqTermChan:
						return
					}
					return
				}
				log.Warn(ctx, "System error occurred",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
			}
			if v.ParseResHasErr {
				if request.Break.ParseError {
					sentLen := len(sentUID)
					for sentLen > 0 {
						select {
						case <-reqTermChan:
							return
						case uid := <-uidChan:
							delete(sentUID, uid)
							sentLen--
						case <-writeErrChan:
							log.Warn(ctx, "write error occurred",
								logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
						}
					}
					log.Warn(ctx, "Term Condition: Response Parse Error",
						logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
					select {
					case termChan <- NewTermChanType(matcher.TerminateTypeByParseResponseError, ""):
					case <-reqTermChan:
						return
					}
					return
				}
				log.Warn(ctx, "Parse error occurred",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
			}
			if v.WithCountLimit {
				sentLen := len(sentUID)
				writeErr := false
				for sentLen > 0 {
					select {
					case <-reqTermChan:
						return
					case uid := <-uidChan:
						delete(sentUID, uid)
						sentLen--
					case <-writeErrChan:
						log.Warn(ctx, "write error occurred",
							logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
						writeErr = true
					}
				}
				if writeErr {
					log.Warn(ctx, "Term Condition: Write Error",
						logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
					select {
					case termChan <- NewTermChanType(matcher.TerminateTypeByWriteError, ""):
					case <-reqTermChan:
						return
					}
					return
				}

				log.Info(ctx, "Term Condition: Count Limit",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
				select {
				case termChan <- NewTermChanType(matcher.TerminateTypeByCount, ""):
				case <-reqTermChan:
					return
				}
				return
			}
			matchID, isMatch, err = request.Break.ResponseBodyMatcher(response)
			if err != nil {
				log.Error(ctx, "failed to search jmespath",
					logger.Value("error", err), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
				sentLen := len(sentUID)
				writeErr := false
				for sentLen > 0 {
					select {
					case <-reqTermChan:
						return
					case uid := <-uidChan:
						delete(sentUID, uid)
						sentLen--
					case <-writeErrChan:
						log.Warn(ctx, "write error occurred",
							logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
						writeErr = true
					}
				}
				if writeErr {
					log.Warn(ctx, "Term Condition: Write Error",
						logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
					select {
					case termChan <- NewTermChanType(matcher.TerminateTypeByWriteError, ""):
					case <-reqTermChan:
						return
					}
					return
				}

				log.Info(ctx, "Term Condition: Response Body Break Filter Error",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
				select {
				case termChan <- NewTermChanType(matcher.TerminateTypeByResponseBodyBreakFilterError, matchID):
				case <-reqTermChan:
					return
				}
				return
			}
			if isMatch {
				sentLen := len(sentUID)
				writeErr := false
				for sentLen > 0 {
					select {
					case <-reqTermChan:
						return
					case uid := <-uidChan:
						delete(sentUID, uid)
						sentLen--
					case <-writeErrChan:
						log.Warn(ctx, "write error occurred",
							logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
						writeErr = true
					}
				}
				if writeErr {
					log.Warn(ctx, "Term Condition: Write Error",
						logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
					select {
					case termChan <- NewTermChanType(matcher.TerminateTypeByWriteError, ""):
					case <-reqTermChan:
						return
					}
					return
				}

				log.Info(ctx, "Term Condition: Response Body",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
				select {
				case termChan <- NewTermChanType(matcher.TerminateTypeByResponseBody, matchID):
				case <-reqTermChan:
					return
				}
				return
			}
			matchID, isMatch = request.Break.StatusCodeMatcher(v.StatusCode)
			if isMatch {
				sentLen := len(sentUID)
				writeErr := false
				for sentLen > 0 {
					select {
					case <-reqTermChan:
						return
					case uid := <-uidChan:
						delete(sentUID, uid)
						sentLen--
					case <-writeErrChan:
						log.Warn(ctx, "write error occurred",
							logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
						writeErr = true
					}
				}
				if writeErr {
					log.Warn(ctx, "Term Condition: Write Error",
						logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
					select {
					case termChan <- NewTermChanType(matcher.TerminateTypeByWriteError, ""):
					case <-reqTermChan:
						return
					}
					return
				}

				log.Info(ctx, "Term Condition: Status Code",
					logger.Value("id", id), logger.Value("on", "runResponseHandler"), logger.Value("count", v.Count))
				select {
				case termChan <- NewTermChanType(matcher.TerminateTypeByStatusCode, matchID):
				case <-reqTermChan:
					return
				}
				return
			}
		}
	}
}

// RunAsyncProcessing runs the async processing
func RunAsyncProcessing(
	ctx context.Context,
	reqTermChan <-chan struct{},
	log logger.Logger,
	id int,
	request ValidMassExecRequest,
	termChan chan<- TermChanType,
	resChan <-chan httpexec.ResponseContent,
	consumer ResponseDataConsumer,
) {
	writeChan := make(chan writeSendData)
	wroteUIDChan := make(chan uuid.UUID)
	writeErrChan := make(chan struct{})
	go func() {
		runResponseHandler(ctx, reqTermChan, log, id, request, termChan, writeErrChan, wroteUIDChan, resChan, writeChan)
	}()

	go func() {
		defer close(wroteUIDChan)
		defer close(writeErrChan)
		defer close(writeChan)
		for {
			select {
			case <-reqTermChan:
				return
			case d := <-writeChan:
				log.Debug(ctx, "Writing data",
					logger.Value("id", id), logger.Value("data", d), logger.Value("on", "runAsyncProcessing"))
				if err := consumer(ctx, log, id, d.writeData); err != nil {
					log.Error(ctx, "failed to write data",
						logger.Value("error", err), logger.Value("on", "runAsyncProcessing"))
					if request.Break.WriteError {
						select {
						case writeErrChan <- struct{}{}:
						case <-reqTermChan:
							return
						}
						select {
						case wroteUIDChan <- d.uid:
						case <-reqTermChan:
							return
						}
						continue
					}
				}
				select {
				case wroteUIDChan <- d.uid:
				case <-reqTermChan:
					return
				}
			}
		}
	}()
}
