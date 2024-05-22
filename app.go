package main

import (
	"context"
	"log"
)

// App struct
type App struct {
	ctx    context.Context
	client *NodeInfo
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.client = &NodeInfo{}
	port, err := GetRandomPort()
	if err != nil {
		log.Fatal(err)
	}
	a.client.MyServerPort = port
	a.client.StartPeerServer(port)
}

// Stop the app
func (a *App) stop(ctx context.Context) {
	if a.client.grpcClient != nil {
		a.client.grpcClient.Close()
	}
	if a.client.grpcServer != nil {
		a.client.grpcServer.GracefulStop()
	}
}

func (a *App) TryConnect(address string, port string) bool {
	log.Println("Trying to connect to:", address+":"+port, "with client:", a.client.grpcClient)
	err := a.client.Connect(address, port)
	if err != nil {
		log.Println(err)
		return false
	}
	pong := a.client.Ping()
	if !pong {
		a.client.grpcClient.Close()
		a.client.grpcClient = nil
	}
	return pong
}

func (a *App) Login(name string, pass string) int32 {
	a.client.Name = name
	a.client.Pass = pass
	if a.client.Login() {
		return int32(a.client.Id)
	} else {
		return -1
	}
}
