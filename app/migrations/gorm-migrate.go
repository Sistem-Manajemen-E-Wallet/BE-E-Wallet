package migrations

import (
	"e-wallet/app/configs"
	"e-wallet/app/databases"
	historyData "e-wallet/features/history/data"
	productData "e-wallet/features/product/data"
	topupdata "e-wallet/features/topups/data"
	transactionData "e-wallet/features/transaction/data"
	userData "e-wallet/features/user/data"
	walletData "e-wallet/features/wallet/data"
)

func InitialMigration() {
	databases.InitDBMysql(configs.InitConfig()).AutoMigrate(&userData.User{}, &productData.Product{}, &walletData.Wallet{}, &topupdata.TopUp{}, &transactionData.Transaction{}, &historyData.History{})
}
