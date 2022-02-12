/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dmr

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_1219 struct{}

var Dm_build_1220 = &dm_build_1219{}

func (Dm_build_1222 *dm_build_1219) Dm_build_1221(dm_build_1223 []byte, dm_build_1224 int, dm_build_1225 byte) int {
	dm_build_1223[dm_build_1224] = dm_build_1225
	return 1
}

func (Dm_build_1227 *dm_build_1219) Dm_build_1226(dm_build_1228 []byte, dm_build_1229 int, dm_build_1230 int8) int {
	dm_build_1228[dm_build_1229] = byte(dm_build_1230)
	return 1
}

func (Dm_build_1232 *dm_build_1219) Dm_build_1231(dm_build_1233 []byte, dm_build_1234 int, dm_build_1235 int16) int {
	dm_build_1233[dm_build_1234] = byte(dm_build_1235)
	dm_build_1234++
	dm_build_1233[dm_build_1234] = byte(dm_build_1235 >> 8)
	return 2
}

func (Dm_build_1237 *dm_build_1219) Dm_build_1236(dm_build_1238 []byte, dm_build_1239 int, dm_build_1240 int32) int {
	dm_build_1238[dm_build_1239] = byte(dm_build_1240)
	dm_build_1239++
	dm_build_1238[dm_build_1239] = byte(dm_build_1240 >> 8)
	dm_build_1239++
	dm_build_1238[dm_build_1239] = byte(dm_build_1240 >> 16)
	dm_build_1239++
	dm_build_1238[dm_build_1239] = byte(dm_build_1240 >> 24)
	dm_build_1239++
	return 4
}

func (Dm_build_1242 *dm_build_1219) Dm_build_1241(dm_build_1243 []byte, dm_build_1244 int, dm_build_1245 int64) int {
	dm_build_1243[dm_build_1244] = byte(dm_build_1245)
	dm_build_1244++
	dm_build_1243[dm_build_1244] = byte(dm_build_1245 >> 8)
	dm_build_1244++
	dm_build_1243[dm_build_1244] = byte(dm_build_1245 >> 16)
	dm_build_1244++
	dm_build_1243[dm_build_1244] = byte(dm_build_1245 >> 24)
	dm_build_1244++
	dm_build_1243[dm_build_1244] = byte(dm_build_1245 >> 32)
	dm_build_1244++
	dm_build_1243[dm_build_1244] = byte(dm_build_1245 >> 40)
	dm_build_1244++
	dm_build_1243[dm_build_1244] = byte(dm_build_1245 >> 48)
	dm_build_1244++
	dm_build_1243[dm_build_1244] = byte(dm_build_1245 >> 56)
	return 8
}

func (Dm_build_1247 *dm_build_1219) Dm_build_1246(dm_build_1248 []byte, dm_build_1249 int, dm_build_1250 float32) int {
	return Dm_build_1247.Dm_build_1266(dm_build_1248, dm_build_1249, math.Float32bits(dm_build_1250))
}

func (Dm_build_1252 *dm_build_1219) Dm_build_1251(dm_build_1253 []byte, dm_build_1254 int, dm_build_1255 float64) int {
	return Dm_build_1252.Dm_build_1271(dm_build_1253, dm_build_1254, math.Float64bits(dm_build_1255))
}

func (Dm_build_1257 *dm_build_1219) Dm_build_1256(dm_build_1258 []byte, dm_build_1259 int, dm_build_1260 uint8) int {
	dm_build_1258[dm_build_1259] = byte(dm_build_1260)
	return 1
}

func (Dm_build_1262 *dm_build_1219) Dm_build_1261(dm_build_1263 []byte, dm_build_1264 int, dm_build_1265 uint16) int {
	dm_build_1263[dm_build_1264] = byte(dm_build_1265)
	dm_build_1264++
	dm_build_1263[dm_build_1264] = byte(dm_build_1265 >> 8)
	return 2
}

func (Dm_build_1267 *dm_build_1219) Dm_build_1266(dm_build_1268 []byte, dm_build_1269 int, dm_build_1270 uint32) int {
	dm_build_1268[dm_build_1269] = byte(dm_build_1270)
	dm_build_1269++
	dm_build_1268[dm_build_1269] = byte(dm_build_1270 >> 8)
	dm_build_1269++
	dm_build_1268[dm_build_1269] = byte(dm_build_1270 >> 16)
	dm_build_1269++
	dm_build_1268[dm_build_1269] = byte(dm_build_1270 >> 24)
	return 3
}

func (Dm_build_1272 *dm_build_1219) Dm_build_1271(dm_build_1273 []byte, dm_build_1274 int, dm_build_1275 uint64) int {
	dm_build_1273[dm_build_1274] = byte(dm_build_1275)
	dm_build_1274++
	dm_build_1273[dm_build_1274] = byte(dm_build_1275 >> 8)
	dm_build_1274++
	dm_build_1273[dm_build_1274] = byte(dm_build_1275 >> 16)
	dm_build_1274++
	dm_build_1273[dm_build_1274] = byte(dm_build_1275 >> 24)
	dm_build_1274++
	dm_build_1273[dm_build_1274] = byte(dm_build_1275 >> 32)
	dm_build_1274++
	dm_build_1273[dm_build_1274] = byte(dm_build_1275 >> 40)
	dm_build_1274++
	dm_build_1273[dm_build_1274] = byte(dm_build_1275 >> 48)
	dm_build_1274++
	dm_build_1273[dm_build_1274] = byte(dm_build_1275 >> 56)
	return 3
}

func (Dm_build_1277 *dm_build_1219) Dm_build_1276(dm_build_1278 []byte, dm_build_1279 int, dm_build_1280 []byte, dm_build_1281 int, dm_build_1282 int) int {
	copy(dm_build_1278[dm_build_1279:dm_build_1279+dm_build_1282], dm_build_1280[dm_build_1281:dm_build_1281+dm_build_1282])
	return dm_build_1282
}

func (Dm_build_1284 *dm_build_1219) Dm_build_1283(dm_build_1285 []byte, dm_build_1286 int, dm_build_1287 []byte, dm_build_1288 int, dm_build_1289 int) int {
	dm_build_1286 += Dm_build_1284.Dm_build_1266(dm_build_1285, dm_build_1286, uint32(dm_build_1289))
	return 4 + Dm_build_1284.Dm_build_1276(dm_build_1285, dm_build_1286, dm_build_1287, dm_build_1288, dm_build_1289)
}

func (Dm_build_1291 *dm_build_1219) Dm_build_1290(dm_build_1292 []byte, dm_build_1293 int, dm_build_1294 []byte, dm_build_1295 int, dm_build_1296 int) int {
	dm_build_1293 += Dm_build_1291.Dm_build_1261(dm_build_1292, dm_build_1293, uint16(dm_build_1296))
	return 2 + Dm_build_1291.Dm_build_1276(dm_build_1292, dm_build_1293, dm_build_1294, dm_build_1295, dm_build_1296)
}

func (Dm_build_1298 *dm_build_1219) Dm_build_1297(dm_build_1299 []byte, dm_build_1300 int, dm_build_1301 string, dm_build_1302 string, dm_build_1303 *DmConnection) int {
	dm_build_1304 := Dm_build_1298.Dm_build_1433(dm_build_1301, dm_build_1302, dm_build_1303)
	dm_build_1300 += Dm_build_1298.Dm_build_1266(dm_build_1299, dm_build_1300, uint32(len(dm_build_1304)))
	return 4 + Dm_build_1298.Dm_build_1276(dm_build_1299, dm_build_1300, dm_build_1304, 0, len(dm_build_1304))
}

func (Dm_build_1306 *dm_build_1219) Dm_build_1305(dm_build_1307 []byte, dm_build_1308 int, dm_build_1309 string, dm_build_1310 string, dm_build_1311 *DmConnection) int {
	dm_build_1312 := Dm_build_1306.Dm_build_1433(dm_build_1309, dm_build_1310, dm_build_1311)

	dm_build_1308 += Dm_build_1306.Dm_build_1261(dm_build_1307, dm_build_1308, uint16(len(dm_build_1312)))
	return 2 + Dm_build_1306.Dm_build_1276(dm_build_1307, dm_build_1308, dm_build_1312, 0, len(dm_build_1312))
}

func (Dm_build_1314 *dm_build_1219) Dm_build_1313(dm_build_1315 []byte, dm_build_1316 int) byte {
	return dm_build_1315[dm_build_1316]
}

func (Dm_build_1318 *dm_build_1219) Dm_build_1317(dm_build_1319 []byte, dm_build_1320 int) int16 {
	var dm_build_1321 int16
	dm_build_1321 = int16(dm_build_1319[dm_build_1320] & 0xff)
	dm_build_1320++
	dm_build_1321 |= int16(dm_build_1319[dm_build_1320]&0xff) << 8
	return dm_build_1321
}

func (Dm_build_1323 *dm_build_1219) Dm_build_1322(dm_build_1324 []byte, dm_build_1325 int) int32 {
	var dm_build_1326 int32
	dm_build_1326 = int32(dm_build_1324[dm_build_1325] & 0xff)
	dm_build_1325++
	dm_build_1326 |= int32(dm_build_1324[dm_build_1325]&0xff) << 8
	dm_build_1325++
	dm_build_1326 |= int32(dm_build_1324[dm_build_1325]&0xff) << 16
	dm_build_1325++
	dm_build_1326 |= int32(dm_build_1324[dm_build_1325]&0xff) << 24
	return dm_build_1326
}

func (Dm_build_1328 *dm_build_1219) Dm_build_1327(dm_build_1329 []byte, dm_build_1330 int) int64 {
	var dm_build_1331 int64
	dm_build_1331 = int64(dm_build_1329[dm_build_1330] & 0xff)
	dm_build_1330++
	dm_build_1331 |= int64(dm_build_1329[dm_build_1330]&0xff) << 8
	dm_build_1330++
	dm_build_1331 |= int64(dm_build_1329[dm_build_1330]&0xff) << 16
	dm_build_1330++
	dm_build_1331 |= int64(dm_build_1329[dm_build_1330]&0xff) << 24
	dm_build_1330++
	dm_build_1331 |= int64(dm_build_1329[dm_build_1330]&0xff) << 32
	dm_build_1330++
	dm_build_1331 |= int64(dm_build_1329[dm_build_1330]&0xff) << 40
	dm_build_1330++
	dm_build_1331 |= int64(dm_build_1329[dm_build_1330]&0xff) << 48
	dm_build_1330++
	dm_build_1331 |= int64(dm_build_1329[dm_build_1330]&0xff) << 56
	return dm_build_1331
}

func (Dm_build_1333 *dm_build_1219) Dm_build_1332(dm_build_1334 []byte, dm_build_1335 int) float32 {
	return math.Float32frombits(Dm_build_1333.Dm_build_1349(dm_build_1334, dm_build_1335))
}

func (Dm_build_1337 *dm_build_1219) Dm_build_1336(dm_build_1338 []byte, dm_build_1339 int) float64 {
	return math.Float64frombits(Dm_build_1337.Dm_build_1354(dm_build_1338, dm_build_1339))
}

func (Dm_build_1341 *dm_build_1219) Dm_build_1340(dm_build_1342 []byte, dm_build_1343 int) uint8 {
	return uint8(dm_build_1342[dm_build_1343] & 0xff)
}

func (Dm_build_1345 *dm_build_1219) Dm_build_1344(dm_build_1346 []byte, dm_build_1347 int) uint16 {
	var dm_build_1348 uint16
	dm_build_1348 = uint16(dm_build_1346[dm_build_1347] & 0xff)
	dm_build_1347++
	dm_build_1348 |= uint16(dm_build_1346[dm_build_1347]&0xff) << 8
	return dm_build_1348
}

func (Dm_build_1350 *dm_build_1219) Dm_build_1349(dm_build_1351 []byte, dm_build_1352 int) uint32 {
	var dm_build_1353 uint32
	dm_build_1353 = uint32(dm_build_1351[dm_build_1352] & 0xff)
	dm_build_1352++
	dm_build_1353 |= uint32(dm_build_1351[dm_build_1352]&0xff) << 8
	dm_build_1352++
	dm_build_1353 |= uint32(dm_build_1351[dm_build_1352]&0xff) << 16
	dm_build_1352++
	dm_build_1353 |= uint32(dm_build_1351[dm_build_1352]&0xff) << 24
	return dm_build_1353
}

func (Dm_build_1355 *dm_build_1219) Dm_build_1354(dm_build_1356 []byte, dm_build_1357 int) uint64 {
	var dm_build_1358 uint64
	dm_build_1358 = uint64(dm_build_1356[dm_build_1357] & 0xff)
	dm_build_1357++
	dm_build_1358 |= uint64(dm_build_1356[dm_build_1357]&0xff) << 8
	dm_build_1357++
	dm_build_1358 |= uint64(dm_build_1356[dm_build_1357]&0xff) << 16
	dm_build_1357++
	dm_build_1358 |= uint64(dm_build_1356[dm_build_1357]&0xff) << 24
	dm_build_1357++
	dm_build_1358 |= uint64(dm_build_1356[dm_build_1357]&0xff) << 32
	dm_build_1357++
	dm_build_1358 |= uint64(dm_build_1356[dm_build_1357]&0xff) << 40
	dm_build_1357++
	dm_build_1358 |= uint64(dm_build_1356[dm_build_1357]&0xff) << 48
	dm_build_1357++
	dm_build_1358 |= uint64(dm_build_1356[dm_build_1357]&0xff) << 56
	return dm_build_1358
}

func (Dm_build_1360 *dm_build_1219) Dm_build_1359(dm_build_1361 []byte, dm_build_1362 int) []byte {
	dm_build_1363 := Dm_build_1360.Dm_build_1349(dm_build_1361, dm_build_1362)

	dm_build_1364 := make([]byte, dm_build_1363)
	copy(dm_build_1364[:int(dm_build_1363)], dm_build_1361[dm_build_1362+4:dm_build_1362+4+int(dm_build_1363)])
	return dm_build_1364
}

func (Dm_build_1366 *dm_build_1219) Dm_build_1365(dm_build_1367 []byte, dm_build_1368 int) []byte {
	dm_build_1369 := Dm_build_1366.Dm_build_1344(dm_build_1367, dm_build_1368)

	dm_build_1370 := make([]byte, dm_build_1369)
	copy(dm_build_1370[:int(dm_build_1369)], dm_build_1367[dm_build_1368+2:dm_build_1368+2+int(dm_build_1369)])
	return dm_build_1370
}

func (Dm_build_1372 *dm_build_1219) Dm_build_1371(dm_build_1373 []byte, dm_build_1374 int, dm_build_1375 int) []byte {

	dm_build_1376 := make([]byte, dm_build_1375)
	copy(dm_build_1376[:dm_build_1375], dm_build_1373[dm_build_1374:dm_build_1374+dm_build_1375])
	return dm_build_1376
}

func (Dm_build_1378 *dm_build_1219) Dm_build_1377(dm_build_1379 []byte, dm_build_1380 int, dm_build_1381 int, dm_build_1382 string, dm_build_1383 *DmConnection) string {
	return Dm_build_1378.Dm_build_1470(dm_build_1379[dm_build_1380:dm_build_1380+dm_build_1381], dm_build_1382, dm_build_1383)
}

func (Dm_build_1385 *dm_build_1219) Dm_build_1384(dm_build_1386 []byte, dm_build_1387 int, dm_build_1388 string, dm_build_1389 *DmConnection) string {
	dm_build_1390 := Dm_build_1385.Dm_build_1349(dm_build_1386, dm_build_1387)
	dm_build_1387 += 4
	return Dm_build_1385.Dm_build_1377(dm_build_1386, dm_build_1387, int(dm_build_1390), dm_build_1388, dm_build_1389)
}

func (Dm_build_1392 *dm_build_1219) Dm_build_1391(dm_build_1393 []byte, dm_build_1394 int, dm_build_1395 string, dm_build_1396 *DmConnection) string {
	dm_build_1397 := Dm_build_1392.Dm_build_1344(dm_build_1393, dm_build_1394)
	dm_build_1394 += 2
	return Dm_build_1392.Dm_build_1377(dm_build_1393, dm_build_1394, int(dm_build_1397), dm_build_1395, dm_build_1396)
}

func (Dm_build_1399 *dm_build_1219) Dm_build_1398(dm_build_1400 byte) []byte {
	return []byte{dm_build_1400}
}

func (Dm_build_1402 *dm_build_1219) Dm_build_1401(dm_build_1403 int16) []byte {
	return []byte{byte(dm_build_1403), byte(dm_build_1403 >> 8)}
}

func (Dm_build_1405 *dm_build_1219) Dm_build_1404(dm_build_1406 int32) []byte {
	return []byte{byte(dm_build_1406), byte(dm_build_1406 >> 8), byte(dm_build_1406 >> 16), byte(dm_build_1406 >> 24)}
}

func (Dm_build_1408 *dm_build_1219) Dm_build_1407(dm_build_1409 int64) []byte {
	return []byte{byte(dm_build_1409), byte(dm_build_1409 >> 8), byte(dm_build_1409 >> 16), byte(dm_build_1409 >> 24), byte(dm_build_1409 >> 32),
		byte(dm_build_1409 >> 40), byte(dm_build_1409 >> 48), byte(dm_build_1409 >> 56)}
}

func (Dm_build_1411 *dm_build_1219) Dm_build_1410(dm_build_1412 float32) []byte {
	return Dm_build_1411.Dm_build_1422(math.Float32bits(dm_build_1412))
}

func (Dm_build_1414 *dm_build_1219) Dm_build_1413(dm_build_1415 float64) []byte {
	return Dm_build_1414.Dm_build_1425(math.Float64bits(dm_build_1415))
}

func (Dm_build_1417 *dm_build_1219) Dm_build_1416(dm_build_1418 uint8) []byte {
	return []byte{byte(dm_build_1418)}
}

func (Dm_build_1420 *dm_build_1219) Dm_build_1419(dm_build_1421 uint16) []byte {
	return []byte{byte(dm_build_1421), byte(dm_build_1421 >> 8)}
}

func (Dm_build_1423 *dm_build_1219) Dm_build_1422(dm_build_1424 uint32) []byte {
	return []byte{byte(dm_build_1424), byte(dm_build_1424 >> 8), byte(dm_build_1424 >> 16), byte(dm_build_1424 >> 24)}
}

func (Dm_build_1426 *dm_build_1219) Dm_build_1425(dm_build_1427 uint64) []byte {
	return []byte{byte(dm_build_1427), byte(dm_build_1427 >> 8), byte(dm_build_1427 >> 16), byte(dm_build_1427 >> 24), byte(dm_build_1427 >> 32), byte(dm_build_1427 >> 40), byte(dm_build_1427 >> 48), byte(dm_build_1427 >> 56)}
}

func (Dm_build_1429 *dm_build_1219) Dm_build_1428(dm_build_1430 []byte, dm_build_1431 string, dm_build_1432 *DmConnection) []byte {
	if dm_build_1431 == "UTF-8" {
		return dm_build_1430
	}

	if dm_build_1432 == nil {
		if e := dm_build_1475(dm_build_1431); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1430), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1432.encodeBuffer == nil {
		dm_build_1432.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1432.encode = dm_build_1475(dm_build_1432.getServerEncoding())
		dm_build_1432.transformReaderDst = make([]byte, 4096)
		dm_build_1432.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1432.encode; e != nil {

		dm_build_1432.encodeBuffer.Reset()

		n, err := dm_build_1432.encodeBuffer.ReadFrom(
			Dm_build_1489(bytes.NewReader(dm_build_1430), e.NewEncoder(), dm_build_1432.transformReaderDst, dm_build_1432.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1432.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1434 *dm_build_1219) Dm_build_1433(dm_build_1435 string, dm_build_1436 string, dm_build_1437 *DmConnection) []byte {
	return Dm_build_1434.Dm_build_1428([]byte(dm_build_1435), dm_build_1436, dm_build_1437)
}

func (Dm_build_1439 *dm_build_1219) Dm_build_1438(dm_build_1440 []byte) byte {
	return Dm_build_1439.Dm_build_1313(dm_build_1440, 0)
}

func (Dm_build_1442 *dm_build_1219) Dm_build_1441(dm_build_1443 []byte) int16 {
	return Dm_build_1442.Dm_build_1317(dm_build_1443, 0)
}

func (Dm_build_1445 *dm_build_1219) Dm_build_1444(dm_build_1446 []byte) int32 {
	return Dm_build_1445.Dm_build_1322(dm_build_1446, 0)
}

func (Dm_build_1448 *dm_build_1219) Dm_build_1447(dm_build_1449 []byte) int64 {
	return Dm_build_1448.Dm_build_1327(dm_build_1449, 0)
}

func (Dm_build_1451 *dm_build_1219) Dm_build_1450(dm_build_1452 []byte) float32 {
	return Dm_build_1451.Dm_build_1332(dm_build_1452, 0)
}

func (Dm_build_1454 *dm_build_1219) Dm_build_1453(dm_build_1455 []byte) float64 {
	return Dm_build_1454.Dm_build_1336(dm_build_1455, 0)
}

func (Dm_build_1457 *dm_build_1219) Dm_build_1456(dm_build_1458 []byte) uint8 {
	return Dm_build_1457.Dm_build_1340(dm_build_1458, 0)
}

func (Dm_build_1460 *dm_build_1219) Dm_build_1459(dm_build_1461 []byte) uint16 {
	return Dm_build_1460.Dm_build_1344(dm_build_1461, 0)
}

func (Dm_build_1463 *dm_build_1219) Dm_build_1462(dm_build_1464 []byte) uint32 {
	return Dm_build_1463.Dm_build_1349(dm_build_1464, 0)
}

func (Dm_build_1466 *dm_build_1219) Dm_build_1465(dm_build_1467 []byte, dm_build_1468 string, dm_build_1469 *DmConnection) []byte {
	if dm_build_1468 == "UTF-8" {
		return dm_build_1467
	}

	if dm_build_1469 == nil {
		if e := dm_build_1475(dm_build_1468); e != nil {

			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1467), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1469.encodeBuffer == nil {
		dm_build_1469.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1469.encode = dm_build_1475(dm_build_1469.getServerEncoding())
		dm_build_1469.transformReaderDst = make([]byte, 4096)
		dm_build_1469.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1469.encode; e != nil {

		dm_build_1469.encodeBuffer.Reset()

		n, err := dm_build_1469.encodeBuffer.ReadFrom(
			Dm_build_1489(bytes.NewReader(dm_build_1467), e.NewDecoder(), dm_build_1469.transformReaderDst, dm_build_1469.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1469.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1471 *dm_build_1219) Dm_build_1470(dm_build_1472 []byte, dm_build_1473 string, dm_build_1474 *DmConnection) string {
	return string(Dm_build_1471.Dm_build_1465(dm_build_1472, dm_build_1473, dm_build_1474))
}

func dm_build_1475(dm_build_1476 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1476); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1477 struct {
	dm_build_1478 io.Reader
	dm_build_1479 transform.Transformer
	dm_build_1480 error

	dm_build_1481                []byte
	dm_build_1482, dm_build_1483 int

	dm_build_1484                []byte
	dm_build_1485, dm_build_1486 int

	dm_build_1487 bool
}

const dm_build_1488 = 4096

func Dm_build_1489(dm_build_1490 io.Reader, dm_build_1491 transform.Transformer, dm_build_1492 []byte, dm_build_1493 []byte) *Dm_build_1477 {
	dm_build_1491.Reset()
	return &Dm_build_1477{
		dm_build_1478: dm_build_1490,
		dm_build_1479: dm_build_1491,
		dm_build_1481: dm_build_1492,
		dm_build_1484: dm_build_1493,
	}
}

func (dm_build_1495 *Dm_build_1477) Read(dm_build_1496 []byte) (int, error) {
	dm_build_1497, dm_build_1498 := 0, error(nil)
	for {

		if dm_build_1495.dm_build_1482 != dm_build_1495.dm_build_1483 {
			dm_build_1497 = copy(dm_build_1496, dm_build_1495.dm_build_1481[dm_build_1495.dm_build_1482:dm_build_1495.dm_build_1483])
			dm_build_1495.dm_build_1482 += dm_build_1497
			if dm_build_1495.dm_build_1482 == dm_build_1495.dm_build_1483 && dm_build_1495.dm_build_1487 {
				return dm_build_1497, dm_build_1495.dm_build_1480
			}
			return dm_build_1497, nil
		} else if dm_build_1495.dm_build_1487 {
			return 0, dm_build_1495.dm_build_1480
		}

		if dm_build_1495.dm_build_1485 != dm_build_1495.dm_build_1486 || dm_build_1495.dm_build_1480 != nil {
			dm_build_1495.dm_build_1482 = 0
			dm_build_1495.dm_build_1483, dm_build_1497, dm_build_1498 = dm_build_1495.dm_build_1479.Transform(dm_build_1495.dm_build_1481, dm_build_1495.dm_build_1484[dm_build_1495.dm_build_1485:dm_build_1495.dm_build_1486], dm_build_1495.dm_build_1480 == io.EOF)
			dm_build_1495.dm_build_1485 += dm_build_1497

			switch {
			case dm_build_1498 == nil:
				if dm_build_1495.dm_build_1485 != dm_build_1495.dm_build_1486 {
					dm_build_1495.dm_build_1480 = nil
				}

				dm_build_1495.dm_build_1487 = dm_build_1495.dm_build_1480 != nil
				continue
			case dm_build_1498 == transform.ErrShortDst && (dm_build_1495.dm_build_1483 != 0 || dm_build_1497 != 0):

				continue
			case dm_build_1498 == transform.ErrShortSrc && dm_build_1495.dm_build_1486-dm_build_1495.dm_build_1485 != len(dm_build_1495.dm_build_1484) && dm_build_1495.dm_build_1480 == nil:

			default:
				dm_build_1495.dm_build_1487 = true

				if dm_build_1495.dm_build_1480 == nil || dm_build_1495.dm_build_1480 == io.EOF {
					dm_build_1495.dm_build_1480 = dm_build_1498
				}
				continue
			}
		}

		if dm_build_1495.dm_build_1485 != 0 {
			dm_build_1495.dm_build_1485, dm_build_1495.dm_build_1486 = 0, copy(dm_build_1495.dm_build_1484, dm_build_1495.dm_build_1484[dm_build_1495.dm_build_1485:dm_build_1495.dm_build_1486])
		}
		dm_build_1497, dm_build_1495.dm_build_1480 = dm_build_1495.dm_build_1478.Read(dm_build_1495.dm_build_1484[dm_build_1495.dm_build_1486:])
		dm_build_1495.dm_build_1486 += dm_build_1497
	}
}
