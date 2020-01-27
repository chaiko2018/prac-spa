package main

import(
  "net"
  "log"
  "context"

  pb "../pb"
  payjp "github.com/payjp/payjp-go/v1"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
)

const (
  port = ":50051"
)

type server struct{}

func (s *server) Charge(ctx context.Context, req *pb.PayRequest) (res *pb.PayResponse, err error) {
  pay := payjp.New("sk_test_c62fade9d045b54cd76d7036", nil)

  charge, err := pay.Charge.Create(int(req.Amount), payjp.Charge{
    Currency: "jpy",
    Card: payjp.Card{
      Number:   "4242424242424242",
      CVC:      "123",
      ExpMonth: "2",
      ExpYear:  "2020",
    },
    Capture: true,
    Description: req.Name + ":" + req.Description,
  })
  if err != nil{
    return nil, err
  }

  res = &pb.PayResponse{
    Paid:     charge.Paid,
    Captured: charge.Captured,
    Amount:   int64(charge.Amount),
  }
  return res, nil
}


func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil{
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterPayManagerServer(s, &server{})

  reflection.Register(s)
  log.Printf("gRPC Server started: localhost%s\n", port)
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
