// Copyright 2019 The go-ethereum Authors
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

package leveldb

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/ethdb/dbtest"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

//testleveldb
func TestLevelDB(t *testing.T) {
	t.Run("DatabaseSuite", func(t *testing.T) {
		dbtest.TestDatabaseSuite(t, func() ethdb.KeyValueStore {
			db, err := leveldb.Open(storage.NewMemStorage(), nil) //此处test打开的是memStorage 不知道目的是？
			if err != nil {
				t.Fatal(err)
			}
			return &Database{
				db: db,
			}
		})
	})
}
func Int32ToBytes(i int32) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint32(buf, uint32(i))
	return buf
}
func Test1(t *testing.T) {
	//打开db
	db, err := leveldb.OpenFile("./", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close() //return 后执行
	//计时
	begin := time.Now()
	//写入
	const SIZE int = 100000
	var keyList [SIZE][]byte
	for i := 0; i < SIZE; i++ {
		var x int32
		binary.Read(rand.Reader, binary.LittleEndian, &x)
		// fmt.Println("random number: ", x)
		key := Int32ToBytes(x)
		value := []byte(key)
		err := db.Put(key, value, nil) //存入db一个kv对
		if err != nil {
			t.Fatal(err)
		}
		//key存入list
		keyList[i] = key
	}

	// //存入列表
	// keyList := list.New()
	// // 尾部添加
	// keyList.PushBack(key)
	// for i := keyList.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// 	key := i.Value
	// 	keyList.Remove(i)
	// 	//读取

	// }
	//读取
	count := 0
	for i := 0; i < SIZE; i++ {
		key := keyList[i]
		data, err := db.Get(key, nil)
		if err != nil {
			t.Fatal(err)
		}
		// fmt.Println("key: ", key)
		// fmt.Println("data: ", data)
		if data != nil {
			count++
		}
	}
	fmt.Println("count: ", count)
	end := time.Now()
	delta := end.Sub(begin)
	fmt.Println("delta time: ", delta)
}

func Test2(t *testing.T) {

	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	} //不使用内存数据，直接取硬盘数据
	db, err := leveldb.OpenFile("./", o)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close() //return 后执行
	//计时
	begin := time.Now()
	//写入
	const SIZE int = 100000
	var keyList [SIZE][]byte
	for i := 0; i < SIZE; i++ {
		var x int32
		binary.Read(rand.Reader, binary.LittleEndian, &x)
		// fmt.Println("random number: ", x)
		key := Int32ToBytes(x)
		value := []byte(key)
		err := db.Put(key, value, nil) //存入db一个kv对
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
		data, err := db.Get(key, nil)
		if err != nil {
			t.Fatal(err)
		}
		// fmt.Println("key: ", key)
		// fmt.Println("data: ", data)
		if data != nil {
			count++
		}
	}
	fmt.Println("count: ", count)
	end := time.Now()
	delta := end.Sub(begin) //计算时间差
	fmt.Println("delta time: ", delta)
}
