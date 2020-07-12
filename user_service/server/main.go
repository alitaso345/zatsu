package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"time"

	pb "github.com/alitaso345/zatsu/user_service/proto"
	"github.com/golang/protobuf/ptypes/empty"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"gopkg.in/gorp.v2"
)

var dbmap *gorp.DbMap

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (service *UserService) CreateUser(ctx context.Context, request *pb.NewUserRequest) (*pb.UserResponse, error) {
	new := newUser(request.Name)
	err := dbmap.Insert(&new)
	errorHandler(err, "Insert failed")

	user := pb.User{Id: 1, Name: request.Name}
	return &pb.UserResponse{User: &user}, nil
}

func (service *UserService) GetUsers(ctx context.Context, empty *empty.Empty) (*pb.UsersResponse, error) {
	var users []User
	_, err := dbmap.Select(&users, "select * from users order by user_id")
	errorHandler(err, "Select failed")

	var pbUsers []*pb.User
	for _, u := range users {
		user := pb.User{Id: u.Id, Name: u.Name}
		pbUsers = append(pbUsers, &user)
	}
	return &pb.UsersResponse{Users: pbUsers}, nil
}

func main() {
	dbmap = initDb()
	defer dbmap.Db.Close()

	lis, err := net.Listen("tcp", ":5051")
	errorHandler(err, "failed to listen")

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &UserService{})
	err = server.Serve(lis)
	errorHandler(err, "failed to serve")
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "./user_db.bin")
	errorHandler(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	errorHandler(err, "Crate table failed")

	return dbmap
}

type User struct {
	Id      int64  `db:"user_id"`
	Name    string `db:",size:50"`
	Created int64
}

func newUser(name string) User {
	return User{
		Name:    name,
		Created: time.Now().UnixNano(),
	}
}

func errorHandler(err error, msg string) {
	if err != nil {
		log.Fatalln(err, msg)
	}
}
