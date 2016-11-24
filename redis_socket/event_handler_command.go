package redis_socket

import (
	"github.com/garyburd/redigo/redis"
	"github.com/giskook/charging_pile_tss/pb"
	"github.com/golang/protobuf/proto"
	"log"
)

func (socket *RedisSocket) ProcessChargingPile() {
	conn := socket.GetConn()
	defer conn.Close()

	var index int = 0
	var pkg *Report.ChargingPileStatus
	for index, pkg = range socket.ChargingPiles {
		conn.Send("GET", pkg.Cpid)
	}

	conn.Flush()

	tobe_commit_cps := make([]*Report.ChargingPileStatus, 1024)
	for i := 0; i < index+1; i++ {
		v_redis, err := conn.Receive()

		if err != nil {
			continue
		}

		v, _ := redis.Bytes(v_redis, nil)

		redis_cps := &Report.ChargingPileStatus{}
		err = proto.Unmarshal(v, redis_cps)
		if err != nil {
			log.Println("unmarshal error")
		} else {
			if !proto.Equal(redis_cps, socket.ChargingPiles[i]) {
				if redis_cps.Timestamp < socket.ChargingPiles[i].Timestamp {
					tobe_commit_cps = append(tobe_commit_cps, socket.ChargingPiles[i])
				}

			}
		}
	}

	for _, new_pkg := range tobe_commit_cps {
		data, _ := proto.Marshal(new_pkg)
		conn.Send("SET", new_pkg.Cpid, data)
	}

	conn.Flush()
	conn.Do("EXEC")
}