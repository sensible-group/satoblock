package model

import (
	"encoding/binary"
	"encoding/json"
)

type TxRequest struct {
	TxHex   string `json:"txHex"`
	ByTxHex string `json:"byTxHex"`
}

type TxResponse struct {
	TxId    string `json:"txId"`
	Index   int    `json:"index"`
	ByTxId  string `json:"byTxId"`
	Sig     string `json:"sigBE"`
	Padding string `json:"padding"`
	Payload string `json:"payload"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (t *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(*t)
}

// redis›
type TxoData struct {
	UTxid       []byte
	Vout        uint32
	BlockHeight uint32
	TxIdx       uint64
	AddressPkh  []byte
	GenesisId   []byte
	Value       uint64
	ScriptType  []byte
	Script      []byte
}

func (d *TxoData) Marshal(buf []byte) {
	binary.LittleEndian.PutUint32(buf, d.BlockHeight) // 4
	binary.LittleEndian.PutUint64(buf[4:], d.TxIdx)   // 8
	binary.LittleEndian.PutUint64(buf[12:], d.Value)  // 8
	// copy(buf[12:], d.AddressPkh)                      // 20
	// copy(buf[32:], d.GenesisId)                       // 20
	// copy(buf[60:], d.ScriptType)                      // 32
	copy(buf[20:], d.Script) // n
}

func (d *TxoData) Unmarshal(buf []byte) {
	d.BlockHeight = binary.LittleEndian.Uint32(buf[:4]) // 4
	d.TxIdx = binary.LittleEndian.Uint64(buf[4:12])     // 8
	d.Value = binary.LittleEndian.Uint64(buf[12:20])    // 8
	// copy(d.AddressPkh, buf[12:32])                      // 20
	// copy(d.GenesisId, buf[32:52])                       // 20
	// copy(d.ScriptType, buf[60:92])                      // 32
	d.Script = buf[20:]
	// d.Script = make([]byte, len(buf)-20)
	// copy(d.Script, buf[20:]) // n
}
