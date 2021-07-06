# mysql 相关
    - 关系型数据库。关系模型表明了数据库中所存储的数据之间的联系（一对一、一对多、多对多）
## 参考资料
- [JavaGuide-mysql](https://github.com/Snailclimb/JavaGuide/blob/master/docs/database/MySQL.md)
- [mysql-全面总结](https://cloud.tencent.com/developer/article/1803404?from=article.detail.1614355)
- [mysql-锁（全面总结）](https://cloud.tencent.com/developer/article/1614355)
- [mysql之锁总结](https://cloud.tencent.com/developer/article/1444243?from=article.detail.1584918)
## mysql基础
1. 启动服务：mysql.server start
2. 登录mysql：mysql -u root -p xxx
3. show databases; show engines;
4. 查看mysql当前默认存储引擎：show variables like '%storage_engine%';
5. 查看表的存储引擎：table status like "table_name" ;
6.
## 核心知识
#### mysql有哪些存储引擎？MyISAM 和 InnoDB的区别？
> 常见的存储引擎有MyISAM,InnoDB,MEMORY,MERGE.

MYISAM和InnoDB的区别：
- MYISAM 支持表级锁 innoDB支持表级锁和行级锁。
- M不支持事务 I支持事务；
- M不支持外键 I支持外键（一般不建议在数据层面做外键，代码层面也可以不过有数据一致性问题）
- M不支持数据异常崩溃后安全恢复，I支持；
- M不支持MVCC I支持（有点废话，mvcc是行级锁的升级，M本身不支持行级锁）
- M索引数据结构B数 I为B+树；M为非聚集索引 I为聚集索引；

其他：
- Memory存储引擎：支持的数据类型有限，不支持text和blob类型；默认使用hash索引，数据都放在内存中，访问速度快。
- Merge存储引擎：是一组MYISAM表的组合，本身没有数据，对其操作就是对内部MYISAM表的操作

### mysql事务
#### mysql的事务？事务的ACID特性？并发事务带来的问题以及解决（隔离级别）？事务的隔离级别？ 
- 事务是逻辑上的一组操作；要么都执行，要么都不执行。
- 四种特性：
    * 原子性
    * 一致性
    * 隔离性
    * 持久性
- 并发带来的问题：
    * 脏读：事务B读取并使用事务A正在修改但未提交的数据，最后A回滚了。
    * 丢失修改：事务A和事务B同时读到数据d，事务A修改提交数据d后；事务B也修改提交覆盖数据d，造成事务A数据修改的丢失。
    * 不可重复读：事务A在未提交前多次读取数据d，在此期间事务B（其他事务）读取并修改了数据d，导致事务A多次读取的数据d不一致。
    * 幻读：幻读和不可重复读类似。事务A读取几行数据期间，事务B插入了几行数据；导致事务A在随后的查询中发现了几行原本不存在的数据。
  > 注意：
  > * 不可重复读在于对单条数据的修改；幻读在于对数据的增删操作；
  > *
  
- 事务的隔离级别：
  * 读取未提交：最低隔离级别。会导致 脏读 不可重复读 幻读。
  * 读取已提交：允许读取并发事务已提交的数据；可以组织脏读，不可阻止不可重复读和幻读。
  * 可重复读：对同一个字段在一个事务里多次读取的结果是一致的；除非数据是本身事务所修改。可阻止脏读，不可重复读；但是幻读仍可能发生。
  * 串行化：最高的隔离级别，完全服从ACID的隔离级别，性能最差；所有事务串行化执行，可以避免脏读，不可重复读以及幻读。
  > 注意：
  > * Mysql InnoDB默认的隔离级别是 可重复读。
  > * 不过Mysql InnoDB 允许使用next-key lock（临键锁）锁算法来避免幻读的发生；
  >

### mysql锁

#### mysql相关锁
锁机制主要用于管理对共享资源并发访问。<br>
在数据库中lock和latch都称为锁，但是两者的意义不一样。
- latch：称为闩锁（shuang suo），其要求锁定的时间必须非常短。若持续的时间长，则应用的性能会非常差。
  <br>在InnoDB存储引擎中，latch又分为mutex互斥锁.
  ![img.png](img/img.png)
- lock：lock的对象是事务，用来锁定的是数据库中的对象，如表、页、行。并且一般lock的对象仅在事务commit或者rollback后进行释放。有死锁检测机制。 
  
通过 ``show engine innodb mutex``可以查看InnoDB存储引擎的中latch
![img.png](img/mutex_img.png)

#### mysql Lock介绍
![img.png](img/lock_img.png)

- 按照细粒度：
  - 表锁：在表读锁和表写锁的环境下：读读不阻塞，读写阻塞，写写阻塞！
    - 表读锁
    - 表写锁
  - 行锁：
    - 共享锁
    - 排它锁

- 不同的存储引擎支持锁的力度不一样
  - myisam 支持表锁
  - InnoDB 支持表锁和行锁

注意：InnoDB只有通过索引条件检索数据才使用行级锁，否则，InnoDB将使用表锁。也就是说，InnoDB的行锁是基于索引的。

#### mysql InnoDB MVCC是啥？实现原理？

#### 一般是怎么优化sql的？
- 通过show status了解各种sql的执行频率
  <br>Com_select 执行select的次数 
  <br>Com_update 执行update的次数
  <br>了解当前数据库是以更新插入为主还是以查询为主；了解各类型的sql比例
- 定位执行效率低的sql
  <br>通过慢查询定位执行效率低的sql
- 通过explain分析低效率sql执行计划
- 通过show profile分析sql
- 通过trace分析优化器如何选择执行计划

#### mysql一般有那几个分区？分别有啥作用？
mysql有四个分区类型；(TODO)
- RANGE分区
- List分区
- Hash分区
- Key分区

## mysql索引相关
#### 什么是索引？索引优缺点
#### mysql索引底层数据结构？B数和B+树？
#### mysql索引类型？创建和使用索引的建议？

## 线上问题

#### 线上大表DDL如何操作（大表加索引，新增字段）？
#### test
    