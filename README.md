# go-diet

`go-diet` will analyze your source code and make the following improvements to both reduce the size of your executable as well as to improve the execution speed:
- Use named return values to eliminate the need to explictly instantiate structs.
- Eliminate `autogenerated` functions
 
## Named return values

Naming return values can lead to less code being generated

```go
func NamedReturnParams(i int) (oi objectInfo) {

	if i == 1 {
		return
	}

	if i == 2 {
		return
	}

	if i == 3 {
		return
	}

	return
```

```
"".NamedReturnParams t=1 size=67 args=0x40 locals=0x0
	0x0000 00000 (named-return-params.go:10)	TEXT	"".NamedReturnParams(SB), $0-64
	0x0000 00000 (named-return-params.go:10)	NOP
	0x0000 00000 (named-return-params.go:10)	NOP
	0x0000 00000 (named-return-params.go:10)	FUNCDATA	$0, gclocals·895d0569a38a56443b84805daa09d838(SB)
	0x0000 00000 (named-return-params.go:10)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (named-return-params.go:10)	MOVQ	$0, "".oi+16(FP)
	0x0009 00009 (named-return-params.go:10)	LEAQ	"".oi+24(FP), DI
	0x000e 00014 (named-return-params.go:10)	XORPS	X0, X0
	0x0011 00017 (named-return-params.go:10)	ADDQ	$-16, DI
	0x0015 00021 (named-return-params.go:10)	DUFFZERO	$288
	0x0028 00040 (named-return-params.go:12)	MOVQ	"".i+8(FP), AX
	0x002d 00045 (named-return-params.go:12)	CMPQ	AX, $1
	0x0031 00049 (named-return-params.go:12)	JEQ	$0, 66
	0x0033 00051 (named-return-params.go:16)	CMPQ	AX, $2
	0x0037 00055 (named-return-params.go:16)	JEQ	$0, 65
	0x0039 00057 (named-return-params.go:20)	CMPQ	AX, $3
	0x003d 00061 (named-return-params.go:20)	JNE	64
	0x003f 00063 (named-return-params.go:21)	RET
	0x0040 00064 (named-return-params.go:24)	RET
	0x0041 00065 (named-return-params.go:17)	RET
	0x0042 00066 (named-return-params.go:13)	RET
```

#### Explicitly instantiating for every return

```go
func NoNamedReturnParams(i int) (objectInfo) {

	if i == 1 {
		return objectInfo{}
	}

	if i == 2 {
		return objectInfo{}
	}

	if i == 3 {
		return objectInfo{}
	}

	return objectInfo{}
}
```

```
"".NoNamedReturnParams t=1 size=243 args=0x40 locals=0x0
	0x0000 00000 (no-named-return-params.go:10)	TEXT	"".NoNamedReturnParams(SB), $0-64
	0x0000 00000 (no-named-return-params.go:10)	NOP
	0x0000 00000 (no-named-return-params.go:10)	NOP
	0x0000 00000 (no-named-return-params.go:10)	FUNCDATA	$0, gclocals·895d0569a38a56443b84805daa09d838(SB)
	0x0000 00000 (no-named-return-params.go:10)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (no-named-return-params.go:10)	MOVQ	$0, "".~r1+16(FP)
	0x0009 00009 (no-named-return-params.go:10)	LEAQ	"".~r1+24(FP), DI
	0x000e 00014 (no-named-return-params.go:10)	XORPS	X0, X0
	0x0011 00017 (no-named-return-params.go:10)	ADDQ	$-16, DI
	0x0015 00021 (no-named-return-params.go:10)	DUFFZERO	$288
	0x0028 00040 (no-named-return-params.go:12)	MOVQ	"".i+8(FP), AX
	0x002d 00045 (no-named-return-params.go:12)	CMPQ	AX, $1
	0x0031 00049 (no-named-return-params.go:12)	JEQ	$0, 199
	0x0037 00055 (no-named-return-params.go:16)	CMPQ	AX, $2
	0x003b 00059 (no-named-return-params.go:16)	JEQ	$0, 155
	0x003d 00061 (no-named-return-params.go:20)	CMPQ	AX, $3
	0x0041 00065 (no-named-return-params.go:20)	JNE	111
	0x0043 00067 (no-named-return-params.go:21)	MOVQ	"".statictmp_2(SB), AX
	0x004a 00074 (no-named-return-params.go:21)	MOVQ	AX, "".~r1+16(FP)
	0x004f 00079 (no-named-return-params.go:21)	LEAQ	"".~r1+24(FP), DI
	0x0054 00084 (no-named-return-params.go:21)	LEAQ	"".statictmp_2+8(SB), SI
	0x005b 00091 (no-named-return-params.go:21)	DUFFCOPY	$854
	0x006e 00110 (no-named-return-params.go:21)	RET
	0x006f 00111 (no-named-return-params.go:24)	MOVQ	"".statictmp_3(SB), AX
	0x0076 00118 (no-named-return-params.go:24)	MOVQ	AX, "".~r1+16(FP)
	0x007b 00123 (no-named-return-params.go:24)	LEAQ	"".~r1+24(FP), DI
	0x0080 00128 (no-named-return-params.go:24)	LEAQ	"".statictmp_3+8(SB), SI
	0x0087 00135 (no-named-return-params.go:24)	DUFFCOPY	$854
	0x009a 00154 (no-named-return-params.go:24)	RET
	0x009b 00155 (no-named-return-params.go:17)	MOVQ	"".statictmp_1(SB), AX
	0x00a2 00162 (no-named-return-params.go:17)	MOVQ	AX, "".~r1+16(FP)
	0x00a7 00167 (no-named-return-params.go:17)	LEAQ	"".~r1+24(FP), DI
	0x00ac 00172 (no-named-return-params.go:17)	LEAQ	"".statictmp_1+8(SB), SI
	0x00b3 00179 (no-named-return-params.go:17)	DUFFCOPY	$854
	0x00c6 00198 (no-named-return-params.go:17)	RET
	0x00c7 00199 (no-named-return-params.go:13)	MOVQ	"".statictmp_0(SB), AX
	0x00ce 00206 (no-named-return-params.go:13)	MOVQ	AX, "".~r1+16(FP)
	0x00d3 00211 (no-named-return-params.go:13)	LEAQ	"".~r1+24(FP), DI
	0x00d8 00216 (no-named-return-params.go:13)	LEAQ	"".statictmp_0+8(SB), SI
	0x00df 00223 (no-named-return-params.go:13)	DUFFCOPY	$854
	0x00f2 00242 (no-named-return-params.go:13)	RET
```

