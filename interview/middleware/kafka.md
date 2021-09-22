## Kafka
- 
#### kafka的工作原理？如何保证顺序？
- 单个partition 单个消费者

#### kafka为啥这么快 高性能？

- 生产者：可以批处理 批量发送；压缩算法（消息压缩）
- IO读写：底层顺序写、零拷贝；mmp 类似mongoDB 索引映射磁盘数据。
- 网络模型：底层基于 Java NIO，采用和 Netty 一样的 Reactor 线程模型。
  - Reactor：把 IO 事件分配给对应的 handler 处理
  - Acceptor：处理客户端连接事件
  - Handler：处理非阻塞的任务
    
- 消息是以 Topic 为单位进行归类，各个 Topic 之间是彼此独立的，互不影响。每个 Topic 又可以分为一个或多个分区。每个分区各自存在一个记录消息数据的日志文件。
- 无锁轻量级Offset

#### kafka如何知道消费进度
- ，Kafka就是用offset来表示消费者的消费进度到哪了，每个消费者会都有自己的offset。说白了offset就是表示消费者的消费进度。

### kafka的ZooKeeper
- Kafka 强依赖于 ZooKeeper ，通过 ZooKeeper 管理自身集群，
  如：Broker 列表管理、Partition 与 Broker 的关系、Partition 与 Consumer 的关系、Producer 与 Consumer 负载均衡、消费进度 Offset 记录、消费者注册 等 ，
  所以为了达到高可用，ZooKeeper 自身也必须是集群。
