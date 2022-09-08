// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: transaction.proto

package generated

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount                  string               `protobuf:"bytes,1,opt,name=Amount,proto3" json:"Amount,omitempty"`
	ToBankAccountNumber     string               `protobuf:"bytes,14,opt,name=ToBankAccountNumber,proto3" json:"ToBankAccountNumber,omitempty"`
	FromAccountNumber       string               `protobuf:"bytes,2,opt,name=FromAccountNumber,proto3" json:"FromAccountNumber,omitempty"`
	Kid                     string               `protobuf:"bytes,3,opt,name=Kid,proto3" json:"Kid,omitempty"`
	Date                    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	Reference               string               `protobuf:"bytes,5,opt,name=Reference,proto3" json:"Reference,omitempty"`
	Payer                   *Payer               `protobuf:"bytes,6,opt,name=Payer,proto3" json:"Payer,omitempty"`
	Currency                string               `protobuf:"bytes,7,opt,name=Currency,proto3" json:"Currency,omitempty"`
	TransactionCode         string               `protobuf:"bytes,8,opt,name=TransactionCode,proto3" json:"TransactionCode,omitempty"`
	IsInternational         bool                 `protobuf:"varint,9,opt,name=IsInternational,proto3" json:"IsInternational,omitempty"`
	TransactionGroupId      string               `protobuf:"bytes,10,opt,name=TransactionGroupId,proto3" json:"TransactionGroupId,omitempty"`
	TransactionInnerGroupId string               `protobuf:"bytes,11,opt,name=TransactionInnerGroupId,proto3" json:"TransactionInnerGroupId,omitempty"`
	IsCreditNote            bool                 `protobuf:"varint,12,opt,name=IsCreditNote,proto3" json:"IsCreditNote,omitempty"`
	Nets                    *Nets                `protobuf:"bytes,13,opt,name=nets,proto3" json:"nets,omitempty"`
	TransactionType         int32                `protobuf:"varint,15,opt,name=TransactionType,proto3" json:"TransactionType,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Transaction) GetToBankAccountNumber() string {
	if x != nil {
		return x.ToBankAccountNumber
	}
	return ""
}

func (x *Transaction) GetFromAccountNumber() string {
	if x != nil {
		return x.FromAccountNumber
	}
	return ""
}

func (x *Transaction) GetKid() string {
	if x != nil {
		return x.Kid
	}
	return ""
}

func (x *Transaction) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Transaction) GetReference() string {
	if x != nil {
		return x.Reference
	}
	return ""
}

func (x *Transaction) GetPayer() *Payer {
	if x != nil {
		return x.Payer
	}
	return nil
}

func (x *Transaction) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Transaction) GetTransactionCode() string {
	if x != nil {
		return x.TransactionCode
	}
	return ""
}

func (x *Transaction) GetIsInternational() bool {
	if x != nil {
		return x.IsInternational
	}
	return false
}

func (x *Transaction) GetTransactionGroupId() string {
	if x != nil {
		return x.TransactionGroupId
	}
	return ""
}

func (x *Transaction) GetTransactionInnerGroupId() string {
	if x != nil {
		return x.TransactionInnerGroupId
	}
	return ""
}

func (x *Transaction) GetIsCreditNote() bool {
	if x != nil {
		return x.IsCreditNote
	}
	return false
}

func (x *Transaction) GetNets() *Nets {
	if x != nil {
		return x.Nets
	}
	return nil
}

func (x *Transaction) GetTransactionType() int32 {
	if x != nil {
		return x.TransactionType
	}
	return 0
}

type Payer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Address1    string `protobuf:"bytes,2,opt,name=Address1,proto3" json:"Address1,omitempty"`
	Address2    string `protobuf:"bytes,3,opt,name=Address2,proto3" json:"Address2,omitempty"`
	ZipCode     string `protobuf:"bytes,4,opt,name=ZipCode,proto3" json:"ZipCode,omitempty"`
	CountryCode string `protobuf:"bytes,5,opt,name=CountryCode,proto3" json:"CountryCode,omitempty"`
}

func (x *Payer) Reset() {
	*x = Payer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payer) ProtoMessage() {}

func (x *Payer) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payer.ProtoReflect.Descriptor instead.
func (*Payer) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *Payer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Payer) GetAddress1() string {
	if x != nil {
		return x.Address1
	}
	return ""
}

func (x *Payer) GetAddress2() string {
	if x != nil {
		return x.Address2
	}
	return ""
}

func (x *Payer) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *Payer) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type Nets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgreementId                string `protobuf:"bytes,12,opt,name=AgreementId,proto3" json:"AgreementId,omitempty"`
	FormNumber                 string `protobuf:"bytes,13,opt,name=FormNumber,proto3" json:"FormNumber,omitempty"`
	TransactionNumber          string `protobuf:"bytes,15,opt,name=TransactionNumber,proto3" json:"TransactionNumber,omitempty"`
	SerialNumber               string `protobuf:"bytes,16,opt,name=SerialNumber,proto3" json:"SerialNumber,omitempty"`
	PartialSettlementNumber    string `protobuf:"bytes,17,opt,name=PartialSettlementNumber,proto3" json:"PartialSettlementNumber,omitempty"`
	TransactionTypeDescription string `protobuf:"bytes,22,opt,name=TransactionTypeDescription,proto3" json:"TransactionTypeDescription,omitempty"`
	// Transmission
	DataTransmitter string `protobuf:"bytes,18,opt,name=DataTransmitter,proto3" json:"DataTransmitter,omitempty"`
	CustomerUnitId  string `protobuf:"bytes,21,opt,name=CustomerUnitId,proto3" json:"CustomerUnitId,omitempty"`
	// Assignment
	AssignmentNumber  string `protobuf:"bytes,19,opt,name=AssignmentNumber,proto3" json:"AssignmentNumber,omitempty"`
	AssignmentAccount string `protobuf:"bytes,20,opt,name=AssignmentAccount,proto3" json:"AssignmentAccount,omitempty"`
}

func (x *Nets) Reset() {
	*x = Nets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nets) ProtoMessage() {}

func (x *Nets) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nets.ProtoReflect.Descriptor instead.
func (*Nets) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *Nets) GetAgreementId() string {
	if x != nil {
		return x.AgreementId
	}
	return ""
}

func (x *Nets) GetFormNumber() string {
	if x != nil {
		return x.FormNumber
	}
	return ""
}

func (x *Nets) GetTransactionNumber() string {
	if x != nil {
		return x.TransactionNumber
	}
	return ""
}

func (x *Nets) GetSerialNumber() string {
	if x != nil {
		return x.SerialNumber
	}
	return ""
}

func (x *Nets) GetPartialSettlementNumber() string {
	if x != nil {
		return x.PartialSettlementNumber
	}
	return ""
}

func (x *Nets) GetTransactionTypeDescription() string {
	if x != nil {
		return x.TransactionTypeDescription
	}
	return ""
}

func (x *Nets) GetDataTransmitter() string {
	if x != nil {
		return x.DataTransmitter
	}
	return ""
}

func (x *Nets) GetCustomerUnitId() string {
	if x != nil {
		return x.CustomerUnitId
	}
	return ""
}

func (x *Nets) GetAssignmentNumber() string {
	if x != nil {
		return x.AssignmentNumber
	}
	return ""
}

func (x *Nets) GetAssignmentAccount() string {
	if x != nil {
		return x.AssignmentAccount
	}
	return ""
}

var File_transaction_proto protoreflect.FileDescriptor

var file_transaction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8,
	0x04, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x54, 0x6f, 0x42, 0x61, 0x6e, 0x6b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x13, 0x54, 0x6f, 0x42, 0x61, 0x6e, 0x6b, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x46, 0x72, 0x6f, 0x6d,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x50, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x50, 0x61, 0x79, 0x65, 0x72, 0x52, 0x05, 0x50, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x49, 0x73, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x49, 0x73,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x2e, 0x0a,
	0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x38, 0x0a,
	0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x6e, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x49, 0x73, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x49,
	0x73, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x6e,
	0x65, 0x74, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x65, 0x74, 0x73, 0x52, 0x04, 0x6e, 0x65, 0x74, 0x73, 0x12,
	0x28, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x05, 0x50, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x12,
	0x18, 0x0a, 0x07, 0x5a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x5a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xc0, 0x03, 0x0a, 0x04,
	0x4e, 0x65, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x67, 0x72, 0x65, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x67, 0x72, 0x65, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x6f, 0x72, 0x6d, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x46, 0x6f, 0x72, 0x6d,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x17, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x61, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x3e, 0x0a, 0x1a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x55, 0x6e,
	0x69, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x2c, 0x0a, 0x11, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x41, 0x73, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_proto_rawDescOnce sync.Once
	file_transaction_proto_rawDescData = file_transaction_proto_rawDesc
)

func file_transaction_proto_rawDescGZIP() []byte {
	file_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_proto_rawDescData)
	})
	return file_transaction_proto_rawDescData
}

var file_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_transaction_proto_goTypes = []interface{}{
	(*Transaction)(nil),         // 0: protobuf.Transaction
	(*Payer)(nil),               // 1: protobuf.Payer
	(*Nets)(nil),                // 2: protobuf.Nets
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_transaction_proto_depIdxs = []int32{
	3, // 0: protobuf.Transaction.date:type_name -> google.protobuf.Timestamp
	1, // 1: protobuf.Transaction.Payer:type_name -> protobuf.Payer
	2, // 2: protobuf.Transaction.nets:type_name -> protobuf.Nets
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_transaction_proto_init() }
func file_transaction_proto_init() {
	if File_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payer); i {
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
		file_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nets); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_proto_depIdxs,
		MessageInfos:      file_transaction_proto_msgTypes,
	}.Build()
	File_transaction_proto = out.File
	file_transaction_proto_rawDesc = nil
	file_transaction_proto_goTypes = nil
	file_transaction_proto_depIdxs = nil
}
