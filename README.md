# 抖音商城大项目

## 项目介绍

抖音商城是一个基于微服务架构的高性能、高可用电商平台，使用Hertz和Kitex构建。项目旨在为用户提供流畅的购物体验，支持商品搜索、购物车管理、订单处理、支付结算等功能。

## 技术选型

| 技术 | 说明 |
| --- | --- |
| Golang | 编程语言 |
| Hertz | 高性能HTTP框架 |
| Kitex | 高性能RPC框架 |
| Gorm | Golang ORM库 |
| MySQL | 关系型数据库 |
| Redis | 缓存 |
| Kafka | 高吞吐量、低延迟的消息队列 |
| XXL-JOB | 分布式定时任务调度平台 |
| Elasticsearch | 分布式搜索引擎 |
| Sentinel | 流量控制和服务降级 |
| Prometheus | 系统监控和报警工具 |
| Grafana | 数据可视化和监控平台 |
| Opentelemetry, Jaeger | 收集和管理遥测数据（链路追踪） |
| 腾讯云CLS日志服务 | 腾讯云提供的日志服务，用于日志的收集、存储、检索和分析 |
| Nacos | 动态服务发现、配置和服务管理平台 |
| Casbin | Golang访问控制库，支持多种访问控制模型（如RBAC、ABAC） |
| Jwt | JSON Web Token |
| Docker | 容器化部署 |
| Harbor | Docker私服 |
| K8S（K3S发行版） | 容器编排平台 |
| Jenkins | 自动化CI/CD |

## 架构设计

### 项目架构图
![输入图片说明](img/%E9%A1%B9%E7%9B%AE%E6%9E%B6%E6%9E%84%E5%9B%BE.png)

### 整体架构说明——一个请求的完整过程
1. Nginx-Ingress将请求负载均衡到网关服务集群
2. 网关rpc调用鉴权服务：
  - 鉴权服务通过双token进行权限校验和token的无感刷新，保证access_token可以在短时间内过期而防止攻击者获得后滥用，用户的登录可以保证长时间不过期，每次access_token过期时可以通过refresh_token同时刷新access_token和refresh_token，实现无感刷新和登录的续期
  - 解析access_token后首先检查配置的用户黑白名单，再通过Casbin的RBAC对用户进行权限校验
3. 限流和熔断（Sentinel）
  - 限流
    - 从两个层面进行限流：HTTP接口和RPC接口
    - 从三个粒度进行限流：
      - 对API的QPS进行限流：这个限流是针对单机的，目的是防止单机负载过高
      - 对IP进行限流：对服务集群进行限流，目的是防止攻击
      - 在单个接口中对userId进行限流：对服务集群进行限流，限制用户对一些接口的使用次数，比如AI相关的服务，因为用多了花的钱也多
  - 熔断
    - 当某个接口的QPS达到设定的阈值后，就响应特定的内容，并暂停提供服务一段时间（配置的时间，如5秒）
4. 通过RPC调用业务服务
  - 重试：为了尽可能避免偶发性的错误，我们在一些接口设置了重试的次数，经过测试，服务的可用性确实得到很大程度的提高
  - 超时：我们也对一些接口设置了超时时间，防止请求时间太长影响用户的体验或者下游服务自身出现问题而对上游服务造成过大的影响，紧接着对上游服务的上游服务也会造成影响
5. 分布式事务
  - 我们并没有采用一些解决分布式事务的一些框架，比如seata，原因如下：
    - 两阶段提交的性能较低，并且商城是C端，很可能影响整个服务的性能
    - 部分场景下（比如扣减库存），一些操作是在redis中完成的，而不是通过数据库，也是为了尽可能提高响应速度和服务的性能
  - 最终一致性：
    - 我们使用了kafka来通知其他服务进行一些操作，如支付成功后的真实库存扣减，取消订单后的释放库存，并且可以重试一定次数直到成功
    - 补偿机制：在锁定商品的库存后，会rpc调用订单服务进行下单操作，在下单前发送一个延时消息，延时时间大于（单次创建订单的时间 * 重试次数），收到延时消息后，检查订单是否创建成功，如果没有创建成功则释放锁定的库存
6. 异步通信
  - 我们部署了一个kafka集群来保障异步通信的可靠性和性能
    - 负载均衡：由于每个业务服务的有三个节点，所以将每个topic的partition设置为3个，并通过比如订单id这样的属性哈希路由到不同的partition，达到每个节点负载均衡
    - 精确一次：在某些场景，我们只希望我们的消息被消费一次，比如扣减库存、释放库存等
      - 生产者：通过配置幂等消费者和事务消息等配置，实现生产者的精确一次
      - 消费者：因为大部分场景消息都带有订单id，所以我们在消费时对订单id进行去重

## 数据库设计

### 表结构设计
![输入图片说明](img/%E6%95%B0%E6%8D%AE%E5%BA%93%E8%A1%A8%E7%BB%93%E6%9E%84.png)

### 逻辑删除

对需要防止误删或删除后可找回的场景（如product、category、order等）添加逻辑删除字段。

### 索引

- **购物车服务**: 建立联合唯一索引（user_id, product_id）。
- **豆包服务**: 建立联合索引（user_id, uuid）。
- **订单服务**: 建立联合索引（user_id, deleted_at, created_at）。
- **商品服务**: 建立联合索引（category_id, deleted_at）。
- **用户服务**: 建立联合唯一索引（user_id, deleted_at）。

## 缓存

1. 鉴权服务
  - 缓存场景：双token（access_token和refresh_token），通过redis的过期来实现token的过期
2. 用户服务
  - 缓存场景：
    - 用户个人信息
    - 用户收货地址
  - 缓存策略：旁路缓存
    - 每次用户登录或刷新access_token时，发消息给kafka异步写入redis，key的过期时间和access_token的过期时间一致，以提高缓存的命中率
    - 调用查询用户个人信息或收货地址时，先查询缓存，缓存命中则缓存，未命中则查询数据库并写入缓存
  - 一致性保证：
    - 用户个人信息：用户修改个人信息后，删除redis中的缓存。如果删除缓存失败，则回滚事务。
3. 豆包AI服务
  - 缓存场景：用户和智能购物助手的临时对话，可以更快获取这次临时对话的聊天记录
  - 缓存策略：Write-Through (写穿透)、旁路缓存
    - 如果临时对话的缓存过期，则从数据库重新加载
  - 一致性保证：
    - 新增聊天消息时，先写入数据库，再写入redis，如果redis写入失败，则回滚事务
4. 商品服务
  - 缓存场景：
    - 搜索结果缓存：使用Elasticsearch查询时，计算DSL的MD5值，如果将MD5作为key，查询到的数据作为value缓存在redis中，缓存命中时直接返回查询的结果
    - 缓存单个商品的基本信息和库存信息：由于商品信息修改频率低，而库存信息修改频率高，所以将商品的基本信息和库存信息分开缓存
  - 缓存策略：旁路缓存
    - 当商品信息不存在时，从数据库中查出来并写入redis；当库存信息不存在时，使用分布式互斥锁加锁后，查询数据库并写入redis
  - 一致性保证：
    - 由于DSL这个缓存无法感知商品信息的变化，所以过期时间较短
    - 对于商品信息的修改，由于商品信息追求的一致性比较强，所以采用延时双删，先删除缓存，再修改商品数据，然后再通过kafka第二次延时删除，并且可以在删除失败时重试直到删除成功

## 项目亮点

1. **订单超时取消**: 采用延时消息+定时任务兜底的方案。
2. **权限校验**: 使用Casbin实现权限校验，支持权限的实时变更。
3. **合理使用事务**: 在多个数据库写操作和写入缓存、发送消息到MQ等场景使用事务。
4. **商品搜索缓存**: 实现不同粒度和层级的缓存，解决热点问题。
5. **字段冗余存储**: 在订单库的order_item中存储相关商品信息，减少RPC请求。
6. **AI相关功能**: 使用豆包大模型和eino实现AI查询订单和模拟用户下单。
7. **乐观锁解决并发安全问题**: 通过乐观锁解决订单状态修改时的并发安全问题。
8. **配置化**: 所有配置通过Nacos存储，包括项目基本配置和限流规则。
9. **链路追踪**: 通过Opentelemetry接入完整的链路追踪。
10. **日志存储与检索**: 使用腾讯云CLS服务进行日志存储和检索。
11. **系统监控**: 使用Prometheus采集服务metrics，并在Grafana上展示。
12. **完善的CI/CD**: 使用Jenkins完成CI/CD。
13. **容器编排**: 使用K8S管理容器

## 测试结果

### 功能测试

- **AI查询用户订单单元测试**
- **支付接口单元测试**
- **下单接口单元测试**

### 性能测试

- **商品搜索接口压测**: 500次请求，通过率100%，平均接口请求耗时772毫秒。

![输入图片说明](img/%E5%95%86%E5%93%81%E6%90%9C%E7%B4%A2%E5%8E%8B%E6%B5%8B%E7%BB%93%E6%9E%9C.png)

## 项目总结与反思

### 目前仍存在的问题

1. 支付宝沙箱请求超时，偶尔返回504状态码。
2. AI功能偶尔返回不符合预期的数据。

### 已识别出的优化项

1. **消息队列消费失败处理**: 对每个场景新建单独的topic作为死信队列。
2. **防止消息队列重复消费**: 使用数据库去重，解决重复消费和重试问题。
3. **AI功能优化**: 通过function call、embedding等特性优化AI回答。

### 架构演进的可能性

1. **服务拆分**: 将商品服务拆分为搜索服务和库存服务，职责更加单一。
2. **高可用架构**: 将MySQL、Redis和Elasticsearch采用主从、集群等高可用架构部署。

# 项目代码生成脚手架
## RPC Client Generator
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

## RPC Server Generator
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

## Http Server Generator
```shell
cd app/api; cwgo server  --type HTTP  --idl ../../idl/api/user_api.proto  --server_name api --module douyin_mall/api
cwgo server  --type HTTP  --idl ../../idl/api/cart_api.proto  --server_name api --module douyin_mall/api
cwgo server  --type HTTP  --idl ../../idl/api/order_api.proto  --server_name api --module douyin_mall/api
cwgo server  --type HTTP  --idl ../../idl/api/product_api.proto  --server_name api --module douyin_mall/api
cwgo server  --type HTTP  --idl ../../idl/api/checkout_api.proto  --server_name api --module douyin_mall/api
cwgo server  --type HTTP  --idl ../../idl/api/payment_api.proto  --server_name api --module douyin_mall/api
```
