"".a1 STEXT size=169 args=0x8 locals=0x68
	0x0000 00000 (asm1.go:8)	TEXT	"".a1(SB), $104-8
	0x0000 00000 (asm1.go:8)	MOVQ	(TLS), CX
	0x0009 00009 (asm1.go:8)	CMPQ	SP, 16(CX)
	0x000d 00013 (asm1.go:8)	JLS	159
	0x0013 00019 (asm1.go:8)	SUBQ	$104, SP
	0x0017 00023 (asm1.go:8)	MOVQ	BP, 96(SP)
	0x001c 00028 (asm1.go:8)	LEAQ	96(SP), BP
	0x0021 00033 (asm1.go:8)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (asm1.go:8)	FUNCDATA	$1, gclocals·86da00badb1a277fd4ad05aca8027ea8(SB)
	0x0021 00033 (asm1.go:8)	FUNCDATA	$3, gclocals·bfec7e55b3f043d1941c093912808913(SB)
	0x0021 00033 (asm1.go:8)	PCDATA	$2, $0
	0x0021 00033 (asm1.go:8)	PCDATA	$0, $0
	0x0021 00033 (asm1.go:8)	MOVQ	$0, "".~r0+112(SP)
	0x002a 00042 (asm1.go:9)	PCDATA	$0, $1
	0x002a 00042 (asm1.go:9)	XORPS	X0, X0
	0x002d 00045 (asm1.go:9)	MOVUPS	X0, ""..autotmp_1+56(SP)
	0x0032 00050 (asm1.go:9)	PCDATA	$2, $1
	0x0032 00050 (asm1.go:9)	LEAQ	""..autotmp_1+56(SP), AX
	0x0037 00055 (asm1.go:9)	MOVQ	AX, ""..autotmp_3+48(SP)
	0x003c 00060 (asm1.go:9)	TESTB	AL, (AX)
	0x003e 00062 (asm1.go:9)	PCDATA	$2, $2
	0x003e 00062 (asm1.go:9)	LEAQ	type.string(SB), CX
	0x0045 00069 (asm1.go:9)	PCDATA	$2, $1
	0x0045 00069 (asm1.go:9)	MOVQ	CX, ""..autotmp_1+56(SP)
	0x004a 00074 (asm1.go:9)	PCDATA	$2, $2
	0x004a 00074 (asm1.go:9)	LEAQ	"".statictmp_0(SB), CX
	0x0051 00081 (asm1.go:9)	PCDATA	$2, $1
	0x0051 00081 (asm1.go:9)	MOVQ	CX, ""..autotmp_1+64(SP)
	0x0056 00086 (asm1.go:9)	TESTB	AL, (AX)
	0x0058 00088 (asm1.go:9)	JMP	90
	0x005a 00090 (asm1.go:9)	MOVQ	AX, ""..autotmp_2+72(SP)
	0x005f 00095 (asm1.go:9)	MOVQ	$1, ""..autotmp_2+80(SP)
	0x0068 00104 (asm1.go:9)	MOVQ	$1, ""..autotmp_2+88(SP)
	0x0071 00113 (asm1.go:9)	PCDATA	$2, $0
	0x0071 00113 (asm1.go:9)	MOVQ	AX, (SP)
	0x0075 00117 (asm1.go:9)	MOVQ	$1, 8(SP)
	0x007e 00126 (asm1.go:9)	MOVQ	$1, 16(SP)
	0x0087 00135 (asm1.go:9)	CALL	fmt.Println(SB)
	0x008c 00140 (asm1.go:10)	PCDATA	$0, $0
	0x008c 00140 (asm1.go:10)	MOVQ	$1, "".~r0+112(SP)
	0x0095 00149 (asm1.go:10)	MOVQ	96(SP), BP
	0x009a 00154 (asm1.go:10)	ADDQ	$104, SP
	0x009e 00158 (asm1.go:10)	RET
	0x009f 00159 (asm1.go:10)	NOP
	0x009f 00159 (asm1.go:8)	PCDATA	$0, $-1
	0x009f 00159 (asm1.go:8)	PCDATA	$2, $-1
	0x009f 00159 (asm1.go:8)	CALL	runtime.morestack_noctxt(SB)
	0x00a4 00164 (asm1.go:8)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 8c  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 68 48 89 6c 24 60 48 8d 6c 24  ...H..hH.l$`H.l$
	0x0020 60 48 c7 44 24 70 00 00 00 00 0f 57 c0 0f 11 44  `H.D$p.....W...D
	0x0030 24 38 48 8d 44 24 38 48 89 44 24 30 84 00 48 8d  $8H.D$8H.D$0..H.
	0x0040 0d 00 00 00 00 48 89 4c 24 38 48 8d 0d 00 00 00  .....H.L$8H.....
	0x0050 00 48 89 4c 24 40 84 00 eb 00 48 89 44 24 48 48  .H.L$@....H.D$HH
	0x0060 c7 44 24 50 01 00 00 00 48 c7 44 24 58 01 00 00  .D$P....H.D$X...
	0x0070 00 48 89 04 24 48 c7 44 24 08 01 00 00 00 48 c7  .H..$H.D$.....H.
	0x0080 44 24 10 01 00 00 00 e8 00 00 00 00 48 c7 44 24  D$..........H.D$
	0x0090 70 01 00 00 00 48 8b 6c 24 60 48 83 c4 68 c3 e8  p....H.l$`H..h..
	0x00a0 00 00 00 00 e9 57 ff ff ff                       .....W...
	rel 5+4 t=16 TLS+0
	rel 65+4 t=15 type.string+0
	rel 77+4 t=15 "".statictmp_0+0
	rel 136+4 t=8 fmt.Println+0
	rel 160+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=282 args=0x0 locals=0x90
	0x0000 00000 (asm1.go:13)	TEXT	"".main(SB), $144-0
	0x0000 00000 (asm1.go:13)	MOVQ	(TLS), CX
	0x0009 00009 (asm1.go:13)	LEAQ	-16(SP), AX
	0x000e 00014 (asm1.go:13)	CMPQ	AX, 16(CX)
	0x0012 00018 (asm1.go:13)	JLS	272
	0x0018 00024 (asm1.go:13)	SUBQ	$144, SP
	0x001f 00031 (asm1.go:13)	MOVQ	BP, 136(SP)
	0x0027 00039 (asm1.go:13)	LEAQ	136(SP), BP
	0x002f 00047 (asm1.go:13)	FUNCDATA	$0, gclocals·3e27b3aa6b89137cce48b3379a2a6610(SB)
	0x002f 00047 (asm1.go:13)	FUNCDATA	$1, gclocals·e181efc4163d33da2446efe6e6fb5653(SB)
	0x002f 00047 (asm1.go:13)	FUNCDATA	$3, gclocals·33b901baab2acec3083d16f1ab81c65a(SB)
	0x002f 00047 (asm1.go:14)	PCDATA	$2, $1
	0x002f 00047 (asm1.go:14)	PCDATA	$0, $1
	0x002f 00047 (asm1.go:14)	LEAQ	"".a1·f(SB), AX
	0x0036 00054 (asm1.go:14)	PCDATA	$2, $0
	0x0036 00054 (asm1.go:14)	MOVQ	AX, "".f+64(SP)
	0x003b 00059 (asm1.go:15)	PCDATA	$0, $2
	0x003b 00059 (asm1.go:15)	XORPS	X0, X0
	0x003e 00062 (asm1.go:15)	MOVUPS	X0, ""..autotmp_1+96(SP)
	0x0043 00067 (asm1.go:15)	PCDATA	$2, $1
	0x0043 00067 (asm1.go:15)	LEAQ	""..autotmp_1+96(SP), AX
	0x0048 00072 (asm1.go:15)	PCDATA	$2, $0
	0x0048 00072 (asm1.go:15)	PCDATA	$0, $3
	0x0048 00072 (asm1.go:15)	MOVQ	AX, ""..autotmp_3+72(SP)
	0x004d 00077 (asm1.go:15)	PCDATA	$2, $1
	0x004d 00077 (asm1.go:15)	LEAQ	"".f+64(SP), AX
	0x0052 00082 (asm1.go:15)	PCDATA	$2, $0
	0x0052 00082 (asm1.go:15)	TESTB	AL, (AX)
	0x0054 00084 (asm1.go:15)	MOVQ	"".f+64(SP), AX
	0x0059 00089 (asm1.go:15)	MOVQ	AX, 8(SP)
	0x005e 00094 (asm1.go:15)	PCDATA	$2, $1
	0x005e 00094 (asm1.go:15)	LEAQ	type.uintptr(SB), AX
	0x0065 00101 (asm1.go:15)	PCDATA	$2, $0
	0x0065 00101 (asm1.go:15)	MOVQ	AX, (SP)
	0x0069 00105 (asm1.go:15)	CALL	runtime.convT2E64(SB)
	0x006e 00110 (asm1.go:15)	PCDATA	$2, $1
	0x006e 00110 (asm1.go:15)	MOVQ	24(SP), AX
	0x0073 00115 (asm1.go:15)	MOVQ	16(SP), CX
	0x0078 00120 (asm1.go:15)	MOVQ	CX, ""..autotmp_4+80(SP)
	0x007d 00125 (asm1.go:15)	MOVQ	AX, ""..autotmp_4+88(SP)
	0x0082 00130 (asm1.go:15)	PCDATA	$2, $2
	0x0082 00130 (asm1.go:15)	MOVQ	""..autotmp_3+72(SP), DX
	0x0087 00135 (asm1.go:15)	TESTB	AL, (DX)
	0x0089 00137 (asm1.go:15)	MOVQ	CX, (DX)
	0x008c 00140 (asm1.go:15)	PCDATA	$2, $3
	0x008c 00140 (asm1.go:15)	LEAQ	8(DX), DI
	0x0090 00144 (asm1.go:15)	PCDATA	$2, $-2
	0x0090 00144 (asm1.go:15)	PCDATA	$0, $-2
	0x0090 00144 (asm1.go:15)	CMPL	runtime.writeBarrier(SB), $0
	0x0097 00151 (asm1.go:15)	JEQ	155
	0x0099 00153 (asm1.go:15)	JMP	265
	0x009b 00155 (asm1.go:15)	MOVQ	AX, 8(DX)
	0x009f 00159 (asm1.go:15)	JMP	161
	0x00a1 00161 (asm1.go:15)	PCDATA	$2, $1
	0x00a1 00161 (asm1.go:15)	PCDATA	$0, $2
	0x00a1 00161 (asm1.go:15)	MOVQ	""..autotmp_3+72(SP), AX
	0x00a6 00166 (asm1.go:15)	TESTB	AL, (AX)
	0x00a8 00168 (asm1.go:15)	JMP	170
	0x00aa 00170 (asm1.go:15)	PCDATA	$2, $0
	0x00aa 00170 (asm1.go:15)	PCDATA	$0, $4
	0x00aa 00170 (asm1.go:15)	MOVQ	AX, ""..autotmp_2+112(SP)
	0x00af 00175 (asm1.go:15)	MOVQ	$1, ""..autotmp_2+120(SP)
	0x00b8 00184 (asm1.go:15)	MOVQ	$1, ""..autotmp_2+128(SP)
	0x00c4 00196 (asm1.go:15)	PCDATA	$2, $1
	0x00c4 00196 (asm1.go:15)	LEAQ	go.string."0x%x\n"(SB), AX
	0x00cb 00203 (asm1.go:15)	PCDATA	$2, $0
	0x00cb 00203 (asm1.go:15)	MOVQ	AX, (SP)
	0x00cf 00207 (asm1.go:15)	MOVQ	$5, 8(SP)
	0x00d8 00216 (asm1.go:15)	PCDATA	$2, $1
	0x00d8 00216 (asm1.go:15)	PCDATA	$0, $2
	0x00d8 00216 (asm1.go:15)	MOVQ	""..autotmp_2+112(SP), AX
	0x00dd 00221 (asm1.go:15)	PCDATA	$2, $0
	0x00dd 00221 (asm1.go:15)	MOVQ	AX, 16(SP)
	0x00e2 00226 (asm1.go:15)	MOVQ	$1, 24(SP)
	0x00eb 00235 (asm1.go:15)	MOVQ	$1, 32(SP)
	0x00f4 00244 (asm1.go:15)	CALL	fmt.Printf(SB)
	0x00f9 00249 (asm1.go:16)	PCDATA	$0, $1
	0x00f9 00249 (asm1.go:16)	MOVQ	136(SP), BP
	0x0101 00257 (asm1.go:16)	ADDQ	$144, SP
	0x0108 00264 (asm1.go:16)	RET
	0x0109 00265 (asm1.go:15)	PCDATA	$2, $-2
	0x0109 00265 (asm1.go:15)	PCDATA	$0, $-2
	0x0109 00265 (asm1.go:15)	CALL	runtime.gcWriteBarrier(SB)
	0x010e 00270 (asm1.go:15)	JMP	161
	0x0110 00272 (asm1.go:15)	NOP
	0x0110 00272 (asm1.go:13)	PCDATA	$0, $-1
	0x0110 00272 (asm1.go:13)	PCDATA	$2, $-1
	0x0110 00272 (asm1.go:13)	CALL	runtime.morestack_noctxt(SB)
	0x0115 00277 (asm1.go:13)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 44 24 f0 48 3b  eH..%....H.D$.H;
	0x0010 41 10 0f 86 f8 00 00 00 48 81 ec 90 00 00 00 48  A.......H......H
	0x0020 89 ac 24 88 00 00 00 48 8d ac 24 88 00 00 00 48  ..$....H..$....H
	0x0030 8d 05 00 00 00 00 48 89 44 24 40 0f 57 c0 0f 11  ......H.D$@.W...
	0x0040 44 24 60 48 8d 44 24 60 48 89 44 24 48 48 8d 44  D$`H.D$`H.D$HH.D
	0x0050 24 40 84 00 48 8b 44 24 40 48 89 44 24 08 48 8d  $@..H.D$@H.D$.H.
	0x0060 05 00 00 00 00 48 89 04 24 e8 00 00 00 00 48 8b  .....H..$.....H.
	0x0070 44 24 18 48 8b 4c 24 10 48 89 4c 24 50 48 89 44  D$.H.L$.H.L$PH.D
	0x0080 24 58 48 8b 54 24 48 84 02 48 89 0a 48 8d 7a 08  $XH.T$H..H..H.z.
	0x0090 83 3d 00 00 00 00 00 74 02 eb 6e 48 89 42 08 eb  .=.....t..nH.B..
	0x00a0 00 48 8b 44 24 48 84 00 eb 00 48 89 44 24 70 48  .H.D$H....H.D$pH
	0x00b0 c7 44 24 78 01 00 00 00 48 c7 84 24 80 00 00 00  .D$x....H..$....
	0x00c0 01 00 00 00 48 8d 05 00 00 00 00 48 89 04 24 48  ....H......H..$H
	0x00d0 c7 44 24 08 05 00 00 00 48 8b 44 24 70 48 89 44  .D$.....H.D$pH.D
	0x00e0 24 10 48 c7 44 24 18 01 00 00 00 48 c7 44 24 20  $.H.D$.....H.D$ 
	0x00f0 01 00 00 00 e8 00 00 00 00 48 8b ac 24 88 00 00  .........H..$...
	0x0100 00 48 81 c4 90 00 00 00 c3 e8 00 00 00 00 eb 91  .H..............
	0x0110 e8 00 00 00 00 e9 e6 fe ff ff                    ..........
	rel 5+4 t=16 TLS+0
	rel 50+4 t=15 "".a1·f+0
	rel 97+4 t=15 type.uintptr+0
	rel 106+4 t=8 runtime.convT2E64+0
	rel 146+4 t=15 runtime.writeBarrier+-1
	rel 199+4 t=15 go.string."0x%x\n"+0
	rel 245+4 t=8 fmt.Printf+0
	rel 266+4 t=8 runtime.gcWriteBarrier+0
	rel 273+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=100 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	93
	0x000f 00015 (<autogenerated>:1)	SUBQ	$8, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, (SP)
	0x0017 00023 (<autogenerated>:1)	LEAQ	(SP), BP
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x001b 00027 (<autogenerated>:1)	PCDATA	$2, $0
	0x001b 00027 (<autogenerated>:1)	PCDATA	$0, $0
	0x001b 00027 (<autogenerated>:1)	CMPB	"".initdone·(SB), $1
	0x0022 00034 (<autogenerated>:1)	JHI	38
	0x0024 00036 (<autogenerated>:1)	JMP	47
	0x0026 00038 (<autogenerated>:1)	PCDATA	$2, $-2
	0x0026 00038 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0026 00038 (<autogenerated>:1)	MOVQ	(SP), BP
	0x002a 00042 (<autogenerated>:1)	ADDQ	$8, SP
	0x002e 00046 (<autogenerated>:1)	RET
	0x002f 00047 (<autogenerated>:1)	PCDATA	$2, $0
	0x002f 00047 (<autogenerated>:1)	PCDATA	$0, $0
	0x002f 00047 (<autogenerated>:1)	CMPB	"".initdone·(SB), $1
	0x0036 00054 (<autogenerated>:1)	JEQ	58
	0x0038 00056 (<autogenerated>:1)	JMP	65
	0x003a 00058 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x003f 00063 (<autogenerated>:1)	UNDEF
	0x0041 00065 (<autogenerated>:1)	MOVB	$1, "".initdone·(SB)
	0x0048 00072 (<autogenerated>:1)	CALL	fmt.init(SB)
	0x004d 00077 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x0054 00084 (<autogenerated>:1)	MOVQ	(SP), BP
	0x0058 00088 (<autogenerated>:1)	ADDQ	$8, SP
	0x005c 00092 (<autogenerated>:1)	RET
	0x005d 00093 (<autogenerated>:1)	NOP
	0x005d 00093 (<autogenerated>:1)	PCDATA	$0, $-1
	0x005d 00093 (<autogenerated>:1)	PCDATA	$2, $-1
	0x005d 00093 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0062 00098 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 4e 48  eH..%....H;a.vNH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 80 3d 00 00 00  ...H.,$H.,$.=...
	0x0020 00 01 77 02 eb 09 48 8b 2c 24 48 83 c4 08 c3 80  ..w...H.,$H.....
	0x0030 3d 00 00 00 00 01 74 02 eb 07 e8 00 00 00 00 0f  =.....t.........
	0x0040 0b c6 05 00 00 00 00 01 e8 00 00 00 00 c6 05 00  ................
	0x0050 00 00 00 02 48 8b 2c 24 48 83 c4 08 c3 e8 00 00  ....H.,$H.......
	0x0060 00 00 eb 9c                                      ....
	rel 5+4 t=16 TLS+0
	rel 29+4 t=15 "".initdone·+-1
	rel 49+4 t=15 "".initdone·+-1
	rel 59+4 t=8 runtime.throwinit+0
	rel 67+4 t=15 "".initdone·+-1
	rel 73+4 t=8 fmt.init+0
	rel 79+4 t=15 "".initdone·+-1
	rel 94+4 t=8 runtime.morestack_noctxt+0
go.string."a1" SRODATA dupok size=2
	0x0000 61 31                                            a1
go.loc."".a1 SDWARFLOC size=0
go.info."".a1 SDWARFINFO size=44
	0x0000 02 22 22 2e 61 31 00 00 00 00 00 00 00 00 00 00  ."".a1..........
	0x0010 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0e 7e  ...............~
	0x0020 72 30 00 01 08 00 00 00 00 01 9c 00              r0..........
	rel 7+8 t=1 "".a1+0
	rel 15+8 t=1 "".a1+169
	rel 25+4 t=29 gofile../Users/yongfuchen/go/src/go_test/asm_go/asm1.go+0
	rel 37+4 t=28 go.info.int+0
go.range."".a1 SDWARFRANGE size=0
go.isstmt."".a1 SDWARFMISC size=0
	0x0000 04 13 04 0e 03 0c 01 5f 02 09 01 0a 02 0a 00     ......._.......
go.string."0x%x\n" SRODATA dupok size=5
	0x0000 30 78 25 78 0a                                   0x%x.
go.loc."".main SDWARFLOC size=0
go.info."".main SDWARFINFO size=45
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 09 66 00 0e 00 00 00 00 03 91 a8 7f 00           .f...........
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+282
	rel 27+4 t=29 gofile../Users/yongfuchen/go/src/go_test/asm_go/asm1.go+0
	rel 36+4 t=28 go.info.func() int+0
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 18 04 17 03 07 01 05 02 03 01 bb 01 02 10 01  ................
	0x0010 07 02 0a 00                                      ....
go.loc."".init SDWARFLOC size=0
go.info."".init SDWARFINFO size=33
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+100
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
go.isstmt."".init SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 07 01 04 02 09 01 10 02 09 01 15  ................
	0x0010 02 07 00                                         ...
"".statictmp_0 SRODATA size=16
	0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."a1"+0
"".initdone· SNOPTRBSS size=1
"".a1·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".a1+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*interface {}- SRODATA dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.02 SRODATA dupok size=1
	0x0000 02                                               .
type.interface {} SRODATA dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*interface {}-+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}- SRODATA dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} SRODATA dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}-+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
type..namedata.*[1]interface {}- SRODATA dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.[1]interface {} SRODATA dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type.*[1]interface {} SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 00 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}-+0
	rel 48+8 t=1 type.[1]interface {}+0
type..namedata.*func() int- SRODATA dupok size=14
	0x0000 00 00 0b 2a 66 75 6e 63 28 29 20 69 6e 74        ...*func() int
type.*func() int SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 51 b2 ed c6 00 08 08 36 00 00 00 00 00 00 00 00  Q......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() int-+0
	rel 48+8 t=1 type.func() int+0
type.func() int SRODATA dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 e5 86 39 e0 02 08 08 33 00 00 00 00 00 00 00 00  ..9....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() int-+0
	rel 44+4 t=6 type.*func() int+0
	rel 56+8 t=1 type.int+0
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
type..importpath.unsafe. SRODATA dupok size=9
	0x0000 00 00 06 75 6e 73 61 66 65                       ...unsafe
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·86da00badb1a277fd4ad05aca8027ea8 SRODATA dupok size=10
	0x0000 02 00 00 00 06 00 00 00 00 04                    ..........
gclocals·bfec7e55b3f043d1941c093912808913 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 01 03                 ...........
gclocals·3e27b3aa6b89137cce48b3379a2a6610 SRODATA dupok size=8
	0x0000 05 00 00 00 00 00 00 00                          ........
gclocals·e181efc4163d33da2446efe6e6fb5653 SRODATA dupok size=18
	0x0000 05 00 00 00 09 00 00 00 00 00 01 00 21 00 23 00  ............!.#.
	0x0010 61 00                                            a.
gclocals·33b901baab2acec3083d16f1ab81c65a SRODATA dupok size=12
	0x0000 04 00 00 00 07 00 00 00 00 01 05 45              ...........E
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
