package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Chokkar-G/go_grpc_math/mathpb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const port = ":9000"

// main works as the main entry point of the client cli
func main() {
	param1 := flag.Int("Param1", 0, "First parameter")
	param2 := flag.Int("Param2", 0, "Second parameter")
	address := flag.String("Address", "localhost:9000", "Address to the calculator server")
	operation := flag.String("Operation", "", "Operation to execute, one of (add, subtract, multiply, divide)")
	flag.Parse()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	conn, err := grpc.Dial(*address, opts...)
	if err != nil {
		log.Fatal(err)
	}
	// close connection later
	defer conn.Close()
	client := mathpb.NewMathClient(conn)

	switch *operation {
	case "add":

		res, err := client.Add(context.Background(),
			&mathpb.Parameters{Param1: int32(*param1),
				Param2: int32(*param2)})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Result)

	case "subtract":
		res, err := client.Subtract(context.Background(),
			&mathpb.Parameters{Param1: int32(*param1),
				Param2: int32(*param2)})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Result)

	case "multiply":
		res, err := client.Multiply(context.Background(),
			&mathpb.Parameters{Param1: int32(*param1),
				Param2: int32(*param2)})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Result)

	case "divide":
		res, err := client.Divide(context.Background(),
			&mathpb.Parameters{Param1: int32(*param1),
				Param2: int32(*param2)})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Result)

	default:
		fmt.Println("Operation must be one of the following (add, subtract, multiply, divide)")
		os.Exit(1)
	}
}
