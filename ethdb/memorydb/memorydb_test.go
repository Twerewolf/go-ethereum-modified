// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package memorydb

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"runtime"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/dbtest"
)

func TestMemoryDB(t *testing.T) {
	t.Run("DatabaseSuite", func(t *testing.T) {
		dbtest.TestDatabaseSuite(t, func() ethdb.KeyValueStore {
			// println("test databaseSuite")
			return New()
		})
	})
}
func Int32ToBytes(i int32) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}

var lastTotalFreed uint64

// Alloc：当前堆上对象占用的内存大小。
// TotalAlloc：堆上总共分配出的内存大小。
// Sys：程序从操作系统总共申请的内存大小。
// NumGC：垃圾回收运行的次数。
func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %vByte TotalAlloc = %vByte  Just Freed = %vByte Sys = %vByte NumGC = %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, ((m.TotalAlloc-m.Alloc)-lastTotalFreed)/1024, m.Sys/1024, m.NumGC)

	lastTotalFreed = m.TotalAlloc - m.Alloc
}

// test执行语句：
// go test  -run  ^Test1$
func Test1(t *testing.T) {
	//打开db
	printMemStats()
	db := New()      // 打开memorydb实例
	defer db.Close() //return 后执行
	//计时
	begin := time.Now()
	//写入
	const SIZE int = 1000000 //100w大小会消耗200MB内存,map
	var keyList [SIZE][]byte
	for i := 0; i < SIZE; i++ {
		var x int32
		binary.Read(rand.Reader, binary.LittleEndian, &x)
		// fmt.Println("random number: ", x)
		key := Int32ToBytes(x)
		value := []byte(key)
		err := db.Put(key, value) //存入db一个kv对
		if err != nil {
			t.Fatal(err)
		}
		//key存入list
		keyList[i] = key
	}
	//读取
	count := 0
	for i := 0; i < SIZE; i++ {
		key := keyList[i]
		data, err := db.Get(key)
		if err != nil {
			t.Fatal(err)
		}
		// fmt.Println("key: ", key)
		if data != nil {
			count++
		}
	}
	// fmt.Println("存储占用的内存字节数：")
	printMemStats()
	fmt.Println("count: ", count)
	end := time.Now()
	delta := end.Sub(begin)
	fmt.Println("delta time: ", delta)
}
