# disasster [WIP] 🚧

[![Go](https://github.com/blacktop/disasster/workflows/Go/badge.svg)](https://github.com/blacktop/disasster/actions) [![GoDoc](https://godoc.org/github.com/blacktop/disasster?status.svg)](https://pkg.go.dev/github.com/blacktop/disasster) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)

> Golang AARCH64 Disassembler.

---

## Install

```bash
$ go get github.com/blacktop/disasster
```

## Getting Started 🚀

```go
package main

import "github.com/blacktop/disasster

func main() {

}
```

## Development 👩‍💻

### Requirments

```bash
$ brew install llvm z3
```

### Or Build

```bash
$ git clone https://github.com/llvm/llvm-project.git
$ cd llvm-project/llvm
$ mkdir build
$ cd build
$ cmake ..
$ make AArch64CommonTableGen -j
```

### Generate `table.json`

```bash
$ git clone https://github.com/llvm/llvm-project.git
$ cd llvm-project/llvm/lib/Target/AArch64
$ /usr/local/opt/llvm/bin/llvm-tblgen AArch64.td -I=../../../include --dump-json > out.json
```

### Update `tables.go`

```bash
$ go generate ./...
```

## TODO ☑️

```bash
❯ jq '.[]? | select(.DecoderMethod? != null) | .DecoderMethod' out.json | sort -n | uniq -c
5448 ""
  12 "DecodeAddSubERegInstruction"
   8 "DecodeAddSubImmShift"
   2 "DecodeAdrInstruction"
  40 "DecodeExclusiveLdStInstruction"
   2 "DecodeFMOVLaneInstruction"
   3 "DecodeFixedPointScaleImm32"
   3 "DecodeFixedPointScaleImm64"
   2 "DecodeImm8OptLsl<16>"
   2 "DecodeImm8OptLsl<32>"
   2 "DecodeImm8OptLsl<64>"
   2 "DecodeImm8OptLsl<8>"
   8 "DecodeLogicalImmInstruction"
   1 "DecodeMRSSystemRegister"
   1 "DecodeMSRSystemRegister"
  10 "DecodeMemExtend"
  21 "DecodeModImmInstruction"
   8 "DecodeModImmTiedInstruction"
   6 "DecodeMoveImmInstruction"
   2 "DecodePCRelLabel19"
  46 "DecodePairLdStInstruction"
   1 "DecodeSImm<10>"
   5 "DecodeSImm<4>"
   2 "DecodeSImm<5>"
   2 "DecodeSImm<6>"
   1 "DecodeSImm<8>"
   2 "DecodeSImm<9>"
   1 "DecodeSVEIncDecImm"
   4 "DecodeSVELogicalImmInstruction"
  96 "DecodeSignedLdStInstruction"
   2 "DecodeSystemPStateInstruction"
   4 "DecodeTestAndBranch"
  24 "DecodeThreeAddrSRegInstruction"
   2 "DecodeUnconditionalBranch"
  24 "DecodeUnsignedLdStInstruction"
   1 "DecodeVecShiftL16Imm"
   1 "DecodeVecShiftL32Imm"
   1 "DecodeVecShiftL64Imm"
   1 "DecodeVecShiftL8Imm"
   2 "DecodeVecShiftR16Imm"
   1 "DecodeVecShiftR16ImmNarrow"
   2 "DecodeVecShiftR32Imm"
   1 "DecodeVecShiftR32ImmNarrow"
   1 "DecodeVecShiftR64Imm"
   1 "DecodeVecShiftR64ImmNarrow"
   2 "DecodeVecShiftR8Imm"
```

## License

MIT Copyright (c) 2020 **blacktop**