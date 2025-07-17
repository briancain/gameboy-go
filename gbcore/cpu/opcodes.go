package gbcore

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// OpcodeInfo represents the information for a single opcode
type OpcodeInfo struct {
	Mnemonic   string                 `json:"mnemonic"`
	Bytes      int                    `json:"bytes"`
	Cycles     []int                  `json:"cycles"`
	Operands   []map[string]interface{} `json:"operands"`
	Immediate  bool                   `json:"immediate"`
	Flags      map[string]string      `json:"flags"`
}

// OpcodesData represents the entire opcodes JSON structure
type OpcodesData struct {
	Unprefixed  map[string]OpcodeInfo `json:"unprefixed"`
	CBPrefixed  map[string]OpcodeInfo `json:"cbprefixed"`
}

// LoadOpcodes loads the opcodes from the JSON file
func LoadOpcodes(filePath string) (*OpcodesData, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading opcodes file: %v", err)
	}

	var opcodes OpcodesData
	err = json.Unmarshal(data, &opcodes)
	if err != nil {
		return nil, fmt.Errorf("error parsing opcodes JSON: %v", err)
	}

	log.Printf("Loaded %d unprefixed opcodes and %d CB-prefixed opcodes", 
		len(opcodes.Unprefixed), len(opcodes.CBPrefixed))

	return &opcodes, nil
}

// GetOpcodeInfo returns the information for a specific opcode
func (o *OpcodesData) GetOpcodeInfo(opcode byte, prefixed bool) *OpcodeInfo {
	opcodeHex := fmt.Sprintf("0x%02X", opcode)
	
	if prefixed {
		if info, ok := o.CBPrefixed[opcodeHex]; ok {
			return &info
		}
	} else {
		if info, ok := o.Unprefixed[opcodeHex]; ok {
			return &info
		}
	}
	
	return nil
}
