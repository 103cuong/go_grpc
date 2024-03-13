package main

import (
	"context"
	pb "github.com/cuongnd9/go-grpc/pkg/pb"
	"google.golang.org/grpc"
	"log"
	"syreclabs.com/go/faker"
	"time"
)

const apiVersion = "pb"

func main() {
	conn, err := grpc.Dial(":50000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection gRPC failed: %v", err)
	}
	defer conn.Close()

	client := pb.NewToDoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	t.Format(time.RFC3339Nano)

	// Create
	req1 := pb.CreateRequest{
		Api: apiVersion,
		ToDo: &pb.ToDo{
			Title:       faker.Name().Title(),
			Description: faker.Name().String(),
		},
	}
	res1, err := client.Create(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create ToDo: <%+v>", &res1)

	id := res1.Id

	// Read
	req2 := pb.ReadRequest{
		Api: apiVersion,
		Id:  id,
	}
	res2, err := client.Read(ctx, &req2)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}
	log.Printf("Read result: <%+v>\n\n", res2)

	// Update
	req3 := pb.UpdateRequest{
		Api: apiVersion,
		ToDo: &pb.ToDo{
			Id:          res2.ToDo.Id,
			Title:       faker.Name().Title(),
			Description: faker.Name().String(),
		},
	}
	res3, err := client.Update(ctx, &req3)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", res3)

	// ReadAll
	req4 := pb.ReadAllRequest{
		Api: apiVersion,
	}
	res4, err := client.ReadAll(ctx, &req4)
	if err != nil {
		log.Fatalf("ReadAll failed: %v", err)
	}
	log.Printf("ReadAll result: <%+v>\n\n", res4)

	// Delete
	req5 := pb.DeleteRequest{
		Api: apiVersion,
		Id:  id,
	}
	res5, err := client.Delete(ctx, &req5)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", res5)
}
