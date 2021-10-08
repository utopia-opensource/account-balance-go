package main

import (
	"errors"
	"flag"
	"log"
	"strings"

	utopiago "github.com/Sagleft/utopialib-go"
)

type launchFlags struct {
	Host  *string
	Port  *int
	Token *string
}

type balances struct {
	Crypton float64
	UUSD    float64
}

func main() {
	fl, err := parseFlags()
	if err != nil {
		log.Fatalln(err)
	}

	client, err := openConnection(fl)
	if err != nil {
		log.Fatalln(err)
	}

	result, err := getBalances(client)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(result)
}

func parseFlags() (*launchFlags, error) {
	fl := launchFlags{}
	fl.Host = flag.String("host", "http://127.0.0.1", "Utopia client host")
	fl.Port = flag.Int("port", 20000, "client port")
	fl.Token = flag.String("token", "", "client account token")
	flag.Parse()

	if fl.Host == nil || fl.Port == nil || fl.Token == nil {
		return nil, errors.New("failed to parse flags")
	}
	return &fl, nil
}

func openConnection(fl *launchFlags) (*utopiago.UtopiaClient, error) {
	host := fl.Host
	hostParts := strings.Split(*host, "://")

	client := utopiago.UtopiaClient{
		Protocol: hostParts[0],
		Token:    *fl.Token,
		Host:     hostParts[1],
		Port:     *fl.Port,
	}
	isConnected := client.CheckClientConnection()
	if !isConnected {
		return nil, errors.New("failed to connect to client")
	}

	return &client, nil
}

func getBalances(client *utopiago.UtopiaClient) (*balances, error) {
	result := balances{}

	var err error
	result.Crypton, err = client.GetBalance()
	if err != nil {
		return nil, err
	}

	result.UUSD, err = client.GetUUSDBalance()
	if err != nil {
		return nil, err
	}

	return &result, nil
}
