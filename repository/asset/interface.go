package asset

import "sirclo/project/capstone/entities"

type AssetRepo interface {
	Create(entities.Asset) error
	Get(string, string, string, int, int) ([]entities.Asset, error)
	GetById(int) (entities.Asset, error)
	Update(entities.Asset, entities.Asset, int) error
	Delete(int) error
	GetSummaryAsset() (entities.SummaryAsset, error)
	GetHistoryUsage(int, int, int) (entities.HistoryUsage, error)
	GetCategory() ([]entities.Categories, error)
}
