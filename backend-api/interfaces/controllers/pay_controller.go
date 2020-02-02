package controllers

import(
  pb "../../../payment-service/proto"
  "../../interfaces/database"
  "../../domain"

  "google.golang.org/grpc"
)

var addr := "localhost:50051"

func Charge(c Context) {
  t := domain.Payment{}
