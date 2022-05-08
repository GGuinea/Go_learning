package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

func getRandom(length int, wg *sync.WaitGroup, responses chan DataAggregator) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	requestUrl := "https://www.random.org/integers/?num=%d&min=1&max=100&col=1&base=10&format=plain&rnd=new"
	resp, err := netClient.Get(fmt.Sprintf(requestUrl, length))

	if err != nil {
		log.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("http request returned code %d\n", resp.StatusCode)
		defer wg.Done()
		return
	}
	data := handleResponse(resp)
	if data.NumList != nil {
		responses <- data
	}

	defer wg.Done()
}

func handleResponse(resp *http.Response) DataAggregator {
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	split := strings.Split(string(body), "\n")
	numbersTable := []int{}

	for i := 0; i < len(split)-1; i++ {
		j, ok := strconv.Atoi(split[i])
		if ok != nil {
			log.Printf("Problem with number parsing, %s\n", ok)
			return DataAggregator{0, nil}
		}
		numbersTable = append(numbersTable, j)
	}
	return DataAggregator{getStdDev(numbersTable), numbersTable}
}

func handleRequest(c *gin.Context) {
	const requestKey = "requests"
	var numOfRequests, error = getParamFromUrl(c, requestKey)
	if error {
		return
	}
	const lengthKey = "length"
	var numOfNumbers = 0
	numOfNumbers, error = getParamFromUrl(c, lengthKey)
	if error {
		return
	}

	var dataAggregatorList = []DataAggregator{}
	var wg sync.WaitGroup
	responses := make(chan DataAggregator, numOfNumbers)
	wg.Add(numOfRequests)
	n := 0
	for n < numOfRequests {
		go getRandom(numOfNumbers, &wg, responses)
		n++
	}
	wg.Wait()
	close(responses)
	dataAggregatorList = processData(dataAggregatorList, responses)
	c.IndentedJSON(http.StatusOK, dataAggregatorList)
}

func getParamFromUrl(c *gin.Context, keyName string) (int, bool) {
	var value, ok = strconv.Atoi(c.Query(keyName))
	if ok != nil || value <= 0 {
		c.IndentedJSON(http.StatusPreconditionFailed, fmt.Sprintf("Problem with URL '%s' param", keyName))
		return 0, true
	}
	return value, false
}

func processData(dataAggregatorList []DataAggregator, responses chan DataAggregator) []DataAggregator {
	globalDataList := []int{}
	globalDataList, dataAggregatorList = createFinalResponse(responses, globalDataList, dataAggregatorList)
	dataAggregatorList = append(dataAggregatorList, DataAggregator{getStdDev(globalDataList), globalDataList})
	return dataAggregatorList
}

func createFinalResponse(responses chan DataAggregator, globalDataList []int, simpleList []DataAggregator) ([]int, []DataAggregator) {
	n := 0
	for n <= len(responses) {
		requestDataList := <-responses
		globalDataList = append(globalDataList, requestDataList.NumList...)
		simpleList = append(simpleList, requestDataList)
		n++
	}
	return globalDataList, simpleList
}

func main() {
	router := gin.Default()
	router.GET("/random/mean", handleRequest)
	router.Run("0.0.0.0:8080")
}
