package service

import (
	"VideoAccessServer/internal/config"
	"VideoAccessServer/internal/dao"
	pb "VideoAccessServer/internal/protocol"
	"context"
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"github.com/willf/bloom"
	"time"
)

func RemovalUserHistory(ctx context.Context, request *pb.UserVidListRequest) (*pb.UserVideoListResponse, error) {
	config.Loger.Println("into RemovalUserHistory...")
	now := time.Now()

	config.Loger.Printf("get req:%v", request)
	if len(request.GetUid()) <= 0 {
		config.Loger.Fatalf("request uid empty")
		return nil , errors.New("request uid empty")
	}
	var retVids = make([]int32,0)

	// 先从Redis 里取 该bloom块
	conn := config.GetRedisClientInstance()
	bloomData, err := dao.GetRedisString(conn, request.GetUid())
	if err != nil {
		config.Loger.Fatalf(" get redis key bloom fiter fail:%v", err)
		return nil, errors.New("get resis key fail")
	}
	if len(bloomData) <= 0 {   // 该bloom 不存在时，全量返回
		config.Loger.Println("no bloom block, return")
		return &pb.UserVideoListResponse{
			Uid:     request.GetUid(),
			VideoId: request.VideoId,
		}, nil
	}
	var g bloom.BloomFilter
	err = json.Unmarshal([]byte(bloomData), &g)
	if err != nil {  // 反序列化失败，也返回全量，避免数据全被过滤掉
		config.Loger.Fatalf(" unmarshal bloom fail:%v", err)
		return &pb.UserVideoListResponse{
			Uid:     request.GetUid(),
			VideoId: request.VideoId,
		}, nil
	}
	for vid := range request.GetVideoId() {  //TODO:
		keyExist := g.Test([]byte(cast.ToString(vid)))
		if !keyExist {
			retVids = append(retVids,cast.ToInt32(vid))
		}
	}
	config.Loger.Printf("send response:%v", retVids)
	config.Loger.Println(" send response success, latency：", time.Now().Sub(now))

	return &pb.UserVideoListResponse{
		Uid:     request.GetUid(),
		VideoId: retVids,
	}, nil

	/*
	// test set key
	var strValue string
	for _,value := range request.GetVideoId() {
		strValue += cast.ToString(value)
	}
	config.Loger.Println("get value:", strValue)

	conn := config.GetRedisClientInstance()
	dao.SetRedisString(conn, request.GetUid(), strValue)

	// test mysql
	c := config.GetMysqlClientInstance()
	dao.InsertToDB(c)
	//查询 数据
	dao.QueryDB(c)
	*/
}