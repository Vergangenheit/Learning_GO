package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// var washPostXML = []byte(`
// <sitemapindex>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
//    </sitemap>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
//    </sitemap>
//    <sitemap>
//       <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
//    </sitemap>
// </sitemapindex>
// `)

// type Sitemapindex struct {
// 	Locations []Location `xml:"sitemap"`
// }

// type Location struct {
// 	Loc string `xml:"loc"`
// }

// type News struct {
// 	Titles    []string `xml:"url>news>title"`
// 	Keywords  []string `xml:"url>news>keywords"`
// 	Locations []string `xml:"url>loc"`
// }

// type NewsMap struct {
// 	Keyword  string
// 	Location string
// }

func main2() {
	var s Sitemapindex
	var n News
	xml.Unmarshal(washPostXML, &s)
	news_map := make(map[string]NewsMap)

	for _, Location := range s.Locations {
		resp, _ := http.Get(Location.Loc)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for idx, _ := range n.Keywords {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}

	}
	// for idx, data := range news_map {
	// 	fmt.Println("\n\n\n\n", idx)
	// 	fmt.Println("\n", data.Keyword)
	// 	fmt.Println("\n", data.Location)
	// }

}
