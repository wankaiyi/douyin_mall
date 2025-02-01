# RPC Client Generator
```shell
cd douyin_mall_rpc
cwgo client --type RPC --module douyin_mall/rpc -I ../idl --service auth --idl ../idl/auth.proto
cwgo client --type RPC --module douyin_mall/rpc -I ../idl --service cart --idl ../idl/cart.proto
cwgo client --type RPC --module douyin_mall/rpc -I ../idl --service checkout --idl ../idl/checkout.proto
cwgo client --type RPC --module douyin_mall/rpc -I ../idl --service order --idl ../idl/order.proto
cwgo client --type RPC --module douyin_mall/rpc -I ../idl --service payment --idl ../idl/payment.proto
cwgo client --type RPC --module douyin_mall/rpc -I ../idl --service product --idl ../idl/product.proto
cwgo client --type RPC --module douyin_mall/rpc -I ../idl --service user --idl ../idl/user.proto
cwgo client --type RPC --module douyin_mall/rpc -I ../idl --service doubao_ai --idl ../idl/doubao_ai.proto
```

# RPC Server Generator
```shell
cd app/auth; cwgo server --type RPC --service auth --module douyin_mall/auth --pass "-use  douyin_mall/rpc/kitex_gen"  -I ../../idl  --idl ../../idl/auth.proto
cd ../cart; cwgo server --type RPC --service cart --module douyin_mall/cart --pass "-use  douyin_mall/rpc/kitex_gen"  -I ../../idl  --idl ../../idl/cart.proto
cd ../checkout; cwgo server --type RPC --service checkout --module douyin_mall/checkout --pass "-use  douyin_mall/rpc/kitex_gen"  -I ../../idl  --idl ../../idl/checkout.proto
cd ../order; cwgo server --type RPC --service order --module douyin_mall/order --pass "-use  douyin_mall/rpc/kitex_gen"  -I ../../idl  --idl ../../idl/order.proto
cd ../payment; cwgo server --type RPC --service payment --module douyin_mall/payment --pass "-use  douyin_mall/rpc/kitex_gen"  -I ../../idl  --idl ../../idl/payment.proto
cd ../product; cwgo server --type RPC --service product --module douyin_mall/product --pass "-use  douyin_mall/rpc/kitex_gen"  -I ../../idl  --idl ../../idl/product.proto
cd ../user; cwgo server --type RPC --service user --module douyin_mall/user --pass "-use  douyin_mall/rpc/kitex_gen"  -I ../../idl  --idl ../../idl/user.proto
cd ../doubao_ai; cwgo server --type RPC --service doubao_ai --module douyin_mall/doubao_ai --pass "-use  douyin_mall/rpc/kitex_gen"  -I ../../idl  --idl ../../idl/doubao_ai.proto
```

# Http Server Generator
```shell
cd app/api; cwgo server  --type HTTP  --idl ../../idl/api/user_api.proto  --server_name api --module douyin_mall/api
cwgo server  --type HTTP  --idl ../../idl/api/cart_api.proto  --server_name api --module douyin_mall/api
cwgo server  --type HTTP  --idl ../../idl/api/order_api.proto  --server_name api --module douyin_mall/api
cwgo server  --type HTTP  --idl ../../idl/api/product_api.proto  --server_name api --module douyin_mall/api
cwgo server  --type HTTP  --idl ../../idl/api/checkout_api.proto  --server_name api --module douyin_mall/api
cwgo server  --type HTTP  --idl ../../idl/api/payment_api.proto  --server_name api --module douyin_mall/api
```
