// Code generated by protoc-gen-gogo.
// source: output.proto
// DO NOT EDIT!

package graph

import proto "github.com/gogo/protobuf/proto"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"
import ann "sourcegraph.com/sourcegraph/srclib/ann"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Output struct {
	Defs []*Def     `protobuf:"bytes,1,rep,name=defs" json:"Defs,omitempty"`
	Refs []*Ref     `protobuf:"bytes,2,rep,name=refs" json:"Refs,omitempty"`
	Docs []*Doc     `protobuf:"bytes,3,rep,name=docs" json:"Docs,omitempty"`
	Anns []*ann.Ann `protobuf:"bytes,4,rep,name=anns" json:"Anns,omitempty"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}

func init() {
}
