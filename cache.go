package main

import (
	"github.com/go-redis/redis/v8"
	"strconv"
)

//
//
//
/*
缓存接口提供如下功能
1. 哈希表
2. 集合
3. 有序集合
4.
 */
type ICache interface {
	Open(host string, port int, dbNum int, pwd string) (err error)
	Close() (err error)
	Pint() (err error)

	//auth(pwd string)
	//selectDb(n int)

	// 删除
	del(key string)(err error)

	// string
	set(key string, value string)(err error)
	get(key string)(err error)

	// set
	sAdd(key string, element string)(err error)
	sMembers()(err error)
	sRem()(err error)

	// hash key value
	hSet()(err error)
	hGet()(err error)

	// hash table
	hGetAll()(err error)
	hMSet()(err error)
	hMGet()(err error)

	// list
	//lPush()
	//lPpop()
	//lPush()

	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64,whence int) (pos int64,err error)
}

type Cache struct {
	rdb *redis.Client
	age int
}

func (c *Cache) Open(host string, port int, dbNum int, pwd string) (err error){
	addr := host + strconv.Itoa(port)
	c.rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       dbNum,  // use default DB
	})
	return nil
}

func (c *Cache) Close() error{
	if c.rdb == nil {
		//return err
	}
	return nil
}

func (c *Cache) Ping() error{
	return nil
}

func (c *Cache) del(key string) error{
	return nil
}

// string
func (c *Cache) set(key string, value string){

}

func (c *Cache) get(key string){

}