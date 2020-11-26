//
// * Copyright 2019 The Feast Authors
// *
// * Licensed under the Apache License, Version 2.0 (the "License");
// * you may not use this file except in compliance with the License.
// * You may obtain a copy of the License at
// *
// *     https://www.apache.org/licenses/LICENSE-2.0
// *
// * Unless required by applicable law or agreed to in writing, software
// * distributed under the License is distributed on an "AS IS" BASIS,
// * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// * See the License for the specific language governing permissions and
// * limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: feast/core/Store.proto

package core

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Store_StoreType int32

const (
	Store_INVALID Store_StoreType = 0
	// Redis stores a FeatureRow element as a key, value pair.
	//
	// The Redis data types used (https://redis.io/topics/data-types):
	// - key: STRING
	// - value: STRING
	//
	// Encodings:
	// - key: byte array of RedisKey (refer to feast.storage.RedisKeyV2)
	// - value: Redis hashmap
	//
	Store_REDIS         Store_StoreType = 1
	Store_REDIS_CLUSTER Store_StoreType = 4
)

// Enum value maps for Store_StoreType.
var (
	Store_StoreType_name = map[int32]string{
		0: "INVALID",
		1: "REDIS",
		4: "REDIS_CLUSTER",
	}
	Store_StoreType_value = map[string]int32{
		"INVALID":       0,
		"REDIS":         1,
		"REDIS_CLUSTER": 4,
	}
)

func (x Store_StoreType) Enum() *Store_StoreType {
	p := new(Store_StoreType)
	*p = x
	return p
}

func (x Store_StoreType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Store_StoreType) Descriptor() protoreflect.EnumDescriptor {
	return file_feast_core_Store_proto_enumTypes[0].Descriptor()
}

func (Store_StoreType) Type() protoreflect.EnumType {
	return &file_feast_core_Store_proto_enumTypes[0]
}

func (x Store_StoreType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Store_StoreType.Descriptor instead.
func (Store_StoreType) EnumDescriptor() ([]byte, []int) {
	return file_feast_core_Store_proto_rawDescGZIP(), []int{0, 0}
}

// Store provides a location where Feast reads and writes feature values.
// Feature values will be written to the Store in the form of FeatureRow elements.
// The way FeatureRow is encoded and decoded when it is written to and read from
// the Store depends on the type of the Store.
//
type Store struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the store.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of store.
	Type Store_StoreType `protobuf:"varint,2,opt,name=type,proto3,enum=feast.core.Store_StoreType" json:"type,omitempty"`
	// Feature sets to subscribe to.
	Subscriptions []*Store_Subscription `protobuf:"bytes,4,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	// Configuration to connect to the store. Required.
	//
	// Types that are assignable to Config:
	//	*Store_RedisConfig_
	//	*Store_RedisClusterConfig_
	Config isStore_Config `protobuf_oneof:"config"`
}

func (x *Store) Reset() {
	*x = Store{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_Store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store) ProtoMessage() {}

func (x *Store) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_Store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store.ProtoReflect.Descriptor instead.
func (*Store) Descriptor() ([]byte, []int) {
	return file_feast_core_Store_proto_rawDescGZIP(), []int{0}
}

func (x *Store) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Store) GetType() Store_StoreType {
	if x != nil {
		return x.Type
	}
	return Store_INVALID
}

func (x *Store) GetSubscriptions() []*Store_Subscription {
	if x != nil {
		return x.Subscriptions
	}
	return nil
}

func (m *Store) GetConfig() isStore_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *Store) GetRedisConfig() *Store_RedisConfig {
	if x, ok := x.GetConfig().(*Store_RedisConfig_); ok {
		return x.RedisConfig
	}
	return nil
}

func (x *Store) GetRedisClusterConfig() *Store_RedisClusterConfig {
	if x, ok := x.GetConfig().(*Store_RedisClusterConfig_); ok {
		return x.RedisClusterConfig
	}
	return nil
}

type isStore_Config interface {
	isStore_Config()
}

type Store_RedisConfig_ struct {
	RedisConfig *Store_RedisConfig `protobuf:"bytes,11,opt,name=redis_config,json=redisConfig,proto3,oneof"`
}

type Store_RedisClusterConfig_ struct {
	RedisClusterConfig *Store_RedisClusterConfig `protobuf:"bytes,14,opt,name=redis_cluster_config,json=redisClusterConfig,proto3,oneof"`
}

func (*Store_RedisConfig_) isStore_Config() {}

func (*Store_RedisClusterConfig_) isStore_Config() {}

type Store_RedisConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Optional. The number of milliseconds to wait before retrying failed Redis connection.
	// By default, Feast uses exponential backoff policy and "initial_backoff_ms" sets the initial wait duration.
	InitialBackoffMs int32 `protobuf:"varint,3,opt,name=initial_backoff_ms,json=initialBackoffMs,proto3" json:"initial_backoff_ms,omitempty"`
	// Optional. Maximum total number of retries for connecting to Redis. Default to zero retries.
	MaxRetries int32 `protobuf:"varint,4,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// Optional. How often flush data to redis
	FlushFrequencySeconds int32 `protobuf:"varint,5,opt,name=flush_frequency_seconds,json=flushFrequencySeconds,proto3" json:"flush_frequency_seconds,omitempty"`
	// Optional. Connect over SSL.
	Ssl bool `protobuf:"varint,6,opt,name=ssl,proto3" json:"ssl,omitempty"`
}

func (x *Store_RedisConfig) Reset() {
	*x = Store_RedisConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_Store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store_RedisConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store_RedisConfig) ProtoMessage() {}

func (x *Store_RedisConfig) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_Store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store_RedisConfig.ProtoReflect.Descriptor instead.
func (*Store_RedisConfig) Descriptor() ([]byte, []int) {
	return file_feast_core_Store_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Store_RedisConfig) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Store_RedisConfig) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Store_RedisConfig) GetInitialBackoffMs() int32 {
	if x != nil {
		return x.InitialBackoffMs
	}
	return 0
}

func (x *Store_RedisConfig) GetMaxRetries() int32 {
	if x != nil {
		return x.MaxRetries
	}
	return 0
}

func (x *Store_RedisConfig) GetFlushFrequencySeconds() int32 {
	if x != nil {
		return x.FlushFrequencySeconds
	}
	return 0
}

func (x *Store_RedisConfig) GetSsl() bool {
	if x != nil {
		return x.Ssl
	}
	return false
}

type Store_RedisClusterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of Redis Uri for all the nodes in Redis Cluster, comma separated. Eg. host1:6379, host2:6379
	ConnectionString string `protobuf:"bytes,1,opt,name=connection_string,json=connectionString,proto3" json:"connection_string,omitempty"`
	InitialBackoffMs int32  `protobuf:"varint,2,opt,name=initial_backoff_ms,json=initialBackoffMs,proto3" json:"initial_backoff_ms,omitempty"`
	MaxRetries       int32  `protobuf:"varint,3,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// Optional. How often flush data to redis
	FlushFrequencySeconds int32 `protobuf:"varint,4,opt,name=flush_frequency_seconds,json=flushFrequencySeconds,proto3" json:"flush_frequency_seconds,omitempty"`
	// Optional. Append a prefix to the Redis Key
	KeyPrefix string `protobuf:"bytes,5,opt,name=key_prefix,json=keyPrefix,proto3" json:"key_prefix,omitempty"`
	// Optional. Enable fallback to another key prefix if the original key is not present.
	// Useful for migrating key prefix without re-ingestion. Disabled by default.
	EnableFallback bool `protobuf:"varint,6,opt,name=enable_fallback,json=enableFallback,proto3" json:"enable_fallback,omitempty"`
	// Optional. This would be the fallback prefix to use if enable_fallback is true.
	FallbackPrefix string `protobuf:"bytes,7,opt,name=fallback_prefix,json=fallbackPrefix,proto3" json:"fallback_prefix,omitempty"`
}

func (x *Store_RedisClusterConfig) Reset() {
	*x = Store_RedisClusterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_Store_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store_RedisClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store_RedisClusterConfig) ProtoMessage() {}

func (x *Store_RedisClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_Store_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store_RedisClusterConfig.ProtoReflect.Descriptor instead.
func (*Store_RedisClusterConfig) Descriptor() ([]byte, []int) {
	return file_feast_core_Store_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Store_RedisClusterConfig) GetConnectionString() string {
	if x != nil {
		return x.ConnectionString
	}
	return ""
}

func (x *Store_RedisClusterConfig) GetInitialBackoffMs() int32 {
	if x != nil {
		return x.InitialBackoffMs
	}
	return 0
}

func (x *Store_RedisClusterConfig) GetMaxRetries() int32 {
	if x != nil {
		return x.MaxRetries
	}
	return 0
}

func (x *Store_RedisClusterConfig) GetFlushFrequencySeconds() int32 {
	if x != nil {
		return x.FlushFrequencySeconds
	}
	return 0
}

func (x *Store_RedisClusterConfig) GetKeyPrefix() string {
	if x != nil {
		return x.KeyPrefix
	}
	return ""
}

func (x *Store_RedisClusterConfig) GetEnableFallback() bool {
	if x != nil {
		return x.EnableFallback
	}
	return false
}

func (x *Store_RedisClusterConfig) GetFallbackPrefix() string {
	if x != nil {
		return x.FallbackPrefix
	}
	return ""
}

type Store_Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of project that the feature sets belongs to. This can be one of
	// - [project_name]
	// - *
	// If an asterisk is provided, filtering on projects will be disabled. All projects will
	// be matched. It is NOT possible to provide an asterisk with a string in order to do
	// pattern matching.
	Project string `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	// Name of the desired feature set. Asterisks can be used as wildcards in the name.
	// Matching on names is only permitted if a specific project is defined. It is disallowed
	// If the project name is set to "*"
	// e.g.
	// - * can be used to match all feature sets
	// - my-feature-set* can be used to match all features prefixed by "my-feature-set"
	// - my-feature-set-6 can be used to select a single feature set
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// All matches with exclude enabled will be filtered out instead of added
	Exclude bool `protobuf:"varint,4,opt,name=exclude,proto3" json:"exclude,omitempty"`
}

func (x *Store_Subscription) Reset() {
	*x = Store_Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feast_core_Store_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Store_Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Store_Subscription) ProtoMessage() {}

func (x *Store_Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_feast_core_Store_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Store_Subscription.ProtoReflect.Descriptor instead.
func (*Store_Subscription) Descriptor() ([]byte, []int) {
	return file_feast_core_Store_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Store_Subscription) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Store_Subscription) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Store_Subscription) GetExclude() bool {
	if x != nil {
		return x.Exclude
	}
	return false
}

var File_feast_core_Store_proto protoreflect.FileDescriptor

var file_feast_core_Store_proto_rawDesc = []byte{
	0x0a, 0x16, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x22, 0xf5, 0x07, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x66, 0x65, 0x61,
	0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x42, 0x0a, 0x0c, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00,
	0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x58, 0x0a,
	0x14, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x65,
	0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x52,
	0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x48, 0x00, 0x52, 0x12, 0x72, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0xce, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x64, 0x69,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x2c, 0x0a, 0x12, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x6f,
	0x66, 0x66, 0x5f, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x4d, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x36,
	0x0a, 0x17, 0x66, 0x6c, 0x75, 0x73, 0x68, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x15, 0x66, 0x6c, 0x75, 0x73, 0x68, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x73, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x73, 0x6c, 0x1a, 0xb9, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x64,
	0x69, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x12,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x5f,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x4d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61,
	0x78, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x66,
	0x6c, 0x75, 0x73, 0x68, 0x5f, 0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x66, 0x6c,
	0x75, 0x73, 0x68, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6b, 0x65, 0x79, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x27, 0x0a, 0x0f, 0x66,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x66, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x1a, 0x5c, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x4a, 0x04, 0x08, 0x02,
	0x10, 0x03, 0x22, 0x4e, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05,
	0x52, 0x45, 0x44, 0x49, 0x53, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x44, 0x49, 0x53,
	0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x10, 0x04, 0x22, 0x04, 0x08, 0x02, 0x10, 0x02,
	0x22, 0x04, 0x08, 0x03, 0x10, 0x03, 0x22, 0x04, 0x08, 0x0c, 0x10, 0x0c, 0x22, 0x04, 0x08, 0x0d,
	0x10, 0x0d, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x53, 0x0a, 0x10,
	0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x42, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x64,
	0x65, 0x76, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feast_core_Store_proto_rawDescOnce sync.Once
	file_feast_core_Store_proto_rawDescData = file_feast_core_Store_proto_rawDesc
)

func file_feast_core_Store_proto_rawDescGZIP() []byte {
	file_feast_core_Store_proto_rawDescOnce.Do(func() {
		file_feast_core_Store_proto_rawDescData = protoimpl.X.CompressGZIP(file_feast_core_Store_proto_rawDescData)
	})
	return file_feast_core_Store_proto_rawDescData
}

var file_feast_core_Store_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_feast_core_Store_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_feast_core_Store_proto_goTypes = []interface{}{
	(Store_StoreType)(0),             // 0: feast.core.Store.StoreType
	(*Store)(nil),                    // 1: feast.core.Store
	(*Store_RedisConfig)(nil),        // 2: feast.core.Store.RedisConfig
	(*Store_RedisClusterConfig)(nil), // 3: feast.core.Store.RedisClusterConfig
	(*Store_Subscription)(nil),       // 4: feast.core.Store.Subscription
}
var file_feast_core_Store_proto_depIdxs = []int32{
	0, // 0: feast.core.Store.type:type_name -> feast.core.Store.StoreType
	4, // 1: feast.core.Store.subscriptions:type_name -> feast.core.Store.Subscription
	2, // 2: feast.core.Store.redis_config:type_name -> feast.core.Store.RedisConfig
	3, // 3: feast.core.Store.redis_cluster_config:type_name -> feast.core.Store.RedisClusterConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_feast_core_Store_proto_init() }
func file_feast_core_Store_proto_init() {
	if File_feast_core_Store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feast_core_Store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feast_core_Store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store_RedisConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feast_core_Store_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store_RedisClusterConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_feast_core_Store_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Store_Subscription); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_feast_core_Store_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Store_RedisConfig_)(nil),
		(*Store_RedisClusterConfig_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_feast_core_Store_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_feast_core_Store_proto_goTypes,
		DependencyIndexes: file_feast_core_Store_proto_depIdxs,
		EnumInfos:         file_feast_core_Store_proto_enumTypes,
		MessageInfos:      file_feast_core_Store_proto_msgTypes,
	}.Build()
	File_feast_core_Store_proto = out.File
	file_feast_core_Store_proto_rawDesc = nil
	file_feast_core_Store_proto_goTypes = nil
	file_feast_core_Store_proto_depIdxs = nil
}
