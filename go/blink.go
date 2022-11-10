package main

import (
  "context"
  "fmt"
  "time"

  "github.com/edaniels/golog"
  "go.viam.com/rdk/robot/client"
  "go.viam.com/rdk/utils"
  "go.viam.com/utils/rpc"
  "go.viam.com/rdk/components/board"
)

func main() {
  logger := golog.NewDevelopmentLogger("client")
  robot, err := client.New(
      context.Background(),
      "[ADD YOUR ROBOT ADDRESS HERE. YOU CAN FIND THIS ON THE CONNECT TAB OF THE VIAM APP]",
      logger,
      client.WithDialOptions(rpc.WithCredentials(rpc.Credentials{
          Type:    utils.CredentialsTypeRobotLocationSecret,
          Payload: "[PLEASE ADD YOUR SECRET HERE. YOU CAN FIND THIS ON THE CONNECT TAB OF THE VIAM APP]",
      })),
  )
  if err != nil {
      logger.Fatal(err)
  }
  defer robot.Close(context.Background())
  logger.Info("Resources:")
  logger.Info(robot.ResourceNames())

  myBoard, err := board.FromRobot(robot, "myBoard")
  if err != nil {
    logger.Fatalf("could not get board: %v", err)
  }

  led, err := myBoard.GPIOPinByName("8")
  if err != nil {
    logger.Fatalf("could not get led: %v", err)
  }

  for {
    err = led.Set(context.Background(), true, nil)
    if err != nil {
      logger.Fatalf("could not set led to on: %v", err)
    }
    fmt.Println("LED is on")

    time.Sleep(1 * time.Second)
    err = led.Set(context.Background(), false, nil)
    if err != nil {
      logger.Fatalf("could not set led to off: %v", err)
    }
    fmt.Println("LED is off")
    time.Sleep(1 * time.Second)
  }

}
