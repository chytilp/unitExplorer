package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/chytilp/unitExplorer/command"
	"github.com/chytilp/unitExplorer/common"
	"github.com/chytilp/unitExplorer/persistence"
)

func main() {
	domainsCmd := flag.NewFlagSet("domains", flag.ExitOnError)
	var domainSourceName string
	domainsCmd.StringVar(&domainSourceName, "sourceName", "", "sourceName")
	domainsCmd.StringVar(&domainSourceName, "s", "", "sourceName")

	//var domainIsLive bool
	//domainsCmd.BoolVar(&domainIsLive ,"isLive", false, "isLive")
	//domainsCmd.BoolVar(&domainIsLive ,"l", false, "isLive")
	//fmt.Println("  isLIve:", domainIsLive)

	eventsCmd := flag.NewFlagSet("events", flag.ExitOnError)
	var eventSourceName string
	eventsCmd.StringVar(&eventSourceName, "sourceName", "", "sourceName")
	eventsCmd.StringVar(&eventSourceName, "s", "", "sourceName")
	var eventDomainId string
	eventsCmd.StringVar(&eventDomainId, "domainId", "", "domainId")
	eventsCmd.StringVar(&eventDomainId, "d", "", "domainId")

	marketsCmd := flag.NewFlagSet("markets", flag.ExitOnError)
	var marketSourceName string
	marketsCmd.StringVar(&marketSourceName, "sourceName", "", "sourceName")
	marketsCmd.StringVar(&marketSourceName, "s", "", "sourceName")
	var marketDomainId string
	marketsCmd.StringVar(&marketDomainId, "domainId", "", "domainId")
	marketsCmd.StringVar(&marketDomainId, "d", "", "domainId")
	var marketEventId string
	marketsCmd.StringVar(&marketEventId, "eventId", "", "eventId")
	marketsCmd.StringVar(&marketEventId, "e", "", "eventId")

	if len(os.Args) < 2 {
		fmt.Println("expected 'domains' or 'events' subcommands")
		os.Exit(1)
	}
	appConfig := common.GetConfig()
	err := persistence.CreateDatabase(appConfig.DatabaseFile)
	if err != nil {
		fmt.Printf("CreateDatabase err: %v\n", err)
		os.Exit(1)
	}

	switch os.Args[1] {

	case "domains":
		domainsCmd.Parse(os.Args[2:])
		domainCommmand := command.ListDomains{
			SourceName: domainSourceName,
			Config:     appConfig,
		}
		err = domainCommmand.Run()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

	case "events":
		eventsCmd.Parse(os.Args[2:])
		eventCommmand := command.ListEvents{
			SourceName: eventSourceName,
			Config:     appConfig,
			DomainId:   eventDomainId,
		}
		err = eventCommmand.Run()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

	case "markets":
		marketsCmd.Parse(os.Args[2:])
		fmt.Printf("command=markets, params: sourceName=%s, domainId: %s, eventId: %s",
			marketSourceName, marketDomainId, marketEventId)

	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
