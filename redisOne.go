package main

/* NOW A PUBLIC MODULE !!! */

import(
	//Native
	"fmt"
	
	//3rd party
	"github.com/go-redis/redis"
	
	//Our Packages
	"github.com/pr0con/go_private_modules/procon_config"
)

var redis_client *redis.Client

func init() {
	redis_client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: procon_config.RedisRP,
		DB:       0, 
	});
	
	pong, err := redis_client.Ping().Result()
	if err != nil { fmt.Println("Redis Error: ", err) } else {
		fmt.Println("Redis Connected: ",pong)
	}	
}

func RedisSet(k string, v string) {
	err := redis_client.Set(k, v, 0).Err()
	if err != nil { fmt.Println(err) }
}

func RedisDel(k string) {
	_, err := redis_client.Del(k).Result()
	if err != nil { fmt.Println(err) }
}


type RedisTask struct { //tt = task_type, tk = task key, tv = task value
	tt string
	tk string
	tv string
}

func NewRedisTask(tt string, tk string, tv string)  *RedisTask {
	return  &RedisTask{tt, tk, tv}
}

func (t *RedisTask) Perform() {
	switch(t.tt) {
		case "set-key":
			RedisSet(t.tk, t.tv)
			break;
		case "del-key":
			RedisDel(t.tk)
			break;
		default:
			break;
	}
}

func main() {
	RedisDel("test")
}