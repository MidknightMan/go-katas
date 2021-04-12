package kata

import (
	"fmt"
	"strconv"
	"strings"
	//   "sort"
)


func StockList(listArt []string, listCat []string) string {
  
  var stocklist = make(map[string]int)
  
    for cat:=0; cat<len(listCat); cat++ {
      for art:=0; art<len(listArt); art++ {
        if strings.HasPrefix(fmt.Sprint(strings.Fields(listArt[art])[0]), listCat[cat])  {
          num, _ := strconv.Atoi(strings.Fields(listArt[art])[1])
          fmt.Println(strings.Fields(listArt[art])[1])
          stocklist[listCat[cat]] += num
        } else {
          stocklist[listCat[cat]] += 0
        }
      }
    }
  
  var trial string
  
  for i:=0;i<len(stocklist);i++ {
    fmt.Println(stocklist[i])
    trial += "here"
  }
  
  fmt.Println(trial)
  
  var z string
    
  keys := make([]string, 0)
    for k, _ := range stocklist {
        keys = append(keys, k)
      }
  fmt.Println(keys)
//   sort.Strings(keys)
  for _, k := range keys {
      z += fmt.Sprintf("(%s : %d) - ", k, stocklist[k])
  }
  
  
  return strings.TrimSuffix(z, " - ")
}