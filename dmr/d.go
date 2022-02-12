/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dmr

import (
	"container/list"
	"io"
)

type Dm_build_1499 struct {
	dm_build_1500 *list.List
	dm_build_1501 *dm_build_1553
	dm_build_1502 int
}

func Dm_build_1503() *Dm_build_1499 {
	return &Dm_build_1499{
		dm_build_1500: list.New(),
		dm_build_1502: 0,
	}
}

func (dm_build_1505 *Dm_build_1499) Dm_build_1504() int {
	return dm_build_1505.dm_build_1502
}

func (dm_build_1507 *Dm_build_1499) Dm_build_1506(dm_build_1508 *Dm_build_0, dm_build_1509 int) int {
	var dm_build_1510 = 0
	var dm_build_1511 = 0
	for dm_build_1510 < dm_build_1509 && dm_build_1507.dm_build_1501 != nil {
		dm_build_1511 = dm_build_1507.dm_build_1501.dm_build_1561(dm_build_1508, dm_build_1509-dm_build_1510)
		if dm_build_1507.dm_build_1501.dm_build_1556 == 0 {
			dm_build_1507.dm_build_1543()
		}
		dm_build_1510 += dm_build_1511
		dm_build_1507.dm_build_1502 -= dm_build_1511
	}
	return dm_build_1510
}

func (dm_build_1513 *Dm_build_1499) Dm_build_1512(dm_build_1514 []byte, dm_build_1515 int, dm_build_1516 int) int {
	var dm_build_1517 = 0
	var dm_build_1518 = 0
	for dm_build_1517 < dm_build_1516 && dm_build_1513.dm_build_1501 != nil {
		dm_build_1518 = dm_build_1513.dm_build_1501.dm_build_1565(dm_build_1514, dm_build_1515, dm_build_1516-dm_build_1517)
		if dm_build_1513.dm_build_1501.dm_build_1556 == 0 {
			dm_build_1513.dm_build_1543()
		}
		dm_build_1517 += dm_build_1518
		dm_build_1513.dm_build_1502 -= dm_build_1518
		dm_build_1515 += dm_build_1518
	}
	return dm_build_1517
}

func (dm_build_1520 *Dm_build_1499) Dm_build_1519(dm_build_1521 io.Writer, dm_build_1522 int) int {
	var dm_build_1523 = 0
	var dm_build_1524 = 0
	for dm_build_1523 < dm_build_1522 && dm_build_1520.dm_build_1501 != nil {
		dm_build_1524 = dm_build_1520.dm_build_1501.dm_build_1570(dm_build_1521, dm_build_1522-dm_build_1523)
		if dm_build_1520.dm_build_1501.dm_build_1556 == 0 {
			dm_build_1520.dm_build_1543()
		}
		dm_build_1523 += dm_build_1524
		dm_build_1520.dm_build_1502 -= dm_build_1524
	}
	return dm_build_1523
}

func (dm_build_1526 *Dm_build_1499) Dm_build_1525(dm_build_1527 []byte, dm_build_1528 int, dm_build_1529 int) {
	if dm_build_1529 == 0 {
		return
	}
	var dm_build_1530 = dm_build_1557(dm_build_1527, dm_build_1528, dm_build_1529)
	if dm_build_1526.dm_build_1501 == nil {
		dm_build_1526.dm_build_1501 = dm_build_1530
	} else {
		dm_build_1526.dm_build_1500.PushBack(dm_build_1530)
	}
	dm_build_1526.dm_build_1502 += dm_build_1529
}

func (dm_build_1532 *Dm_build_1499) dm_build_1531(dm_build_1533 int) byte {
	var dm_build_1534 = dm_build_1533
	var dm_build_1535 = dm_build_1532.dm_build_1501
	for dm_build_1534 > 0 && dm_build_1535 != nil {
		if dm_build_1535.dm_build_1556 == 0 {
			continue
		}
		if dm_build_1534 > dm_build_1535.dm_build_1556-1 {
			dm_build_1534 -= dm_build_1535.dm_build_1556
			dm_build_1535 = dm_build_1532.dm_build_1500.Front().Value.(*dm_build_1553)
		} else {
			break
		}
	}
	return dm_build_1535.dm_build_1574(dm_build_1534)
}
func (dm_build_1537 *Dm_build_1499) Dm_build_1536(dm_build_1538 *Dm_build_1499) {
	if dm_build_1538.dm_build_1502 == 0 {
		return
	}
	var dm_build_1539 = dm_build_1538.dm_build_1501
	for dm_build_1539 != nil {
		dm_build_1537.dm_build_1540(dm_build_1539)
		dm_build_1538.dm_build_1543()
		dm_build_1539 = dm_build_1538.dm_build_1501
	}
	dm_build_1538.dm_build_1502 = 0
}
func (dm_build_1541 *Dm_build_1499) dm_build_1540(dm_build_1542 *dm_build_1553) {
	if dm_build_1542.dm_build_1556 == 0 {
		return
	}
	if dm_build_1541.dm_build_1501 == nil {
		dm_build_1541.dm_build_1501 = dm_build_1542
	} else {
		dm_build_1541.dm_build_1500.PushBack(dm_build_1542)
	}
	dm_build_1541.dm_build_1502 += dm_build_1542.dm_build_1556
}

func (dm_build_1544 *Dm_build_1499) dm_build_1543() {
	var dm_build_1545 = dm_build_1544.dm_build_1500.Front()
	if dm_build_1545 == nil {
		dm_build_1544.dm_build_1501 = nil
	} else {
		dm_build_1544.dm_build_1501 = dm_build_1545.Value.(*dm_build_1553)
		dm_build_1544.dm_build_1500.Remove(dm_build_1545)
	}
}

func (dm_build_1547 *Dm_build_1499) Dm_build_1546() []byte {
	var dm_build_1548 = make([]byte, dm_build_1547.dm_build_1502)
	var dm_build_1549 = dm_build_1547.dm_build_1501
	var dm_build_1550 = 0
	var dm_build_1551 = len(dm_build_1548)
	var dm_build_1552 = 0
	for dm_build_1549 != nil {
		if dm_build_1549.dm_build_1556 > 0 {
			if dm_build_1551 > dm_build_1549.dm_build_1556 {
				dm_build_1552 = dm_build_1549.dm_build_1556
			} else {
				dm_build_1552 = dm_build_1551
			}
			copy(dm_build_1548[dm_build_1550:dm_build_1550+dm_build_1552], dm_build_1549.dm_build_1554[dm_build_1549.dm_build_1555:dm_build_1549.dm_build_1555+dm_build_1552])
			dm_build_1550 += dm_build_1552
			dm_build_1551 -= dm_build_1552
		}
		if dm_build_1547.dm_build_1500.Front() == nil {
			dm_build_1549 = nil
		} else {
			dm_build_1549 = dm_build_1547.dm_build_1500.Front().Value.(*dm_build_1553)
		}
	}
	return dm_build_1548
}

type dm_build_1553 struct {
	dm_build_1554 []byte
	dm_build_1555 int
	dm_build_1556 int
}

func dm_build_1557(dm_build_1558 []byte, dm_build_1559 int, dm_build_1560 int) *dm_build_1553 {
	return &dm_build_1553{
		dm_build_1558,
		dm_build_1559,
		dm_build_1560,
	}
}

func (dm_build_1562 *dm_build_1553) dm_build_1561(dm_build_1563 *Dm_build_0, dm_build_1564 int) int {
	if dm_build_1562.dm_build_1556 <= dm_build_1564 {
		dm_build_1564 = dm_build_1562.dm_build_1556
	}
	dm_build_1563.Dm_build_79(dm_build_1562.dm_build_1554[dm_build_1562.dm_build_1555 : dm_build_1562.dm_build_1555+dm_build_1564])
	dm_build_1562.dm_build_1555 += dm_build_1564
	dm_build_1562.dm_build_1556 -= dm_build_1564
	return dm_build_1564
}

func (dm_build_1566 *dm_build_1553) dm_build_1565(dm_build_1567 []byte, dm_build_1568 int, dm_build_1569 int) int {
	if dm_build_1566.dm_build_1556 <= dm_build_1569 {
		dm_build_1569 = dm_build_1566.dm_build_1556
	}
	copy(dm_build_1567[dm_build_1568:dm_build_1568+dm_build_1569], dm_build_1566.dm_build_1554[dm_build_1566.dm_build_1555:dm_build_1566.dm_build_1555+dm_build_1569])
	dm_build_1566.dm_build_1555 += dm_build_1569
	dm_build_1566.dm_build_1556 -= dm_build_1569
	return dm_build_1569
}

func (dm_build_1571 *dm_build_1553) dm_build_1570(dm_build_1572 io.Writer, dm_build_1573 int) int {
	if dm_build_1571.dm_build_1556 <= dm_build_1573 {
		dm_build_1573 = dm_build_1571.dm_build_1556
	}
	dm_build_1572.Write(dm_build_1571.dm_build_1554[dm_build_1571.dm_build_1555 : dm_build_1571.dm_build_1555+dm_build_1573])
	dm_build_1571.dm_build_1555 += dm_build_1573
	dm_build_1571.dm_build_1556 -= dm_build_1573
	return dm_build_1573
}
func (dm_build_1575 *dm_build_1553) dm_build_1574(dm_build_1576 int) byte {
	return dm_build_1575.dm_build_1554[dm_build_1575.dm_build_1555+dm_build_1576]
}
