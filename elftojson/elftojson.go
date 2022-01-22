package elftojson

import (
	"encoding/json"
	"fmt"

	"github.com/kinpoko/rdelf/readelf"
)

type ElfHeadersInfoJson struct {
	ElfHeader      ElfHeaderJson       `json:"elfheader"`
	ProgramHeaders []ProgramHeaderJson `json:"programheaders"`
	SectionHeaders []SectionHeaderJson `json:"sectionheaders"`
}

type ElfHeaderJson struct {
	Magic            string `json:"magic"`
	Class            string `json:"class"`
	Data             string `json:"data"`
	Version          string `json:"version"`
	Type             string `json:"type"`
	Machine          string `json:"machine"`
	EntryPoint       string `json:"entrypoint"`
	StartOfPHeader   string `json:"startofprogramheaders"`
	StartOfSHeader   string `json:"startofsectionheaders"`
	SizeOfPHeader    string `json:"sizeofprogramheaders"`
	NumOfPHeader     string `json:"numberofprogramheaders"`
	SizeOfSHeader    string `json:"sizeofsectionheaders"`
	NumOfSHeader     string `json:"numberofsectionheaders"`
	StringTableIndex string `json:"sectionheaderstringtableindex"`
}

type ProgramHeaderJson struct {
	Index  string `json:"index"`
	Type   string `json:"type"`
	Flags  string `json:"flags"`
	Offset string `json:"offset"`
	VAddr  string `json:"virtaddr"`
	PAddr  string `json:"physaddr"`
	FSize  string `json:"filesiz"`
	MSize  string `json:"memsiz"`
}

type SectionHeaderJson struct {
	Index     string `json:"index"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Flags     string `json:"flags"`
	Address   string `json:"address"`
	Offset    string `json:"offset"`
	Size      string `json:"size"`
	Link      string `json:"link"`
	Info      string `json:"info"`
	Alignment string `json:"align"`
	EntrySize string `json:"entsize"`
}

func ELFToJson(file []byte) ([]byte, error) {
	info, err := readelf.ReadELFHeader(file)
	if err != nil {
		return nil, err
	}
	var magicStr string
	for _, n := range info.Magic {
		magicStr += fmt.Sprintf("%x ", n)
	}

	headerJson := ElfHeaderJson{
		Magic:            magicStr,
		Class:            info.Class,
		Data:             info.Data,
		Version:          fmt.Sprintf("%x", info.Version),
		Type:             info.Type,
		Machine:          info.Machine,
		EntryPoint:       fmt.Sprintf("0x%x", info.EntryPoint),
		StartOfPHeader:   fmt.Sprintf("%d (bytes)", info.StartOfPHeader),
		StartOfSHeader:   fmt.Sprintf("%d (bytes)", info.StartOfSHeader),
		SizeOfPHeader:    fmt.Sprintf("%d", info.SizeOfPHeader),
		NumOfPHeader:     fmt.Sprintf("%d", info.NumOfPHeader),
		SizeOfSHeader:    fmt.Sprintf("%d", info.SizeOfSHeader),
		NumOfSHeader:     fmt.Sprintf("%d", info.NumOfSHeader),
		StringTableIndex: fmt.Sprintf("%d", info.StringTableIndex),
	}

	phinfos, err := readelf.ReadProgramHeaders(file, info.StartOfPHeader, info.NumOfPHeader, info.SizeOfPHeader)
	if err != nil {
		return nil, err
	}
	var programHeadersJson []ProgramHeaderJson
	for i, phinfo := range phinfos {
		programHeaderJson := ProgramHeaderJson{
			Index:  fmt.Sprintf("%d", i),
			Type:   phinfo.Type,
			Flags:  phinfo.Flags,
			Offset: fmt.Sprintf("0x%x", phinfo.Offset),
			VAddr:  fmt.Sprintf("0x%x", phinfo.VAddr),
			PAddr:  fmt.Sprintf("0x%x", phinfo.PAddr),
			FSize:  fmt.Sprintf("0x%x", phinfo.FSize),
			MSize:  fmt.Sprintf("0x%x", phinfo.MSize),
		}
		programHeadersJson = append(programHeadersJson, programHeaderJson)
	}

	shinfos, err := readelf.ReadSectionHeaders(file, info.StartOfSHeader, info.NumOfSHeader, info.SizeOfSHeader)
	if err != nil {
		return nil, err
	}
	var sectionHeadersJson []SectionHeaderJson
	for i, shinfo := range shinfos {
		sectionHeaderJson := SectionHeaderJson{
			Index:     fmt.Sprintf("%d", i),
			Name:      shinfo.NameString,
			Type:      shinfo.Type,
			Flags:     shinfo.Flags,
			Address:   fmt.Sprintf("0x%x", shinfo.Address),
			Offset:    fmt.Sprintf("0x%x", shinfo.Offset),
			Size:      fmt.Sprintf("0x%x", shinfo.Size),
			Link:      fmt.Sprintf("%d", shinfo.Link),
			Info:      fmt.Sprintf("%d", shinfo.Info),
			Alignment: fmt.Sprintf("0x%x", shinfo.Alignment),
			EntrySize: fmt.Sprintf("0x%x", shinfo.EntrySize),
		}
		sectionHeadersJson = append(sectionHeadersJson, sectionHeaderJson)
	}

	infojson := ElfHeadersInfoJson{
		ElfHeader:      headerJson,
		ProgramHeaders: programHeadersJson,
		SectionHeaders: sectionHeadersJson,
	}

	jsonBytes, err := json.Marshal(infojson)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil

}
