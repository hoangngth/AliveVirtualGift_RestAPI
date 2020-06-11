package connection

import (
	"AliveVirtualGift_RestAPI/src/proto"
	"os"

	"google.golang.org/grpc"
)

//ServiceConnection ...
type ServiceConnection struct {
	ClientSessionService proto.SessionServiceClient
	ClientAccountService proto.AccountServiceClient
}

//GetSessionServiceClient ...
func (sc *ServiceConnection) GetSessionServiceClient() proto.SessionServiceClient {
	if sc != nil {
		return sc.ClientSessionService
	}
	return nil
}

//GetAccountServiceClient ...
func (sc *ServiceConnection) GetAccountServiceClient() proto.AccountServiceClient {
	if sc != nil {
		return sc.ClientAccountService
	}
	return nil
}

//DialToSessionServiceServer ...
func DialToSessionServiceServer() *ServiceConnection {

	port := os.Getenv("SESSION_SERVICE_PORT")

	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &ServiceConnection{
		ClientSessionService: proto.NewSessionServiceClient(conn),
	}
}

//DialToAccountServiceServer ...
func DialToAccountServiceServer() *ServiceConnection {

	port := os.Getenv("ACCOUNT_SERVICE_PORT")

	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return &ServiceConnection{
		ClientAccountService: proto.NewAccountServiceClient(conn),
	}
}
