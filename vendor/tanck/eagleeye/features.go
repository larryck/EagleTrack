// created by tanck, track feature supported by the system
package eagleeye

type CkFeature struct {
  Author string
  Name string
  Date string
  Desc string
}

var CkFeatureList = []CkFeature { {Author: "tanck1", Name: "memoryUsageWithoutCache", Date: "20180606", Desc: "add the memory usage without cache metric"}, {Author: "tanck1", Name: "1000/10000NetworkRates", Date: "20180606", Desc: "split network usages of 1000 and 10000Mb devices apart"}, {Author: "tanck1", Name: "customLongTermMetrics", Date: "20180606", Desc: "add MemoryUsageWithoutCache,1000/10000NetworkRates to longterm store"}, /*#FEATURE_MAGIC#*/}


