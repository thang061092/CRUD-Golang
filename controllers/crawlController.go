package controllers

import (
	"api/helper"
	"log"
	"net/http"
	"net/url"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/gocolly/colly"
)

type CrawlController struct {
	beego.Controller
}

func (cr *CrawlController) GetData() {
	// create a new collector
	c := colly.NewCollector()

	// authenticate
	err := c.Post("http://cpanelndt.tienngay.vn/login", map[string]string{"email": "thangbm@tienngay.vn", "password": "thangvip01"})
	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
		setCookie := r.Headers.Get("Set-Cookie")
		if setCookie != "" {
			u, err := url.Parse("http://cpanelndt.tienngay.vn/login")
			if err != nil {
				log.Fatal(err)
			}
			cookies := strings.Split(setCookie, "; ")
			for _, cookie := range cookies {
				c.SetCookies(u.String(), []*http.Cookie{{
					Name:  strings.Split(cookie, "=")[0],
					Value: strings.Split(cookie, "=")[1],
				}})
			}
			r.Request.Visit("http://cpanelndt.tienngay.vn/transaction/process")
		}
	})

	var result [][]string
	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			row := []string{}
			el.ForEach("td", func(_ int, td *colly.HTMLElement) {
				row = append(row, td.Text)
			})
			result = append(result, row)
		})
	})

	helper.SendResponse(&cr.Controller, helper.HTTP_OK, helper.SUCCESS, result)

	// start scraping
	c.Visit("http://cpanelndt.tienngay.vn")
}
