// main.go
package main

import ( 
    "log" 
    "fmt" 
    "net/http" 
    "html/template" 
    "encoding/csv" 
    "os" 
)

type PageVariables struct { 
    user string 
} f
unc mainPage(w http.ResponseWriter, r *http.Request){

    var data = [][]string{{"Name","Bankid","Country"},{}}
    if r.Method == "POST" {

    r.ParseForm()

    file, err := os.Create("result.csv") // store in result.csv
    if err != nil {
            log.Fatal("Erro creating csv file!", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _ , val  := range data {

    err := writer.Write(val)
      if err != nil {
              log.Fatal("Error Writing to CSV!", err)
      }

    }

    // To-Do's 
    s := []string{} 
    s = append(s, r.FormValue("username"))
    s = append(s, r.FormValue("id"))
    s = append(s, r.FormValue("country"))
    s = append(s, r.FormValue("department"))
    fmt.Println(s)
    for _, value := range r.Form["new_data"] {
       fmt.Println(value)
       s = append(s, value)

       fmt.Println(s)
       writer.Write(s)
    }
    }


    t, err := template.ParseFiles("index.html")
    if err != nil {
      log.Println("Template parsing error:", err)
    }

    err = t.Execute(w, "No data for now")
    if err != nil {
            log.Fatal("Template execution error", err)
    }
} func main() { http.HandleFunc("/", mainPage) http.ListenAndServe(":9111", nil)

}

func checkError(message string, err error) {

    if err != nil {
    log.Fatal(message, err)
    }
}
