package main

import (
	"challenge2019/src/models"
	"challenge2019/src/service"
)

const (
	PARTNERSPATH = "./resource/partners.csv"
	INPUTPATH    = "./resource/input.csv"
	OUTPUTPATH   = "./resource/output-p1.csv"
)

// findMin finds the minimum cost for each `input`
func findMin(inputList []models.Input, partnerList []models.PartnerRecord) []models.Output {
	var minCost int
	var choosenPartner *models.PartnerRecord
	outputList := make([]models.Output, 0, len(inputList))
	for _, input := range inputList {
		minCost = 0
		choosenPartner = nil
		for _, partner := range partnerList {
			if partner.TheatreID == input.TheatreID {
				if input.Volume > partner.Min && input.Volume <= partner.Max {
					costOfDelivery := input.Volume * partner.CostPerGB
					if costOfDelivery < partner.MinCost {
						costOfDelivery = partner.MinCost
					}
					if minCost == 0 || minCost > costOfDelivery {
						minCost = costOfDelivery
						choosenPartner = &partner
					}
				}
			}
		}
		if choosenPartner != nil {
			output := models.NewOutput(input.DeliveryID, true, choosenPartner.PartnerID, minCost)
			outputList = append(outputList, output)
		} else {
			output := models.NewOutput(input.DeliveryID, false, "", 0)
			outputList = append(outputList, output)
		}

	}
	return outputList
}

func main() {
	partnerList := service.ReadPartnerCsv(PARTNERSPATH)
	inputList := service.ReadInput(INPUTPATH)
	outputList := findMin(inputList, partnerList)
	service.WriteOutput(OUTPUTPATH, outputList)
}
