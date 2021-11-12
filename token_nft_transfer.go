package hedera

import (
	"github.com/hashgraph/hedera-protobufs-go/services"
	protobuf "google.golang.org/protobuf/proto"
)

type TokenNftTransfer struct {
	SenderAccountID   AccountID
	ReceiverAccountID AccountID
	SerialNumber      int64
}

func _NftTransferFromProtobuf(pb *services.NftTransfer) TokenNftTransfer {
	if pb == nil {
		return TokenNftTransfer{}
	}

	senderAccountID := AccountID{}
	if pb.SenderAccountID != nil {
		senderAccountID = *_AccountIDFromProtobuf(pb.SenderAccountID)
	}

	receiverAccountID := AccountID{}
	if pb.ReceiverAccountID != nil {
		receiverAccountID = *_AccountIDFromProtobuf(pb.ReceiverAccountID)
	}

	return TokenNftTransfer{
		SenderAccountID:   senderAccountID,
		ReceiverAccountID: receiverAccountID,
		SerialNumber:      pb.SerialNumber,
	}
}

func (transfer *TokenNftTransfer) _ToProtobuf() *services.NftTransfer {
	return &services.NftTransfer{
		SenderAccountID:   transfer.SenderAccountID._ToProtobuf(),
		ReceiverAccountID: transfer.ReceiverAccountID._ToProtobuf(),
		SerialNumber:      transfer.SerialNumber,
	}
}

func (transfer TokenNftTransfer) ToBytes() []byte {
	data, err := protobuf.Marshal(transfer._ToProtobuf())
	if err != nil {
		return make([]byte, 0)
	}

	return data
}

func NftTransferFromBytes(data []byte) (TokenNftTransfer, error) {
	if data == nil {
		return TokenNftTransfer{}, errByteArrayNull
	}
	pb := services.NftTransfer{}
	err := protobuf.Unmarshal(data, &pb)
	if err != nil {
		return TokenNftTransfer{}, err
	}

	return _NftTransferFromProtobuf(&pb), nil
}
