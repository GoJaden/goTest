package main

import (
	"encoding/json"
	"fmt"
)

type DashboardInFolder []struct {
	ID          int           `json:"id"`
	UID         string        `json:"uid"`
	Title       string        `json:"title"`
	URI         string        `json:"uri"`
	URL         string        `json:"url"`
	Slug        string        `json:"slug"`
	Type        string        `json:"type"`
	Tags        []interface{} `json:"tags"`
	IsStarred   bool          `json:"isStarred"`
	FolderID    int           `json:"folderId"`
	FolderUID   string        `json:"folderUid"`
	FolderTitle string        `json:"folderTitle"`
	FolderURL   string        `json:"folderUrl"`
}

func main() {

	tt := `[{"id":224,"uid":"RXn3fVRGk","title":"sdyxmall-api-auto","uri":"db/sdyxmall-api-auto","url":"/d/RXn3fVRGk/sdyxmall-api-auto","slug":"","type":"dash-db","tags":[],"isStarred":false,"folderId":223,"folderUid":"Hh3uB4gMk","folderTitle":"sdyxmall-auto","folderUrl":"/dashboards/f/Hh3uB4gMk/sdyxmall-auto"},{"id":225,"uid":"U14gL4RMk","title":"sdyxmall-pod-auto","uri":"db/sdyxmall-pod-auto","url":"/d/U14gL4RMk/sdyxmall-pod-auto","slug":"","type":"dash-db","tags":[],"isStarred":false,"folderId":223,"folderUid":"Hh3uB4gMk","folderTitle":"sdyxmall-auto","folderUrl":"/dashboards/f/Hh3uB4gMk/sdyxmall-auto"}]`

	var data = new(DashboardInFolder)
	//res := new(DashboardInFolder)
	json.Unmarshal([]byte(tt), data)

	fmt.Println((*data)[0])

}
