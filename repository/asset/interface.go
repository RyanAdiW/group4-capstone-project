package asset

import "sirclo/project/capstone/entities"

type AssetRepo interface {
	Create(entities.Asset) error
	Get(string, string) ([]entities.Asset, error)
	GetById(int) (entities.Asset, error)
	Update(entities.Asset, entities.Asset, int) error
	Delete(int) error
	GetSummaryAsset() (entities.SummaryAsset, error)
	GetHistoryUsage(int) (entities.HistoryUsage, error)
}
