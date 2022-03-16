nohup go run ./service/book/rpc/book.go -f ./service/book/rpc/etc/book.yaml > ./book.log 2>&1 &
nohup go run ./service/lend/rpc/lend.go -f ./service/lend/rpc/etc/lend.yaml > ./lend.log 2>&1 &
nohup go run ./service/systemlog/rpc/systemlog.go -f ./service/systemlog/rpc/etc/systemlog.yaml > ./systemlog.log 2>&1 &
nohup go run ./service/user/rpc/user.go -f ./service/user/rpc/etc/user.yaml > ./user.log 2>&1 &
nohup go run ./service/vip/rpc/vip.go -f ./service/vip/rpc/etc/vip.yaml > ./vip.log 2>&1 &
nohup go run ./api/admin.go -f ./api/etc/admin-api.yaml > ./admin.log 2>&1 &
