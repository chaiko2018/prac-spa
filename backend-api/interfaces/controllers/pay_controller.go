package controllers

import(
  pb "../../../payment-service/pb"
  "../../interfaces/database"
  "../../domain"

  "google.golang.org/grpc"
)

var addr = "localhost:50051"

func Charge(c Context) {
  t := domain.Payment{}
  c.Bind(&t)
  identifier, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusInternalServerError, err)
  }

  res, err := database.FindById(int64(identifier))
  if err != nil{
    c.JSON(http.StatusInternalServerError, err)
  }

  greq := &pb.PayRequest{
    Id:           int64(identifier),
    Name:         res.Name,
    Description:  res.Description,
    Amount:       res.Amount,
  }

  conn, err := grpc.Dial(addr, grpc.WithInsecure())
  if err != nil{
    c.JSON(http.StatusForbidden, err)
  }
  defer conn.Close()
  client := pb.NewPayManagerClient(conn)

  gres, err := client.Charge(context.Background(), greq)
  if err != nil{
    c.JSON(http.StatusForbidden, err)
  }
  c.JSON(http.StatusOK, gres)
}


