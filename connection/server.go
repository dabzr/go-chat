package main

import (
	"fmt"
	"net"
	"os"
  "net/http"
  "io"
)

func verifyError(err error){
  if(err != nil){
    fmt.Println(err)
    os.Exit(3)
  }
}
func getPrivateIPV4(){
    conn, err := net.Dial("udp", "8.8.8.8:80")
    verifyError(err)
    defer conn.Close()

    localAddr := conn.LocalAddr().(*net.UDPAddr)

    ipv4 := localAddr.IP.String()
    fmt.Println("Endereço IPv4 privado da máquina:", ipv4)
}

func getPublicIPV4(){
  url := "https://api.ipify.org?format=text"
  resp, err := http.Get(url)
  if err != nil {
      fmt.Println("Erro ao fazer a solicitação:", err)
      return
  }
  defer resp.Body.Close()

  ip, err := io.ReadAll(resp.Body)
  if err != nil {
      fmt.Println("Erro ao ler a resposta:", err)
      return
  }

  fmt.Printf("Seu IP público é: %s\n", ip)
}

func startServer() {
  fmt.Println("Aguardando Conexão")
  ln, err := net.Listen("tcp", ":5000")
  verifyError(err)
  getPrivateIPV4()
  getPublicIPV4()
  connection, err := ln.Accept()
  verifyError(err)
  defer connection.Close()
  defer ln.Close()

} 

func main() {
  startServer()
}
