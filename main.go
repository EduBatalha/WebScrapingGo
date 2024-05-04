package main

import "WebScrapingGo/controllers"

func main() {
	controller := controllers.ScraperController{
		SpreadsheetPath: "urls.xlsx",
	}
	controller.ExecuteScraping()
}
