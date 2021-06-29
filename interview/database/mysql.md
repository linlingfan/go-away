# mysql 相关
    - 关系型数据库。关系模型表明了数据库中所存储的数据之间的联系（一对一、一对多、多对多）
## mysql基础
1. 启动服务：mysql.server start
2. 登录mysql：mysql -u root -p xxx
3. show databases; show engines;
4. 查看mysql当前默认存储引擎：show variables like '%storage_engine%';
5. 查看表的存储引擎：table status like "table_name" ;
6.
## 核心知识

#### mysql有哪些存储引擎？MyISAM 和 InnoDB的区别？
#### mysql锁的机制？InnoDB锁的算法？
#### mysql的事务？事务的ACID特性？并发事务带来的问题以及解决（隔离级别）？事务的隔离级别？ 
#### 

## mysql规范和建议

## mysql索引相关
#### 什么是索引？索引优缺点
#### mysql索引底层数据结构？B数和B+树？
#### mysql索引类型？创建和使用索引的建议？

## 线上问题

#### 线上大表DDL如何操作（大表加索引，新增字段）？
#### 
    