package orders

func main() {
    grpcServer := NewGRPCServer(":9000")
    grpcServer.Run()
}
