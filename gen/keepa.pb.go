// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: keepa.proto

package proto

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_keepa_proto protoreflect.FileDescriptor

var file_keepa_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0c, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x61, 0x72,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xf8, 0x06, 0x0a, 0x05, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x12, 0x2e, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x12,
	0x18, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x64,
	0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43,
	0x61, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72,
	0x64, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x42,
	0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72,
	0x64, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42,
	0x61, 0x6e, 0x6b, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x14, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x54,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x74, 0x65, 0x78, 0x74, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x17, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x78,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x3d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x18, 0x2e,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x43, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12,
	0x1b, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x2e, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x1f,
	0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6c, 0x6c,
	0x76, 0x6c, 0x6c, 0x2f, 0x6b, 0x65, 0x65, 0x70, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_keepa_proto_goTypes = []interface{}{
	(*AuthRequest)(nil),           // 0: auth.AuthRequest
	(*GetBankCardRequest)(nil),    // 1: card.GetBankCardRequest
	(*CreateBankCardRequest)(nil), // 2: card.CreateBankCardRequest
	(*UpdateBankCardRequest)(nil), // 3: card.UpdateBankCardRequest
	(*DeleteBankCardRequest)(nil), // 4: card.DeleteBankCardRequest
	(*GetTextRequest)(nil),        // 5: text.GetTextRequest
	(*CreateTextRequest)(nil),     // 6: text.CreateTextRequest
	(*UpdateTextRequest)(nil),     // 7: text.UpdateTextRequest
	(*DeleteTextRequest)(nil),     // 8: text.DeleteTextRequest
	(*GetBinaryRequest)(nil),      // 9: binary.GetBinaryRequest
	(*CreateBinaryRequest)(nil),   // 10: binary.CreateBinaryRequest
	(*UpdateBinaryRequest)(nil),   // 11: binary.UpdateBinaryRequest
	(*DeleteBinaryRequest)(nil),   // 12: binary.DeleteBinaryRequest
	(*AuthResponse)(nil),          // 13: auth.AuthResponse
	(*BankCardResponse)(nil),      // 14: card.BankCardResponse
	(*emptypb.Empty)(nil),         // 15: google.protobuf.Empty
	(*TextResponse)(nil),          // 16: text.TextResponse
	(*BinaryResponse)(nil),        // 17: binary.BinaryResponse
}
var file_keepa_proto_depIdxs = []int32{
	0,  // 0: proto.Keepa.Login:input_type -> auth.AuthRequest
	0,  // 1: proto.Keepa.Register:input_type -> auth.AuthRequest
	1,  // 2: proto.Keepa.GetBankCard:input_type -> card.GetBankCardRequest
	2,  // 3: proto.Keepa.CreateBankCard:input_type -> card.CreateBankCardRequest
	3,  // 4: proto.Keepa.UpdateBankCard:input_type -> card.UpdateBankCardRequest
	4,  // 5: proto.Keepa.DeleteBankCard:input_type -> card.DeleteBankCardRequest
	5,  // 6: proto.Keepa.GetText:input_type -> text.GetTextRequest
	6,  // 7: proto.Keepa.CreateText:input_type -> text.CreateTextRequest
	7,  // 8: proto.Keepa.UpdateText:input_type -> text.UpdateTextRequest
	8,  // 9: proto.Keepa.DeleteText:input_type -> text.DeleteTextRequest
	9,  // 10: proto.Keepa.GetBinary:input_type -> binary.GetBinaryRequest
	10, // 11: proto.Keepa.CreateBinary:input_type -> binary.CreateBinaryRequest
	11, // 12: proto.Keepa.UpdateBinary:input_type -> binary.UpdateBinaryRequest
	12, // 13: proto.Keepa.DeleteBinary:input_type -> binary.DeleteBinaryRequest
	13, // 14: proto.Keepa.Login:output_type -> auth.AuthResponse
	13, // 15: proto.Keepa.Register:output_type -> auth.AuthResponse
	14, // 16: proto.Keepa.GetBankCard:output_type -> card.BankCardResponse
	14, // 17: proto.Keepa.CreateBankCard:output_type -> card.BankCardResponse
	14, // 18: proto.Keepa.UpdateBankCard:output_type -> card.BankCardResponse
	15, // 19: proto.Keepa.DeleteBankCard:output_type -> google.protobuf.Empty
	16, // 20: proto.Keepa.GetText:output_type -> text.TextResponse
	16, // 21: proto.Keepa.CreateText:output_type -> text.TextResponse
	16, // 22: proto.Keepa.UpdateText:output_type -> text.TextResponse
	15, // 23: proto.Keepa.DeleteText:output_type -> google.protobuf.Empty
	17, // 24: proto.Keepa.GetBinary:output_type -> binary.BinaryResponse
	17, // 25: proto.Keepa.CreateBinary:output_type -> binary.BinaryResponse
	17, // 26: proto.Keepa.UpdateBinary:output_type -> binary.BinaryResponse
	15, // 27: proto.Keepa.DeleteBinary:output_type -> google.protobuf.Empty
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_keepa_proto_init() }
func file_keepa_proto_init() {
	if File_keepa_proto != nil {
		return
	}
	file_binary_proto_init()
	file_text_proto_init()
	file_card_proto_init()
	file_auth_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_keepa_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keepa_proto_goTypes,
		DependencyIndexes: file_keepa_proto_depIdxs,
	}.Build()
	File_keepa_proto = out.File
	file_keepa_proto_rawDesc = nil
	file_keepa_proto_goTypes = nil
	file_keepa_proto_depIdxs = nil
}
