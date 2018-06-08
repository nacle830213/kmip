package packetprocessors

import (
	"fmt"
	"kmipserver/kmip"
	"errors"
	"context"
)

type BatchCount struct {}



func init() {
	kmip.Kmpiprocessor[4325389] = new(BatchCount)
}


func (r * BatchCount) ProcessPacket(ctx context.Context, t *kmip.TTLV, req []byte, response []byte , processor kmip.Processor) ([]byte,error) {

	fmt.Println("BatchCount",t.Tag, t.Type , t.Length, t.Value)

	if(len(req)) <= 0 {
		return nil,errors.New("Incomplete Packet")
	}

	f,s := kmip.ReadTTLV(req)
	p := kmip.GetProcessor(s.Tag)

	if p!= nil {
		p.ProcessPacket(ctx, &s,req[f:], nil, nil)
	}
	return nil,errors.New("Invalid Packet")
}
