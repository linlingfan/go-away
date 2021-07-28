# 计算机网络

###. http和https的区别
    https://www.cnblogs.com/wqhwe/p/5407468.html
    
### https加密过程
    
###. TCP流量控制和拥塞控制 
    - [链接1](https://www.cnblogs.com/LloydDracarys/articles/9032696.html)
    - [链接2](https://blog.csdn.net/dangzhangjing97/article/details/81008836)
    - 流量控制（数据单位是字节）：通过滑动窗口实现，接收方在返回的ack中告知发送发自己接受窗口的大小，进而控制发送发数据发送的大小；
    防止发送方发送数据过快过多导致接收方处理不过，进而数据溢出导致丢包。
        - 死锁原因和解决：
            - 当某个ACK报文(接收方B返回的)丢失了，就会出现发送发A等待接受方B确认，并且B等待A发送数据的死锁状态。
            - 为了解决这种问题，TCP引入了持续计时器（Persistence timer），当A收到rwnd=0时，就启用该计时器，时间到了则发送一个1字节的探测报文，
        询问B是很忙还是上个ACK丢失了，然后B回应自身的接收窗口大小，如果返回仍为0（A重设持续计时器继续等待）或者会重发rwnd=x；
        一般设置重发3次，达到3次有的TCP实现就会发RST把链接断了
        - 滑动窗口的优点：使用滑动窗口可以一次发送多段数据，提高性能；（高效可靠的发送大量数据。）
    
    - 拥塞控制：防止过多的数据注入到网络中，这样可以使网络中的路由器或链路不致过载。出现网络拥塞会导致大量数据丢失。
        - 拥塞控制的机制（四种算法）：慢开始和拥塞避免结合，快重传和快恢复。
        - 如何判断对包
            - 定时器超时
            - 收到三次ack
            
###. 三次握手 四次挥手
[参见](https://mp.weixin.qq.com/s/NL7Jzh0lYoA395yzaGxBHw)
    - 三次握手
    - 四次挥手
### DNS域名解析的过程？

### TCP和UDP的区别？

### 如何保证TCP链接的稳定性

### 打开一个URL的过程
[参考](https://mp.weixin.qq.com/s/I6BLwbIpfGEJnxjDcPXc1A)
### TCP 半连接队列和全连接队列满了会发生什么？又该如何应对？
[参见](https://mp.weixin.qq.com/s/tRXlq1hErqKQLMMLcxoXvg)