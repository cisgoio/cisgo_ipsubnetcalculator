package main

import (
	"cisgo_ipsubnetcalculator/ipsubnet"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

//import "github.com/brotherpowers/ipsubnet"

func init() {
	//==Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	//==Output to stdout instead of the default stderr
	//==Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	//==Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
	//log.SetLevel(log.InfoLevel)
	//log.SetLevel(log.WarnLevel)

	/*
		logflag.Parse() // Call after regular flag.Parse()
		logrus.Debug("Debug log")
		logrus.Info("Info log")
		logrus.Warn("Warn log")
		logrus.Error("Error log")
		logrus.Fatal("Fatal log")
	*/
}

func main() {
	fmt.Println("...in main()")

	/*
		log.WithFields(log.Fields{
			"animal": "walrus-Debug",
			"size":   10,
		}).Debug("...walrus-Debug")

		log.WithFields(log.Fields{
			"animal": "walrus-Info",
			"size":   10,
		}).Info("...walrus-Info")
	*/

	myIP := "192.168.112.203"
	mySubNetBlock := 23
	sub := ipsubnet.SubnetCalculator(myIP, mySubNetBlock)
	log.WithFields(log.Fields{
		"myIP = ":                myIP,
		"mySubNetBlock":          mySubNetBlock,
		"nr. of IP Addresses = ": sub.GetNumberIPAddresses(),
	}).Debug(".. INPUT(myIP,mySubNetBlock)")

	fmt.Println("======= ANS: ")
	fmt.Println(sub)
	fmt.Println(sub.GetNumberIPAddresses())

	//fmt.Println(ipsubnet.GetValue())

}
