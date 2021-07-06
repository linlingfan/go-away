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
  - 表锁：开销小、加锁快、不会产生死锁；在表读锁和表写锁的环境下：读读不阻塞，读写阻塞，写写阻塞！
    - 表读锁
    - 表写锁
  - 行锁：开销大、加锁慢、会产生死锁；锁粒度小,锁冲突概率低,并发率高（有一下两种行锁模式）
    - 共享锁(s Lock):允许事务读取一行数据
    - 排它锁(x Lock):允许事务更新或删除一行数据
  > 1. 普通 select 语句默认不加锁，而CUD操作默认加排他锁
  > 2. 不同的存储引擎支持锁的力度不一样
  >   - myisam 支持表锁
  >   - InnoDB 支持表锁和行锁
  > 
  - 页面锁：开销，加锁，以及锁冲突介于表级锁和行级锁之间，会出现死锁

- 按照锁的使用方式（行锁)
  - 共享锁
  - 排他锁 - （悲观锁的一种实现）
  
- 两种思想上的锁
  - 乐观锁
  - 悲观锁
  
- InnoDB的几种行级锁类型（行锁的实现方式）：
  - Record Lock：对单行索引项加锁
  - Gap Lock：锁定一个范围，但不包含记录本身
    - 间隙锁只会在Repeatable read隔离级别下使用(关闭gap锁：将事务隔离级别变为read committed（事务降级）)
    - 使用间隙锁的目的：防止幻读。
    - eg：select * from user where id > 1 for update;
  - Next-key Lock：等于record lock + Gap lock;对记录的前后间隙加锁，包含记录本身
  

>注意: 
> 1. InnoDB只有通过索引条件检索数据才使用行级锁，否则，InnoDB将使用表锁。也就是说，InnoDB的行锁是基于索引的。
> 2. 临界锁next-key lock当查询的索引是唯一索引的时候，InnoDB会将临键锁优化成记录锁，从而提高并发.

#### 乐观锁和悲观锁?
无论是Read committed还是Repeatable read隔离级别，都是为了解决读写冲突的问题。乐观锁和悲观锁亦是。<br>
乐观锁：是一种思想，认为不会锁定的情况下去更新数据，如果发现不对劲，才不更新(回滚)。在数据库中往往添加一个version字段来实现。在更新时校验该字段的值是否和第一次查出来的一样。如果一样更新，反之拒绝。
<br>之所以叫乐观，因为这个模式没有从数据库加锁，等到更新的时候再判断是否可以更新。<br>
悲观锁：认为数据库会发生并发冲突，直接上来就把数据锁住，其他事务不能修改，直至提交了当前事务；是数据库层面加锁，都会阻塞去等待锁。如：for update。

#### mysql InnoDB MVCC是啥？实现原理？
- MVCC（多版本并发控制），又称为一致性非锁定读。指InnoDB通过行多版本控制的方式来读取当前数据库中行的数据。<br>
  > 在事务隔离界别read committed 和 repeatable read（InnoDB默认的事务隔离界别）下，InnoDB使用非锁定一致性读。
  > - 在read committed隔离级别下，非一致性读总是读取被锁定行的最新一份快照数据（如果没有被锁定，则读取行的最新数据；如果行锁定了，则读取该行的最新一个快照）。
  > - 在repeatable read事务隔离级别下，对于快照数据，非一致性读总是读取事务开始时的快照。
  - 优点：MVCC在大多数情况下代替了行锁，实现了对读的非阻塞，**读不加锁，读写不冲突**，极大的提高了读效率。
  - 缺点：每行记录都需要额外的存储空间，需要做更多的行维护和检查工作。
  - InnoDB中隐藏列：DB_TRX_ID（事务ID）、 DB_ROLL_PTR（回滚指针）、 DB_ROW_ID（行自增ID）；初始为null
- 实现原理:
  > 基于undo log实现。undo log本身是为回滚而用的。具体内容就是复制事务开始前的行到undo buffer，在适合的时间把undo buffer中的内容刷新到磁盘。<br>
  > undo log的简单工作过程：
  > - 开始事务
  > - 记录数据行数据备份到undo log
  > - 更新数据
  > - 将undo log写到磁盘
  > - 提交事务
  > 
> 注意：
> <br>undo log必须在数据持久化前持久化到磁盘，这样才能保证在系统崩溃时，可以使用undo log来回滚事务；
> <br>Innodb通过undo log保存了已更改行的旧版本的快照。
> <br>提交事务做的事情有：写redo log和binlog，并且把数据持久化到磁盘（可以通过参数控制）

- MVCC的sql规则
> MVCC只在READ COMMITTED 和 REPEATABLE READ 两个隔离级别下工作。READ UNCOMMITTED总是读取最新的数据行，而不是符合当前事务版本的数据行。而SERIALIZABLE 则会对所有读取的行都加锁。另外事务的版本号是递增的。
> - SELECT
> <br>InnoDB只查找行的事务ID 小于当前事务ID 的数据行（避免幻读）
> - INSERT
> <br>新插入的每一行保存当前事务ID作为行的事务ID
> - DELETE
> <br>删除的每一行保存当前事务ID作为行的事务ID
> - UPDATE
> <br>实际上是删除旧行，插入新行。
> <br>保存当前的事务ID作为新行的事务ID，同时保存当前事务ID到旧行的事务ID（上一个版本的事务ID即回滚ID）。

- 一致性锁定度（MVCC为一致性非锁定读）
> 1. 在默认的配置下，即事务的隔离级别为可重复度，InnoDB存储引擎的select操作使用一致性非锁定读（即MVCC）。但是在某些情况下，用户需要显示的对数据库读取操作进行加锁，以保证数据逻辑的一致性。而这要求数据库支持加锁语句，InnoDB存储引擎对select支持两种一致性的锁定读操作：
> ``` 
> select …  for update
> select … lock in share mode;
> select … for update 对读取的行记录加一个X锁，其他事务不能对已锁定的行加上任何锁。
> select…lock in share mode对读取的行记录加一个S锁，其他事务可以向被锁定的行加S锁，但是如果加X锁，则会被阻塞。
> ```
> 2. **对于一致性非锁定读，即时读取的行已经被执行了select…for update，也是可以进行读取的。**
> 3. select…for update或者select…lock in share mode必须在事务中，因为当事务提交了，锁也就释放了。从而避免锁没有释放，可能导致死锁的情况。
> 4. 如果不加筛选条件（或者筛选条件不走索引），会升级为表锁
> 5. 索引数据重复率太高会导致全表扫描：当表中索引字段数据重复率太高，则MySQL可能会忽略索引，进行全表扫描，此时使用表锁。可使用 force index 强制使用索引。
> 

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
    * 幻读：幻读和不可重复读类似。事务A读取几行数并准备更新期间，事务B插入了一行数据；导致事务A在结束后发现了还有一行数据未更新,感觉出现幻觉。
  > 注意：
  > * 不可重复读在于对单条数据的修改；幻读在于对数据的增删操作；
  > *
  
- 事务的隔离级别：
  * 读取未提交：最低隔离级别。会导致 脏读 不可重复读 幻读。
  * 读取已提交：允许读取并发事务已提交的数据；可以组织脏读，不可阻止不可重复读和幻读。
  * 可重复读：对同一个字段在一个事务里多次读取的结果是一致的；除非数据是本身事务所修改。可阻止脏读，不可重复读；但是幻读仍可能发生。
  * 串行化：最高的隔离级别，完全服从ACID的隔离级别，性能最差；所有事务串行化执行，可以避免脏读，不可重复读以及幻读。
  > 注意：
  > * Mysql InnoDB默认的隔离级别是 可重复读。查看隔离级别：``select @@tx_isolation;``
  > * 不过Mysql InnoDB 允许使用next-key lock（临键锁）锁算法来避免幻读的发生；
  > * 设置隔离级别：``set session transaction isolation level xxxx;``
  > * InnoDB实现的Repeatable read隔离级别配合GAP间隙锁已经避免了幻读！*

- 解决并发带来的问题：
  > 1. 适当根据实际业务需要提高事务的隔离级别。隔离级别越高，越能保证数据的完整性和一致性；但是对并发性能影响也越大。一般建议使用Read Committed来避免脏读，提高并发性能。
  > 2. 应用层可以结合乐观锁和悲观锁来解决脏读、不可重复读和幻读的问题。(排它锁)
  >

- 阻塞：事务因为等待其他事务释放锁而等待

- 超时：等待其他事务释放锁，超过超时时间，就认为是超时。innodb_lock_wait_timeout：用来控制超时时间，默认是50秒。

#### mysql死锁的产生和避免？
- 概念：死锁是指两个或者两个以上的事务，因争夺资源而造成的一种互相等待的现象。若无外力作用，所有事务都将无法推进下去。````

- 处理死锁
  1.（FIFO处理死锁）解决数据库死锁最简单的方法：设置超时时间。
  2. 等待图：因为FIFO处理死锁可能不适用，所以数据库普遍采用了wait-for graph（等待图）的方式来进行死锁检测。和超时机制比较，这是一种更为主动的死锁检测方式，InnoDB也采用了这种方式。
     ![img_1.png](img/img_1.png)
     ![img.png](img/wait_graph.png)
  3. 锁升级：比如行锁升级为表锁（或串行化），但是并发性能降低
- 减少或避免死锁  
  1. 以固定的顺序访问表和行：简单方法是对id列表先排序，后执行，这样就避免了交叉等待锁的情形；将两个事务的sql顺序调整为一致，也能避免死锁。 
  2. 大事务拆成小事务：大事务更倾向于死锁，如果业务允许，将大事务拆小。
  3. 在同一个事务中，尽可能做到一次锁定所需要的所有资源，减少死锁概率。
  4. 降低隔离级别。如果业务允许，将隔离级别调低也是较好的选择，比如将隔离级别从RR调整为RC，可以避免掉很多因为gap锁造成的死锁。
  5. 为表添加合理的索引。

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

### Mysql bin log、undo log、redo log作用和区别

## 线上问题

#### 线上大表DDL如何操作（大表加索引，新增字段）？
#### test
    