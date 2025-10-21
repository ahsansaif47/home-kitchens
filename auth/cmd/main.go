package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ahsansaif47/home-kitchens/auth/config"
	"github.com/ahsansaif47/home-kitchens/auth/http/routes"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"

	pb "github.com/ahsansaif47/home-kitchens/common/gRPC/generated"
)

func main() {

	go func() {
		startGRPC()
	}()

	startHTTP()
}

func startHTTP() {
	app := fiber.New()
	routes.InitRoutes(app)

	port := config.GetConfig().GlobalCfg.Port
	log.Printf("Fiber server listening on port: %s", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))

}

func startGRPC() {

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("error starting gRPC server: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEmailServiceServer(grpcServer, nil)

	log.Println("gRPC server started on port 50052")

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("Failed to start gROC server: %w", err)
	}
}
