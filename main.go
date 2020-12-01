package main

import (
	"VideoAccessServer/internal/config"
	pb "VideoAccessServer/internal/protocol"
	"VideoAccessServer/internal/service"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"time"
)

const (
	port = ":50051"
)

type videoAccessImp struct {}

func init() {
	// check log dir exist
	if _, err := os.Stat(config.LogPath); os.IsNotExist(err) {
		_ = os.Mkdir(config.LogPath, 0777)
		_ = os.Chmod(config.LogPath, 0777)
	}
	file := "./" + config.LogPath + config.APP_NAME + "_" + time.Now().Format("20060102") + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	config.Loger = log.New(logFile, "", log.LstdFlags|log.Lshortfile|log.Ldate|log.Ltime)
}

func (s *videoAccessImp) RemovalUserHistory(ctx context.Context,
	request *pb.UserVidListRequest) (*pb.UserVideoListResponse, error) {
	return service.RemovalUserHistory(ctx, request)
}

func (s *videoAccessImp) SetUserHistory(ctx context.Context,
	request *pb.UserVidListRequest) (*pb.OperResponse, error) {
	return service.SetUserHitory(ctx, request)
}


func main() {
	config.Loger.Println(">>>>>>>>>>>>>service begin>>>>>>>>>>>")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		config.Loger.Fatalf("fail to listen :%v", err)
	}

	s := grpc.NewServer()
	pb.RegisterVideoAccessServer(s, &videoAccessImp{})

	defer config.GetRedisClientInstance().Close()
	defer config.GetMysqlClientInstance().Close()

	// not need
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		config.Loger.Fatalf("fail to server :%v", err)
	}


}
