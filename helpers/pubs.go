package helpers

import (
	"database/sql"
	"log"
)

type Sex struct {
	Male   int `json:"male"`
	Female int `json:"female"`
}
type Img struct {
	Src  string `json:"src"`
	Blur string `json:"img_blur"`
}

type Pub struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Thematic     string `json:"thematic"`
	Price        int    `json:"price"`
	Status       int    `json:"status"`
	PubID        int    `json:"pub_id"`
	ScheduleLink string `json:"schedule_link"`
	Subs         int    `json:"subs"`
	Sex          Sex    `json:"sex"`
	Img          Img    `json:"img"`
	WeekStat     int    `json:"week_stat"`
	Reach        int    `json:"reach"`
}

func GetPubsList() ([]*Pub, bool) {
	pubs := make([]*Pub, 0)
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3050)/panel")
	if err != nil {
		log.Print("Error to connect to database")
		log.Print(err)
	}
	rows, err := db.Query("SELECT id,title,thematic,price,status,pub_id,subs,male,female,img,img_blur,week_stat,reach,schedule_link FROM pubs")
	if err != nil {
		log.Print("Error query")
		log.Print(err)
	}
	if err != nil {
		return pubs, true
	}
	defer rows.Close()
	for rows.Next() {
		pub := new(Pub)
		err := rows.Scan(
			&pub.ID,
			&pub.Title,
			&pub.Thematic,
			&pub.Price,
			&pub.Status,
			&pub.PubID,
			&pub.Subs,
			&pub.Sex.Male,
			&pub.Sex.Female,
			&pub.Img.Src,
			&pub.Img.Blur,
			&pub.WeekStat,
			&pub.Reach,
			&pub.ScheduleLink,
		)
		if err != nil {
			log.Print("Error scan row")
			log.Print(err)
		}
		pubs = append(pubs, pub)
	}
	if err = rows.Err(); err != nil {
		log.Print("zx")
		log.Print(err)
	}

	return pubs, false
}
