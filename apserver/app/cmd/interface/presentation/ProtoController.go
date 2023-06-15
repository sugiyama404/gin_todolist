package presentaion

import (
	"context"
	"log"
	"net/http"

	pb "app/cmd/domain/pb"

	"github.com/gin-gonic/gin"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func Greek(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
	})
}
