package getBookData

import (
        "log"
        "reflect"
        "getCsvData/getCsvData"
    )

func GetBookDetails(bookId string)(string){
    getBookCollections := getCsvData.ReadCsv()
    log.Println(reflect.TypeOf(getBookCollections))
    log.Printf("---(GetBookDetails)---- Received Book Id :: ", bookId)
    return bookId
}


