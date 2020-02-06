package main

import(
  pb "./pb"
  "log"
  "net"
  service "./service"
  "google.golang.org/grpc"
)

func main() {
  listenPort, err := net.Listen("tcp", ":19003")
  if err != nil {
    log.Fatalln(err)
  }
  server := grpc.NewServer()
  catService := &service.MyCatService{}

  pb.RegisterCatServer(server, catService)
  server.Serve(listenPort)
}
