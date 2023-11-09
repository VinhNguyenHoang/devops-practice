package collectors

import (
	"log"

	"cs/internal/tom/models"

	"github.com/gocolly/colly/v2"
)

type CellphonesCollector struct {
	BaseURL string
	c       *colly.Collector
	Phones  []*models.Phone
}

func (clr *CellphonesCollector) initLogic() {
	// simple logic
	clr.c.OnHTML("a.product__link.button__link", func(h *colly.HTMLElement) {
		name := h.ChildText("div.product__name>h3")
		price := h.ChildText("p.product__price--show")

		phone := &models.Phone{
			Name:  name,
			Price: price,
		}

		clr.Phones = append(clr.Phones, phone)
	})

	clr.c.OnError(func(r *colly.Response, err error) {
		log.Println(err)
	})
}

func (clr *CellphonesCollector) RunCollect() error {
	return clr.c.Visit(clr.BaseURL)
}

func (clr *CellphonesCollector) GetCollection() []*models.Phone {
	return clr.Phones
}

func NewCellphonesCollector(baseURL string) Collector {
	col := colly.NewCollector()

	collector := &CellphonesCollector{
		c: col,
	}

	if baseURL == "" {
		collector.BaseURL = "https://cellphones.com.vn/mobile.html"
	}

	collector.initLogic()

	return collector
}
