// Code generated by protoc-gen-go.
// source: google/cloud/dialogflow/v2beta1/agent.proto
// DO NOT EDIT!

/*
Package dialogflow is a generated protocol buffer package.

It is generated from these files:
	google/cloud/dialogflow/v2beta1/agent.proto

It has these top-level messages:
	Agent
	GetAgentRequest
	SearchAgentsRequest
	SearchAgentsResponse
	TrainAgentRequest
	ExportAgentRequest
	ExportAgentResponse
	ImportAgentRequest
	RestoreAgentRequest
*/
package dialogflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_longrunning "google.golang.org/genproto/googleapis/longrunning"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/protobuf/field_mask"
import _ "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Match mode determines how intents are detected from user queries.
type Agent_MatchMode int32

const (
	// Not specified.
	Agent_MATCH_MODE_UNSPECIFIED Agent_MatchMode = 0
	// Best for agents with a small number of examples in intents and/or wide
	// use of templates syntax and composite entities.
	Agent_MATCH_MODE_HYBRID Agent_MatchMode = 1
	// Can be used for agents with a large number of examples in intents,
	// especially the ones using @sys.any or very large developer entities.
	Agent_MATCH_MODE_ML_ONLY Agent_MatchMode = 2
)

var Agent_MatchMode_name = map[int32]string{
	0: "MATCH_MODE_UNSPECIFIED",
	1: "MATCH_MODE_HYBRID",
	2: "MATCH_MODE_ML_ONLY",
}
var Agent_MatchMode_value = map[string]int32{
	"MATCH_MODE_UNSPECIFIED": 0,
	"MATCH_MODE_HYBRID":      1,
	"MATCH_MODE_ML_ONLY":     2,
}

func (x Agent_MatchMode) String() string {
	return proto.EnumName(Agent_MatchMode_name, int32(x))
}
func (Agent_MatchMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Represents a conversational agent.
type Agent struct {
	// Required. The project of this agent.
	// Format: `projects/<Project ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Required. The name of this agent.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	// Required. The default language of the agent as a language tag. See
	// [Language Support](https://dialogflow.com/docs/reference/language) for a
	// list of the currently supported language codes.
	// This field cannot be set by the `Update` method.
	DefaultLanguageCode string `protobuf:"bytes,3,opt,name=default_language_code,json=defaultLanguageCode" json:"default_language_code,omitempty"`
	// Optional. The list of all languages supported by this agent (except for the
	// `default_language_code`).
	SupportedLanguageCodes []string `protobuf:"bytes,4,rep,name=supported_language_codes,json=supportedLanguageCodes" json:"supported_language_codes,omitempty"`
	// Required. The time zone of this agent from the
	// [time zone database](https://www.iana.org/time-zones), e.g.,
	// America/New_York, Europe/Paris.
	TimeZone string `protobuf:"bytes,5,opt,name=time_zone,json=timeZone" json:"time_zone,omitempty"`
	// Optional. The description of this agent.
	// The maximum length is 500 characters. If exceeded, the request is rejected.
	Description string `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	// Optional. The URI of the agent's avatar.
	// Avatars are used throughout API.AI console and in the self-hosted
	// [Web Demo](https://dialogflow.com/docs/integrations/web-demo) integration.
	AvatarUri string `protobuf:"bytes,7,opt,name=avatar_uri,json=avatarUri" json:"avatar_uri,omitempty"`
	// Optional. Determines whether this agent should log conversation queries.
	EnableLogging bool `protobuf:"varint,8,opt,name=enable_logging,json=enableLogging" json:"enable_logging,omitempty"`
	// Optional. Determines how intents are detected from user queries.
	MatchMode Agent_MatchMode `protobuf:"varint,9,opt,name=match_mode,json=matchMode,enum=google.cloud.dialogflow.v2beta1.Agent_MatchMode" json:"match_mode,omitempty"`
	// Optional. To filter out false positive results and still get variety in
	// matched natural language inputs for your agent, you can tune the machine
	// learning classification threshold. If the returned score value is less than
	// the threshold value, then a fallback intent is be triggered or, if there
	// are no fallback intents defined, no intent will be triggered. The score
	// values range from 0.0 (completely uncertain) to 1.0 (completely certain).
	// If set to 0.0, the default of 0.3 is used.
	ClassificationThreshold float32 `protobuf:"fixed32,10,opt,name=classification_threshold,json=classificationThreshold" json:"classification_threshold,omitempty"`
}

func (m *Agent) Reset()                    { *m = Agent{} }
func (m *Agent) String() string            { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()               {}
func (*Agent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The request message for [Agents.GetAgent].
type GetAgentRequest struct {
	// Required. The project that the agent to fetch is associated with.
	// Format: `projects/<Project ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
}

func (m *GetAgentRequest) Reset()                    { *m = GetAgentRequest{} }
func (m *GetAgentRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAgentRequest) ProtoMessage()               {}
func (*GetAgentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// The request message for [Agents.SearchAgents].
type SearchAgentsRequest struct {
	// Required. The project to list agents from.
	// Format: `projects/<Project ID or '-'>`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. The maximum number of items to return in a single page. By
	// default 100 and at most 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// Optional. The next_page_token value returned from a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
}

func (m *SearchAgentsRequest) Reset()                    { *m = SearchAgentsRequest{} }
func (m *SearchAgentsRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchAgentsRequest) ProtoMessage()               {}
func (*SearchAgentsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// The response message for [Agents.SearchAgents].
type SearchAgentsResponse struct {
	// The list of agents. There will be a maximum number of items returned based
	// on the page_size field in the request.
	Agents []*Agent `protobuf:"bytes,1,rep,name=agents" json:"agents,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
}

func (m *SearchAgentsResponse) Reset()                    { *m = SearchAgentsResponse{} }
func (m *SearchAgentsResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchAgentsResponse) ProtoMessage()               {}
func (*SearchAgentsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SearchAgentsResponse) GetAgents() []*Agent {
	if m != nil {
		return m.Agents
	}
	return nil
}

// The request message for [Agents.TrainAgent].
type TrainAgentRequest struct {
	// Required. The project that the agent to train is associated with.
	// Format: `projects/<Project ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
}

func (m *TrainAgentRequest) Reset()                    { *m = TrainAgentRequest{} }
func (m *TrainAgentRequest) String() string            { return proto.CompactTextString(m) }
func (*TrainAgentRequest) ProtoMessage()               {}
func (*TrainAgentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// The request message for [Agents.ExportAgent].
type ExportAgentRequest struct {
	// Required. The project that the agent to export is associated with.
	// Format: `projects/<Project ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Optional. The URI to export the agent to. Note: The URI must start with
	// "gs://". If left unspecified, the serialized agent is returned inline.
	AgentUri string `protobuf:"bytes,2,opt,name=agent_uri,json=agentUri" json:"agent_uri,omitempty"`
}

func (m *ExportAgentRequest) Reset()                    { *m = ExportAgentRequest{} }
func (m *ExportAgentRequest) String() string            { return proto.CompactTextString(m) }
func (*ExportAgentRequest) ProtoMessage()               {}
func (*ExportAgentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// The response message for [Agents.ExportAgent].
type ExportAgentResponse struct {
	// Required. The exported agent.
	//
	// Types that are valid to be assigned to Agent:
	//	*ExportAgentResponse_AgentUri
	//	*ExportAgentResponse_AgentContent
	Agent isExportAgentResponse_Agent `protobuf_oneof:"agent"`
}

func (m *ExportAgentResponse) Reset()                    { *m = ExportAgentResponse{} }
func (m *ExportAgentResponse) String() string            { return proto.CompactTextString(m) }
func (*ExportAgentResponse) ProtoMessage()               {}
func (*ExportAgentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isExportAgentResponse_Agent interface {
	isExportAgentResponse_Agent()
}

type ExportAgentResponse_AgentUri struct {
	AgentUri string `protobuf:"bytes,1,opt,name=agent_uri,json=agentUri,oneof"`
}
type ExportAgentResponse_AgentContent struct {
	AgentContent []byte `protobuf:"bytes,2,opt,name=agent_content,json=agentContent,proto3,oneof"`
}

func (*ExportAgentResponse_AgentUri) isExportAgentResponse_Agent()     {}
func (*ExportAgentResponse_AgentContent) isExportAgentResponse_Agent() {}

func (m *ExportAgentResponse) GetAgent() isExportAgentResponse_Agent {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *ExportAgentResponse) GetAgentUri() string {
	if x, ok := m.GetAgent().(*ExportAgentResponse_AgentUri); ok {
		return x.AgentUri
	}
	return ""
}

func (m *ExportAgentResponse) GetAgentContent() []byte {
	if x, ok := m.GetAgent().(*ExportAgentResponse_AgentContent); ok {
		return x.AgentContent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ExportAgentResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ExportAgentResponse_OneofMarshaler, _ExportAgentResponse_OneofUnmarshaler, _ExportAgentResponse_OneofSizer, []interface{}{
		(*ExportAgentResponse_AgentUri)(nil),
		(*ExportAgentResponse_AgentContent)(nil),
	}
}

func _ExportAgentResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ExportAgentResponse)
	// agent
	switch x := m.Agent.(type) {
	case *ExportAgentResponse_AgentUri:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AgentUri)
	case *ExportAgentResponse_AgentContent:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.AgentContent)
	case nil:
	default:
		return fmt.Errorf("ExportAgentResponse.Agent has unexpected type %T", x)
	}
	return nil
}

func _ExportAgentResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ExportAgentResponse)
	switch tag {
	case 1: // agent.agent_uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Agent = &ExportAgentResponse_AgentUri{x}
		return true, err
	case 2: // agent.agent_content
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Agent = &ExportAgentResponse_AgentContent{x}
		return true, err
	default:
		return false, nil
	}
}

func _ExportAgentResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ExportAgentResponse)
	// agent
	switch x := m.Agent.(type) {
	case *ExportAgentResponse_AgentUri:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AgentUri)))
		n += len(x.AgentUri)
	case *ExportAgentResponse_AgentContent:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AgentContent)))
		n += len(x.AgentContent)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The request message for [Agents.ImportAgent].
type ImportAgentRequest struct {
	// Required. The project that the agent to import is associated with.
	// Format: `projects/<Project ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Required. The agent to import.
	//
	// Types that are valid to be assigned to Agent:
	//	*ImportAgentRequest_AgentUri
	//	*ImportAgentRequest_AgentContent
	Agent isImportAgentRequest_Agent `protobuf_oneof:"agent"`
}

func (m *ImportAgentRequest) Reset()                    { *m = ImportAgentRequest{} }
func (m *ImportAgentRequest) String() string            { return proto.CompactTextString(m) }
func (*ImportAgentRequest) ProtoMessage()               {}
func (*ImportAgentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isImportAgentRequest_Agent interface {
	isImportAgentRequest_Agent()
}

type ImportAgentRequest_AgentUri struct {
	AgentUri string `protobuf:"bytes,2,opt,name=agent_uri,json=agentUri,oneof"`
}
type ImportAgentRequest_AgentContent struct {
	AgentContent []byte `protobuf:"bytes,3,opt,name=agent_content,json=agentContent,proto3,oneof"`
}

func (*ImportAgentRequest_AgentUri) isImportAgentRequest_Agent()     {}
func (*ImportAgentRequest_AgentContent) isImportAgentRequest_Agent() {}

func (m *ImportAgentRequest) GetAgent() isImportAgentRequest_Agent {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *ImportAgentRequest) GetAgentUri() string {
	if x, ok := m.GetAgent().(*ImportAgentRequest_AgentUri); ok {
		return x.AgentUri
	}
	return ""
}

func (m *ImportAgentRequest) GetAgentContent() []byte {
	if x, ok := m.GetAgent().(*ImportAgentRequest_AgentContent); ok {
		return x.AgentContent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ImportAgentRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ImportAgentRequest_OneofMarshaler, _ImportAgentRequest_OneofUnmarshaler, _ImportAgentRequest_OneofSizer, []interface{}{
		(*ImportAgentRequest_AgentUri)(nil),
		(*ImportAgentRequest_AgentContent)(nil),
	}
}

func _ImportAgentRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ImportAgentRequest)
	// agent
	switch x := m.Agent.(type) {
	case *ImportAgentRequest_AgentUri:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AgentUri)
	case *ImportAgentRequest_AgentContent:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.AgentContent)
	case nil:
	default:
		return fmt.Errorf("ImportAgentRequest.Agent has unexpected type %T", x)
	}
	return nil
}

func _ImportAgentRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ImportAgentRequest)
	switch tag {
	case 2: // agent.agent_uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Agent = &ImportAgentRequest_AgentUri{x}
		return true, err
	case 3: // agent.agent_content
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Agent = &ImportAgentRequest_AgentContent{x}
		return true, err
	default:
		return false, nil
	}
}

func _ImportAgentRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ImportAgentRequest)
	// agent
	switch x := m.Agent.(type) {
	case *ImportAgentRequest_AgentUri:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AgentUri)))
		n += len(x.AgentUri)
	case *ImportAgentRequest_AgentContent:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AgentContent)))
		n += len(x.AgentContent)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// The request message for [Agents.RestoreAgent].
type RestoreAgentRequest struct {
	// Required. The project that the agent to restore is associated with.
	// Format: `projects/<Project ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// Required. The agent to restore.
	//
	// Types that are valid to be assigned to Agent:
	//	*RestoreAgentRequest_AgentUri
	//	*RestoreAgentRequest_AgentContent
	Agent isRestoreAgentRequest_Agent `protobuf_oneof:"agent"`
}

func (m *RestoreAgentRequest) Reset()                    { *m = RestoreAgentRequest{} }
func (m *RestoreAgentRequest) String() string            { return proto.CompactTextString(m) }
func (*RestoreAgentRequest) ProtoMessage()               {}
func (*RestoreAgentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type isRestoreAgentRequest_Agent interface {
	isRestoreAgentRequest_Agent()
}

type RestoreAgentRequest_AgentUri struct {
	AgentUri string `protobuf:"bytes,2,opt,name=agent_uri,json=agentUri,oneof"`
}
type RestoreAgentRequest_AgentContent struct {
	AgentContent []byte `protobuf:"bytes,3,opt,name=agent_content,json=agentContent,proto3,oneof"`
}

func (*RestoreAgentRequest_AgentUri) isRestoreAgentRequest_Agent()     {}
func (*RestoreAgentRequest_AgentContent) isRestoreAgentRequest_Agent() {}

func (m *RestoreAgentRequest) GetAgent() isRestoreAgentRequest_Agent {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *RestoreAgentRequest) GetAgentUri() string {
	if x, ok := m.GetAgent().(*RestoreAgentRequest_AgentUri); ok {
		return x.AgentUri
	}
	return ""
}

func (m *RestoreAgentRequest) GetAgentContent() []byte {
	if x, ok := m.GetAgent().(*RestoreAgentRequest_AgentContent); ok {
		return x.AgentContent
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RestoreAgentRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RestoreAgentRequest_OneofMarshaler, _RestoreAgentRequest_OneofUnmarshaler, _RestoreAgentRequest_OneofSizer, []interface{}{
		(*RestoreAgentRequest_AgentUri)(nil),
		(*RestoreAgentRequest_AgentContent)(nil),
	}
}

func _RestoreAgentRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RestoreAgentRequest)
	// agent
	switch x := m.Agent.(type) {
	case *RestoreAgentRequest_AgentUri:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AgentUri)
	case *RestoreAgentRequest_AgentContent:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.AgentContent)
	case nil:
	default:
		return fmt.Errorf("RestoreAgentRequest.Agent has unexpected type %T", x)
	}
	return nil
}

func _RestoreAgentRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RestoreAgentRequest)
	switch tag {
	case 2: // agent.agent_uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Agent = &RestoreAgentRequest_AgentUri{x}
		return true, err
	case 3: // agent.agent_content
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Agent = &RestoreAgentRequest_AgentContent{x}
		return true, err
	default:
		return false, nil
	}
}

func _RestoreAgentRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RestoreAgentRequest)
	// agent
	switch x := m.Agent.(type) {
	case *RestoreAgentRequest_AgentUri:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AgentUri)))
		n += len(x.AgentUri)
	case *RestoreAgentRequest_AgentContent:
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.AgentContent)))
		n += len(x.AgentContent)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Agent)(nil), "google.cloud.dialogflow.v2beta1.Agent")
	proto.RegisterType((*GetAgentRequest)(nil), "google.cloud.dialogflow.v2beta1.GetAgentRequest")
	proto.RegisterType((*SearchAgentsRequest)(nil), "google.cloud.dialogflow.v2beta1.SearchAgentsRequest")
	proto.RegisterType((*SearchAgentsResponse)(nil), "google.cloud.dialogflow.v2beta1.SearchAgentsResponse")
	proto.RegisterType((*TrainAgentRequest)(nil), "google.cloud.dialogflow.v2beta1.TrainAgentRequest")
	proto.RegisterType((*ExportAgentRequest)(nil), "google.cloud.dialogflow.v2beta1.ExportAgentRequest")
	proto.RegisterType((*ExportAgentResponse)(nil), "google.cloud.dialogflow.v2beta1.ExportAgentResponse")
	proto.RegisterType((*ImportAgentRequest)(nil), "google.cloud.dialogflow.v2beta1.ImportAgentRequest")
	proto.RegisterType((*RestoreAgentRequest)(nil), "google.cloud.dialogflow.v2beta1.RestoreAgentRequest")
	proto.RegisterEnum("google.cloud.dialogflow.v2beta1.Agent_MatchMode", Agent_MatchMode_name, Agent_MatchMode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Agents service

type AgentsClient interface {
	// Retrieves the specified agent.
	GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*Agent, error)
	// Returns the list of agents.
	//
	// Since there is at most one conversational agent per project, this method is
	// useful primarily for listing all agents across projects the caller has
	// access to. One can achieve that with a wildcard project collection id "-".
	// Refer to [List
	// Sub-Collections](https://cloud.google.com/apis/design/design_patterns#list_sub-collections).
	SearchAgents(ctx context.Context, in *SearchAgentsRequest, opts ...grpc.CallOption) (*SearchAgentsResponse, error)
	// Trains the specified agent.
	//
	//
	// Operation<response: google.protobuf.Empty,
	//           metadata: google.protobuf.Struct>
	TrainAgent(ctx context.Context, in *TrainAgentRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error)
	// Exports the specified agent to a ZIP file.
	//
	//
	// Operation<response: ExportAgentResponse,
	//           metadata: google.protobuf.Struct>
	ExportAgent(ctx context.Context, in *ExportAgentRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error)
	// Imports the specified agent from a ZIP file.
	//
	// Uploads new intents and entity types without deleting the existing ones.
	// Intents and entity types with the same name are replaced with the new
	// versions from ImportAgentRequest.
	//
	//
	// Operation<response: google.protobuf.Empty,
	//           metadata: google.protobuf.Struct>
	ImportAgent(ctx context.Context, in *ImportAgentRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error)
	// Restores the specified agent from a ZIP file.
	//
	// Replaces the current agent version with a new one. All the intents and
	// entity types in the older version are deleted.
	//
	//
	// Operation<response: google.protobuf.Empty,
	//           metadata: google.protobuf.Struct>
	RestoreAgent(ctx context.Context, in *RestoreAgentRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error)
}

type agentsClient struct {
	cc *grpc.ClientConn
}

func NewAgentsClient(cc *grpc.ClientConn) AgentsClient {
	return &agentsClient{cc}
}

func (c *agentsClient) GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*Agent, error) {
	out := new(Agent)
	err := grpc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.Agents/GetAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) SearchAgents(ctx context.Context, in *SearchAgentsRequest, opts ...grpc.CallOption) (*SearchAgentsResponse, error) {
	out := new(SearchAgentsResponse)
	err := grpc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.Agents/SearchAgents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) TrainAgent(ctx context.Context, in *TrainAgentRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error) {
	out := new(google_longrunning.Operation)
	err := grpc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.Agents/TrainAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) ExportAgent(ctx context.Context, in *ExportAgentRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error) {
	out := new(google_longrunning.Operation)
	err := grpc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.Agents/ExportAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) ImportAgent(ctx context.Context, in *ImportAgentRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error) {
	out := new(google_longrunning.Operation)
	err := grpc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.Agents/ImportAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) RestoreAgent(ctx context.Context, in *RestoreAgentRequest, opts ...grpc.CallOption) (*google_longrunning.Operation, error) {
	out := new(google_longrunning.Operation)
	err := grpc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.Agents/RestoreAgent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Agents service

type AgentsServer interface {
	// Retrieves the specified agent.
	GetAgent(context.Context, *GetAgentRequest) (*Agent, error)
	// Returns the list of agents.
	//
	// Since there is at most one conversational agent per project, this method is
	// useful primarily for listing all agents across projects the caller has
	// access to. One can achieve that with a wildcard project collection id "-".
	// Refer to [List
	// Sub-Collections](https://cloud.google.com/apis/design/design_patterns#list_sub-collections).
	SearchAgents(context.Context, *SearchAgentsRequest) (*SearchAgentsResponse, error)
	// Trains the specified agent.
	//
	//
	// Operation<response: google.protobuf.Empty,
	//           metadata: google.protobuf.Struct>
	TrainAgent(context.Context, *TrainAgentRequest) (*google_longrunning.Operation, error)
	// Exports the specified agent to a ZIP file.
	//
	//
	// Operation<response: ExportAgentResponse,
	//           metadata: google.protobuf.Struct>
	ExportAgent(context.Context, *ExportAgentRequest) (*google_longrunning.Operation, error)
	// Imports the specified agent from a ZIP file.
	//
	// Uploads new intents and entity types without deleting the existing ones.
	// Intents and entity types with the same name are replaced with the new
	// versions from ImportAgentRequest.
	//
	//
	// Operation<response: google.protobuf.Empty,
	//           metadata: google.protobuf.Struct>
	ImportAgent(context.Context, *ImportAgentRequest) (*google_longrunning.Operation, error)
	// Restores the specified agent from a ZIP file.
	//
	// Replaces the current agent version with a new one. All the intents and
	// entity types in the older version are deleted.
	//
	//
	// Operation<response: google.protobuf.Empty,
	//           metadata: google.protobuf.Struct>
	RestoreAgent(context.Context, *RestoreAgentRequest) (*google_longrunning.Operation, error)
}

func RegisterAgentsServer(s *grpc.Server, srv AgentsServer) {
	s.RegisterService(&_Agents_serviceDesc, srv)
}

func _Agents_GetAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).GetAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.Agents/GetAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).GetAgent(ctx, req.(*GetAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_SearchAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).SearchAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.Agents/SearchAgents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).SearchAgents(ctx, req.(*SearchAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_TrainAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).TrainAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.Agents/TrainAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).TrainAgent(ctx, req.(*TrainAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_ExportAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).ExportAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.Agents/ExportAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).ExportAgent(ctx, req.(*ExportAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_ImportAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).ImportAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.Agents/ImportAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).ImportAgent(ctx, req.(*ImportAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_RestoreAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).RestoreAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.Agents/RestoreAgent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).RestoreAgent(ctx, req.(*RestoreAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dialogflow.v2beta1.Agents",
	HandlerType: (*AgentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgent",
			Handler:    _Agents_GetAgent_Handler,
		},
		{
			MethodName: "SearchAgents",
			Handler:    _Agents_SearchAgents_Handler,
		},
		{
			MethodName: "TrainAgent",
			Handler:    _Agents_TrainAgent_Handler,
		},
		{
			MethodName: "ExportAgent",
			Handler:    _Agents_ExportAgent_Handler,
		},
		{
			MethodName: "ImportAgent",
			Handler:    _Agents_ImportAgent_Handler,
		},
		{
			MethodName: "RestoreAgent",
			Handler:    _Agents_RestoreAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dialogflow/v2beta1/agent.proto",
}

func init() { proto.RegisterFile("google/cloud/dialogflow/v2beta1/agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 975 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x41, 0x6f, 0xdc, 0x44,
	0x14, 0xae, 0x37, 0x4d, 0xb2, 0xfb, 0x92, 0xb4, 0xe9, 0x84, 0x06, 0x6b, 0xdb, 0xa8, 0x8b, 0x4b,
	0xab, 0xed, 0x46, 0xd8, 0x74, 0xd3, 0x4a, 0x10, 0x04, 0x52, 0xb3, 0x9b, 0x36, 0x2b, 0x65, 0x93,
	0xc8, 0x49, 0x2a, 0xb5, 0x17, 0x6b, 0x62, 0xcf, 0x3a, 0x43, 0xed, 0x19, 0xe3, 0x99, 0x2d, 0x6d,
	0x0a, 0x1c, 0xf8, 0x05, 0x48, 0x20, 0x21, 0x38, 0x72, 0x42, 0x1c, 0x38, 0xf1, 0x4f, 0xf8, 0x0b,
	0xfc, 0x08, 0x8e, 0xc8, 0x63, 0x6f, 0xd6, 0x9b, 0xb4, 0xb5, 0x91, 0x90, 0xb8, 0xd9, 0xdf, 0xfb,
	0xde, 0x7b, 0xdf, 0xcc, 0x7c, 0x4f, 0x33, 0xb0, 0xea, 0x73, 0xee, 0x07, 0xc4, 0x72, 0x03, 0x3e,
	0xf4, 0x2c, 0x8f, 0xe2, 0x80, 0xfb, 0x83, 0x80, 0x7f, 0x69, 0x3d, 0x6f, 0x1f, 0x11, 0x89, 0xef,
	0x5a, 0xd8, 0x27, 0x4c, 0x9a, 0x51, 0xcc, 0x25, 0x47, 0x37, 0x52, 0xb2, 0xa9, 0xc8, 0xe6, 0x98,
	0x6c, 0x66, 0xe4, 0xfa, 0xf5, 0xac, 0x1a, 0x8e, 0xa8, 0x85, 0x19, 0xe3, 0x12, 0x4b, 0xca, 0x99,
	0x48, 0xd3, 0xeb, 0x37, 0xb3, 0x68, 0xc0, 0x99, 0x1f, 0x0f, 0x19, 0xa3, 0xcc, 0xb7, 0x78, 0x44,
	0xe2, 0x09, 0xd2, 0xb5, 0x8c, 0xa4, 0xfe, 0x8e, 0x86, 0x03, 0x8b, 0x84, 0x91, 0x7c, 0x99, 0x05,
	0x1b, 0x67, 0x83, 0x03, 0x4a, 0x02, 0xcf, 0x09, 0xb1, 0x78, 0x96, 0x31, 0xae, 0x9f, 0x65, 0x08,
	0x19, 0x0f, 0xdd, 0x6c, 0x01, 0xc6, 0x4f, 0x17, 0x61, 0xfa, 0x41, 0xb2, 0x20, 0xb4, 0x0c, 0x33,
	0x11, 0x8e, 0x09, 0x93, 0xba, 0xd6, 0xd0, 0x9a, 0x35, 0x3b, 0xfb, 0x43, 0xef, 0xc1, 0xbc, 0x47,
	0x45, 0x14, 0xe0, 0x97, 0x0e, 0xc3, 0x21, 0xd1, 0x2b, 0x2a, 0x3a, 0x97, 0x61, 0x3b, 0x38, 0x24,
	0xa8, 0x0d, 0x57, 0x3d, 0x32, 0xc0, 0xc3, 0x40, 0x3a, 0x01, 0x66, 0xfe, 0x10, 0xfb, 0xc4, 0x71,
	0xb9, 0x47, 0xf4, 0x29, 0xc5, 0x5d, 0xca, 0x82, 0xdb, 0x59, 0xac, 0xc3, 0x3d, 0x82, 0x3e, 0x02,
	0x5d, 0x0c, 0xa3, 0x88, 0xc7, 0x92, 0x78, 0x93, 0x59, 0x42, 0xbf, 0xd8, 0x98, 0x6a, 0xd6, 0xec,
	0xe5, 0xd3, 0x78, 0x3e, 0x51, 0xa0, 0x6b, 0x50, 0x93, 0x34, 0x24, 0xce, 0x09, 0x67, 0x44, 0x9f,
	0x56, 0x1d, 0xaa, 0x09, 0xf0, 0x94, 0x33, 0x82, 0x1a, 0x30, 0xe7, 0x11, 0xe1, 0xc6, 0x34, 0x4a,
	0xb6, 0x50, 0x9f, 0xc9, 0xc4, 0x8e, 0x21, 0xb4, 0x02, 0x80, 0x9f, 0x63, 0x89, 0x63, 0x67, 0x18,
	0x53, 0x7d, 0x56, 0x11, 0x6a, 0x29, 0x72, 0x18, 0x53, 0x74, 0x0b, 0x2e, 0x11, 0x86, 0x8f, 0x02,
	0xe2, 0x04, 0xdc, 0xf7, 0x29, 0xf3, 0xf5, 0x6a, 0x43, 0x6b, 0x56, 0xed, 0x85, 0x14, 0xdd, 0x4e,
	0x41, 0xb4, 0x0b, 0x10, 0x62, 0xe9, 0x1e, 0x3b, 0x61, 0xb2, 0xce, 0x5a, 0x43, 0x6b, 0x5e, 0x6a,
	0x7f, 0x68, 0x16, 0xb8, 0xc1, 0x54, 0x3b, 0x6d, 0xf6, 0x93, 0xc4, 0x3e, 0xf7, 0x88, 0x5d, 0x0b,
	0x47, 0x9f, 0xe8, 0x63, 0xd0, 0xdd, 0x00, 0x0b, 0x41, 0x07, 0xd4, 0x55, 0xc7, 0xef, 0xc8, 0xe3,
	0x98, 0x88, 0x63, 0x1e, 0x78, 0x3a, 0x34, 0xb4, 0x66, 0xc5, 0x7e, 0x77, 0x32, 0x7e, 0x30, 0x0a,
	0x1b, 0x8f, 0xa1, 0x76, 0x5a, 0x12, 0xd5, 0x61, 0xb9, 0xff, 0xe0, 0xa0, 0xb3, 0xe5, 0xf4, 0x77,
	0xbb, 0x9b, 0xce, 0xe1, 0xce, 0xfe, 0xde, 0x66, 0xa7, 0xf7, 0xb0, 0xb7, 0xd9, 0x5d, 0xbc, 0x80,
	0xae, 0xc2, 0x95, 0x5c, 0x6c, 0xeb, 0xc9, 0x86, 0xdd, 0xeb, 0x2e, 0x6a, 0x68, 0x19, 0x50, 0x0e,
	0xee, 0x6f, 0x3b, 0xbb, 0x3b, 0xdb, 0x4f, 0x16, 0x2b, 0xc6, 0x1d, 0xb8, 0xfc, 0x88, 0x48, 0xa5,
	0xd9, 0x26, 0x5f, 0x0c, 0x89, 0x78, 0xa3, 0x49, 0x0c, 0x0a, 0x4b, 0xfb, 0x04, 0xc7, 0xee, 0xb1,
	0x62, 0x8b, 0x02, 0x7a, 0x72, 0x84, 0x51, 0x72, 0xdc, 0x82, 0x9e, 0xa4, 0x86, 0x9a, 0xb6, 0xab,
	0x09, 0xb0, 0x4f, 0x4f, 0x48, 0x72, 0x40, 0x2a, 0x28, 0xf9, 0x33, 0xc2, 0x32, 0x0b, 0x29, 0xfa,
	0x41, 0x02, 0x18, 0xdf, 0xc0, 0x3b, 0x93, 0xad, 0x44, 0xc4, 0x99, 0x20, 0xe8, 0x33, 0x98, 0x51,
	0x93, 0x29, 0x74, 0xad, 0x31, 0xd5, 0x9c, 0x6b, 0xdf, 0x2e, 0x77, 0x1a, 0x76, 0x96, 0x85, 0x6e,
	0xc3, 0x65, 0x46, 0x5e, 0x48, 0x27, 0xd7, 0x3b, 0xb5, 0xfa, 0x42, 0x02, 0xef, 0x9d, 0xf6, 0x5f,
	0x85, 0x2b, 0x07, 0x31, 0xa6, 0xac, 0xd4, 0xbe, 0xf4, 0x00, 0x6d, 0xbe, 0x48, 0x4c, 0x5c, 0x86,
	0x9d, 0x6c, 0x8b, 0x12, 0xa3, 0x9c, 0x99, 0x36, 0xaf, 0x2a, 0xe0, 0x30, 0xa6, 0x86, 0x07, 0x4b,
	0x13, 0xa5, 0xb2, 0x65, 0xaf, 0xe4, 0x73, 0x54, 0xb9, 0xad, 0x0b, 0xe3, 0x2c, 0x74, 0x0b, 0x16,
	0xd2, 0xb0, 0xcb, 0x99, 0x4c, 0x3a, 0x26, 0x65, 0xe7, 0xb7, 0x2e, 0xd8, 0xf3, 0x0a, 0xee, 0xa4,
	0xe8, 0xc6, 0x2c, 0x4c, 0xab, 0x7f, 0xe3, 0x15, 0xa0, 0x5e, 0x58, 0x5a, 0xf0, 0xca, 0x39, 0xc1,
	0x6f, 0x6f, 0x3e, 0xf5, 0xf6, 0xe6, 0x5f, 0xc1, 0x92, 0x4d, 0x84, 0xe4, 0x31, 0xf9, 0x1f, 0xba,
	0xb7, 0xff, 0x98, 0x85, 0x99, 0xd4, 0x53, 0xe8, 0x3b, 0x0d, 0xaa, 0x23, 0xeb, 0xa3, 0xe2, 0xb1,
	0x3e, 0x33, 0x25, 0xf5, 0x92, 0xd6, 0x33, 0x5a, 0xdf, 0xfe, 0xf9, 0xd7, 0xf7, 0x95, 0xf7, 0x91,
	0x71, 0x7a, 0xb7, 0xbc, 0x4a, 0x97, 0xf6, 0x69, 0x14, 0xf3, 0xcf, 0x89, 0x2b, 0x85, 0xd5, 0xfa,
	0x3a, 0xbd, 0x6f, 0xd0, 0xef, 0x1a, 0xcc, 0xe7, 0x7d, 0x8f, 0xee, 0x15, 0x36, 0x79, 0xcd, 0x44,
	0xd6, 0xef, 0xff, 0xcb, 0xac, 0xd4, 0x65, 0xc6, 0x5d, 0xa5, 0x74, 0x15, 0xdd, 0x29, 0x56, 0xba,
	0x2e, 0x54, 0x01, 0xf4, 0x83, 0x06, 0x30, 0x1e, 0x14, 0xd4, 0x2e, 0x6c, 0x7c, 0x6e, 0xaa, 0xea,
	0x2b, 0xa3, 0x9c, 0xdc, 0xfd, 0x68, 0xee, 0x8e, 0xee, 0x47, 0x63, 0x4d, 0x89, 0xfa, 0xc0, 0x68,
	0x96, 0x10, 0x25, 0x93, 0xe2, 0xeb, 0x5a, 0x0b, 0xfd, 0xa8, 0xc1, 0x5c, 0x6e, 0x8e, 0xd0, 0x5a,
	0xa1, 0xae, 0xf3, 0x03, 0x5c, 0x24, 0xec, 0x9e, 0x12, 0x66, 0x1a, 0x65, 0x76, 0x8b, 0xa8, 0xea,
	0x23, 0x65, 0xb9, 0xd9, 0x2b, 0xa1, 0xec, 0xfc, 0xa4, 0xfe, 0x97, 0xca, 0x68, 0x38, 0x52, 0xf6,
	0xb3, 0x06, 0xf3, 0xf9, 0xc1, 0x2c, 0xe1, 0xbd, 0xd7, 0xcc, 0x71, 0x91, 0xb6, 0xfb, 0x4a, 0x9b,
	0x65, 0xb4, 0x4a, 0x68, 0x8b, 0xd3, 0xf2, 0xeb, 0x5a, 0x6b, 0xe3, 0x57, 0x0d, 0x6e, 0xba, 0x3c,
	0x2c, 0x52, 0xb4, 0x01, 0x4a, 0xcb, 0x5e, 0xf2, 0xea, 0xd9, 0xd3, 0x9e, 0xf6, 0x32, 0xba, 0xcf,
	0x93, 0xb7, 0x87, 0xc9, 0x63, 0xdf, 0xf2, 0x09, 0x53, 0x6f, 0x22, 0x2b, 0x0d, 0xe1, 0x88, 0x8a,
	0x37, 0x3e, 0x02, 0x3f, 0x19, 0x43, 0x7f, 0x6b, 0xda, 0x2f, 0x95, 0x4a, 0xf7, 0xe1, 0x6f, 0x95,
	0x1b, 0x8f, 0xd2, 0x9a, 0x1d, 0x25, 0xa1, 0x3b, 0x96, 0xf0, 0x38, 0x4d, 0x3a, 0x9a, 0x51, 0xf5,
	0xd7, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x35, 0x4d, 0xe2, 0x83, 0x63, 0x0a, 0x00, 0x00,
}
