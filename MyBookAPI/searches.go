package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "strconv"
    "errors"
    "reflect"
)

var BooksCollections []BookStructure

type BookStructure struct {
    BookID        string  `json:"book_id"`
    Title         string  `json:"title"`
    Author        string  `json:"author"`
    AverageRating float64 `json: "average_rating"`
    ISBN          string  `json: "isbn"`
    ISBN13        string  `json: "isbn13" `
    LanguageCode  string  `json: "language_code"`
    NumofPages    int     `json: "num_of_pages"`
    Ratings       int     `json: "ratings"`
    Reviews       int     `json: "reviews"`
}



func TotalBookCollections(records [][]string) []BookStructure {
    for _, eachRecord := range records {
        averageRating, _ := strconv.ParseFloat(eachRecord[3], 64)
        numofPages, _ := strconv.Atoi(eachRecord[7])
        ratings, _ := strconv.Atoi(eachRecord[8])
        reviews, _ := strconv.Atoi(eachRecord[9])

        book := BookStructure{BookID: eachRecord[0], Title: eachRecord[1],
            Author: eachRecord[2], AverageRating: averageRating,
            ISBN: eachRecord[4], ISBN13: eachRecord[5],
            LanguageCode: eachRecord[6], NumofPages: numofPages,
            Ratings: ratings, Reviews: reviews}

        BooksCollections = append(BooksCollections, book)
    }
    return BooksCollections

}
func ReadCsv() []BookStructure {
    csvFilePath := "C:\\Users\\kamalsai\\Desktop\\My practices\\MyBookAPI\\BooksDB\\totalBooks.csv"
    fmt.Println("Given File Path ::", csvFilePath)
    csvFile, err := os.Open(csvFilePath)
    //fmt.Println(reflect.TypeOf(csvFile))
    if err != nil {
        log.Fatalln(err)
    }
    csvReaderObj := csv.NewReader(csvFile)
    records, err := csvReaderObj.ReadAll()
    if err != nil {
        log.Fatalln(err)
    }
    //fmt.Println(len(records), reflect.TypeOf(records), reflect.ValueOf(records).Kind())
    return TotalBookCollections(records)
}



func SearchByISBN(isbn string){
    found :=false
    getBookCollections := ReadCsv()
    for i:=0;i<=len(getBookCollections)-1;i++{
        if getBookCollections[i].ISBN == isbn{
            found = true
            fmt.Println(getBookCollections[i])
        }
    }
    if found != true{
       fmt.Println("ISBN received in the request is NOT Found :: ", isbn)
    }
}

func SearchByISBN13(isbn13 string){
    found :=false
    getBookCollections := ReadCsv()
    for i:=0;i<=len(getBookCollections)-1;i++{
        if getBookCollections[i].ISBN13 == isbn13{
            found = true
            fmt.Println(getBookCollections[i])
        }
    }
    if found != true{
       fmt.Println("SearchByISBN13 received in the request is NOT Found :: ", isbn13)
    }
}

func SearchByBookID(bookId string)(BookStructure, error){
    getBookCollections := ReadCsv()
    for i:=0;i<=len(getBookCollections)-1;i++{
        if getBookCollections[i].BookID == bookId{
            return getBookCollections[i], nil
        }
    }
    return BookStructure{}, errors.New("BookID received in the request is NOT Found")
}

type errResponse struct{
        //Timestamp time.Time `json:"timestamp"`
        //Status int `json:"status"`
        ErrorType string `json:"errorType"`
        //Path string `json:"path"`
}

func MarshalData(dataobj interface{}){
    fmt.Println(dataobj, reflect.TypeOf(dataobj), reflect.TypeOf(dataobj).Kind())
    data, _ := json.Marshal(dataobj)
    fmt.Println(string(data), reflect.TypeOf(data), reflect.TypeOf(data).Kind())
}


func main(){
    a,b := SearchByBookID("56")
    if b != nil{
        errObj := errResponse{ErrorType: "Not Found"}
        MarshalData(errObj)
    }else{
        MarshalData(a)
    }
}