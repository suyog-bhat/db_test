package main

import (
	"fmt"
	"github.com/linxGnu/grocksdb"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
	"os"
	"strconv"
	"time"
)

var (
	numberOfWrites int
)

func main() {
	//wg := sync.WaitGroup{}
	//wg.Add(2)
	//wg *sync.WaitGroup
	checkRocksDb()
	checklevelDb()
	//wg.Wait()

}

func checkRocksDb() {
	numberOfWritesStr := os.Getenv("NUMBER")
	numberOfWrites, _ = strconv.Atoi(numberOfWritesStr)
	bbto := grocksdb.NewDefaultBlockBasedTableOptions()
	opts := grocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true)
	db, err := grocksdb.OpenDb(opts, "./rocksdb")
	if err != nil {
		fmt.Println(err)
	}
	wo := grocksdb.NewDefaultWriteOptions()
	ro := grocksdb.NewDefaultReadOptions()
	wb := grocksdb.NewWriteBatch()
	t1 := time.Now()
	putBatchrocks(wb)
	err = db.Write(wo, wb)
	t2 := time.Now()
	fmt.Println("RocksDB----")
	fmt.Println("time took for write:", t2.Sub(t1))

	ro.SetPrefixSameAsStart(true)
	ro.SetIterateUpperBound([]byte("00f"))
	//ro.SetIterateLowerBound([]byte("00f"))
	fmt.Println(ro.PrefixSameAsStart())
	it := db.NewIterator(ro)
	t3 := time.Now()
	it.Seek([]byte("00e"))

	recordCount := countRecordsRocks(it)
	it.Close()
	t4 := time.Now()
	fmt.Println("Num of records:", recordCount, t4.Sub(t3))
	//err = db.Delete(wo, wb)
	db.Close()
	//wg.Done()
}

func putBatchrocks(wb *grocksdb.WriteBatch) {
	//for start := time.Now(); time.Since(start) < time.Second*5; {
	//	wb.Put([]byte(fmt.Sprintf("%d", i)), []byte("qwe"))
	//	i++
	//}
	buf := []byte{'0','0', 4 + 16: 0}
	var i int64
	for i = 0; i < 1000; i++ {
		buf = strconv.AppendInt(buf[:2], i, 16)
		wb.Put(buf, []byte("random"))
	}
}

func countRecordsRocks(it *grocksdb.Iterator) int {
	var count int
	for ; it.Valid(); it.Next() {
		count++
		//fmt.Println(string(it.Key().Data()))
	}
	return count
}


func checklevelDb() {
	db, err := leveldb.OpenFile("./lvldb", nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	batch := new(leveldb.Batch)
	t1 := time.Now()
	putBatchlevel(batch)
	err = db.Write(batch, nil)
	t2 := time.Now()

	fmt.Println("LevelDb-----")
	fmt.Println("time took for write", t2.Sub(t1))

	prefix:=[]byte("00e")
	iter := db.NewIterator(util.BytesPrefix(prefix), nil)
	t3 := time.Now()
	count := countRecordsLevel(&iter)
	t4 := time.Now()
	fmt.Println("Num of records:", count, t4.Sub(t3))
	iter.Release()
	//wg.Done()

}

func putBatchlevel(batch *leveldb.Batch) {
	//for start := time.Now(); time.Since(start) < time.Second*5; {
	//	batch.Put([]byte(fmt.Sprintf("%d", i)), []byte("qwe"))
	//	i++
	//}
	buf := []byte{'0','0', 4 + 16: 0}
	var i int64
	for i = 0; i < 1000; i++ {
		buf = strconv.AppendInt(buf[:2], i, 16)
		batch.Put(buf, []byte("random"))
	}
}

func countRecordsLevel(iter *iterator.Iterator) int {
	var count int
	for (*iter).Next() {
		count++
	}
	return count
}