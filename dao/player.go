package dao

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
	"strings"
	"time"
)

func (d *Dao) ZRevRange(addr string, start, end int) (p []string, err error) {
	redisConn := d.Redis.GetRedisClientByAddr(addr).Get()
	defer func() {
		if err := redisConn.Close(); err != nil {
			log.Println("ZRevRange redisConn.Close err:", err)
		}
	}()
	key := "consumer:cache"
	if res, err := redis.Strings(redisConn.Do("ZRevRange", key, start, end-1)); err != nil {
		return p, err
	} else {
		return res, nil
	}
}

func (d *Dao) ZAdd(pid int) (err error) {
	redisConn := d.Redis.GetRedisPool(pid).Get()
	defer func() {
		if err := redisConn.Close(); err != nil {
			log.Println("ZAdd redisConn.Close err:", err)
		}
	}()
	key := "consumer:cache"
	if _, err = redisConn.Do("ZAdd", key, time.Now().Unix(), pid); err != nil {
		return err
	} else {
		return nil
	}
}

func (d *Dao) LRange(addr string, length int) (p []string, err error) {
	redisConn := d.Redis.GetRedisClientByAddr(addr).Get()
	defer func() {
		if err := redisConn.Close(); err != nil {
			log.Println("LRange redisConn.Close err:", err)
		}
	}()
	key := "consumer_list"
	if res, err := redis.Strings(redisConn.Do("lRange", key, 0, length-1)); err != nil {
		return p, err
	} else {
		return res, nil
	}
}

func (d *Dao) LTRIM(addr string, length int) (err error) {
	redisConn := d.Redis.GetRedisClientByAddr(addr).Get()
	defer func() {
		if err := redisConn.Close(); err != nil {
			log.Println("LTRIM redisConn.Close err:", err)
		}
	}()
	key := "consumer_list"
	if _, err := redisConn.Do("LTRIM", key, length, -1); err != nil {
		return err
	}
	return nil
}

func (d *Dao) LPopOne(addr string) (p string, err error) {
	redisConn := d.Redis.GetRedisClientByAddr(addr).Get()
	defer func() {
		if err := redisConn.Close(); err != nil {
			log.Println("LPopOne redisConn.Close err:", err)
		}
	}()
	key := "consumer_list"
	if res, err := redis.String(redisConn.Do("lPop", key)); err != nil {
		return "", err
	} else {
		return res, nil
	}
}

func (d *Dao) GetSinglePlayerList(pid int) (p []string, err error) {
	redisConn := d.Redis.GetRedisPool(pid).Get()
	defer func() {
		if err := redisConn.Close(); err != nil {
			log.Println("GetSinglePlayerList redisConn.Close err:", err)
		}
	}()
	key := fmt.Sprintf("consumer:one:%d", pid)
	log.Println("key:", key)
	if res, err := redis.Strings(redisConn.Do("lRange", key, 0, -1)); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func (d *Dao) UpdateSinglePlayerList(pid int, m map[string]int) (err error) {
	redisConn := d.Redis.GetRedisPool(pid).Get()
	defer func() {
		if err := redisConn.Close(); err != nil {
			log.Println("GetSinglePlayerList redisConn.Close err:", err)
		}
	}()
	key := fmt.Sprintf("consumer:one:%d", pid)
	if _, err := redisConn.Do("multi"); err != nil {
		return err
	}
	for k, num := range m {
		if _, err := redisConn.Do("lRem", key, num, k); err != nil {
			return err
		}
	}
	if _, err := redisConn.Do("exec"); err != nil {
		return err
	}
	return nil
}

func (d *Dao) GetPlayerSettings(pid int) (cid string, close int, err error) {
	redisConn := d.Redis.GetRedisPool(pid).Get()
	defer func() {
		if err := redisConn.Close(); err != nil {
			log.Println("GetPlayerSettings redisConn.Close err:", err)
		}
	}()
	var (
		res map[string]string
	)
	settingKey := fmt.Sprintf("push_setting:%d", pid)
	if res, err = redis.StringMap(redisConn.Do("hGetAll", settingKey)); err != nil {
		return "", 0, err
	}
	if t, ok := res["cid"]; ok {
		cid = t
	}
	if settings, ok := res["settings"]; ok {
		tmp := strings.Split(settings, ",")
		if len(tmp) >= 8 {
			tmp1, err := strconv.Atoi(tmp[7])
			if err != nil {
				return cid, 0, err
			}
			if tmp1 == 0 {
				return cid, 1, nil
			}
		}
		if len(tmp) >= 7 {
			tmp2, err := strconv.Atoi(tmp[6])
			if err != nil {
				return cid, 0, err
			}
			close = tmp2
		}
	}
	return cid, close, nil
}

func (d *Dao) GetAllPlayer() {

}
