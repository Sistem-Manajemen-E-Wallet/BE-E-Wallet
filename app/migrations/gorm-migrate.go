package migrations

import (
	"e-wallet/app/configs"
	"e-wallet/app/databases"
	productData "e-wallet/features/product/data"
	userData "e-wallet/features/user/data"
	walletData "e-wallet/features/wallet/data"
)

func InitialMigration() {
	databases.InitDBMysql(configs.InitConfig()).AutoMigrate(&userData.User{}, &productData.Product{}, &walletData.Wallet{})
}
