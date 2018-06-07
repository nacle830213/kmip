package packetprocessors

import "fmt"

type NetworkIdentifier struct {}



func init() {
	Kmpiprocessor["4200AB"] = new(NetworkIdentifier)
}


func (r * NetworkIdentifier) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("NetworkIdentifier",string(t.Tag), string(t.Type) , t.Length, t.Value)

	f,s := ReadTTLV(req)
	p := GetProcessor(string(s.Tag))

	if(len(req[f:])) <= 0 {
		return nil, nil
	}

	if p!= nil {
		p.ProcessPacket(&s,req[f:], nil, nil)
	}
	return nil,nil
}