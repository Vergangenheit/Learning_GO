// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// // var washPostXML = []byte(`
// // <sitemapindex>
// //    <sitemap>
// //       <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
// //    </sitemap>
// //    <sitemap>
// //       <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
// //    </sitemap>
// //    <sitemap>
// //       <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
// //    </sitemap>
// // </sitemapindex>
// // `)

// type Sitemapindex struct {
// 	Locations []string `xml:"sitemap>loc"`
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

// func main() {
// 	var s Sitemapindex
// 	//	var n News
// 	//	news_map := make(map[string]NewsMap)
// 	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
// 	bytes, _ := ioutil.ReadAll(resp.Body)
// 	xml.Unmarshal(bytes, &s)

// 	for _, Location := range s.Locations {
// 		resp2, _ := http.Get(Location)
// 		fmt.Println(resp2)
// 	}
// }
