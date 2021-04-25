package main

import (
	"fmt"
	"mchampaneri/fyne/01/lib"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var data []lib.Feed
var w fyne.Window
var a fyne.App
var err error

func main() {
	a = app.New()
	w = a.NewWindow("GET RSS Feeds")
	w.Resize(fyne.NewSize(800, 400))

	FeedList := container.NewVBox()

	progressbar := dialog.NewCustom("Please wait while loading",
		"Cancel",
		widget.NewProgressBarInfinite(), w)
	progressbar.Hide()

	Scroller := container.NewVScroll(FeedList)
	feedFrom := widget.NewEntry()
	feedFrom.SetPlaceHolder("RSS feed endpoint")

	feedResults := widget.NewLabel("")

	feedFrom.OnSubmitted = func(text string) {
		progressbar.Show()
		FeedList.Objects = nil
		go func() {
			data, err = lib.ReadFeed(feedFrom.Text)
			if err != nil {
				feedResults.SetText(err.Error())
				progressbar.Hide()
				return
			}
			for _, d := range data {
				FeedItem := container.NewVBox(
					widget.NewLabel(d.Title),
					container.NewHBox(
						widget.NewButton("Read more", func() {
							urlObj, _ := url.Parse(d.Link)
							a.OpenURL(urlObj)
						})))
				FeedList.Add(FeedItem)
			}
			feedResults.SetText(fmt.Sprintf("Got %d results",
				len(FeedList.Objects)))
			progressbar.Hide()
			FeedList.Refresh()
		}()
	}

	w.SetContent(
		container.NewBorder(
			container.NewVBox(
				widget.NewForm(widget.NewFormItem("Feed From", feedFrom)),
				container.NewHBox(
					widget.NewButton("Get Feeds", func() {

						progressbar.Show()
						FeedList.Objects = nil
						go func() {
							data, err = lib.ReadFeed(feedFrom.Text)
							if err != nil {
								feedResults.SetText(err.Error())
								progressbar.Hide()
								return
							}
							for _, d := range data {
								FeedItem := container.NewVBox(
									widget.NewLabel(d.Title),
									container.NewHBox(
										widget.NewButton("Read more", func() {
											urlObj, _ := url.Parse(d.Link)
											a.OpenURL(urlObj)
										})))
								FeedList.Add(FeedItem)
							}
							feedResults.SetText(fmt.Sprintf("Got %d results",
								len(FeedList.Objects)))
							progressbar.Hide()
							FeedList.Refresh()
						}()
					}),
					widget.NewButton("Clear Feeds", func() {
						go func() {
							FeedList.Objects = nil
							feedResults.SetText("Feed cleared")
							FeedList.Refresh()
						}()
					}),
				),
			),
			feedResults,
			nil,
			nil,
			Scroller,
		))

	w.ShowAndRun()
}
