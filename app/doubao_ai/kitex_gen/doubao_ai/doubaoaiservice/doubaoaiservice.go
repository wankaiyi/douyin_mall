// Code generated by Kitex v0.9.1. DO NOT EDIT.

package doubaoaiservice

import (
	"context"
	doubao_ai "douyin_mall/doubao_ai/kitex_gen/doubao_ai"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"AnalyzeSearchOrderQuestion": kitex.NewMethodInfo(
		analyzeSearchOrderQuestionHandler,
		newAnalyzeSearchOrderQuestionArgs,
		newAnalyzeSearchOrderQuestionResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	doubaoAiServiceServiceInfo                = NewServiceInfo()
	doubaoAiServiceServiceInfoForClient       = NewServiceInfoForClient()
	doubaoAiServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return doubaoAiServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return doubaoAiServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return doubaoAiServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "DoubaoAiService"
	handlerType := (*doubao_ai.DoubaoAiService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "doubao_ai",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func analyzeSearchOrderQuestionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(doubao_ai.SearchOrderQuestionReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(doubao_ai.DoubaoAiService).AnalyzeSearchOrderQuestion(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AnalyzeSearchOrderQuestionArgs:
		success, err := handler.(doubao_ai.DoubaoAiService).AnalyzeSearchOrderQuestion(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AnalyzeSearchOrderQuestionResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAnalyzeSearchOrderQuestionArgs() interface{} {
	return &AnalyzeSearchOrderQuestionArgs{}
}

func newAnalyzeSearchOrderQuestionResult() interface{} {
	return &AnalyzeSearchOrderQuestionResult{}
}

type AnalyzeSearchOrderQuestionArgs struct {
	Req *doubao_ai.SearchOrderQuestionReq
}

func (p *AnalyzeSearchOrderQuestionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(doubao_ai.SearchOrderQuestionReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AnalyzeSearchOrderQuestionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AnalyzeSearchOrderQuestionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AnalyzeSearchOrderQuestionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AnalyzeSearchOrderQuestionArgs) Unmarshal(in []byte) error {
	msg := new(doubao_ai.SearchOrderQuestionReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AnalyzeSearchOrderQuestionArgs_Req_DEFAULT *doubao_ai.SearchOrderQuestionReq

func (p *AnalyzeSearchOrderQuestionArgs) GetReq() *doubao_ai.SearchOrderQuestionReq {
	if !p.IsSetReq() {
		return AnalyzeSearchOrderQuestionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AnalyzeSearchOrderQuestionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AnalyzeSearchOrderQuestionArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AnalyzeSearchOrderQuestionResult struct {
	Success *doubao_ai.SearchOrderQuestionResp
}

var AnalyzeSearchOrderQuestionResult_Success_DEFAULT *doubao_ai.SearchOrderQuestionResp

func (p *AnalyzeSearchOrderQuestionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(doubao_ai.SearchOrderQuestionResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AnalyzeSearchOrderQuestionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AnalyzeSearchOrderQuestionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AnalyzeSearchOrderQuestionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AnalyzeSearchOrderQuestionResult) Unmarshal(in []byte) error {
	msg := new(doubao_ai.SearchOrderQuestionResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AnalyzeSearchOrderQuestionResult) GetSuccess() *doubao_ai.SearchOrderQuestionResp {
	if !p.IsSetSuccess() {
		return AnalyzeSearchOrderQuestionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AnalyzeSearchOrderQuestionResult) SetSuccess(x interface{}) {
	p.Success = x.(*doubao_ai.SearchOrderQuestionResp)
}

func (p *AnalyzeSearchOrderQuestionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AnalyzeSearchOrderQuestionResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) AnalyzeSearchOrderQuestion(ctx context.Context, Req *doubao_ai.SearchOrderQuestionReq) (r *doubao_ai.SearchOrderQuestionResp, err error) {
	var _args AnalyzeSearchOrderQuestionArgs
	_args.Req = Req
	var _result AnalyzeSearchOrderQuestionResult
	if err = p.c.Call(ctx, "AnalyzeSearchOrderQuestion", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
