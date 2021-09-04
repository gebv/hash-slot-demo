# hash-slot-demo
Hash slot demo

Distribution of small text into N nodes.
Hash function - murmur3.

Manual checks

```bash
# 100 nodes
$ go run ./main.go -num 100 -val demo4 | grep "assigned to a node"
# assigned to a node: 89
$ go run ./main.go -num 100 -val demo3 | grep "assigned to a node"
# assigned to a node: 63
$ go run ./main.go -num 100 -val demo2 | grep "assigned to a node"
# assigned to a node: 78
$ go run ./main.go -num 100 -val demo1 | grep "assigned to a node"
# assigned to a node: 27

# 6 nodes
$ go run ./main.go -num 6 -val demo4 | grep "assigned to a node"
# assigned to a node: 5
$ go run ./main.go -num 6 -val demo3 | grep "assigned to a node"
# assigned to a node: 3
$ go run ./main.go -num 6 -val demo2 | grep "assigned to a node"
# assigned to a node: 4
$ go run ./main.go -num 6 -val demo1 | grep "assigned to a node"
# assigned to a node: 1
```

From a file (each line as an incoming value)

```bash
# 6 nodes
$ cat twitterusernames.csv | go run main.go
# bit size: 32
# max value by bit: 4294967295
# denominator: 134217727
# hasher: murmur3
# num of nodes: 6
# node boundaries ( length = 22369621 ):
# 	#%d: 1 0 22369622
# 	#%d: 2 22369622 44739243
# 	#%d: 3 44739243 67108864
# 	#%d: 4 67108864 89478485
# 	#%d: 5 89478485 111848106
# 	#%d: 6 111848106 134217727
# distribution histogram: map[0:26891 1:26526 2:26677 3:26508 4:26849 5:26549]

# 100 nodes
$ cat twitterusernames.csv | go run main.go -num 100
# bit size: 32
# max value by bit: 4294967295
# denominator: 134217727
# hasher: murmur3
# num of nodes: 100
# node boundaries ( length = 1342177 ):
# 	#%d: 1 0 1342178
# ...
# 	#%d: 100 132875524 134217701
# distribution histogram: map[0:1551 1:1557 2:1639 3:1596 4:1709 5:1626 6:1636 7:1621 8:1694 9:1658 10:1589 11:1578 12:1623 13:1562 14:1645 15:1560 16:1568 17:1647 18:1643 19:1608 20:1553 21:1700 22:1560 23:1662 24:1561 25:1531 26:1588 27:1594 28:1581 29:1570 30:1505 31:1577 32:1598 33:1568 34:1582 35:1483 36:1649 37:1588 38:1646 39:1622 40:1583 41:1629 42:1662 43:1585 44:1542 45:1577 46:1695 47:1667 48:1569 49:1557 50:1605 51:1650 52:1533 53:1507 54:1721 55:1628 56:1590 57:1651 58:1529 59:1535 60:1654 61:1534 62:1635 63:1582 64:1622 65:1493 66:1611 67:1504 68:1600 69:1637 70:1639 71:1522 72:1614 73:1625 74:1695 75:1636 76:1603 77:1612 78:1619 79:1625 80:1634 81:1550 82:1612 83:1492 84:1614 85:1686 86:1644 87:1498 88:1615 89:1639 90:1555 91:1578 92:1568 93:1573 94:1632 95:1687 96:1537 97:1550 98:1641 99:1590]
```

refs
* https://medium.com/miro-engineering/choosing-a-hash-function-to-solve-a-data-sharding-problem-c656259e2b54
* https://severalnines.com/database-blog/hash-slot-vs-consistent-hashing-redis
* twitter usernames datasets https://www.usna.edu/Users/cs/nchamber/data/twitternames/ only usernames `cat usernames-train.txt | cut -d ';' -f2 | awk '{print tolower($0)}' > twitterusernames.csv` (all usernames)
