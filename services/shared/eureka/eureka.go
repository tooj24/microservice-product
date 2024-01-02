package eureka

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

const (
	EurekaUri = "http://localhost:8761/eureka/v2/apps/"
	AppId     = "product-service"
)

type DataCenterInfo struct {
	XMLName xml.Name `xml:"dataCenterInfo"`
	Name    string   `xml:"name"`
}

type Instance struct {
	XMLName       xml.Name       `xml:"instance"`
	InstanceId    string         `xml:"instanceId"`
	HostName      string         `xml:"hostName"`
	App           string         `xml:"app"`
	IPAddr        string         `xml:"ipAddr"`
	Port          int            `xml:"port"`
	VipAddr       string         `xml:"vipAddress"`
	SecureVipAddr string         `xml:"secureVipAddress"`
	Status        string         `xml:"status"`
	HomePageURL   string         `xml:"homePageUrl"`
	DataCenter    DataCenterInfo `xml:"dataCenterInfo"`
}

type EurekaInfo struct {
	InstanceID string
	Port       int
}

type EurekaService interface {
	Register(serviceName string, host string, port int) EurekaInfo
	Unregister(serviceName string)
}

type EurekaServiceImpl struct {
}

func NewEurekaServiceImpl() EurekaService {
	return &EurekaServiceImpl{}
}

func (service *EurekaServiceImpl) Register(serviceName string, host string, port int) EurekaInfo {
	println("======== EUREKA registration ========")
	instance := Instance{
		InstanceId:    fmt.Sprintf("%s:%s:%d", host, serviceName, port),
		HostName:      host,
		App:           serviceName,
		IPAddr:        host,
		Port:          port,
		VipAddr:       serviceName,
		SecureVipAddr: serviceName,
		Status:        "UP",
		HomePageURL:   fmt.Sprintf("http://%s:%d", host, port),
		DataCenter: DataCenterInfo{
			Name: "MyOwn",
		},
	}

	xmlPayload, err := xml.Marshal(instance)
	HandleError("Error encoding XML :", err)

	SendHttpRequest("POST", EurekaUri+serviceName, bytes.NewBuffer(xmlPayload))

	return EurekaInfo{
		InstanceID: instance.InstanceId,
		Port:       instance.Port,
	}
}

func (service *EurekaServiceImpl) Unregister(serviceName string) {
	println("======== EUREKA unregistration ========")
	url := fmt.Sprintf("%s/%s/%s", EurekaUri, AppId, serviceName)
	SendHttpRequest("DELETE", url, nil)
}

func HandleError(message string, err error) {
	if err != nil {
		fmt.Println(fmt.Sprintf("[EUREKA] %s", message), err)
		return
	}
}

func SendHttpRequest(method string, url string, body io.Reader) {
	req, err := http.NewRequest(method, url, body)
	HandleError("Error creating request", err)

	encodedBasic := base64.StdEncoding.EncodeToString([]byte("eureka:password"))
	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedBasic))

	client := &http.Client{}
	resp, err := client.Do(req)
	HandleError("Error sending request: ", err)

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("[EUREKA] Operation finished")
	} else {
		fmt.Println("[EUREKA] Failed failed \n Status code: ", resp.StatusCode)
	}
}

func RandomPort() int {
	maxValue, minValue := 1000, 6000
	return 3000 + rand.Intn((maxValue-minValue)+minValue)
}
