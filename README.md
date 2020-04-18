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

- [ ] Integrate with <https://github.com/llir/llvm>
- [ ] Impliment the **DecoderMethods**

```bash
❯ jq '.[]? | select(.DecoderMethod? != null) | .DecoderMethod' table.json | sort -n | uniq -c | sort
   1 "DecodeMRSSystemRegister"
   1 "DecodeMSRSystemRegister"
   1 "DecodeSImm<10>"
   1 "DecodeSImm<8>"
   1 "DecodeSVEIncDecImm"
   1 "DecodeVecShiftL16Imm"
   1 "DecodeVecShiftL32Imm"
   1 "DecodeVecShiftL64Imm"
   1 "DecodeVecShiftL8Imm"
   1 "DecodeVecShiftR16ImmNarrow"
   1 "DecodeVecShiftR32ImmNarrow"
   1 "DecodeVecShiftR64Imm"
   1 "DecodeVecShiftR64ImmNarrow"
   2 "DecodeAdrInstruction"
   2 "DecodeFMOVLaneInstruction"
   2 "DecodeImm8OptLsl<16>"
   2 "DecodeImm8OptLsl<32>"
   2 "DecodeImm8OptLsl<64>"
   2 "DecodeImm8OptLsl<8>"
   2 "DecodePCRelLabel19"
   2 "DecodeSImm<5>"
   2 "DecodeSImm<6>"
   2 "DecodeSImm<9>"
   2 "DecodeSystemPStateInstruction"
   2 "DecodeUnconditionalBranch"
   2 "DecodeVecShiftR16Imm"
   2 "DecodeVecShiftR32Imm"
   2 "DecodeVecShiftR8Imm"
   3 "DecodeFixedPointScaleImm32"
   3 "DecodeFixedPointScaleImm64"
   4 "DecodeSVELogicalImmInstruction"
   4 "DecodeTestAndBranch"
   5 "DecodeSImm<4>"
   6 "DecodeMoveImmInstruction"
   8 "DecodeAddSubImmShift"
   8 "DecodeLogicalImmInstruction"
   8 "DecodeModImmTiedInstruction"
  10 "DecodeMemExtend"
  12 "DecodeAddSubERegInstruction"
  21 "DecodeModImmInstruction"
  24 "DecodeThreeAddrSRegInstruction"
  24 "DecodeUnsignedLdStInstruction"
  40 "DecodeExclusiveLdStInstruction"
  46 "DecodePairLdStInstruction"
  96 "DecodeSignedLdStInstruction"
```

## License

MIT Copyright (c) 2020 **blacktop**