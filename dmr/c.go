/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dmr

import (
	"io"
	"math"
)

type Dm_build_0 struct {
	dm_build_1 []byte
	dm_build_2 int
}

func Dm_build_3(dm_build_4 int) *Dm_build_0 {
	return &Dm_build_0{make([]byte, 0, dm_build_4), 0}
}

func Dm_build_5(dm_build_6 []byte) *Dm_build_0 {
	return &Dm_build_0{dm_build_6, 0}
}

func (dm_build_8 *Dm_build_0) dm_build_7(dm_build_9 int) *Dm_build_0 {

	dm_build_10 := len(dm_build_8.dm_build_1)
	dm_build_11 := cap(dm_build_8.dm_build_1)

	if dm_build_10+dm_build_9 <= dm_build_11 {
		dm_build_8.dm_build_1 = dm_build_8.dm_build_1[:dm_build_10+dm_build_9]
	} else {
		remain := dm_build_9 + dm_build_10 - dm_build_11
		nbuf := make([]byte, dm_build_9+dm_build_10, 2*dm_build_11+remain)
		copy(nbuf, dm_build_8.dm_build_1)
		dm_build_8.dm_build_1 = nbuf
	}

	return dm_build_8
}

func (dm_build_13 *Dm_build_0) Dm_build_12() int {
	return len(dm_build_13.dm_build_1)
}

func (dm_build_15 *Dm_build_0) Dm_build_14(dm_build_16 int) *Dm_build_0 {
	for i := dm_build_16; i < len(dm_build_15.dm_build_1); i++ {
		dm_build_15.dm_build_1[i] = 0
	}
	dm_build_15.dm_build_1 = dm_build_15.dm_build_1[:dm_build_16]
	return dm_build_15
}

func (dm_build_18 *Dm_build_0) Dm_build_17(dm_build_19 int) *Dm_build_0 {
	dm_build_18.dm_build_2 = dm_build_19
	return dm_build_18
}

func (dm_build_21 *Dm_build_0) Dm_build_20() int {
	return dm_build_21.dm_build_2
}

func (dm_build_23 *Dm_build_0) Dm_build_22(dm_build_24 bool) int {
	return len(dm_build_23.dm_build_1) - dm_build_23.dm_build_2
}

func (dm_build_26 *Dm_build_0) Dm_build_25(dm_build_27 int, dm_build_28 bool, dm_build_29 bool) *Dm_build_0 {

	if dm_build_28 {
		if dm_build_29 {
			dm_build_26.dm_build_7(dm_build_27)
		} else {
			dm_build_26.dm_build_1 = dm_build_26.dm_build_1[:len(dm_build_26.dm_build_1)-dm_build_27]
		}
	} else {
		if dm_build_29 {
			dm_build_26.dm_build_2 += dm_build_27
		} else {
			dm_build_26.dm_build_2 -= dm_build_27
		}
	}

	return dm_build_26
}

func (dm_build_31 *Dm_build_0) Dm_build_30(dm_build_32 io.Reader, dm_build_33 int) int {
	dm_build_34 := len(dm_build_31.dm_build_1)
	dm_build_31.dm_build_7(dm_build_33)
	dm_build_35 := 0
	for dm_build_33 > 0 {
		n, err := dm_build_32.Read(dm_build_31.dm_build_1[dm_build_34+dm_build_35:])
		if n > 0 && err == io.EOF {
			dm_build_35 += n
			dm_build_31.dm_build_1 = dm_build_31.dm_build_1[:dm_build_34+dm_build_35]
			return dm_build_35
		} else if n > 0 && err == nil {
			dm_build_33 -= n
			dm_build_35 += n
		} else if n == 0 && err != nil {
			panic("load err")
		}
	}

	return dm_build_35
}

func (dm_build_37 *Dm_build_0) Dm_build_36(dm_build_38 io.Writer) *Dm_build_0 {
	dm_build_38.Write(dm_build_37.dm_build_1)
	return dm_build_37
}

func (dm_build_40 *Dm_build_0) Dm_build_39(dm_build_41 bool) int {
	dm_build_42 := len(dm_build_40.dm_build_1)
	dm_build_40.dm_build_7(1)

	if dm_build_41 {
		return copy(dm_build_40.dm_build_1[dm_build_42:], []byte{1})
	} else {
		return copy(dm_build_40.dm_build_1[dm_build_42:], []byte{0})
	}
}

func (dm_build_44 *Dm_build_0) Dm_build_43(dm_build_45 byte) int {
	dm_build_46 := len(dm_build_44.dm_build_1)
	dm_build_44.dm_build_7(1)

	return copy(dm_build_44.dm_build_1[dm_build_46:], Dm_build_1220.Dm_build_1398(dm_build_45))
}

func (dm_build_48 *Dm_build_0) Dm_build_47(dm_build_49 int16) int {
	dm_build_50 := len(dm_build_48.dm_build_1)
	dm_build_48.dm_build_7(2)

	return copy(dm_build_48.dm_build_1[dm_build_50:], Dm_build_1220.Dm_build_1401(dm_build_49))
}

func (dm_build_52 *Dm_build_0) Dm_build_51(dm_build_53 int32) int {
	dm_build_54 := len(dm_build_52.dm_build_1)
	dm_build_52.dm_build_7(4)

	return copy(dm_build_52.dm_build_1[dm_build_54:], Dm_build_1220.Dm_build_1404(dm_build_53))
}

func (dm_build_56 *Dm_build_0) Dm_build_55(dm_build_57 uint8) int {
	dm_build_58 := len(dm_build_56.dm_build_1)
	dm_build_56.dm_build_7(1)

	return copy(dm_build_56.dm_build_1[dm_build_58:], Dm_build_1220.Dm_build_1416(dm_build_57))
}

func (dm_build_60 *Dm_build_0) Dm_build_59(dm_build_61 uint16) int {
	dm_build_62 := len(dm_build_60.dm_build_1)
	dm_build_60.dm_build_7(2)

	return copy(dm_build_60.dm_build_1[dm_build_62:], Dm_build_1220.Dm_build_1419(dm_build_61))
}

func (dm_build_64 *Dm_build_0) Dm_build_63(dm_build_65 uint32) int {
	dm_build_66 := len(dm_build_64.dm_build_1)
	dm_build_64.dm_build_7(4)

	return copy(dm_build_64.dm_build_1[dm_build_66:], Dm_build_1220.Dm_build_1422(dm_build_65))
}

func (dm_build_68 *Dm_build_0) Dm_build_67(dm_build_69 uint64) int {
	dm_build_70 := len(dm_build_68.dm_build_1)
	dm_build_68.dm_build_7(8)

	return copy(dm_build_68.dm_build_1[dm_build_70:], Dm_build_1220.Dm_build_1425(dm_build_69))
}

func (dm_build_72 *Dm_build_0) Dm_build_71(dm_build_73 float32) int {
	dm_build_74 := len(dm_build_72.dm_build_1)
	dm_build_72.dm_build_7(4)

	return copy(dm_build_72.dm_build_1[dm_build_74:], Dm_build_1220.Dm_build_1422(math.Float32bits(dm_build_73)))
}

func (dm_build_76 *Dm_build_0) Dm_build_75(dm_build_77 float64) int {
	dm_build_78 := len(dm_build_76.dm_build_1)
	dm_build_76.dm_build_7(8)

	return copy(dm_build_76.dm_build_1[dm_build_78:], Dm_build_1220.Dm_build_1425(math.Float64bits(dm_build_77)))
}

func (dm_build_80 *Dm_build_0) Dm_build_79(dm_build_81 []byte) int {
	dm_build_82 := len(dm_build_80.dm_build_1)
	dm_build_80.dm_build_7(len(dm_build_81))
	return copy(dm_build_80.dm_build_1[dm_build_82:], dm_build_81)
}

func (dm_build_84 *Dm_build_0) Dm_build_83(dm_build_85 []byte) int {
	return dm_build_84.Dm_build_51(int32(len(dm_build_85))) + dm_build_84.Dm_build_79(dm_build_85)
}

func (dm_build_87 *Dm_build_0) Dm_build_86(dm_build_88 []byte) int {
	return dm_build_87.Dm_build_55(uint8(len(dm_build_88))) + dm_build_87.Dm_build_79(dm_build_88)
}

func (dm_build_90 *Dm_build_0) Dm_build_89(dm_build_91 []byte) int {
	return dm_build_90.Dm_build_59(uint16(len(dm_build_91))) + dm_build_90.Dm_build_79(dm_build_91)
}

func (dm_build_93 *Dm_build_0) Dm_build_92(dm_build_94 []byte) int {
	return dm_build_93.Dm_build_79(dm_build_94) + dm_build_93.Dm_build_43(0)
}

func (dm_build_96 *Dm_build_0) Dm_build_95(dm_build_97 string, dm_build_98 string, dm_build_99 *DmConnection) int {
	dm_build_100 := Dm_build_1220.Dm_build_1433(dm_build_97, dm_build_98, dm_build_99)
	return dm_build_96.Dm_build_83(dm_build_100)
}

func (dm_build_102 *Dm_build_0) Dm_build_101(dm_build_103 string, dm_build_104 string, dm_build_105 *DmConnection) int {
	dm_build_106 := Dm_build_1220.Dm_build_1433(dm_build_103, dm_build_104, dm_build_105)
	return dm_build_102.Dm_build_86(dm_build_106)
}

func (dm_build_108 *Dm_build_0) Dm_build_107(dm_build_109 string, dm_build_110 string, dm_build_111 *DmConnection) int {
	dm_build_112 := Dm_build_1220.Dm_build_1433(dm_build_109, dm_build_110, dm_build_111)
	return dm_build_108.Dm_build_89(dm_build_112)
}

func (dm_build_114 *Dm_build_0) Dm_build_113(dm_build_115 string, dm_build_116 string, dm_build_117 *DmConnection) int {
	dm_build_118 := Dm_build_1220.Dm_build_1433(dm_build_115, dm_build_116, dm_build_117)
	return dm_build_114.Dm_build_92(dm_build_118)
}

func (dm_build_120 *Dm_build_0) Dm_build_119() byte {
	dm_build_121 := Dm_build_1220.Dm_build_1313(dm_build_120.dm_build_1, dm_build_120.dm_build_2)
	dm_build_120.dm_build_2++
	return dm_build_121
}

func (dm_build_123 *Dm_build_0) Dm_build_122() int16 {
	dm_build_124 := Dm_build_1220.Dm_build_1317(dm_build_123.dm_build_1, dm_build_123.dm_build_2)
	dm_build_123.dm_build_2 += 2
	return dm_build_124
}

func (dm_build_126 *Dm_build_0) Dm_build_125() int32 {
	dm_build_127 := Dm_build_1220.Dm_build_1322(dm_build_126.dm_build_1, dm_build_126.dm_build_2)
	dm_build_126.dm_build_2 += 4
	return dm_build_127
}

func (dm_build_129 *Dm_build_0) Dm_build_128() int64 {
	dm_build_130 := Dm_build_1220.Dm_build_1327(dm_build_129.dm_build_1, dm_build_129.dm_build_2)
	dm_build_129.dm_build_2 += 8
	return dm_build_130
}

func (dm_build_132 *Dm_build_0) Dm_build_131() float32 {
	dm_build_133 := Dm_build_1220.Dm_build_1332(dm_build_132.dm_build_1, dm_build_132.dm_build_2)
	dm_build_132.dm_build_2 += 4
	return dm_build_133
}

func (dm_build_135 *Dm_build_0) Dm_build_134() float64 {
	dm_build_136 := Dm_build_1220.Dm_build_1336(dm_build_135.dm_build_1, dm_build_135.dm_build_2)
	dm_build_135.dm_build_2 += 8
	return dm_build_136
}

func (dm_build_138 *Dm_build_0) Dm_build_137() uint8 {
	dm_build_139 := Dm_build_1220.Dm_build_1340(dm_build_138.dm_build_1, dm_build_138.dm_build_2)
	dm_build_138.dm_build_2 += 1
	return dm_build_139
}

func (dm_build_141 *Dm_build_0) Dm_build_140() uint16 {
	dm_build_142 := Dm_build_1220.Dm_build_1344(dm_build_141.dm_build_1, dm_build_141.dm_build_2)
	dm_build_141.dm_build_2 += 2
	return dm_build_142
}

func (dm_build_144 *Dm_build_0) Dm_build_143() uint32 {
	dm_build_145 := Dm_build_1220.Dm_build_1349(dm_build_144.dm_build_1, dm_build_144.dm_build_2)
	dm_build_144.dm_build_2 += 4
	return dm_build_145
}

func (dm_build_147 *Dm_build_0) Dm_build_146(dm_build_148 int) []byte {
	dm_build_149 := Dm_build_1220.Dm_build_1371(dm_build_147.dm_build_1, dm_build_147.dm_build_2, dm_build_148)
	dm_build_147.dm_build_2 += dm_build_148
	return dm_build_149
}

func (dm_build_151 *Dm_build_0) Dm_build_150() []byte {
	return dm_build_151.Dm_build_146(int(dm_build_151.Dm_build_125()))
}

func (dm_build_153 *Dm_build_0) Dm_build_152() []byte {
	return dm_build_153.Dm_build_146(int(dm_build_153.Dm_build_119()))
}

func (dm_build_155 *Dm_build_0) Dm_build_154() []byte {
	return dm_build_155.Dm_build_146(int(dm_build_155.Dm_build_122()))
}

func (dm_build_157 *Dm_build_0) Dm_build_156(dm_build_158 int) []byte {
	return dm_build_157.Dm_build_146(dm_build_158)
}

func (dm_build_160 *Dm_build_0) Dm_build_159() []byte {
	dm_build_161 := 0
	for dm_build_160.Dm_build_119() != 0 {
		dm_build_161++
	}
	dm_build_160.Dm_build_25(dm_build_161, false, false)
	return dm_build_160.Dm_build_146(dm_build_161)
}

func (dm_build_163 *Dm_build_0) Dm_build_162(dm_build_164 int, dm_build_165 string, dm_build_166 *DmConnection) string {
	return Dm_build_1220.Dm_build_1470(dm_build_163.Dm_build_146(dm_build_164), dm_build_165, dm_build_166)
}

func (dm_build_168 *Dm_build_0) Dm_build_167(dm_build_169 string, dm_build_170 *DmConnection) string {
	return Dm_build_1220.Dm_build_1470(dm_build_168.Dm_build_150(), dm_build_169, dm_build_170)
}

func (dm_build_172 *Dm_build_0) Dm_build_171(dm_build_173 string, dm_build_174 *DmConnection) string {
	return Dm_build_1220.Dm_build_1470(dm_build_172.Dm_build_152(), dm_build_173, dm_build_174)
}

func (dm_build_176 *Dm_build_0) Dm_build_175(dm_build_177 string, dm_build_178 *DmConnection) string {
	return Dm_build_1220.Dm_build_1470(dm_build_176.Dm_build_154(), dm_build_177, dm_build_178)
}

func (dm_build_180 *Dm_build_0) Dm_build_179(dm_build_181 string, dm_build_182 *DmConnection) string {
	return Dm_build_1220.Dm_build_1470(dm_build_180.Dm_build_159(), dm_build_181, dm_build_182)
}

func (dm_build_184 *Dm_build_0) Dm_build_183(dm_build_185 int, dm_build_186 byte) int {
	return dm_build_184.Dm_build_219(dm_build_185, Dm_build_1220.Dm_build_1398(dm_build_186))
}

func (dm_build_188 *Dm_build_0) Dm_build_187(dm_build_189 int, dm_build_190 int16) int {
	return dm_build_188.Dm_build_219(dm_build_189, Dm_build_1220.Dm_build_1401(dm_build_190))
}

func (dm_build_192 *Dm_build_0) Dm_build_191(dm_build_193 int, dm_build_194 int32) int {
	return dm_build_192.Dm_build_219(dm_build_193, Dm_build_1220.Dm_build_1404(dm_build_194))
}

func (dm_build_196 *Dm_build_0) Dm_build_195(dm_build_197 int, dm_build_198 int64) int {
	return dm_build_196.Dm_build_219(dm_build_197, Dm_build_1220.Dm_build_1407(dm_build_198))
}

func (dm_build_200 *Dm_build_0) Dm_build_199(dm_build_201 int, dm_build_202 float32) int {
	return dm_build_200.Dm_build_219(dm_build_201, Dm_build_1220.Dm_build_1410(dm_build_202))
}

func (dm_build_204 *Dm_build_0) Dm_build_203(dm_build_205 int, dm_build_206 float64) int {
	return dm_build_204.Dm_build_219(dm_build_205, Dm_build_1220.Dm_build_1413(dm_build_206))
}

func (dm_build_208 *Dm_build_0) Dm_build_207(dm_build_209 int, dm_build_210 uint8) int {
	return dm_build_208.Dm_build_219(dm_build_209, Dm_build_1220.Dm_build_1416(dm_build_210))
}

func (dm_build_212 *Dm_build_0) Dm_build_211(dm_build_213 int, dm_build_214 uint16) int {
	return dm_build_212.Dm_build_219(dm_build_213, Dm_build_1220.Dm_build_1419(dm_build_214))
}

func (dm_build_216 *Dm_build_0) Dm_build_215(dm_build_217 int, dm_build_218 uint32) int {
	return dm_build_216.Dm_build_219(dm_build_217, Dm_build_1220.Dm_build_1422(dm_build_218))
}

func (dm_build_220 *Dm_build_0) Dm_build_219(dm_build_221 int, dm_build_222 []byte) int {
	return copy(dm_build_220.dm_build_1[dm_build_221:], dm_build_222)
}

func (dm_build_224 *Dm_build_0) Dm_build_223(dm_build_225 int, dm_build_226 []byte) int {
	return dm_build_224.Dm_build_191(dm_build_225, int32(len(dm_build_226))) + dm_build_224.Dm_build_219(dm_build_225+4, dm_build_226)
}

func (dm_build_228 *Dm_build_0) Dm_build_227(dm_build_229 int, dm_build_230 []byte) int {
	return dm_build_228.Dm_build_183(dm_build_229, byte(len(dm_build_230))) + dm_build_228.Dm_build_219(dm_build_229+1, dm_build_230)
}

func (dm_build_232 *Dm_build_0) Dm_build_231(dm_build_233 int, dm_build_234 []byte) int {
	return dm_build_232.Dm_build_187(dm_build_233, int16(len(dm_build_234))) + dm_build_232.Dm_build_219(dm_build_233+2, dm_build_234)
}

func (dm_build_236 *Dm_build_0) Dm_build_235(dm_build_237 int, dm_build_238 []byte) int {
	return dm_build_236.Dm_build_219(dm_build_237, dm_build_238) + dm_build_236.Dm_build_183(dm_build_237+len(dm_build_238), 0)
}

func (dm_build_240 *Dm_build_0) Dm_build_239(dm_build_241 int, dm_build_242 string, dm_build_243 string, dm_build_244 *DmConnection) int {
	return dm_build_240.Dm_build_223(dm_build_241, Dm_build_1220.Dm_build_1433(dm_build_242, dm_build_243, dm_build_244))
}

func (dm_build_246 *Dm_build_0) Dm_build_245(dm_build_247 int, dm_build_248 string, dm_build_249 string, dm_build_250 *DmConnection) int {
	return dm_build_246.Dm_build_227(dm_build_247, Dm_build_1220.Dm_build_1433(dm_build_248, dm_build_249, dm_build_250))
}

func (dm_build_252 *Dm_build_0) Dm_build_251(dm_build_253 int, dm_build_254 string, dm_build_255 string, dm_build_256 *DmConnection) int {
	return dm_build_252.Dm_build_231(dm_build_253, Dm_build_1220.Dm_build_1433(dm_build_254, dm_build_255, dm_build_256))
}

func (dm_build_258 *Dm_build_0) Dm_build_257(dm_build_259 int, dm_build_260 string, dm_build_261 string, dm_build_262 *DmConnection) int {
	return dm_build_258.Dm_build_235(dm_build_259, Dm_build_1220.Dm_build_1433(dm_build_260, dm_build_261, dm_build_262))
}

func (dm_build_264 *Dm_build_0) Dm_build_263(dm_build_265 int) byte {
	return Dm_build_1220.Dm_build_1438(dm_build_264.Dm_build_290(dm_build_265, 1))
}

func (dm_build_267 *Dm_build_0) Dm_build_266(dm_build_268 int) int16 {
	return Dm_build_1220.Dm_build_1441(dm_build_267.Dm_build_290(dm_build_268, 2))
}

func (dm_build_270 *Dm_build_0) Dm_build_269(dm_build_271 int) int32 {
	return Dm_build_1220.Dm_build_1444(dm_build_270.Dm_build_290(dm_build_271, 4))
}

func (dm_build_273 *Dm_build_0) Dm_build_272(dm_build_274 int) int64 {
	return Dm_build_1220.Dm_build_1447(dm_build_273.Dm_build_290(dm_build_274, 8))
}

func (dm_build_276 *Dm_build_0) Dm_build_275(dm_build_277 int) float32 {
	return Dm_build_1220.Dm_build_1450(dm_build_276.Dm_build_290(dm_build_277, 4))
}

func (dm_build_279 *Dm_build_0) Dm_build_278(dm_build_280 int) float64 {
	return Dm_build_1220.Dm_build_1453(dm_build_279.Dm_build_290(dm_build_280, 8))
}

func (dm_build_282 *Dm_build_0) Dm_build_281(dm_build_283 int) uint8 {
	return Dm_build_1220.Dm_build_1456(dm_build_282.Dm_build_290(dm_build_283, 1))
}

func (dm_build_285 *Dm_build_0) Dm_build_284(dm_build_286 int) uint16 {
	return Dm_build_1220.Dm_build_1459(dm_build_285.Dm_build_290(dm_build_286, 2))
}

func (dm_build_288 *Dm_build_0) Dm_build_287(dm_build_289 int) uint32 {
	return Dm_build_1220.Dm_build_1462(dm_build_288.Dm_build_290(dm_build_289, 4))
}

func (dm_build_291 *Dm_build_0) Dm_build_290(dm_build_292 int, dm_build_293 int) []byte {
	return dm_build_291.dm_build_1[dm_build_292 : dm_build_292+dm_build_293]
}

func (dm_build_295 *Dm_build_0) Dm_build_294(dm_build_296 int) []byte {
	dm_build_297 := dm_build_295.Dm_build_269(dm_build_296)
	return dm_build_295.Dm_build_290(dm_build_296+4, int(dm_build_297))
}

func (dm_build_299 *Dm_build_0) Dm_build_298(dm_build_300 int) []byte {
	dm_build_301 := dm_build_299.Dm_build_263(dm_build_300)
	return dm_build_299.Dm_build_290(dm_build_300+1, int(dm_build_301))
}

func (dm_build_303 *Dm_build_0) Dm_build_302(dm_build_304 int) []byte {
	dm_build_305 := dm_build_303.Dm_build_266(dm_build_304)
	return dm_build_303.Dm_build_290(dm_build_304+2, int(dm_build_305))
}

func (dm_build_307 *Dm_build_0) Dm_build_306(dm_build_308 int) []byte {
	dm_build_309 := 0
	for dm_build_307.Dm_build_263(dm_build_308) != 0 {
		dm_build_308++
		dm_build_309++
	}

	return dm_build_307.Dm_build_290(dm_build_308-dm_build_309, int(dm_build_309))
}

func (dm_build_311 *Dm_build_0) Dm_build_310(dm_build_312 int, dm_build_313 string, dm_build_314 *DmConnection) string {
	return Dm_build_1220.Dm_build_1470(dm_build_311.Dm_build_294(dm_build_312), dm_build_313, dm_build_314)
}

func (dm_build_316 *Dm_build_0) Dm_build_315(dm_build_317 int, dm_build_318 string, dm_build_319 *DmConnection) string {
	return Dm_build_1220.Dm_build_1470(dm_build_316.Dm_build_298(dm_build_317), dm_build_318, dm_build_319)
}

func (dm_build_321 *Dm_build_0) Dm_build_320(dm_build_322 int, dm_build_323 string, dm_build_324 *DmConnection) string {
	return Dm_build_1220.Dm_build_1470(dm_build_321.Dm_build_302(dm_build_322), dm_build_323, dm_build_324)
}

func (dm_build_326 *Dm_build_0) Dm_build_325(dm_build_327 int, dm_build_328 string, dm_build_329 *DmConnection) string {
	return Dm_build_1220.Dm_build_1470(dm_build_326.Dm_build_306(dm_build_327), dm_build_328, dm_build_329)
}
