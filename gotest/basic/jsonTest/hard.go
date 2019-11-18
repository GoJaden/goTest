package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const url string = "https://itunes.apple.com/cn/lookup?id=1247810530"

func main() {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	//contentLength , _:=strconv.Atoi(resp.Header.Get("content-length"))
	b, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("文件", string(b))

	/*f , err := os.Open("1.txt")
	defer f.Close()
	if err != nil{
		fmt.Println(err)
	}
	buf := make([]byte,getFileSize("1.txt"))
	f.Read(buf)
	fmt.Println("文件：")
	fmt.Println(string(buf))
	v :=new(Version)
	err1 := json.Unmarshal(buf,v)
	if err != nil{
		fmt.Println("格式化异常:",err1)
	}

	fmt.Println("版本：",v.Results[0].Version)*/
}

//获取文件大小
func getFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

type Version struct {
	ResultCount int       `json:"resultCount"`
	Results     []Results `json:"results"`
}
type Results struct {
	IsGameCenterEnabled                bool          `json:"isGameCenterEnabled"`
	ScreenshotUrls                     []string      `json:"screenshotUrls"`
	IpadScreenshotUrls                 []interface{} `json:"ipadScreenshotUrls"`
	AppletvScreenshotUrls              []interface{} `json:"appletvScreenshotUrls"`
	ArtworkURL60                       string        `json:"artworkUrl60"`
	ArtworkURL512                      string        `json:"artworkUrl512"`
	ArtworkURL100                      string        `json:"artworkUrl100"`
	ArtistViewURL                      string        `json:"artistViewUrl"`
	SupportedDevices                   []string      `json:"supportedDevices"`
	Advisories                         []interface{} `json:"advisories"`
	Kind                               string        `json:"kind"`
	Features                           []interface{} `json:"features"`
	TrackCensoredName                  string        `json:"trackCensoredName"`
	LanguageCodesISO2A                 []string      `json:"languageCodesISO2A"`
	FileSizeBytes                      string        `json:"fileSizeBytes"`
	ContentAdvisoryRating              string        `json:"contentAdvisoryRating"`
	AverageUserRatingForCurrentVersion int           `json:"averageUserRatingForCurrentVersion"`
	UserRatingCountForCurrentVersion   int           `json:"userRatingCountForCurrentVersion"`
	TrackViewURL                       string        `json:"trackViewUrl"`
	TrackContentRating                 string        `json:"trackContentRating"`
	MinimumOsVersion                   string        `json:"minimumOsVersion"`
	CurrentVersionReleaseDate          time.Time     `json:"currentVersionReleaseDate"`
	ReleaseNotes                       string        `json:"releaseNotes"`
	TrackID                            int           `json:"trackId"`
	TrackName                          string        `json:"trackName"`
	SellerName                         string        `json:"sellerName"`
	PrimaryGenreID                     int           `json:"primaryGenreId"`
	IsVppDeviceBasedLicensingEnabled   bool          `json:"isVppDeviceBasedLicensingEnabled"`
	FormattedPrice                     string        `json:"formattedPrice"`
	ReleaseDate                        time.Time     `json:"releaseDate"`
	PrimaryGenreName                   string        `json:"primaryGenreName"`
	GenreIds                           []string      `json:"genreIds"`
	Currency                           string        `json:"currency"`
	Version                            string        `json:"version"`
	WrapperType                        string        `json:"wrapperType"`
	ArtistID                           int           `json:"artistId"`
	ArtistName                         string        `json:"artistName"`
	Genres                             []string      `json:"genres"`
	Price                              int           `json:"price"`
	Description                        string        `json:"description"`
	BundleID                           string        `json:"bundleId"`
	AverageUserRating                  int           `json:"averageUserRating"`
	UserRatingCount                    int           `json:"userRatingCount"`
}
