package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/prashantkumardagur/grpc-chat/chat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//==============================================================================

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Input(str *string) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", 1)
	*str = text
}

//==============================================================================

func main() {
	// Create a gRPC connection with the server
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	HandleError(err)
	defer conn.Close()

	// Create a gRPC client
	client := pb.NewChatClient(conn)

	// ============================================================================

	fmt.Println("BOT> Welcome to gRPC Chat")
	var username string

	for {
		fmt.Print("BOT> Enter new username: ")
		Input(&username)

		res, err := client.CheckUser(context.Background(), &pb.User{Username: username})
		HandleError(err)

		if res.GetSuccess() {
			fmt.Println("BOT> User already exists")
		} else {
			fmt.Println("BOT> User logged in")
			break
		}
	}

	stream, err := client.Messaging(context.Background())
	HandleError(err)
	Chat(client, stream, username)

}
