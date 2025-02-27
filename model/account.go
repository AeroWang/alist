package model

import (
	"github.com/Xhofe/alist/conf"
	"time"
)

type Account struct {
	Name           string `json:"name" gorm:"primaryKey" validate:"required"`
	Index          int    `json:"index" validate:"required"`
	Type           string `json:"type"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	RefreshToken   string `json:"refresh_token"`
	AccessToken    string `json:"access_token"`
	RootFolder     string `json:"root_folder"`
	Status         string `json:"status"`
	CronId         int
	DriveId        string
	Limit          int        `json:"limit"`
	OrderBy        string     `json:"order_by"`
	OrderDirection string     `json:"order_direction"`
	Proxy          bool       `json:"proxy"`
	UpdatedAt      *time.Time `json:"updated_at"`
	Search         bool       `json:"search"`
	ClientId       string     `json:"client_id"`
	ClientSecret   string     `json:"client_secret"`
	Zone           string     `json:"zone"`
	RedirectUri    string     `json:"redirect_uri"`
	SiteUrl        string     `json:"site_url"`
	SiteId         string
	OnedriveType   string `json:"onedrive_type"`
}

var accountsMap = map[string]Account{}

// SaveAccount save account to database
func SaveAccount(account Account) error {
	if err := conf.DB.Save(account).Error; err != nil {
		return err
	}
	RegisterAccount(account)
	return nil
}

func DeleteAccount(name string) error {
	account := Account{
		Name: name,
	}
	if err := conf.DB.Delete(&account).Error; err != nil {
		return err
	}
	delete(accountsMap, name)
	return nil
}

func AccountsCount() int {
	return len(accountsMap)
}

func RegisterAccount(account Account) {
	accountsMap[account.Name] = account
}

func GetAccount(name string) (Account, bool) {
	if len(accountsMap) == 1 {
		for _, v := range accountsMap {
			return v, true
		}
	}
	account, ok := accountsMap[name]
	return account, ok
}

func GetAccountFiles() ([]*File, error) {
	files := make([]*File, 0)
	var accounts []Account
	if err := conf.DB.Order("`index`").Find(&accounts).Error; err != nil {
		return nil, err
	}
	for _, v := range accounts {
		files = append(files, &File{
			Name:      v.Name,
			Size:      0,
			Type:      conf.FOLDER,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return files, nil
}

func GetAccounts() ([]Account, error) {
	var accounts []Account
	if err := conf.DB.Order("`index`").Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}
