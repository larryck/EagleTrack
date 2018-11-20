// created by tanck, track feature supported by the system

package eagleeye

import (
  "net/http"
  "encoding/json"
)

func Handler() http.Handler {
  return http.HandlerFunc(handleGetFeatures)
}

type SimpleFeatures struct{
  Features []string
}

type FeaturesDetails struct{
  Features []CkFeature
}

func getSimplefeatures(cf []CkFeature) []string {
  var result []string
  for _, f := range cf {
    result = append(result, f.Name)
  }
  return result
}

func handleGetFeatures(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Connection, User-Agent, Cookie")

  if len(r.Form["detail"]) >0 && r.Form["detail"][0]=="true" {
    //get feature details
    jn, _ := json.Marshal(CkFeatureList)
    w.Write(jn)
    return
  }

  features := getSimplefeatures(CkFeatureList)
  jn, _ := json.Marshal(features)
  w.Write(jn)
  return
}
