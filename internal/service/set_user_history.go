package service

import (
	"VideoAccessServer/internal/config"
	"VideoAccessServer/internal/dao"
	"VideoAccessServer/internal/error_def"
	pb "VideoAccessServer/internal/protocol"
	"context"
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"github.com/willf/bloom"
	"time"
)

func SetUserHitory(ctx context.Context, request *pb.UserVidListRequest) (*pb.OperResponse, error) {
	config.Loger.Println("into RemovalUserHistory...")
	now := time.Now()

	config.Loger.Printf("get req:%v", request)
	if len(request.GetUid()) <= 0 || len(request.GetVideoId()) <= 0 {
		config.Loger.Printf("request uid or vids empty")
		return &pb.OperResponse{
			RetCode: error_def.CODE_ERR_PARAM_INVALID,
			RetMsg:  "request uid or vids empty",
		} , errors.New("request uid or vids empty")
	}
	// 写入bloomfilter
	// 先从Redis 里取 该bloom块
	conn := config.GetRedisClientInstance()
	bloomData, err := dao.GetRedisString(conn, request.GetUid())
	if err != nil {
		config.Loger.Printf(" get redis key bloom fiter fail:%v", err)
		return nil, errors.New("get resis key fail")
	}
	var data []byte
	if len(bloomData) <= 0 {
		f := bloom.New(1000, 4)
		for vid := range request.GetVideoId() {
			f.Add([]byte(cast.ToString(vid)))
		}
		data, err = json.Marshal(f)
		if err != nil {
			config.Loger.Printf(" marshal data fail:%v", data)
			return &pb.OperResponse{
				RetCode: error_def.CODE_ERR_MARSHAL_FAIL,
				RetMsg:  "marshal data fail",
			}, err
		}
	} else {
		var g bloom.BloomFilter
		err = json.Unmarshal([]byte(bloomData), &g)
		if err != nil {
			config.Loger.Printf(" unmarshal data fail:%v", err)
			return &pb.OperResponse{
				RetCode: error_def.CODE_ERR_UNMARSHAL_FAIL,
				RetMsg:  "unmarshal data fail",
			}, err
		}
		for vid := range request.GetVideoId() {
			g.Add([]byte(cast.ToString(vid)))
		}
		data, err = json.Marshal(g)
		if err != nil {
			config.Loger.Printf(" marshal data fail:%v", data)
			return &pb.OperResponse{
				RetCode: error_def.CODE_ERR_MARSHAL_FAIL,
				RetMsg:  "marshal data fail",
			}, err
		}
	}
	// 写入bloom filter 后，再写入redis
	if err := dao.SetRedisString(conn, request.GetUid(), data); err != nil {
		config.Loger.Printf(" write redis fail:%v", err)
		return &pb.OperResponse{
			RetCode: error_def.CODE_ERR_SET_REDIS_FAIL,
			RetMsg:  "set redis fail",
		}, err
	}

	config.Loger.Println("send response , latency:", time.Now().Sub(now))
	return &pb.OperResponse{
		RetCode: error_def.CODE_SUCCESS,
		RetMsg:  "success",
	}, nil

}