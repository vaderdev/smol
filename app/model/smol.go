package model

func GetAllSmols() ([]Smol, error) {
	var smols []Smol

	tx := db.Find(&smols)
	if tx.Error != nil {
		return []Smol{}, tx.Error
	}

	return smols, nil
}

func GetSmol(id uint64) (Smol, error) {
	var smol Smol

	tx := db.Where("id = ?", id).First(&smol)

	if tx.Error != nil {
		return Smol{}, tx.Error
	}

	return smol, nil
}

func CreateSmol(smol Smol) error {
	tx := db.Create(&smol)
	return tx.Error
}

func UpdateSmol(smol Smol) error {
	tx := db.Save(&smol)
	return tx.Error
}

func DeleteSmol(id uint64) error {
	tx := db.Unscoped().Delete(&Smol{}, id)
	return tx.Error
}

func FindBySmolUrl(url string) (Smol, error) {
	var smol Smol
	tx := db.Where("smol = ?", url).First(&smol)
	return smol, tx.Error
}
