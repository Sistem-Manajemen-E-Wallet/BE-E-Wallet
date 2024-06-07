test-all:
	go test ./features/user/service -coverprofile=cover.out
	go test ./features/history/service -coverprofile=cover.out
	go test ./features/product/service -coverprofile=cover.out
	go test ./features/topups/service -coverprofile=cover.out
	go test ./features/transaction/service -coverprofile=cover.out
	go test ./features/wallet/service -coverprofile=cover.out
test-user:
	go test ./features/user/service -coverprofile=cover.out
test-history:
	go test ./features/history/service -coverprofile=cover.out
test-product:
	go test ./features/product/service -coverprofile=cover.out
test-topups:
	go test ./features/topups/service -coverprofile=cover.out
test-transaction:
	go test ./features/transaction/service -coverprofile=cover.out
test-wallet:
	go test ./features/wallet/service -coverprofile=cover.out
