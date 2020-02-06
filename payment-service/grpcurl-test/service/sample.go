package service

import(
  "context"
  "errors"
  pb "../pb"
)

type MyCatService struct {}

func (s *MyCatService) GetMyCat(ctx context.Context, request *pb.GetMyCatMessage) (response *pb.MyCatResponse, err error) {
  if request.TargetCat == "tama" {
    return &pb.MyCatResponse{
      Name: "tama",
      Kind: "meinecoon",
    }, nil
  }

  if request.TargetCat == "mike" {
    return &pb.MyCatResponse{
      Name: "mike",
      Kind: "Norwegian Forest Cat",
    }, nil
  }

  return nil, errors.New("Not Found YourCat")
}
