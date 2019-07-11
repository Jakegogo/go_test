"".time4_b STEXT size=556 args=0x10 locals=0xb8
	0x0000 00000 (target_func.go:10)	TEXT	"".time4_b(SB), $184-16
	0x0000 00000 (target_func.go:10)	MOVQ	(TLS), CX
	0x0009 00009 (target_func.go:10)	LEAQ	-56(SP), AX
	0x000e 00014 (target_func.go:10)	CMPQ	AX, 16(CX)
	0x0012 00018 (target_func.go:10)	JLS	546
	0x0018 00024 (target_func.go:10)	SUBQ	$184, SP
	0x001f 00031 (target_func.go:10)	MOVQ	BP, 176(SP)
	0x0027 00039 (target_func.go:10)	LEAQ	176(SP), BP
	0x002f 00047 (target_func.go:10)	FUNCDATA	$0, gclocals·f6bd6b3389b872033d462029172c8612(SB)
	0x002f 00047 (target_func.go:10)	FUNCDATA	$1, gclocals·ed20391e33635a5bf94ce2d547e0f407(SB)
	0x002f 00047 (target_func.go:10)	FUNCDATA	$3, gclocals·fedbdc973195b17703414cfc4d862ea8(SB)
	0x002f 00047 (target_func.go:10)	PCDATA	$2, $0
	0x002f 00047 (target_func.go:10)	PCDATA	$0, $0
	0x002f 00047 (target_func.go:10)	MOVQ	$0, "".~r1+200(SP)
	0x003b 00059 (target_func.go:11)	MOVQ	"".a+192(SP), AX
	0x0043 00067 (target_func.go:11)	LEAQ	(AX)(AX*4), CX
	0x0047 00071 (target_func.go:11)	LEAQ	(AX)(CX*2), AX
	0x004b 00075 (target_func.go:11)	MOVQ	AX, "".result+48(SP)
	0x0050 00080 (target_func.go:12)	CMPQ	AX, $50
	0x0054 00084 (target_func.go:12)	JLT	88
	0x0056 00086 (target_func.go:12)	JMP	116
	0x0058 00088 (target_func.go:13)	MOVQ	$3, "".~r1+200(SP)
	0x0064 00100 (target_func.go:13)	MOVQ	176(SP), BP
	0x006c 00108 (target_func.go:13)	ADDQ	$184, SP
	0x0073 00115 (target_func.go:13)	RET
	0x0074 00116 (target_func.go:15)	CMPQ	AX, $100
	0x0078 00120 (target_func.go:15)	JLT	124
	0x007a 00122 (target_func.go:15)	JMP	152
	0x007c 00124 (target_func.go:16)	MOVQ	$6, "".~r1+200(SP)
	0x0088 00136 (target_func.go:16)	MOVQ	176(SP), BP
	0x0090 00144 (target_func.go:16)	ADDQ	$184, SP
	0x0097 00151 (target_func.go:16)	RET
	0x0098 00152 (target_func.go:18)	CMPQ	AX, $200
	0x009e 00158 (target_func.go:18)	JLT	162
	0x00a0 00160 (target_func.go:18)	JMP	190
	0x00a2 00162 (target_func.go:19)	MOVQ	$9, "".~r1+200(SP)
	0x00ae 00174 (target_func.go:19)	MOVQ	176(SP), BP
	0x00b6 00182 (target_func.go:19)	ADDQ	$184, SP
	0x00bd 00189 (target_func.go:19)	RET
	0x00be 00190 (target_func.go:21)	MOVQ	$3, "".C_s(SB)
	0x00c9 00201 (target_func.go:22)	PCDATA	$0, $1
	0x00c9 00201 (target_func.go:22)	XORPS	X0, X0
	0x00cc 00204 (target_func.go:22)	MOVUPS	X0, ""..autotmp_3+96(SP)
	0x00d1 00209 (target_func.go:22)	PCDATA	$2, $1
	0x00d1 00209 (target_func.go:22)	LEAQ	""..autotmp_3+96(SP), AX
	0x00d6 00214 (target_func.go:22)	MOVQ	AX, ""..autotmp_7+72(SP)
	0x00db 00219 (target_func.go:22)	TESTB	AL, (AX)
	0x00dd 00221 (target_func.go:22)	PCDATA	$2, $2
	0x00dd 00221 (target_func.go:22)	LEAQ	type.string(SB), CX
	0x00e4 00228 (target_func.go:22)	PCDATA	$2, $1
	0x00e4 00228 (target_func.go:22)	MOVQ	CX, ""..autotmp_3+96(SP)
	0x00e9 00233 (target_func.go:22)	PCDATA	$2, $2
	0x00e9 00233 (target_func.go:22)	LEAQ	"".statictmp_0(SB), CX
	0x00f0 00240 (target_func.go:22)	PCDATA	$2, $1
	0x00f0 00240 (target_func.go:22)	MOVQ	CX, ""..autotmp_3+104(SP)
	0x00f5 00245 (target_func.go:22)	TESTB	AL, (AX)
	0x00f7 00247 (target_func.go:22)	JMP	249
	0x00f9 00249 (target_func.go:22)	MOVQ	AX, ""..autotmp_6+152(SP)
	0x0101 00257 (target_func.go:22)	MOVQ	$1, ""..autotmp_6+160(SP)
	0x010d 00269 (target_func.go:22)	MOVQ	$1, ""..autotmp_6+168(SP)
	0x0119 00281 (target_func.go:22)	PCDATA	$2, $0
	0x0119 00281 (target_func.go:22)	MOVQ	AX, (SP)
	0x011d 00285 (target_func.go:22)	MOVQ	$1, 8(SP)
	0x0126 00294 (target_func.go:22)	MOVQ	$1, 16(SP)
	0x012f 00303 (target_func.go:22)	CALL	fmt.Println(SB)
	0x0134 00308 (target_func.go:23)	PCDATA	$0, $2
	0x0134 00308 (target_func.go:23)	XORPS	X0, X0
	0x0137 00311 (target_func.go:23)	MOVUPS	X0, ""..autotmp_4+80(SP)
	0x013c 00316 (target_func.go:23)	PCDATA	$2, $1
	0x013c 00316 (target_func.go:23)	LEAQ	""..autotmp_4+80(SP), AX
	0x0141 00321 (target_func.go:23)	PCDATA	$2, $0
	0x0141 00321 (target_func.go:23)	PCDATA	$0, $3
	0x0141 00321 (target_func.go:23)	MOVQ	AX, ""..autotmp_9+64(SP)
	0x0146 00326 (target_func.go:23)	PCDATA	$2, $1
	0x0146 00326 (target_func.go:23)	LEAQ	type.int(SB), AX
	0x014d 00333 (target_func.go:23)	PCDATA	$2, $0
	0x014d 00333 (target_func.go:23)	MOVQ	AX, (SP)
	0x0151 00337 (target_func.go:23)	MOVQ	"".C_s(SB), AX
	0x0158 00344 (target_func.go:23)	MOVQ	AX, 8(SP)
	0x015d 00349 (target_func.go:23)	CALL	runtime.convT2E64(SB)
	0x0162 00354 (target_func.go:23)	PCDATA	$2, $1
	0x0162 00354 (target_func.go:23)	MOVQ	24(SP), AX
	0x0167 00359 (target_func.go:23)	MOVQ	16(SP), CX
	0x016c 00364 (target_func.go:23)	MOVQ	CX, ""..autotmp_10+112(SP)
	0x0171 00369 (target_func.go:23)	MOVQ	AX, ""..autotmp_10+120(SP)
	0x0176 00374 (target_func.go:23)	PCDATA	$2, $3
	0x0176 00374 (target_func.go:23)	MOVQ	""..autotmp_9+64(SP), DX
	0x017b 00379 (target_func.go:23)	TESTB	AL, (DX)
	0x017d 00381 (target_func.go:23)	MOVQ	CX, (DX)
	0x0180 00384 (target_func.go:23)	PCDATA	$2, $4
	0x0180 00384 (target_func.go:23)	LEAQ	8(DX), DI
	0x0184 00388 (target_func.go:23)	PCDATA	$2, $-2
	0x0184 00388 (target_func.go:23)	PCDATA	$0, $-2
	0x0184 00388 (target_func.go:23)	CMPL	runtime.writeBarrier(SB), $0
	0x018b 00395 (target_func.go:23)	JEQ	402
	0x018d 00397 (target_func.go:23)	JMP	536
	0x0192 00402 (target_func.go:23)	MOVQ	AX, 8(DX)
	0x0196 00406 (target_func.go:23)	JMP	408
	0x0198 00408 (target_func.go:23)	PCDATA	$2, $1
	0x0198 00408 (target_func.go:23)	PCDATA	$0, $2
	0x0198 00408 (target_func.go:23)	MOVQ	""..autotmp_9+64(SP), AX
	0x019d 00413 (target_func.go:23)	TESTB	AL, (AX)
	0x019f 00415 (target_func.go:23)	JMP	417
	0x01a1 00417 (target_func.go:23)	MOVQ	AX, ""..autotmp_8+128(SP)
	0x01a9 00425 (target_func.go:23)	MOVQ	$1, ""..autotmp_8+136(SP)
	0x01b5 00437 (target_func.go:23)	MOVQ	$1, ""..autotmp_8+144(SP)
	0x01c1 00449 (target_func.go:23)	PCDATA	$2, $0
	0x01c1 00449 (target_func.go:23)	MOVQ	AX, (SP)
	0x01c5 00453 (target_func.go:23)	MOVQ	$1, 8(SP)
	0x01ce 00462 (target_func.go:23)	MOVQ	$1, 16(SP)
	0x01d7 00471 (target_func.go:23)	CALL	fmt.Println(SB)
	0x01dc 00476 (target_func.go:26)	PCDATA	$0, $0
	0x01dc 00476 (target_func.go:26)	MOVQ	$3, (SP)
	0x01e4 00484 (target_func.go:26)	CALL	"".time6_b(SB)
	0x01e9 00489 (target_func.go:26)	MOVQ	8(SP), AX
	0x01ee 00494 (target_func.go:26)	MOVQ	AX, ""..autotmp_5+56(SP)
	0x01f3 00499 (target_func.go:26)	MOVQ	"".result+48(SP), CX
	0x01f8 00504 (target_func.go:26)	IMULQ	AX, CX
	0x01fc 00508 (target_func.go:26)	LEAQ	1(CX), AX
	0x0200 00512 (target_func.go:26)	MOVQ	AX, "".~r1+200(SP)
	0x0208 00520 (target_func.go:26)	MOVQ	176(SP), BP
	0x0210 00528 (target_func.go:26)	ADDQ	$184, SP
	0x0217 00535 (target_func.go:26)	RET
	0x0218 00536 (target_func.go:23)	PCDATA	$2, $-2
	0x0218 00536 (target_func.go:23)	PCDATA	$0, $-2
	0x0218 00536 (target_func.go:23)	CALL	runtime.gcWriteBarrier(SB)
	0x021d 00541 (target_func.go:23)	JMP	408
	0x0222 00546 (target_func.go:23)	NOP
	0x0222 00546 (target_func.go:10)	PCDATA	$0, $-1
	0x0222 00546 (target_func.go:10)	PCDATA	$2, $-1
	0x0222 00546 (target_func.go:10)	CALL	runtime.morestack_noctxt(SB)
	0x0227 00551 (target_func.go:10)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 8d 44 24 c8 48 3b  eH..%....H.D$.H;
	0x0010 41 10 0f 86 0a 02 00 00 48 81 ec b8 00 00 00 48  A.......H......H
	0x0020 89 ac 24 b0 00 00 00 48 8d ac 24 b0 00 00 00 48  ..$....H..$....H
	0x0030 c7 84 24 c8 00 00 00 00 00 00 00 48 8b 84 24 c0  ..$........H..$.
	0x0040 00 00 00 48 8d 0c 80 48 8d 04 48 48 89 44 24 30  ...H...H..HH.D$0
	0x0050 48 83 f8 32 7c 02 eb 1c 48 c7 84 24 c8 00 00 00  H..2|...H..$....
	0x0060 03 00 00 00 48 8b ac 24 b0 00 00 00 48 81 c4 b8  ....H..$....H...
	0x0070 00 00 00 c3 48 83 f8 64 7c 02 eb 1c 48 c7 84 24  ....H..d|...H..$
	0x0080 c8 00 00 00 06 00 00 00 48 8b ac 24 b0 00 00 00  ........H..$....
	0x0090 48 81 c4 b8 00 00 00 c3 48 3d c8 00 00 00 7c 02  H.......H=....|.
	0x00a0 eb 1c 48 c7 84 24 c8 00 00 00 09 00 00 00 48 8b  ..H..$........H.
	0x00b0 ac 24 b0 00 00 00 48 81 c4 b8 00 00 00 c3 48 c7  .$....H.......H.
	0x00c0 05 00 00 00 00 03 00 00 00 0f 57 c0 0f 11 44 24  ..........W...D$
	0x00d0 60 48 8d 44 24 60 48 89 44 24 48 84 00 48 8d 0d  `H.D$`H.D$H..H..
	0x00e0 00 00 00 00 48 89 4c 24 60 48 8d 0d 00 00 00 00  ....H.L$`H......
	0x00f0 48 89 4c 24 68 84 00 eb 00 48 89 84 24 98 00 00  H.L$h....H..$...
	0x0100 00 48 c7 84 24 a0 00 00 00 01 00 00 00 48 c7 84  .H..$........H..
	0x0110 24 a8 00 00 00 01 00 00 00 48 89 04 24 48 c7 44  $........H..$H.D
	0x0120 24 08 01 00 00 00 48 c7 44 24 10 01 00 00 00 e8  $.....H.D$......
	0x0130 00 00 00 00 0f 57 c0 0f 11 44 24 50 48 8d 44 24  .....W...D$PH.D$
	0x0140 50 48 89 44 24 40 48 8d 05 00 00 00 00 48 89 04  PH.D$@H......H..
	0x0150 24 48 8b 05 00 00 00 00 48 89 44 24 08 e8 00 00  $H......H.D$....
	0x0160 00 00 48 8b 44 24 18 48 8b 4c 24 10 48 89 4c 24  ..H.D$.H.L$.H.L$
	0x0170 70 48 89 44 24 78 48 8b 54 24 40 84 02 48 89 0a  pH.D$xH.T$@..H..
	0x0180 48 8d 7a 08 83 3d 00 00 00 00 00 74 05 e9 86 00  H.z..=.....t....
	0x0190 00 00 48 89 42 08 eb 00 48 8b 44 24 40 84 00 eb  ..H.B...H.D$@...
	0x01a0 00 48 89 84 24 80 00 00 00 48 c7 84 24 88 00 00  .H..$....H..$...
	0x01b0 00 01 00 00 00 48 c7 84 24 90 00 00 00 01 00 00  .....H..$.......
	0x01c0 00 48 89 04 24 48 c7 44 24 08 01 00 00 00 48 c7  .H..$H.D$.....H.
	0x01d0 44 24 10 01 00 00 00 e8 00 00 00 00 48 c7 04 24  D$..........H..$
	0x01e0 03 00 00 00 e8 00 00 00 00 48 8b 44 24 08 48 89  .........H.D$.H.
	0x01f0 44 24 38 48 8b 4c 24 30 48 0f af c8 48 8d 41 01  D$8H.L$0H...H.A.
	0x0200 48 89 84 24 c8 00 00 00 48 8b ac 24 b0 00 00 00  H..$....H..$....
	0x0210 48 81 c4 b8 00 00 00 c3 e8 00 00 00 00 e9 76 ff  H.............v.
	0x0220 ff ff e8 00 00 00 00 e9 d4 fd ff ff              ............
	rel 5+4 t=16 TLS+0
	rel 193+4 t=15 "".C_s+-4
	rel 224+4 t=15 type.string+0
	rel 236+4 t=15 "".statictmp_0+0
	rel 304+4 t=8 fmt.Println+0
	rel 329+4 t=15 type.int+0
	rel 340+4 t=15 "".C_s+0
	rel 350+4 t=8 runtime.convT2E64+0
	rel 390+4 t=15 runtime.writeBarrier+-1
	rel 472+4 t=8 fmt.Println+0
	rel 485+4 t=8 "".time6_b+0
	rel 537+4 t=8 runtime.gcWriteBarrier+0
	rel 547+4 t=8 runtime.morestack_noctxt+0
"".trace_b STEXT size=236 args=0x8 locals=0x80
	0x0000 00000 (target_func.go:29)	TEXT	"".trace_b(SB), $128-8
	0x0000 00000 (target_func.go:29)	MOVQ	(TLS), CX
	0x0009 00009 (target_func.go:29)	CMPQ	SP, 16(CX)
	0x000d 00013 (target_func.go:29)	JLS	226
	0x0013 00019 (target_func.go:29)	SUBQ	$128, SP
	0x001a 00026 (target_func.go:29)	MOVQ	BP, 120(SP)
	0x001f 00031 (target_func.go:29)	LEAQ	120(SP), BP
	0x0024 00036 (target_func.go:29)	FUNCDATA	$0, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x0024 00036 (target_func.go:29)	FUNCDATA	$1, gclocals·1d53b44fb3211f44318e421290131d18(SB)
	0x0024 00036 (target_func.go:29)	FUNCDATA	$3, gclocals·3ee7973a6249c7a7f2fab45f78192248(SB)
	0x0024 00036 (target_func.go:30)	PCDATA	$2, $0
	0x0024 00036 (target_func.go:30)	PCDATA	$0, $0
	0x0024 00036 (target_func.go:30)	MOVQ	"".i+136(SP), AX
	0x002c 00044 (target_func.go:30)	MOVQ	AX, ""..autotmp_2+48(SP)
	0x0031 00049 (target_func.go:30)	PCDATA	$0, $1
	0x0031 00049 (target_func.go:30)	XORPS	X0, X0
	0x0034 00052 (target_func.go:30)	MOVUPS	X0, ""..autotmp_1+80(SP)
	0x0039 00057 (target_func.go:30)	PCDATA	$2, $1
	0x0039 00057 (target_func.go:30)	LEAQ	""..autotmp_1+80(SP), AX
	0x003e 00062 (target_func.go:30)	PCDATA	$2, $0
	0x003e 00062 (target_func.go:30)	PCDATA	$0, $2
	0x003e 00062 (target_func.go:30)	MOVQ	AX, ""..autotmp_4+56(SP)
	0x0043 00067 (target_func.go:30)	PCDATA	$2, $1
	0x0043 00067 (target_func.go:30)	LEAQ	type.int(SB), AX
	0x004a 00074 (target_func.go:30)	PCDATA	$2, $0
	0x004a 00074 (target_func.go:30)	MOVQ	AX, (SP)
	0x004e 00078 (target_func.go:30)	MOVQ	""..autotmp_2+48(SP), AX
	0x0053 00083 (target_func.go:30)	MOVQ	AX, 8(SP)
	0x0058 00088 (target_func.go:30)	CALL	runtime.convT2E64(SB)
	0x005d 00093 (target_func.go:30)	MOVQ	16(SP), AX
	0x0062 00098 (target_func.go:30)	PCDATA	$2, $2
	0x0062 00098 (target_func.go:30)	MOVQ	24(SP), CX
	0x0067 00103 (target_func.go:30)	MOVQ	AX, ""..autotmp_5+64(SP)
	0x006c 00108 (target_func.go:30)	MOVQ	CX, ""..autotmp_5+72(SP)
	0x0071 00113 (target_func.go:30)	PCDATA	$2, $3
	0x0071 00113 (target_func.go:30)	MOVQ	""..autotmp_4+56(SP), DX
	0x0076 00118 (target_func.go:30)	TESTB	AL, (DX)
	0x0078 00120 (target_func.go:30)	MOVQ	AX, (DX)
	0x007b 00123 (target_func.go:30)	PCDATA	$2, $4
	0x007b 00123 (target_func.go:30)	LEAQ	8(DX), DI
	0x007f 00127 (target_func.go:30)	PCDATA	$2, $-2
	0x007f 00127 (target_func.go:30)	PCDATA	$0, $-2
	0x007f 00127 (target_func.go:30)	CMPL	runtime.writeBarrier(SB), $0
	0x0086 00134 (target_func.go:30)	JEQ	138
	0x0088 00136 (target_func.go:30)	JMP	216
	0x008a 00138 (target_func.go:30)	MOVQ	CX, 8(DX)
	0x008e 00142 (target_func.go:30)	JMP	144
	0x0090 00144 (target_func.go:30)	PCDATA	$2, $1
	0x0090 00144 (target_func.go:30)	PCDATA	$0, $1
	0x0090 00144 (target_func.go:30)	MOVQ	""..autotmp_4+56(SP), AX
	0x0095 00149 (target_func.go:30)	TESTB	AL, (AX)
	0x0097 00151 (target_func.go:30)	JMP	153
	0x0099 00153 (target_func.go:30)	MOVQ	AX, ""..autotmp_3+96(SP)
	0x009e 00158 (target_func.go:30)	MOVQ	$1, ""..autotmp_3+104(SP)
	0x00a7 00167 (target_func.go:30)	MOVQ	$1, ""..autotmp_3+112(SP)
	0x00b0 00176 (target_func.go:30)	PCDATA	$2, $0
	0x00b0 00176 (target_func.go:30)	MOVQ	AX, (SP)
	0x00b4 00180 (target_func.go:30)	MOVQ	$1, 8(SP)
	0x00bd 00189 (target_func.go:30)	MOVQ	$1, 16(SP)
	0x00c6 00198 (target_func.go:30)	CALL	fmt.Println(SB)
	0x00cb 00203 (target_func.go:31)	PCDATA	$0, $0
	0x00cb 00203 (target_func.go:31)	MOVQ	120(SP), BP
	0x00d0 00208 (target_func.go:31)	ADDQ	$128, SP
	0x00d7 00215 (target_func.go:31)	RET
	0x00d8 00216 (target_func.go:30)	PCDATA	$2, $-2
	0x00d8 00216 (target_func.go:30)	PCDATA	$0, $-2
	0x00d8 00216 (target_func.go:30)	MOVQ	CX, AX
	0x00db 00219 (target_func.go:30)	CALL	runtime.gcWriteBarrier(SB)
	0x00e0 00224 (target_func.go:30)	JMP	144
	0x00e2 00226 (target_func.go:30)	NOP
	0x00e2 00226 (target_func.go:29)	PCDATA	$0, $-1
	0x00e2 00226 (target_func.go:29)	PCDATA	$2, $-1
	0x00e2 00226 (target_func.go:29)	CALL	runtime.morestack_noctxt(SB)
	0x00e7 00231 (target_func.go:29)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 cf  eH..%....H;a....
	0x0010 00 00 00 48 81 ec 80 00 00 00 48 89 6c 24 78 48  ...H......H.l$xH
	0x0020 8d 6c 24 78 48 8b 84 24 88 00 00 00 48 89 44 24  .l$xH..$....H.D$
	0x0030 30 0f 57 c0 0f 11 44 24 50 48 8d 44 24 50 48 89  0.W...D$PH.D$PH.
	0x0040 44 24 38 48 8d 05 00 00 00 00 48 89 04 24 48 8b  D$8H......H..$H.
	0x0050 44 24 30 48 89 44 24 08 e8 00 00 00 00 48 8b 44  D$0H.D$......H.D
	0x0060 24 10 48 8b 4c 24 18 48 89 44 24 40 48 89 4c 24  $.H.L$.H.D$@H.L$
	0x0070 48 48 8b 54 24 38 84 02 48 89 02 48 8d 7a 08 83  HH.T$8..H..H.z..
	0x0080 3d 00 00 00 00 00 74 02 eb 4e 48 89 4a 08 eb 00  =.....t..NH.J...
	0x0090 48 8b 44 24 38 84 00 eb 00 48 89 44 24 60 48 c7  H.D$8....H.D$`H.
	0x00a0 44 24 68 01 00 00 00 48 c7 44 24 70 01 00 00 00  D$h....H.D$p....
	0x00b0 48 89 04 24 48 c7 44 24 08 01 00 00 00 48 c7 44  H..$H.D$.....H.D
	0x00c0 24 10 01 00 00 00 e8 00 00 00 00 48 8b 6c 24 78  $..........H.l$x
	0x00d0 48 81 c4 80 00 00 00 c3 48 89 c8 e8 00 00 00 00  H.......H.......
	0x00e0 eb ae e8 00 00 00 00 e9 14 ff ff ff              ............
	rel 5+4 t=16 TLS+0
	rel 70+4 t=15 type.int+0
	rel 89+4 t=8 runtime.convT2E64+0
	rel 129+4 t=15 runtime.writeBarrier+-1
	rel 199+4 t=8 fmt.Println+0
	rel 220+4 t=8 runtime.gcWriteBarrier+0
	rel 227+4 t=8 runtime.morestack_noctxt+0
"".time5_b STEXT size=190 args=0x10 locals=0x70
	0x0000 00000 (target_func.go:33)	TEXT	"".time5_b(SB), $112-16
	0x0000 00000 (target_func.go:33)	MOVQ	(TLS), CX
	0x0009 00009 (target_func.go:33)	CMPQ	SP, 16(CX)
	0x000d 00013 (target_func.go:33)	JLS	180
	0x0013 00019 (target_func.go:33)	SUBQ	$112, SP
	0x0017 00023 (target_func.go:33)	MOVQ	BP, 104(SP)
	0x001c 00028 (target_func.go:33)	LEAQ	104(SP), BP
	0x0021 00033 (target_func.go:33)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (target_func.go:33)	FUNCDATA	$1, gclocals·86da00badb1a277fd4ad05aca8027ea8(SB)
	0x0021 00033 (target_func.go:33)	FUNCDATA	$3, gclocals·bfec7e55b3f043d1941c093912808913(SB)
	0x0021 00033 (target_func.go:33)	PCDATA	$2, $0
	0x0021 00033 (target_func.go:33)	PCDATA	$0, $0
	0x0021 00033 (target_func.go:33)	MOVQ	$0, "".~r1+128(SP)
	0x002d 00045 (target_func.go:34)	MOVQ	"".a+120(SP), AX
	0x0032 00050 (target_func.go:34)	IMUL3Q	$22, AX, AX
	0x0036 00054 (target_func.go:34)	MOVQ	AX, "".result+48(SP)
	0x003b 00059 (target_func.go:35)	PCDATA	$0, $1
	0x003b 00059 (target_func.go:35)	XORPS	X0, X0
	0x003e 00062 (target_func.go:35)	MOVUPS	X0, ""..autotmp_3+64(SP)
	0x0043 00067 (target_func.go:35)	PCDATA	$2, $1
	0x0043 00067 (target_func.go:35)	LEAQ	""..autotmp_3+64(SP), AX
	0x0048 00072 (target_func.go:35)	MOVQ	AX, ""..autotmp_5+56(SP)
	0x004d 00077 (target_func.go:35)	TESTB	AL, (AX)
	0x004f 00079 (target_func.go:35)	PCDATA	$2, $2
	0x004f 00079 (target_func.go:35)	LEAQ	type.string(SB), CX
	0x0056 00086 (target_func.go:35)	PCDATA	$2, $1
	0x0056 00086 (target_func.go:35)	MOVQ	CX, ""..autotmp_3+64(SP)
	0x005b 00091 (target_func.go:35)	PCDATA	$2, $2
	0x005b 00091 (target_func.go:35)	LEAQ	"".statictmp_1(SB), CX
	0x0062 00098 (target_func.go:35)	PCDATA	$2, $1
	0x0062 00098 (target_func.go:35)	MOVQ	CX, ""..autotmp_3+72(SP)
	0x0067 00103 (target_func.go:35)	TESTB	AL, (AX)
	0x0069 00105 (target_func.go:35)	JMP	107
	0x006b 00107 (target_func.go:35)	MOVQ	AX, ""..autotmp_4+80(SP)
	0x0070 00112 (target_func.go:35)	MOVQ	$1, ""..autotmp_4+88(SP)
	0x0079 00121 (target_func.go:35)	MOVQ	$1, ""..autotmp_4+96(SP)
	0x0082 00130 (target_func.go:35)	PCDATA	$2, $0
	0x0082 00130 (target_func.go:35)	MOVQ	AX, (SP)
	0x0086 00134 (target_func.go:35)	MOVQ	$1, 8(SP)
	0x008f 00143 (target_func.go:35)	MOVQ	$1, 16(SP)
	0x0098 00152 (target_func.go:35)	CALL	fmt.Println(SB)
	0x009d 00157 (target_func.go:36)	PCDATA	$0, $0
	0x009d 00157 (target_func.go:36)	MOVQ	"".result+48(SP), AX
	0x00a2 00162 (target_func.go:36)	MOVQ	AX, "".~r1+128(SP)
	0x00aa 00170 (target_func.go:36)	MOVQ	104(SP), BP
	0x00af 00175 (target_func.go:36)	ADDQ	$112, SP
	0x00b3 00179 (target_func.go:36)	RET
	0x00b4 00180 (target_func.go:36)	NOP
	0x00b4 00180 (target_func.go:33)	PCDATA	$0, $-1
	0x00b4 00180 (target_func.go:33)	PCDATA	$2, $-1
	0x00b4 00180 (target_func.go:33)	CALL	runtime.morestack_noctxt(SB)
	0x00b9 00185 (target_func.go:33)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 a1  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 70 48 89 6c 24 68 48 8d 6c 24  ...H..pH.l$hH.l$
	0x0020 68 48 c7 84 24 80 00 00 00 00 00 00 00 48 8b 44  hH..$........H.D
	0x0030 24 78 48 6b c0 16 48 89 44 24 30 0f 57 c0 0f 11  $xHk..H.D$0.W...
	0x0040 44 24 40 48 8d 44 24 40 48 89 44 24 38 84 00 48  D$@H.D$@H.D$8..H
	0x0050 8d 0d 00 00 00 00 48 89 4c 24 40 48 8d 0d 00 00  ......H.L$@H....
	0x0060 00 00 48 89 4c 24 48 84 00 eb 00 48 89 44 24 50  ..H.L$H....H.D$P
	0x0070 48 c7 44 24 58 01 00 00 00 48 c7 44 24 60 01 00  H.D$X....H.D$`..
	0x0080 00 00 48 89 04 24 48 c7 44 24 08 01 00 00 00 48  ..H..$H.D$.....H
	0x0090 c7 44 24 10 01 00 00 00 e8 00 00 00 00 48 8b 44  .D$..........H.D
	0x00a0 24 30 48 89 84 24 80 00 00 00 48 8b 6c 24 68 48  $0H..$....H.l$hH
	0x00b0 83 c4 70 c3 e8 00 00 00 00 e9 42 ff ff ff        ..p.......B...
	rel 5+4 t=16 TLS+0
	rel 82+4 t=15 type.string+0
	rel 94+4 t=15 "".statictmp_1+0
	rel 153+4 t=8 fmt.Println+0
	rel 181+4 t=8 runtime.morestack_noctxt+0
"".time6_b STEXT size=219 args=0x10 locals=0x78
	0x0000 00000 (target_func.go:39)	TEXT	"".time6_b(SB), $120-16
	0x0000 00000 (target_func.go:39)	MOVQ	(TLS), CX
	0x0009 00009 (target_func.go:39)	CMPQ	SP, 16(CX)
	0x000d 00013 (target_func.go:39)	JLS	209
	0x0013 00019 (target_func.go:39)	SUBQ	$120, SP
	0x0017 00023 (target_func.go:39)	MOVQ	BP, 112(SP)
	0x001c 00028 (target_func.go:39)	LEAQ	112(SP), BP
	0x0021 00033 (target_func.go:39)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (target_func.go:39)	FUNCDATA	$1, gclocals·86da00badb1a277fd4ad05aca8027ea8(SB)
	0x0021 00033 (target_func.go:39)	FUNCDATA	$3, gclocals·bfec7e55b3f043d1941c093912808913(SB)
	0x0021 00033 (target_func.go:39)	PCDATA	$2, $0
	0x0021 00033 (target_func.go:39)	PCDATA	$0, $0
	0x0021 00033 (target_func.go:39)	MOVQ	$0, "".~r1+136(SP)
	0x002d 00045 (target_func.go:40)	MOVQ	"".a+128(SP), AX
	0x0035 00053 (target_func.go:40)	LEAQ	(AX)(AX*2), AX
	0x0039 00057 (target_func.go:40)	MOVQ	AX, "".result+48(SP)
	0x003e 00062 (target_func.go:41)	PCDATA	$0, $1
	0x003e 00062 (target_func.go:41)	XORPS	X0, X0
	0x0041 00065 (target_func.go:41)	MOVUPS	X0, ""..autotmp_3+72(SP)
	0x0046 00070 (target_func.go:41)	PCDATA	$2, $1
	0x0046 00070 (target_func.go:41)	LEAQ	""..autotmp_3+72(SP), AX
	0x004b 00075 (target_func.go:41)	MOVQ	AX, ""..autotmp_6+64(SP)
	0x0050 00080 (target_func.go:41)	TESTB	AL, (AX)
	0x0052 00082 (target_func.go:41)	PCDATA	$2, $2
	0x0052 00082 (target_func.go:41)	LEAQ	type.string(SB), CX
	0x0059 00089 (target_func.go:41)	PCDATA	$2, $1
	0x0059 00089 (target_func.go:41)	MOVQ	CX, ""..autotmp_3+72(SP)
	0x005e 00094 (target_func.go:41)	PCDATA	$2, $2
	0x005e 00094 (target_func.go:41)	LEAQ	"".statictmp_2(SB), CX
	0x0065 00101 (target_func.go:41)	PCDATA	$2, $1
	0x0065 00101 (target_func.go:41)	MOVQ	CX, ""..autotmp_3+80(SP)
	0x006a 00106 (target_func.go:41)	TESTB	AL, (AX)
	0x006c 00108 (target_func.go:41)	JMP	110
	0x006e 00110 (target_func.go:41)	MOVQ	AX, ""..autotmp_5+88(SP)
	0x0073 00115 (target_func.go:41)	MOVQ	$1, ""..autotmp_5+96(SP)
	0x007c 00124 (target_func.go:41)	MOVQ	$1, ""..autotmp_5+104(SP)
	0x0085 00133 (target_func.go:41)	PCDATA	$2, $0
	0x0085 00133 (target_func.go:41)	MOVQ	AX, (SP)
	0x0089 00137 (target_func.go:41)	MOVQ	$1, 8(SP)
	0x0092 00146 (target_func.go:41)	MOVQ	$1, 16(SP)
	0x009b 00155 (target_func.go:41)	CALL	fmt.Println(SB)
	0x00a0 00160 (target_func.go:42)	PCDATA	$0, $0
	0x00a0 00160 (target_func.go:42)	MOVQ	$2, (SP)
	0x00a8 00168 (target_func.go:42)	CALL	"".time7_b(SB)
	0x00ad 00173 (target_func.go:42)	MOVQ	8(SP), AX
	0x00b2 00178 (target_func.go:42)	MOVQ	AX, ""..autotmp_4+56(SP)
	0x00b7 00183 (target_func.go:42)	MOVQ	"".result+48(SP), CX
	0x00bc 00188 (target_func.go:42)	ADDQ	CX, AX
	0x00bf 00191 (target_func.go:42)	MOVQ	AX, "".~r1+136(SP)
	0x00c7 00199 (target_func.go:42)	MOVQ	112(SP), BP
	0x00cc 00204 (target_func.go:42)	ADDQ	$120, SP
	0x00d0 00208 (target_func.go:42)	RET
	0x00d1 00209 (target_func.go:42)	NOP
	0x00d1 00209 (target_func.go:39)	PCDATA	$0, $-1
	0x00d1 00209 (target_func.go:39)	PCDATA	$2, $-1
	0x00d1 00209 (target_func.go:39)	CALL	runtime.morestack_noctxt(SB)
	0x00d6 00214 (target_func.go:39)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 be  eH..%....H;a....
	0x0010 00 00 00 48 83 ec 78 48 89 6c 24 70 48 8d 6c 24  ...H..xH.l$pH.l$
	0x0020 70 48 c7 84 24 88 00 00 00 00 00 00 00 48 8b 84  pH..$........H..
	0x0030 24 80 00 00 00 48 8d 04 40 48 89 44 24 30 0f 57  $....H..@H.D$0.W
	0x0040 c0 0f 11 44 24 48 48 8d 44 24 48 48 89 44 24 40  ...D$HH.D$HH.D$@
	0x0050 84 00 48 8d 0d 00 00 00 00 48 89 4c 24 48 48 8d  ..H......H.L$HH.
	0x0060 0d 00 00 00 00 48 89 4c 24 50 84 00 eb 00 48 89  .....H.L$P....H.
	0x0070 44 24 58 48 c7 44 24 60 01 00 00 00 48 c7 44 24  D$XH.D$`....H.D$
	0x0080 68 01 00 00 00 48 89 04 24 48 c7 44 24 08 01 00  h....H..$H.D$...
	0x0090 00 00 48 c7 44 24 10 01 00 00 00 e8 00 00 00 00  ..H.D$..........
	0x00a0 48 c7 04 24 02 00 00 00 e8 00 00 00 00 48 8b 44  H..$.........H.D
	0x00b0 24 08 48 89 44 24 38 48 8b 4c 24 30 48 01 c8 48  $.H.D$8H.L$0H..H
	0x00c0 89 84 24 88 00 00 00 48 8b 6c 24 70 48 83 c4 78  ..$....H.l$pH..x
	0x00d0 c3 e8 00 00 00 00 e9 25 ff ff ff                 .......%...
	rel 5+4 t=16 TLS+0
	rel 85+4 t=15 type.string+0
	rel 97+4 t=15 "".statictmp_2+0
	rel 156+4 t=8 fmt.Println+0
	rel 169+4 t=8 "".time7_b+0
	rel 210+4 t=8 runtime.morestack_noctxt+0
"".time7_b STEXT size=248 args=0x10 locals=0x80
	0x0000 00000 (target_func.go:45)	TEXT	"".time7_b(SB), $128-16
	0x0000 00000 (target_func.go:45)	MOVQ	(TLS), CX
	0x0009 00009 (target_func.go:45)	CMPQ	SP, 16(CX)
	0x000d 00013 (target_func.go:45)	JLS	238
	0x0013 00019 (target_func.go:45)	SUBQ	$128, SP
	0x001a 00026 (target_func.go:45)	MOVQ	BP, 120(SP)
	0x001f 00031 (target_func.go:45)	LEAQ	120(SP), BP
	0x0024 00036 (target_func.go:45)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0024 00036 (target_func.go:45)	FUNCDATA	$1, gclocals·86da00badb1a277fd4ad05aca8027ea8(SB)
	0x0024 00036 (target_func.go:45)	FUNCDATA	$3, gclocals·bfec7e55b3f043d1941c093912808913(SB)
	0x0024 00036 (target_func.go:45)	PCDATA	$2, $0
	0x0024 00036 (target_func.go:45)	PCDATA	$0, $0
	0x0024 00036 (target_func.go:45)	MOVQ	$0, "".~r1+144(SP)
	0x0030 00048 (target_func.go:46)	MOVQ	"".b+136(SP), AX
	0x0038 00056 (target_func.go:46)	IMUL3Q	$101, AX, AX
	0x003c 00060 (target_func.go:46)	MOVQ	AX, "".result+64(SP)
	0x0041 00065 (target_func.go:47)	PCDATA	$2, $1
	0x0041 00065 (target_func.go:47)	LEAQ	go.string."tcp"(SB), AX
	0x0048 00072 (target_func.go:47)	PCDATA	$2, $0
	0x0048 00072 (target_func.go:47)	MOVQ	AX, (SP)
	0x004c 00076 (target_func.go:47)	MOVQ	$3, 8(SP)
	0x0055 00085 (target_func.go:47)	PCDATA	$2, $1
	0x0055 00085 (target_func.go:47)	LEAQ	go.string."0.0.0.0:8090"(SB), AX
	0x005c 00092 (target_func.go:47)	PCDATA	$2, $0
	0x005c 00092 (target_func.go:47)	MOVQ	AX, 16(SP)
	0x0061 00097 (target_func.go:47)	MOVQ	$12, 24(SP)
	0x006a 00106 (target_func.go:47)	CALL	net.Listen(SB)
	0x006f 00111 (target_func.go:48)	PCDATA	$0, $1
	0x006f 00111 (target_func.go:48)	XORPS	X0, X0
	0x0072 00114 (target_func.go:48)	MOVUPS	X0, ""..autotmp_3+80(SP)
	0x0077 00119 (target_func.go:48)	PCDATA	$2, $1
	0x0077 00119 (target_func.go:48)	LEAQ	""..autotmp_3+80(SP), AX
	0x007c 00124 (target_func.go:48)	MOVQ	AX, ""..autotmp_5+72(SP)
	0x0081 00129 (target_func.go:48)	TESTB	AL, (AX)
	0x0083 00131 (target_func.go:48)	PCDATA	$2, $2
	0x0083 00131 (target_func.go:48)	LEAQ	type.string(SB), CX
	0x008a 00138 (target_func.go:48)	PCDATA	$2, $1
	0x008a 00138 (target_func.go:48)	MOVQ	CX, ""..autotmp_3+80(SP)
	0x008f 00143 (target_func.go:48)	PCDATA	$2, $2
	0x008f 00143 (target_func.go:48)	LEAQ	"".statictmp_3(SB), CX
	0x0096 00150 (target_func.go:48)	PCDATA	$2, $1
	0x0096 00150 (target_func.go:48)	MOVQ	CX, ""..autotmp_3+88(SP)
	0x009b 00155 (target_func.go:48)	TESTB	AL, (AX)
	0x009d 00157 (target_func.go:48)	JMP	159
	0x009f 00159 (target_func.go:48)	MOVQ	AX, ""..autotmp_4+96(SP)
	0x00a4 00164 (target_func.go:48)	MOVQ	$1, ""..autotmp_4+104(SP)
	0x00ad 00173 (target_func.go:48)	MOVQ	$1, ""..autotmp_4+112(SP)
	0x00b6 00182 (target_func.go:48)	PCDATA	$2, $0
	0x00b6 00182 (target_func.go:48)	MOVQ	AX, (SP)
	0x00ba 00186 (target_func.go:48)	MOVQ	$1, 8(SP)
	0x00c3 00195 (target_func.go:48)	MOVQ	$1, 16(SP)
	0x00cc 00204 (target_func.go:48)	CALL	fmt.Println(SB)
	0x00d1 00209 (target_func.go:49)	PCDATA	$0, $0
	0x00d1 00209 (target_func.go:49)	MOVQ	"".result+64(SP), AX
	0x00d6 00214 (target_func.go:49)	INCQ	AX
	0x00d9 00217 (target_func.go:49)	MOVQ	AX, "".~r1+144(SP)
	0x00e1 00225 (target_func.go:49)	MOVQ	120(SP), BP
	0x00e6 00230 (target_func.go:49)	ADDQ	$128, SP
	0x00ed 00237 (target_func.go:49)	RET
	0x00ee 00238 (target_func.go:49)	NOP
	0x00ee 00238 (target_func.go:45)	PCDATA	$0, $-1
	0x00ee 00238 (target_func.go:45)	PCDATA	$2, $-1
	0x00ee 00238 (target_func.go:45)	CALL	runtime.morestack_noctxt(SB)
	0x00f3 00243 (target_func.go:45)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 db  eH..%....H;a....
	0x0010 00 00 00 48 81 ec 80 00 00 00 48 89 6c 24 78 48  ...H......H.l$xH
	0x0020 8d 6c 24 78 48 c7 84 24 90 00 00 00 00 00 00 00  .l$xH..$........
	0x0030 48 8b 84 24 88 00 00 00 48 6b c0 65 48 89 44 24  H..$....Hk.eH.D$
	0x0040 40 48 8d 05 00 00 00 00 48 89 04 24 48 c7 44 24  @H......H..$H.D$
	0x0050 08 03 00 00 00 48 8d 05 00 00 00 00 48 89 44 24  .....H......H.D$
	0x0060 10 48 c7 44 24 18 0c 00 00 00 e8 00 00 00 00 0f  .H.D$...........
	0x0070 57 c0 0f 11 44 24 50 48 8d 44 24 50 48 89 44 24  W...D$PH.D$PH.D$
	0x0080 48 84 00 48 8d 0d 00 00 00 00 48 89 4c 24 50 48  H..H......H.L$PH
	0x0090 8d 0d 00 00 00 00 48 89 4c 24 58 84 00 eb 00 48  ......H.L$X....H
	0x00a0 89 44 24 60 48 c7 44 24 68 01 00 00 00 48 c7 44  .D$`H.D$h....H.D
	0x00b0 24 70 01 00 00 00 48 89 04 24 48 c7 44 24 08 01  $p....H..$H.D$..
	0x00c0 00 00 00 48 c7 44 24 10 01 00 00 00 e8 00 00 00  ...H.D$.........
	0x00d0 00 48 8b 44 24 40 48 ff c0 48 89 84 24 90 00 00  .H.D$@H..H..$...
	0x00e0 00 48 8b 6c 24 78 48 81 c4 80 00 00 00 c3 e8 00  .H.l$xH.........
	0x00f0 00 00 00 e9 08 ff ff ff                          ........
	rel 5+4 t=16 TLS+0
	rel 68+4 t=15 go.string."tcp"+0
	rel 88+4 t=15 go.string."0.0.0.0:8090"+0
	rel 107+4 t=8 net.Listen+0
	rel 134+4 t=15 type.string+0
	rel 146+4 t=15 "".statictmp_3+0
	rel 205+4 t=8 fmt.Println+0
	rel 239+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=105 args=0x0 locals=0x8
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $8-0
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	JLS	98
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
	0x004d 00077 (<autogenerated>:1)	CALL	net.init(SB)
	0x0052 00082 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x0059 00089 (<autogenerated>:1)	MOVQ	(SP), BP
	0x005d 00093 (<autogenerated>:1)	ADDQ	$8, SP
	0x0061 00097 (<autogenerated>:1)	RET
	0x0062 00098 (<autogenerated>:1)	NOP
	0x0062 00098 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0062 00098 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0062 00098 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0067 00103 (<autogenerated>:1)	JMP	0
	0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 53 48  eH..%....H;a.vSH
	0x0010 83 ec 08 48 89 2c 24 48 8d 2c 24 80 3d 00 00 00  ...H.,$H.,$.=...
	0x0020 00 01 77 02 eb 09 48 8b 2c 24 48 83 c4 08 c3 80  ..w...H.,$H.....
	0x0030 3d 00 00 00 00 01 74 02 eb 07 e8 00 00 00 00 0f  =.....t.........
	0x0040 0b c6 05 00 00 00 00 01 e8 00 00 00 00 e8 00 00  ................
	0x0050 00 00 c6 05 00 00 00 00 02 48 8b 2c 24 48 83 c4  .........H.,$H..
	0x0060 08 c3 e8 00 00 00 00 eb 97                       .........
	rel 5+4 t=16 TLS+0
	rel 29+4 t=15 "".initdone·+-1
	rel 49+4 t=15 "".initdone·+-1
	rel 59+4 t=8 runtime.throwinit+0
	rel 67+4 t=15 "".initdone·+-1
	rel 73+4 t=8 fmt.init+0
	rel 78+4 t=8 net.init+0
	rel 84+4 t=15 "".initdone·+-1
	rel 99+4 t=8 runtime.morestack_noctxt+0
go.string."yes" SRODATA dupok size=3
	0x0000 79 65 73                                         yes
go.loc."".time4_b SDWARFLOC size=0
go.info."".time4_b SDWARFINFO size=78
	0x0000 02 22 22 2e 74 69 6d 65 34 5f 62 00 00 00 00 00  ."".time4_b.....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00  ................
	0x0020 00 00 01 0e 61 00 00 0a 00 00 00 00 01 9c 09 72  ....a..........r
	0x0030 65 73 75 6c 74 00 0b 00 00 00 00 03 91 f0 7e 0e  esult.........~.
	0x0040 7e 72 31 00 01 0a 00 00 00 00 02 91 08 00        ~r1...........
	rel 12+8 t=1 "".time4_b+0
	rel 20+8 t=1 "".time4_b+556
	rel 30+4 t=29 gofile../Users/yongfuchen/go1/src/go_test/asm_go/target_func.go+0
	rel 40+4 t=28 go.info.int+0
	rel 55+4 t=28 go.info.int+0
	rel 70+4 t=28 go.info.int+0
go.range."".time4_b SDWARFRANGE size=0
go.isstmt."".time4_b SDWARFMISC size=0
	0x0000 04 18 04 17 03 14 01 0d 02 04 01 04 02 0c 01 10  ................
	0x0010 02 04 01 04 02 0c 01 10 02 06 01 04 02 0c 01 10  ................
	0x0020 02 0e 01 68 02 03 01 a5 01 02 08 01 3e 02 0a 00  ...h........>...
go.loc."".trace_b SDWARFLOC size=0
go.info."".trace_b SDWARFINFO size=47
	0x0000 02 22 22 2e 74 72 61 63 65 5f 62 00 00 00 00 00  ."".trace_b.....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00  ................
	0x0020 00 00 01 0e 69 00 00 1d 00 00 00 00 01 9c 00     ....i..........
	rel 12+8 t=1 "".trace_b+0
	rel 20+8 t=1 "".trace_b+236
	rel 30+4 t=29 gofile../Users/yongfuchen/go1/src/go_test/asm_go/target_func.go+0
	rel 40+4 t=28 go.info.int+0
go.range."".trace_b SDWARFRANGE size=0
go.isstmt."".trace_b SDWARFMISC size=0
	0x0000 04 13 04 11 03 08 01 9f 01 02 0d 01 0a 02 0a 00  ................
go.string."origin time5 called" SRODATA dupok size=19
	0x0000 6f 72 69 67 69 6e 20 74 69 6d 65 35 20 63 61 6c  origin time5 cal
	0x0010 6c 65 64                                         led
go.loc."".time5_b SDWARFLOC size=0
go.info."".time5_b SDWARFINFO size=78
	0x0000 02 22 22 2e 74 69 6d 65 35 5f 62 00 00 00 00 00  ."".time5_b.....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00  ................
	0x0020 00 00 01 0e 61 00 00 21 00 00 00 00 01 9c 09 72  ....a..!.......r
	0x0030 65 73 75 6c 74 00 22 00 00 00 00 03 91 b8 7f 0e  esult.".........
	0x0040 7e 72 31 00 01 21 00 00 00 00 02 91 08 00        ~r1..!........
	rel 12+8 t=1 "".time5_b+0
	rel 20+8 t=1 "".time5_b+190
	rel 30+4 t=29 gofile../Users/yongfuchen/go1/src/go_test/asm_go/target_func.go+0
	rel 40+4 t=28 go.info.int+0
	rel 55+4 t=28 go.info.int+0
	rel 70+4 t=28 go.info.int+0
go.range."".time5_b SDWARFRANGE size=0
go.isstmt."".time5_b SDWARFMISC size=0
	0x0000 04 13 04 0e 03 11 01 09 02 03 01 5f 02 05 01 12  ..........._....
	0x0010 02 0a 00                                         ...
go.string."time6 called" SRODATA dupok size=12
	0x0000 74 69 6d 65 36 20 63 61 6c 6c 65 64              time6 called
go.loc."".time6_b SDWARFLOC size=0
go.info."".time6_b SDWARFINFO size=78
	0x0000 02 22 22 2e 74 69 6d 65 36 5f 62 00 00 00 00 00  ."".time6_b.....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00  ................
	0x0020 00 00 01 0e 61 00 00 27 00 00 00 00 01 9c 09 72  ....a..'.......r
	0x0030 65 73 75 6c 74 00 28 00 00 00 00 03 91 b0 7f 0e  esult.(.........
	0x0040 7e 72 31 00 01 27 00 00 00 00 02 91 08 00        ~r1..'........
	rel 12+8 t=1 "".time6_b+0
	rel 20+8 t=1 "".time6_b+219
	rel 30+4 t=29 gofile../Users/yongfuchen/go1/src/go_test/asm_go/target_func.go+0
	rel 40+4 t=28 go.info.int+0
	rel 55+4 t=28 go.info.int+0
	rel 70+4 t=28 go.info.int+0
go.range."".time6_b SDWARFRANGE size=0
go.isstmt."".time6_b SDWARFMISC size=0
	0x0000 04 13 04 0e 03 14 01 09 02 03 01 5f 02 08 01 29  ..........._...)
	0x0010 02 0a 00                                         ...
go.string."time7 called" SRODATA dupok size=12
	0x0000 74 69 6d 65 37 20 63 61 6c 6c 65 64              time7 called
go.string."tcp" SRODATA dupok size=3
	0x0000 74 63 70                                         tcp
go.string."0.0.0.0:8090" SRODATA dupok size=12
	0x0000 30 2e 30 2e 30 2e 30 3a 38 30 39 30              0.0.0.0:8090
go.loc."".time7_b SDWARFLOC size=0
go.info."".time7_b SDWARFINFO size=78
	0x0000 02 22 22 2e 74 69 6d 65 37 5f 62 00 00 00 00 00  ."".time7_b.....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00  ................
	0x0020 00 00 01 0e 62 00 00 2d 00 00 00 00 01 9c 09 72  ....b..-.......r
	0x0030 65 73 75 6c 74 00 2e 00 00 00 00 03 91 b8 7f 0e  esult...........
	0x0040 7e 72 31 00 01 2d 00 00 00 00 02 91 08 00        ~r1..-........
	rel 12+8 t=1 "".time7_b+0
	rel 20+8 t=1 "".time7_b+248
	rel 30+4 t=29 gofile../Users/yongfuchen/go1/src/go_test/asm_go/target_func.go+0
	rel 40+4 t=28 go.info.int+0
	rel 55+4 t=28 go.info.int+0
	rel 70+4 t=28 go.info.int+0
go.range."".time7_b SDWARFRANGE size=0
go.isstmt."".time7_b SDWARFMISC size=0
	0x0000 04 13 04 11 03 14 01 09 02 07 01 27 02 03 01 5f  ...........'..._
	0x0010 02 05 01 18 02 0a 00                             .......
go.loc."".init SDWARFLOC size=0
go.info."".init SDWARFINFO size=33
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01  ................
	0x0020 00                                               .
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+105
	rel 27+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
go.isstmt."".init SDWARFMISC size=0
	0x0000 04 0f 04 0c 03 07 01 04 02 09 01 10 02 09 01 1a  ................
	0x0010 02 07 00                                         ...
"".C_s SNOPTRDATA size=8
	0x0000 09 00 00 00 00 00 00 00                          ........
"".statictmp_0 SRODATA size=16
	0x0000 00 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."yes"+0
"".statictmp_1 SRODATA size=16
	0x0000 00 00 00 00 00 00 00 00 13 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."origin time5 called"+0
"".statictmp_2 SRODATA size=16
	0x0000 00 00 00 00 00 00 00 00 0c 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."time6 called"+0
"".statictmp_3 SRODATA size=16
	0x0000 00 00 00 00 00 00 00 00 0c 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."time7 called"+0
"".initdone· SNOPTRBSS size=1
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
type..importpath.fmt. SRODATA dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
type..importpath.net. SRODATA dupok size=6
	0x0000 00 00 03 6e 65 74                                ...net
gclocals·f6bd6b3389b872033d462029172c8612 SRODATA dupok size=8
	0x0000 04 00 00 00 00 00 00 00                          ........
gclocals·ed20391e33635a5bf94ce2d547e0f407 SRODATA dupok size=16
	0x0000 04 00 00 00 0e 00 00 00 00 00 20 00 08 00 09 00  .......... .....
gclocals·fedbdc973195b17703414cfc4d862ea8 SRODATA dupok size=13
	0x0000 05 00 00 00 07 00 00 00 00 01 03 05 45           ............E
gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
	0x0000 03 00 00 00 00 00 00 00                          ........
gclocals·1d53b44fb3211f44318e421290131d18 SRODATA dupok size=11
	0x0000 03 00 00 00 08 00 00 00 00 10 11                 ...........
gclocals·3ee7973a6249c7a7f2fab45f78192248 SRODATA dupok size=13
	0x0000 05 00 00 00 07 00 00 00 00 01 02 06 46           ............F
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·86da00badb1a277fd4ad05aca8027ea8 SRODATA dupok size=10
	0x0000 02 00 00 00 06 00 00 00 00 04                    ..........
gclocals·bfec7e55b3f043d1941c093912808913 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 01 03                 ...........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
