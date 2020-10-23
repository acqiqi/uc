package main

import (
	"log"
	"uc/models"
)

func WorkIn() {
	list, err := models.MaGetAllList()
	if err != nil {
		log.Println(err.Error())
		return
	}

	for i, v := range list {

	}

}
