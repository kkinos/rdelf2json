# rdelf2json

CLI application for parsing ELF and converting to json

## Install

```bash
go install github.com/kinpoko/rdelf2json@latest
```

## Usage

```bash
rdelf2json sample/a.out | jq '.elfheader'
{
  "magic": "7f 45 4c 46 2 1 1 0 0 0 0 0 0 0 0 0 ",
  "class": "ELF64",
  "data": "little endian",
  "version": "1",
  "type": "A shared object",
  "machine": "AMD x86-64",
  "entrypoint": "0x1060",
  "startofprogramheaders": "64 (bytes)",
  "startofsectionheaders": "14712 (bytes)",
  "sizeofprogramheaders": "56",
  "numberofprogramheaders": "13",
  "sizeofsectionheaders": "64",
  "numberofsectionheaders": "31",
  "sectionheaderstringtableindex": "30"
}
```
