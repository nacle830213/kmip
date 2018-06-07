package packetprocessors

import "fmt"

type CryptographicAlgorithm struct {}



func init() {
	Kmpiprocessor["420028"] = new(CryptographicAlgorithm)
}


func (r * CryptographicAlgorithm) ProcessPacket(t *TTLV, req []byte, response []byte , processor Processor) ([]byte,error) {

	fmt.Println("CryptographicAlgorithm",string(t.Tag), string(t.Type) , t.Length, t.Value)

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
