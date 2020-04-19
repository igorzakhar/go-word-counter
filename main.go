package main

import (
    "bufio"
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "sync"
)

// структура для хранения результатов каждого запроса
// включая кол-во вхождений строки "Go" в теле ответа
type result struct {
    url   string
    count int
    err   error
}

func main() {
    maxGoroutines := 5
    resultChannel := make(chan *result)
    done := make(chan bool)

    go RunConcurrentTasks(done, resultChannel, maxGoroutines)

    go func() {
        <-done
        close(resultChannel)
    }()

    OutputOfResults(resultChannel)
}

// RunConcurrentTasks запускает параллельные задачи GetWordCount, число одновременно
// выполняемыx задач ограничено величиной передаваемой в аргументе concurrencyLimit
func RunConcurrentTasks(done chan bool, resultChan chan *result, concurrencyLimit int) {

    var waitgroup sync.WaitGroup

    semaphoreChan := make(chan struct{}, concurrencyLimit)

    defer func() {
        close(semaphoreChan)
    }()

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {

        semaphoreChan <- struct{}{}
        waitgroup.Add(1)

        go func(url string) {
            defer waitgroup.Done()

            res := GetWordCount(url)
            resultChan <- &res

            <-semaphoreChan

        }(scanner.Text())
    }
    waitgroup.Wait()
    done <- true
}

// GetWordCount отправляет http get запрос на указанный url-адрес и считает
// кол-во вхождений строки "Go" в теле ответа, возвращает значение типа result
func GetWordCount(url string) result {
    response, err := http.Get(url)
    if err != nil {
        return result{url, 0, err}
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return result{url, 0, err}
    }

    wordCount := bytes.Count(body, []byte("Go"))

    return result{url, wordCount, err}
}

// OutputOfResults слушает канал resultChan, при получении результата вычислений каждой задачи
// GetWordCount выводит его в stdout, аккумулирует значение поля count экземпляра структуры result
// в переменной totalCount и по окончании цикла чтения из канала выводит его в stdout
func OutputOfResults(resultChan chan *result) {
    totalCount := 0

    for res := range resultChan {
        if res.err != nil {
            continue
        }
        fmt.Printf("Count for %s: %d\n", res.url, res.count)

        totalCount += res.count
    }
    fmt.Printf("Total: %v\n", totalCount)
}
