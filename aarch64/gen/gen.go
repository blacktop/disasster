package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Table map[string]Inst

type field struct {
	Def       string `json:"def,omitempty"`
	Kind      string `json:"kind,omitempty"`
	Printable string `json:"printable,omitempty"`
}

type InstField struct {
	Index     int    `json:"index,omitempty"`
	Kind      string `json:"kind,omitempty"`
	Printable string `json:"printable,omitempty"`
	Var       string `json:"var,omitempty"`
}

type operand struct {
	Args      [][]interface{} `json:"args,omitempty"`
	Kind      string          `json:"kind,omitempty"`
	Operator  field           `json:"operator,omitempty"`
	Printable string          `json:"printable,omitempty"`
}

type Register struct {
	Namespace         string        `json:"Namespace,omitempty"`
	AsmName           string        `json:"AsmName,omitempty"`
	AltNames          []string      `json:"AltNames,omitempty"`
	Aliases           []Register    `json:"Aliases,omitempty"`
	SubRegs           []Register    `json:"SubRegs,omitempty"`
	SubRegIndices     []interface{} `json:"SubRegIndices,omitempty"`
	RegAltNameIndices []interface{} `json:"RegAltNameIndices,omitempty"`
	DwarfNumbers      []int         `json:"DwarfNumbers,omitempty"`
	CostPerUse        int           `json:"CostPerUse,omitempty"`
	CoveredBySubRegs  int           `json:"CoveredBySubRegs,omitempty"`
	HWEncoding        []int         `json:"HWEncoding,omitempty"`
	IsArtificial      int           `json:"isArtificial,omitempty"`
}

// Inst object
type Inst struct {
	Anonymous                   bool          `json:"!anonymous"`
	Fields                      []string      `json:"!fields"`
	Name                        string        `json:"!name"`
	SuperClasses                []string      `json:"!superclasses"`
	AddedComplexity             int           `json:"AddedComplexity"`
	AsmMatchConverter           string        `json:"AsmMatchConverter"`
	AsmString                   string        `json:"AsmString"`
	AsmVariantName              string        `json:"AsmVariantName"`
	CodeSize                    int           `json:"CodeSize"`
	Constraints                 string        `json:"Constraints"`
	DecoderMethod               string        `json:"DecoderMethod"`
	DecoderNamespace            string        `json:"DecoderNamespace"`
	Defs                        []Register    `json:"Defs"`
	DestructiveInstType         field         `json:"DestructiveInstType"`
	DisableEncoding             string        `json:"DisableEncoding"`
	ElementSize                 field         `json:"ElementSize"`
	EncodingInfos               interface{}   `json:"EncodingInfos"`
	F                           field         `json:"F"`
	FastISelShouldIgnore        int           `json:"FastISelShouldIgnore"`
	Form                        []int         `json:"Form"`
	InOperandList               operand       `json:"InOperandList"`
	Inst                        []interface{} `json:"Inst"`
	Itinerary                   field         `json:"Itinerary"`
	Namespace                   string        `json:"Namespace"`
	OutOperandList              operand       `json:"OutOperandList"`
	Pattern                     operand       `json:"Pattern"`
	PostEncoderMethod           string        `json:"PostEncoderMethod"`
	Predicates                  []interface{} `json:"Predicates"`
	SchedRW                     []field       `json:"SchedRW"`
	Size                        int           `json:"Size"`
	SoftFail                    []int         `json:"SoftFail"`
	TSFlags                     []int         `json:"TSFlags"`
	TwoOperandAliasConstraint   string        `json:"TwoOperandAliasConstraint"`
	Unpredictable               []int         `json:"Unpredictable"`
	UseNamedOperandTable        int           `json:"UseNamedOperandTable"`
	Uses                        []Register    `json:"Uses"`
	CanFoldAsLoad               int           `json:"canFoldAsLoad"`
	HasCompleteDecoder          int           `json:"hasCompleteDecoder"`
	HasCtrlDep                  int           `json:"hasCtrlDep"`
	HasDelaySlot                int           `json:"hasDelaySlot"`
	HasExtraDefRegAllocReq      int           `json:"hasExtraDefRegAllocReq"`
	HasExtraSrcRegAllocReq      int           `json:"hasExtraSrcRegAllocReq"`
	HasNoSchedulingInfo         int           `json:"hasNoSchedulingInfo"`
	HasPostISelHook             int           `json:"hasPostISelHook"`
	HasSideEffects              int           `json:"hasSideEffects"`
	IsAdd                       int           `json:"isAdd"`
	IsAsCheapAsAMove            int           `json:"isAsCheapAsAMove"`
	IsAsmParserOnly             int           `json:"isAsmParserOnly"`
	IsAuthenticated             int           `json:"isAuthenticated"`
	IsBarrier                   int           `json:"isBarrier"`
	IsBitcast                   int           `json:"isBitcast"`
	IsBranch                    int           `json:"isBranch"`
	IsCall                      int           `json:"isCall"`
	IsCodeGenOnly               int           `json:"isCodeGenOnly"`
	IsCommutable                int           `json:"isCommutable"`
	IsCompare                   int           `json:"isCompare"`
	IsConvergent                int           `json:"isConvergent"`
	IsConvertibleToThreeAddress int           `json:"isConvertibleToThreeAddress"`
	IsEhScopeReturn             int           `json:"isEHScopeReturn"`
	IsExtractSubreg             int           `json:"isExtractSubreg"`
	IsIndirectBranch            int           `json:"isIndirectBranch"`
	IsInsertSubreg              int           `json:"isInsertSubreg"`
	IsMoveImm                   int           `json:"isMoveImm"`
	IsMoveReg                   int           `json:"isMoveReg"`
	IsNotDuplicable             int           `json:"isNotDuplicable"`
	IsPreISelOpcode             int           `json:"isPreISelOpcode"`
	IsPredicable                int           `json:"isPredicable"`
	IsPseudo                    int           `json:"isPseudo"`
	IsReMaterializable          int           `json:"isReMaterializable"`
	IsRegSequence               int           `json:"isRegSequence"`
	IsReturn                    int           `json:"isReturn"`
	IsSelect                    int           `json:"isSelect"`
	IsTerminator                int           `json:"isTerminator"`
	IsTrap                      int           `json:"isTrap"`
	IsUnpredicable              int           `json:"isUnpredicable"`
	MayLoad                     int           `json:"mayLoad"`
	MayRaiseFpException         int           `json:"mayRaiseFPException"`
	MayStore                    int           `json:"mayStore"`
	UsesCustomInserter          int           `json:"usesCustomInserter"`
	VariadicOpsAreDefs          int           `json:"variadicOpsAreDefs"`
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {

	filename := flag.String("o", "tables2.go", "the name of the automatically generated file")
	jsonfile := flag.String("i", "./table.json", "the llvm-tblgen JSON file")
	flag.Parse()
	out, err := os.Create(*filename)
	check(err)
	defer out.Close()

	tableJSON, err := ioutil.ReadFile(*jsonfile)
	check(err)

	var table Table
	err = json.Unmarshal(tableJSON, &table)
	adrp := table["ADRP"]
	fmt.Println(adrp)
}
