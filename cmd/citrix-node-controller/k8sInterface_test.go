package main

import (
	"fmt"
	"k8s.io/klog"
	"runtime"
	"testing"
)

func TestConvertPrefixLenToMask(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println("Current test filename: " + filename)
	t.Log("Prefix Length 24")
	mask := ConvertPrefixLenToMask("24")
	if mask != "255.255.255.0" {
		t.Error("Expected 255.255.255.0, got ", mask)
	}
	mask = ConvertPrefixLenToMask("8")
	if mask != "255.0.0.0" {
		t.Error("Expected 255.0.0.0, got ", mask)
	}
	mask = ConvertPrefixLenToMask("16")
	if mask != "255.255.0.0" {
		t.Error("Expected 255.255.0.0, got ", mask)
	}
	mask = ConvertPrefixLenToMask("30")
	if mask != "255.255.255.252" {
		t.Error("Expected 255.255.255.252, got ", mask)
	}
	mask = ConvertPrefixLenToMask("29")
	if mask != "255.255.255.248" {
		t.Error("Expected 255.255.255.248, got ", mask)
	}
	mask = ConvertPrefixLenToMask("25")
	if mask != "255.255.255.128" {
		t.Error("Expected 255.255.255.128, got ", mask)
	}
	mask = ConvertPrefixLenToMask("19")
	if mask != "255.255.224.0" {
		t.Error("Expected 255.255.224.0, got ", mask)
	}
	mask = ConvertPrefixLenToMask("17")
	if mask != "255.255.128.0" {
		t.Error("Expected 255.255.128.0, got ", mask)
	}
	mask = ConvertPrefixLenToMask("30")
	if mask != "255.255.255.252" {
		t.Error("Expected 255.255.255.252, got ", mask)
	}
}

func TestConfigDecider(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println("Current test filename: " + filename)
	ControllerInputObj := FetchCitrixNodeControllerInput()
	nsobj, api := getClientAndDeviceInfo()
	if err != nil {
		klog.Fatal("K8s Client Error", err, nsobj)
	}
	IngressDeviceClient := createIngressDeviceClient(ControllerInputObj)
	ConfigDecider(api, IngressDeviceClient, ControllerInputObj)
}
func TestGenerateNextPodAddr(t *testing.T) {
	nextIP := GenerateNextPodAddr("10.10.10.10")
     	if (nextIP != "10.10.10.11"){
		t.Error("Expected 10.10.10.11, got ", nextIP)
	}
	nextIP = GenerateNextPodAddr("10.10.10.244")
     	if (nextIP != "10.10.10.245"){
		t.Error("Expected 10.10.10.245, got ", nextIP)
	}
	nextIP = GenerateNextPodAddr("0.0.0.0")
     	if (nextIP != "0.0.0.1"){
		t.Error("Expected 0.0.0.1, got ", nextIP)
	}
	nextIP = GenerateNextPodAddr("10.10.10.300")
     	if (nextIP != "Error"){
		t.Error("Expected Error, got ", nextIP)
	}
}
/*
func TestCreateK8sApiserverClient(t *testing.T){
	func() {
                defer func() {
                        if r := recover(); r == nil {
                                t.Errorf("CreateK8sApiserverClient Input should have panicked!")
                        }
                }()
                // This function should cause a panic
		CreateK8sApiserverClient()
        }()

}
func TestCitrixNodeWatcher(t *testing.T){
	
	controllerInput, api := getClientAndDeviceInfo()
	ingressDevice := createIngressDeviceClient(controllerInput)
	CitrixNodeWatcher(api, ingressDevice, controllerInput)
}*/
/*
func TestCoreHandler(t *testing.T){
	controllerInput, api := getClientAndDeviceInfo()
	ingressDevice := createIngressDeviceClient(controllerInput)

	event := "ADD"
	obj := api.GetDummyNode(controllerInput)	
	newobj := api.GetDummyNode(controllerInput)	
	CoreHandler(api, obj, newobj, event, ingressDevice, controllerInput)
	event = "ADD"	
	CoreHandler(api, obj, newobj, event, ingressDevice, controllerInput)
	event = "ADD"	
	CoreHandler(api, obj, newobj, event, ingressDevice, controllerInput)
	event = "DELETE"	
	CoreHandler(api, obj, newobj, event, ingressDevice, controllerInput)
	event = "DELETE"	
	CoreHandler(api, obj, newobj, event, ingressDevice, controllerInput)
	event = "DELETE"	
	CoreHandler(api, obj, newobj, event, ingressDevice, controllerInput)
	event = "UPDATE"	
	CoreHandler(api, obj, newobj, event, ingressDevice, controllerInput)
	event = "UPDATE"	
	CoreHandler(api, obj, newobj, event, ingressDevice, controllerInput)
	event = "UPDATE"	
	CoreHandler(api, obj, newobj, event, ingressDevice, controllerInput)
}*/
