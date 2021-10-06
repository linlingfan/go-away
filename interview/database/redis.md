## redis

- [基本知识](https://mp.weixin.qq.com/s/aOiadiWG2nNaZowmoDQPMQ)
- [参考](https://github.com/Snailclimb/JavaGuide/blob/master/docs/database/Redis/redis-all.md)

### redis的常用数据结构？使用场景？bitmap布隆过滤器、HyperLogLog?

### redis分布式锁 

- setNex
- set key value ex nx
- redlock

### redis持久化 RDB和AOF

### redis的hash槽？一致性hash和hash槽的区别？

- redis的hash槽：redis cluster采用数据分片的哈希槽来进行数据存储和数据的读取。redis cluster一共有2^14（16384）个槽，所有的master节点都会有一个槽区比如0～1000，槽数是可以迁移的。master节点的slave节点不分配槽，只拥有读权限。但是注意在代码中redis cluster执行读写操作的都是master节点，并不是你想 的读是从节点，写是主节点。第一次新建redis cluster时，16384个槽是被master节点均匀分布的。

和一致性哈希相比

1. 它并不是闭合的，key的定位规则是根据CRC-16(key)%16384的值来判断属于哪个槽区，从而判断该key属于哪个节点，而一致性哈希是根据hash(key)的值来顺时针找第一个hash(ip)的节点，从而确定key存储在哪个节点。
2. 一致性哈希是创建虚拟节点来实现节点宕机后的数据转移并保证数据的安全性和集群的可用性的。redis cluster是采用master节点有多个slave节点机制来保证数据的完整性的,master节点写入数据，slave节点同步数据。当master节点挂机后，slave节点会通过选举机制选举出一个节点变成master节点，实现高可用。但是这里有一点需要考虑，如果master节点存在热点缓存，某一个时刻某个key的访问急剧增高，这时该mater节点可能操劳过度而死，随后从节点选举为主节点后，同样宕机，一次类推，造成缓存雪崩。解决这个问题请看我的另一篇文章如何应对热点缓存问题
3. 扩容和缩容  
可以看到一致性哈希算法在新增和删除节点后，数据会按照顺时针来重新分布节点。而redis cluster的新增和删除节点都需要手动来分配槽区。


### redis ZSET的底层实现？跳表的实现？

### redis 如何保证高可用？哨兵模式

### redis 主从同步

- [参考](https://blog.csdn.net/qq_44590469/article/details/111932736)

- 开始RDB全量同步 + 增量同步
- 一主多从。减少主节点的同步压力 可以使用 主-从-从同步模式。通过从节点级联的方式进行同步。
- 同步期间网络中断：网络断连重新连接后，主从节点通过分别维护的偏移量来同步写命令。即：当网络断开连接时，从节点不再进行同步，此时主节点由于不断接收新的写操作的偏移量会大于从节点的偏移量。当连接恢复时，从节点向主节点发送带有偏移量的psync 命令，主节点根据偏移量来进行比较，只需将未同步写命令同步给从节点即可。
