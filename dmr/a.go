/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dmr

import (
	"bytes"
	"crypto/tls"
	"net"
	"strconv"
	"time"

	"github.com/nfjBill/gorm-driver-dm/dmr/security"
)

const (
	Dm_build_330 = 8192
	Dm_build_331 = 2 * time.Second
)

type dm_build_332 struct {
	dm_build_333 *net.TCPConn
	dm_build_334 *tls.Conn
	dm_build_335 *Dm_build_0
	dm_build_336 *DmConnection
	dm_build_337 security.Cipher
	dm_build_338 bool
	dm_build_339 bool
	dm_build_340 *security.DhKey

	dm_build_341 bool
	dm_build_342 string
	dm_build_343 bool
}

func dm_build_344(dm_build_345 *DmConnection) (*dm_build_332, error) {
	dm_build_346, dm_build_347 := dm_build_349(dm_build_345.dmConnector.host+":"+strconv.Itoa(int(dm_build_345.dmConnector.port)), time.Duration(dm_build_345.dmConnector.socketTimeout)*time.Second)
	if dm_build_347 != nil {
		return nil, dm_build_347
	}

	dm_build_348 := dm_build_332{}
	dm_build_348.dm_build_333 = dm_build_346
	dm_build_348.dm_build_335 = Dm_build_3(Dm_build_604)
	dm_build_348.dm_build_336 = dm_build_345
	dm_build_348.dm_build_338 = false
	dm_build_348.dm_build_339 = false
	dm_build_348.dm_build_341 = false
	dm_build_348.dm_build_342 = ""
	dm_build_348.dm_build_343 = false
	dm_build_345.Access = &dm_build_348

	return &dm_build_348, nil
}

func dm_build_349(dm_build_350 string, dm_build_351 time.Duration) (*net.TCPConn, error) {
	dm_build_352, dm_build_353 := net.DialTimeout("tcp", dm_build_350, dm_build_351)
	if dm_build_353 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_350).throw()
	}

	if tcpConn, ok := dm_build_352.(*net.TCPConn); ok {

		tcpConn.SetKeepAlive(true)
		tcpConn.SetKeepAlivePeriod(Dm_build_331)
		tcpConn.SetNoDelay(true)

		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_355 *dm_build_332) dm_build_354(dm_build_356 dm_build_725) bool {
	var dm_build_357 = dm_build_355.dm_build_336.dmConnector.compress
	if dm_build_356.dm_build_740() == Dm_build_632 || dm_build_357 == Dm_build_681 {
		return false
	}

	if dm_build_357 == Dm_build_679 {
		return true
	} else if dm_build_357 == Dm_build_680 {
		return !dm_build_355.dm_build_336.Local && dm_build_356.dm_build_738() > Dm_build_678
	}

	return false
}

func (dm_build_359 *dm_build_332) dm_build_358(dm_build_360 dm_build_725) bool {
	var dm_build_361 = dm_build_359.dm_build_336.dmConnector.compress
	if dm_build_360.dm_build_740() == Dm_build_632 || dm_build_361 == Dm_build_681 {
		return false
	}

	if dm_build_361 == Dm_build_679 {
		return true
	} else if dm_build_361 == Dm_build_680 {
		return dm_build_359.dm_build_335.Dm_build_263(Dm_build_640) == 1
	}

	return false
}

func (dm_build_363 *dm_build_332) dm_build_362(dm_build_364 dm_build_725) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_366 := dm_build_364.dm_build_738()

	if dm_build_366 > 0 {

		if dm_build_363.dm_build_354(dm_build_364) {
			var retBytes, err = Compress(dm_build_363.dm_build_335, Dm_build_633, int(dm_build_366), int(dm_build_363.dm_build_336.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_363.dm_build_335.Dm_build_14(Dm_build_633)

			dm_build_363.dm_build_335.Dm_build_51(dm_build_366)

			dm_build_363.dm_build_335.Dm_build_79(retBytes)

			dm_build_364.dm_build_739(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_363.dm_build_335.Dm_build_183(Dm_build_640, 1)
		}

		if dm_build_363.dm_build_339 {
			dm_build_366 = dm_build_364.dm_build_738()
			var retBytes = dm_build_363.dm_build_337.Encrypt(dm_build_363.dm_build_335.Dm_build_290(Dm_build_633, int(dm_build_366)), true)

			dm_build_363.dm_build_335.Dm_build_14(Dm_build_633)

			dm_build_363.dm_build_335.Dm_build_79(retBytes)

			dm_build_364.dm_build_739(int32(len(retBytes)))
		}
	}

	if dm_build_363.dm_build_335.Dm_build_12() > Dm_build_605 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_364.dm_build_734()
	if dm_build_363.dm_build_594(dm_build_364) {
		if dm_build_363.dm_build_334 != nil {
			dm_build_363.dm_build_335.Dm_build_17(0)
			dm_build_363.dm_build_335.Dm_build_36(dm_build_363.dm_build_334)
		}
	} else {
		dm_build_363.dm_build_335.Dm_build_17(0)
		dm_build_363.dm_build_335.Dm_build_36(dm_build_363.dm_build_333)
	}
	return nil
}

func (dm_build_368 *dm_build_332) dm_build_367(dm_build_369 dm_build_725) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				panic(p)
			}
		}
	}()

	dm_build_371 := int32(0)
	if dm_build_368.dm_build_594(dm_build_369) {
		if dm_build_368.dm_build_334 != nil {
			dm_build_368.dm_build_335.Dm_build_14(0)
			dm_build_368.dm_build_335.Dm_build_30(dm_build_368.dm_build_334, Dm_build_633)
			dm_build_371 = dm_build_369.dm_build_738()
			if dm_build_371 > 0 {
				dm_build_368.dm_build_335.Dm_build_30(dm_build_368.dm_build_334, int(dm_build_371))
			}
		}
	} else {

		dm_build_368.dm_build_335.Dm_build_14(0)
		dm_build_368.dm_build_335.Dm_build_30(dm_build_368.dm_build_333, Dm_build_633)
		dm_build_371 = dm_build_369.dm_build_738()

		if dm_build_371 > 0 {
			dm_build_368.dm_build_335.Dm_build_30(dm_build_368.dm_build_333, int(dm_build_371))
		}
	}

	dm_build_369.dm_build_735()

	dm_build_371 = dm_build_369.dm_build_738()
	if dm_build_371 <= 0 {
		return nil
	}

	if dm_build_368.dm_build_339 {
		ebytes := dm_build_368.dm_build_335.Dm_build_290(Dm_build_633, int(dm_build_371))
		bytes, err := dm_build_368.dm_build_337.Decrypt(ebytes, true)
		if err != nil {
			return err
		}
		dm_build_368.dm_build_335.Dm_build_14(Dm_build_633)
		dm_build_368.dm_build_335.Dm_build_79(bytes)
		dm_build_369.dm_build_739(int32(len(bytes)))
	}

	if dm_build_368.dm_build_358(dm_build_369) {

		dm_build_371 = dm_build_369.dm_build_738()
		cbytes := dm_build_368.dm_build_335.Dm_build_290(Dm_build_633+ULINT_SIZE, int(dm_build_371-ULINT_SIZE))
		bytes, err := UnCompress(cbytes, int(dm_build_368.dm_build_336.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_368.dm_build_335.Dm_build_14(Dm_build_633)
		dm_build_368.dm_build_335.Dm_build_79(bytes)
		dm_build_369.dm_build_739(int32(len(bytes)))
	}
	return nil
}

func (dm_build_373 *dm_build_332) dm_build_372(dm_build_374 dm_build_725) (dm_build_375 interface{}, dm_build_376 error) {
	dm_build_376 = dm_build_374.dm_build_729(dm_build_374)
	if dm_build_376 != nil {
		return nil, dm_build_376
	}

	dm_build_376 = dm_build_373.dm_build_362(dm_build_374)
	if dm_build_376 != nil {
		return nil, dm_build_376
	}

	dm_build_376 = dm_build_373.dm_build_367(dm_build_374)
	if dm_build_376 != nil {
		return nil, dm_build_376
	}

	return dm_build_374.dm_build_733(dm_build_374)
}

func (dm_build_378 *dm_build_332) dm_build_377() (*dm_build_1162, error) {

	Dm_build_379 := dm_build_1168(dm_build_378)
	_, dm_build_380 := dm_build_378.dm_build_372(Dm_build_379)
	if dm_build_380 != nil {
		return nil, dm_build_380
	}

	return Dm_build_379, nil
}

func (dm_build_382 *dm_build_332) dm_build_381() error {

	dm_build_383 := dm_build_1030(dm_build_382)
	_, dm_build_384 := dm_build_382.dm_build_372(dm_build_383)
	if dm_build_384 != nil {
		return dm_build_384
	}

	return nil
}

func (dm_build_386 *dm_build_332) dm_build_385() error {

	var dm_build_387 *dm_build_1162
	var err error
	if dm_build_387, err = dm_build_386.dm_build_377(); err != nil {
		return err
	}

	if dm_build_386.dm_build_336.sslEncrypt == 2 {
		if err = dm_build_386.dm_build_590(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_386.dm_build_336.sslEncrypt == 1 {
		if err = dm_build_386.dm_build_590(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_386.dm_build_339 || dm_build_386.dm_build_338 {
		k, err := dm_build_386.dm_build_580()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_387.Dm_build_1166)
		encryptType := dm_build_387.dm_build_1164
		hashType := int(dm_build_387.Dm_build_1165)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_386.dm_build_583(encryptType, sessionKey, dm_build_386.dm_build_336.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_386.dm_build_381(); err != nil {
		return err
	}
	return nil
}

func (dm_build_390 *dm_build_332) Dm_build_389(dm_build_391 *DmStatement) error {
	dm_build_392 := dm_build_1191(dm_build_390, dm_build_391)
	_, dm_build_393 := dm_build_390.dm_build_372(dm_build_392)
	if dm_build_393 != nil {
		return dm_build_393
	}

	return nil
}

func (dm_build_395 *dm_build_332) Dm_build_394(dm_build_396 int32) error {
	dm_build_397 := dm_build_1201(dm_build_395, dm_build_396)
	_, dm_build_398 := dm_build_395.dm_build_372(dm_build_397)
	if dm_build_398 != nil {
		return dm_build_398
	}

	return nil
}

func (dm_build_400 *dm_build_332) Dm_build_399(dm_build_401 *DmStatement, dm_build_402 bool, dm_build_403 int16) (*execRetInfo, error) {
	dm_build_404 := dm_build_1068(dm_build_400, dm_build_401, dm_build_402, dm_build_403)
	dm_build_405, dm_build_406 := dm_build_400.dm_build_372(dm_build_404)
	if dm_build_406 != nil {
		return nil, dm_build_406
	}
	return dm_build_405.(*execRetInfo), nil
}

func (dm_build_408 *dm_build_332) Dm_build_407(dm_build_409 *DmStatement, dm_build_410 int16) (*execRetInfo, error) {
	return dm_build_408.Dm_build_399(dm_build_409, false, Dm_build_685)
}

func (dm_build_412 *dm_build_332) Dm_build_411(dm_build_413 *DmStatement, dm_build_414 []OptParameter) (*execRetInfo, error) {
	dm_build_415, dm_build_416 := dm_build_412.dm_build_372(dm_build_827(dm_build_412, dm_build_413, dm_build_414))
	if dm_build_416 != nil {
		return nil, dm_build_416
	}

	return dm_build_415.(*execRetInfo), nil
}

func (dm_build_418 *dm_build_332) Dm_build_417(dm_build_419 *DmStatement, dm_build_420 int16) (*execRetInfo, error) {
	return dm_build_418.Dm_build_399(dm_build_419, true, dm_build_420)
}

func (dm_build_422 *dm_build_332) Dm_build_421(dm_build_423 *DmStatement, dm_build_424 [][]interface{}) (*execRetInfo, error) {
	dm_build_425 := dm_build_850(dm_build_422, dm_build_423, dm_build_424)
	dm_build_426, dm_build_427 := dm_build_422.dm_build_372(dm_build_425)
	if dm_build_427 != nil {
		return nil, dm_build_427
	}
	return dm_build_426.(*execRetInfo), nil
}

func (dm_build_429 *dm_build_332) Dm_build_428(dm_build_430 *DmStatement, dm_build_431 [][]interface{}, dm_build_432 bool) (*execRetInfo, error) {
	var dm_build_433, dm_build_434 = 0, 0
	var dm_build_435 = len(dm_build_431)
	var dm_build_436 [][]interface{}
	var dm_build_437 = NewExceInfo()
	dm_build_437.updateCounts = make([]int64, dm_build_435)
	var dm_build_438 = false
	for dm_build_433 < dm_build_435 {
		for dm_build_434 = dm_build_433; dm_build_434 < dm_build_435; dm_build_434++ {
			paramData := dm_build_431[dm_build_434]
			bindData := make([]interface{}, dm_build_430.paramCount)
			dm_build_438 = false
			for icol := 0; icol < int(dm_build_430.paramCount); icol++ {
				if dm_build_430.params[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_429.dm_build_563(bindData, paramData, icol) {
					dm_build_438 = true
					break
				}
			}

			if dm_build_438 {
				break
			}
			dm_build_436 = append(dm_build_436, bindData)
		}

		if dm_build_434 != dm_build_433 {
			tmpExecInfo, err := dm_build_429.Dm_build_421(dm_build_430, dm_build_436)
			if err != nil {
				return nil, err
			}
			dm_build_436 = dm_build_436[0:0]
			dm_build_437.union(tmpExecInfo, dm_build_433, dm_build_434-dm_build_433)
		}

		if dm_build_434 < dm_build_435 {
			tmpExecInfo, err := dm_build_429.Dm_build_439(dm_build_430, dm_build_431[dm_build_434], dm_build_432)
			if err != nil {
				return nil, err
			}

			dm_build_432 = true
			dm_build_437.union(tmpExecInfo, dm_build_434, 1)
		}

		dm_build_433 = dm_build_434 + 1
	}
	for _, i := range dm_build_437.updateCounts {
		if i > 0 {
			dm_build_437.updateCount += i
		}
	}
	return dm_build_437, nil
}

func (dm_build_440 *dm_build_332) Dm_build_439(dm_build_441 *DmStatement, dm_build_442 []interface{}, dm_build_443 bool) (*execRetInfo, error) {

	var dm_build_444 = make([]interface{}, dm_build_441.paramCount)
	for icol := 0; icol < int(dm_build_441.paramCount); icol++ {
		if dm_build_441.params[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_440.dm_build_563(dm_build_444, dm_build_442, icol) {

			if !dm_build_443 {
				preExecute := dm_build_1058(dm_build_440, dm_build_441, dm_build_441.params)
				dm_build_440.dm_build_372(preExecute)
				dm_build_443 = true
			}

			dm_build_440.dm_build_569(dm_build_441, dm_build_441.params[icol], icol, dm_build_442[icol].(iOffRowBinder))
			dm_build_444[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_445 = make([][]interface{}, 1, 1)
	dm_build_445[0] = dm_build_444

	dm_build_446 := dm_build_850(dm_build_440, dm_build_441, dm_build_445)
	dm_build_447, dm_build_448 := dm_build_440.dm_build_372(dm_build_446)
	if dm_build_448 != nil {
		return nil, dm_build_448
	}
	return dm_build_447.(*execRetInfo), nil
}

func (dm_build_450 *dm_build_332) Dm_build_449(dm_build_451 *DmStatement, dm_build_452 int16) (*execRetInfo, error) {
	dm_build_453 := dm_build_1045(dm_build_450, dm_build_451, dm_build_452)

	dm_build_454, dm_build_455 := dm_build_450.dm_build_372(dm_build_453)
	if dm_build_455 != nil {
		return nil, dm_build_455
	}
	return dm_build_454.(*execRetInfo), nil
}

func (dm_build_457 *dm_build_332) Dm_build_456(dm_build_458 *innerRows, dm_build_459 int64) (*execRetInfo, error) {
	dm_build_460 := dm_build_950(dm_build_457, dm_build_458, dm_build_459, INT64_MAX)
	dm_build_461, dm_build_462 := dm_build_457.dm_build_372(dm_build_460)
	if dm_build_462 != nil {
		return nil, dm_build_462
	}
	return dm_build_461.(*execRetInfo), nil
}

func (dm_build_464 *dm_build_332) Commit() error {
	dm_build_465 := dm_build_813(dm_build_464)
	_, dm_build_466 := dm_build_464.dm_build_372(dm_build_465)
	if dm_build_466 != nil {
		return dm_build_466
	}

	return nil
}

func (dm_build_468 *dm_build_332) Rollback() error {
	dm_build_469 := dm_build_1106(dm_build_468)
	_, dm_build_470 := dm_build_468.dm_build_372(dm_build_469)
	if dm_build_470 != nil {
		return dm_build_470
	}

	return nil
}

func (dm_build_472 *dm_build_332) Dm_build_471(dm_build_473 *DmConnection) error {
	dm_build_474 := dm_build_1111(dm_build_472, dm_build_473.IsoLevel)
	_, dm_build_475 := dm_build_472.dm_build_372(dm_build_474)
	if dm_build_475 != nil {
		return dm_build_475
	}

	return nil
}

func (dm_build_477 *dm_build_332) Dm_build_476(dm_build_478 *DmStatement, dm_build_479 string) error {
	dm_build_480 := dm_build_818(dm_build_477, dm_build_478, dm_build_479)
	_, dm_build_481 := dm_build_477.dm_build_372(dm_build_480)
	if dm_build_481 != nil {
		return dm_build_481
	}

	return nil
}

func (dm_build_483 *dm_build_332) Dm_build_482(dm_build_484 []uint32) ([]int64, error) {
	dm_build_485 := dm_build_1209(dm_build_483, dm_build_484)
	dm_build_486, dm_build_487 := dm_build_483.dm_build_372(dm_build_485)
	if dm_build_487 != nil {
		return nil, dm_build_487
	}
	return dm_build_486.([]int64), nil
}

func (dm_build_489 *dm_build_332) Close() error {
	if dm_build_489.dm_build_343 {
		return nil
	}

	dm_build_490 := dm_build_489.dm_build_333.Close()
	if dm_build_490 != nil {
		return dm_build_490
	}

	dm_build_489.dm_build_336 = nil
	dm_build_489.dm_build_343 = true
	return nil
}

func (dm_build_492 *dm_build_332) dm_build_491(dm_build_493 *lob) (int64, error) {
	dm_build_494 := dm_build_981(dm_build_492, dm_build_493)
	dm_build_495, dm_build_496 := dm_build_492.dm_build_372(dm_build_494)
	if dm_build_496 != nil {
		return 0, dm_build_496
	}
	return dm_build_495.(int64), nil
}

func (dm_build_498 *dm_build_332) dm_build_497(dm_build_499 *lob, dm_build_500 int32, dm_build_501 int32) ([]byte, error) {
	dm_build_502 := dm_build_968(dm_build_498, dm_build_499, int(dm_build_500), int(dm_build_501))
	dm_build_503, dm_build_504 := dm_build_498.dm_build_372(dm_build_502)
	if dm_build_504 != nil {
		return nil, dm_build_504
	}
	return dm_build_503.([]byte), nil
}

func (dm_build_506 *dm_build_332) dm_build_505(dm_build_507 *DmBlob, dm_build_508 int32, dm_build_509 int32) ([]byte, error) {
	var dm_build_510 = make([]byte, dm_build_509)
	var dm_build_511 int32 = 0
	var dm_build_512 int32 = 0
	var dm_build_513 []byte
	var dm_build_514 error
	for dm_build_511 < dm_build_509 {
		dm_build_512 = dm_build_509 - dm_build_511
		if dm_build_512 > Dm_build_718 {
			dm_build_512 = Dm_build_718
		}
		dm_build_513, dm_build_514 = dm_build_506.dm_build_497(&dm_build_507.lob, dm_build_508, dm_build_512)
		if dm_build_514 != nil {
			return nil, dm_build_514
		}
		if dm_build_513 == nil || len(dm_build_513) == 0 {
			break
		}
		Dm_build_1220.Dm_build_1276(dm_build_510, int(dm_build_511), dm_build_513, 0, len(dm_build_513))
		dm_build_511 += int32(len(dm_build_513))
		dm_build_508 += int32(len(dm_build_513))
		if dm_build_507.readOver {
			break
		}
	}
	return dm_build_510, nil
}

func (dm_build_516 *dm_build_332) dm_build_515(dm_build_517 *DmClob, dm_build_518 int32, dm_build_519 int32) (string, error) {
	var dm_build_520 bytes.Buffer
	var dm_build_521 int32 = 0
	var dm_build_522 int32 = 0
	var dm_build_523 []byte
	var dm_build_524 string
	var dm_build_525 error
	for dm_build_521 < dm_build_519 {
		dm_build_522 = dm_build_519 - dm_build_521
		if dm_build_522 > Dm_build_718/2 {
			dm_build_522 = Dm_build_718 / 2
		}
		dm_build_523, dm_build_525 = dm_build_516.dm_build_497(&dm_build_517.lob, dm_build_518, dm_build_522)
		if dm_build_525 != nil {
			return "", dm_build_525
		}
		if dm_build_523 == nil || len(dm_build_523) == 0 {
			break
		}
		dm_build_524 = Dm_build_1220.Dm_build_1377(dm_build_523, 0, len(dm_build_523), dm_build_517.serverEncoding, dm_build_516.dm_build_336)

		dm_build_520.WriteString(dm_build_524)
		dm_build_521 += int32(len(dm_build_524))
		dm_build_518 += int32(len(dm_build_524))
		if dm_build_517.readOver {
			break
		}
	}
	return dm_build_520.String(), nil
}

func (dm_build_527 *dm_build_332) dm_build_526(dm_build_528 *DmClob, dm_build_529 int, dm_build_530 string, dm_build_531 string) (int, error) {
	var dm_build_532 = Dm_build_1220.Dm_build_1433(dm_build_530, dm_build_531, dm_build_527.dm_build_336)
	var dm_build_533 = 0
	var dm_build_534 = len(dm_build_532)
	var dm_build_535 = 0
	var dm_build_536 = 0
	var dm_build_537 = 0
	var dm_build_538 = dm_build_534/Dm_build_717 + 1
	var dm_build_539 byte = 0
	var dm_build_540 byte = 0x01
	var dm_build_541 byte = 0x02
	for i := 0; i < dm_build_538; i++ {
		dm_build_539 = 0
		if i == 0 {
			dm_build_539 |= dm_build_540
		}
		if i == dm_build_538-1 {
			dm_build_539 |= dm_build_541
		}
		dm_build_537 = dm_build_534 - dm_build_536
		if dm_build_537 > Dm_build_717 {
			dm_build_537 = Dm_build_717
		}

		setLobData := dm_build_1125(dm_build_527, &dm_build_528.lob, dm_build_539, dm_build_529, dm_build_532, dm_build_533, dm_build_537)
		ret, err := dm_build_527.dm_build_372(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_535, nil
		} else {
			dm_build_529 += int(tmp)
			dm_build_535 += int(tmp)
			dm_build_536 += dm_build_537
			dm_build_533 += dm_build_537
		}
	}
	return dm_build_535, nil
}

func (dm_build_543 *dm_build_332) dm_build_542(dm_build_544 *DmBlob, dm_build_545 int, dm_build_546 []byte) (int, error) {
	var dm_build_547 = 0
	var dm_build_548 = len(dm_build_546)
	var dm_build_549 = 0
	var dm_build_550 = 0
	var dm_build_551 = 0
	var dm_build_552 = dm_build_548/Dm_build_717 + 1
	var dm_build_553 byte = 0
	var dm_build_554 byte = 0x01
	var dm_build_555 byte = 0x02
	for i := 0; i < dm_build_552; i++ {
		dm_build_553 = 0
		if i == 0 {
			dm_build_553 |= dm_build_554
		}
		if i == dm_build_552-1 {
			dm_build_553 |= dm_build_555
		}
		dm_build_551 = dm_build_548 - dm_build_550
		if dm_build_551 > Dm_build_717 {
			dm_build_551 = Dm_build_717
		}

		setLobData := dm_build_1125(dm_build_543, &dm_build_544.lob, dm_build_553, dm_build_545, dm_build_546, dm_build_547, dm_build_551)
		ret, err := dm_build_543.dm_build_372(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_549, nil
		} else {
			dm_build_545 += int(tmp)
			dm_build_549 += int(tmp)
			dm_build_550 += dm_build_551
			dm_build_547 += dm_build_551
		}
	}
	return dm_build_549, nil
}

func (dm_build_557 *dm_build_332) dm_build_556(dm_build_558 *lob, dm_build_559 int) (int64, error) {
	dm_build_560 := dm_build_992(dm_build_557, dm_build_558, dm_build_559)
	dm_build_561, dm_build_562 := dm_build_557.dm_build_372(dm_build_560)
	if dm_build_562 != nil {
		return dm_build_558.length, dm_build_562
	}
	return dm_build_561.(int64), nil
}

func (dm_build_564 *dm_build_332) dm_build_563(dm_build_565 []interface{}, dm_build_566 []interface{}, dm_build_567 int) bool {
	var dm_build_568 = false
	if dm_build_567 >= len(dm_build_566) || dm_build_566[dm_build_567] == nil {
		dm_build_565[dm_build_567] = ParamDataEnum_Null
	} else if binder, ok := dm_build_566[dm_build_567].(iOffRowBinder); ok {
		dm_build_568 = true
		dm_build_565[dm_build_567] = ParamDataEnum_OFF_ROW
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_564.dm_build_336) {
			dm_build_565[dm_build_567] = &lobCtl{lob.buildCtlData()}
			dm_build_568 = false
		}
	} else {
		dm_build_565[dm_build_567] = dm_build_566[dm_build_567]
	}
	return dm_build_568
}

func (dm_build_570 *dm_build_332) dm_build_569(dm_build_571 *DmStatement, dm_build_572 parameter, dm_build_573 int, dm_build_574 iOffRowBinder) error {
	var dm_build_575 = Dm_build_1503()
	dm_build_574.read(dm_build_575)
	var dm_build_576 = 0
	for !dm_build_574.isReadOver() || dm_build_575.Dm_build_1504() > 0 {
		if !dm_build_574.isReadOver() && dm_build_575.Dm_build_1504() < Dm_build_717 {
			dm_build_574.read(dm_build_575)
		}
		if dm_build_575.Dm_build_1504() > Dm_build_717 {
			dm_build_576 = Dm_build_717
		} else {
			dm_build_576 = dm_build_575.Dm_build_1504()
		}

		putData := dm_build_1096(dm_build_570, dm_build_571, int16(dm_build_573), dm_build_575, int32(dm_build_576))
		_, err := dm_build_570.dm_build_372(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_578 *dm_build_332) dm_build_577() ([]byte, error) {
	var dm_build_579 error
	if dm_build_578.dm_build_340 == nil {
		if dm_build_578.dm_build_340, dm_build_579 = security.NewClientKeyPair(); dm_build_579 != nil {
			return nil, dm_build_579
		}
	}
	return security.Bn2Bytes(dm_build_578.dm_build_340.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_581 *dm_build_332) dm_build_580() (*security.DhKey, error) {
	var dm_build_582 error
	if dm_build_581.dm_build_340 == nil {
		if dm_build_581.dm_build_340, dm_build_582 = security.NewClientKeyPair(); dm_build_582 != nil {
			return nil, dm_build_582
		}
	}
	return dm_build_581.dm_build_340, nil
}

func (dm_build_584 *dm_build_332) dm_build_583(dm_build_585 int, dm_build_586 []byte, dm_build_587 string, dm_build_588 int) (dm_build_589 error) {
	if dm_build_585 > 0 && dm_build_585 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_586 != nil {
		dm_build_584.dm_build_337, dm_build_589 = security.NewSymmCipher(dm_build_585, dm_build_586)
	} else if dm_build_585 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_584.dm_build_337, dm_build_589 = security.NewThirdPartCipher(dm_build_585, dm_build_586, dm_build_587, dm_build_588); dm_build_589 != nil {
			dm_build_589 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_589.Error()).throw()
		}
	}
	return
}

func (dm_build_591 *dm_build_332) dm_build_590(dm_build_592 bool) (dm_build_593 error) {
	if dm_build_591.dm_build_334, dm_build_593 = security.NewTLSFromTCP(dm_build_591.dm_build_333, dm_build_591.dm_build_336.dmConnector.sslCertPath, dm_build_591.dm_build_336.dmConnector.sslKeyPath, dm_build_591.dm_build_336.dmConnector.user); dm_build_593 != nil {
		return
	}
	if !dm_build_592 {
		dm_build_591.dm_build_334 = nil
	}
	return
}

func (dm_build_595 *dm_build_332) dm_build_594(dm_build_596 dm_build_725) bool {
	return dm_build_596.dm_build_740() != Dm_build_632 && dm_build_595.dm_build_336.sslEncrypt == 1
}
