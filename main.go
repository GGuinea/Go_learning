package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type response struct {
	StandardDeviation float64 `json:"stddev"`
	NumList           []int   `json:"data"`
}

func getRandom(length int, wg *sync.WaitGroup, responses chan response) int {
	var netClient = &http.Client{}

	resp, err := netClient.Get("https://www.random.org/integers/?num=10&min=1&max=100&col=1&base=10&format=plain&rnd=new")

	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("http request returned code %d\n", resp.StatusCode)
	}

	responses <- handleResponse(resp) // Return error if there is one, nil if not.
	defer wg.Done()
	return 1
}

func handleResponse(resp *http.Response) response {
	log.Println("handle response")
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	split := strings.Split(string(body), "\n")
	t2 := []int{}

	for i := 0; i < len(split)-1; i++ {
		j, _ := strconv.Atoi(split[i])
		t2 = append(t2, j)
	}
	log.Println(split)
	return response{getStdDev(t2), t2}
}

func getStdDev(t2 []int) float64 {
	mean := getMean(t2)
	var standardDeviation float64 = 0
	for _, i := range t2 {
		standardDeviation += math.Pow(float64(i)-mean, 2)
	}
	return math.Sqrt(standardDeviation / float64(len(t2)))
}

func getMean(t2 []int) float64 {
	var mean float64 = 0
	for _, i := range t2 {
		mean += float64(i)
	}
	return mean / float64(len(t2))
}

func serverRequest(w http.ResponseWriter, r *http.Request) {
	var simpleList = []response{}

	requestParam, ok := r.URL.Query()["requests"]
	if !ok || len(requestParam[0]) < 1 {
		log.Println("Problem with URL param 'rquests'")
		return
	}

	lengthParam, ok := r.URL.Query()["length"]
	if !ok || len(lengthParam[0]) < 1 {
		log.Println("Problem with URL 'length' param")
		return
	}

	requestNumber, _ := strconv.Atoi(requestParam[0])
	length, _ := strconv.Atoi(lengthParam[0])
	w.Header().Set("Content-Type", "application/json")
	var wg sync.WaitGroup
	responses := make(chan response, length)
	wg.Add(requestNumber)
	n := 0
	for n < requestNumber {
		go getRandom(length, &wg, responses)
		n++
	}
	wg.Wait()
	n = 0
	close(responses)
	globalDataList := []int{}
	for n <= len(responses) {
		requestDataList := <-responses
		globalDataList = append(globalDataList, requestDataList.NumList...)
		simpleList = append(simpleList, requestDataList)
		n++
	}
	simpleList = append(simpleList, response{getStdDev(globalDataList), globalDataList})

	json.NewEncoder(w).Encode(simpleList)

}

func main() {

	http.HandleFunc("/random/mean", serverRequest)
	http.ListenAndServe(":8080", nil)
}
