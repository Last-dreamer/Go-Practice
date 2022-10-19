package main

import "fmt"

const (
	Online = 0
	Offline = 1
	Maintenance = 2
	Retired = 3
)

func printServerStatus(servers map[string]int){
	fmt.Println("\nThere are: ", len(servers), "Servers")

	stats := make(map[int]int)

	for _, status := range servers {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1 
		case Maintenance:
			stats[Maintenance] += 1 
		case Retired:
			stats[Retired] += 1 
		default:
			panic("unhandled server status")
		}	
	}

	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are offline")
	fmt.Println(stats[Maintenance], "servers are under maintenance")
	fmt.Println(stats[Retired], "servers are retired")


}

func main(){

	servers := []string{"barikot", "Peshawar", "Swat", "Karachi"}

	serverStatus := make(map[string]int)

	for _, server:= range servers {
		serverStatus[server] = Online
	}

	printServerStatus(serverStatus)

	serverStatus["barikot"] = Retired
	serverStatus["Swat"] = Offline

	printServerStatus(serverStatus)

	for server := range serverStatus {
		serverStatus[server] = Maintenance
	}

	printServerStatus(serverStatus)
}