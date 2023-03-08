---
dg-home: false
dg-publish: true
---
#java #mysql

---
## 介绍
MySQL主从复制是一个异步的复制过程，底层是基于MySQL数据库自带的**二进制日志**功能。就是一台或多台MySQL数据库（slave，即**从库**）从另一台MySQL数据库（master，即**主库**）进行日志的幅值然后再解析日志并应用到自身，最终实现**从库**的数据和**主库**的数据保持一致。MySQL主从复制是MySQL数据库自带的功能，无需借助第三方工具。

**MySQL主从复制过程分成三步：**
- master将改变记录到二进制日志（binary log）
- slave将master的binary log拷贝到他的中继日志（relay log）
- slave重做中继日志的中继事件，将改变应用到自己的数据库中
![[Pasted image 20220928210852.png]]

## 配置-前置条件
提前准备两台服务器，分别安装MySQL并启动服务成功

## 配置-主库Master
1. 第一步：修改mysql数据库的配置文件/etc/my.cnf
```ini
[mysqld]
log-bin=mysql-bin    #[必须]启用二进制日志
server-id=100        #[必须]服务器唯一ID
```
![[Pasted image 20220929160126.png]]
2. 第二步：重启mysql数据库
3. 第三步：登录mysql数据库，执行以下SQL
```sql
GRANT REPLICATION SLAVE ON *.* to 'user'@'%' identified by 'password';
```
**注**：作用是创建一个用户user,密码为password，并且给user用户授予REPLICATION SLAVE权限。常用于建立复制时所需要用到的用户权限，也就是slave必须被master授权具有该权限的用户，才能通过该用户复制。
![[Pasted image 20220929161148.png]]
4. 第四步：登录mysql数据库，执行下面的SQL语句，记录下file和position的值
```sql
show master status;
```
![[Pasted image 20220929161345.png]]
## 配置-从库slave
1. 修改mysql数据库配置文件my.cnf
```ini
[mysql]
server-id=101    #[必须]服务器唯一id
```
2. 重启mysql服务
3. 登录mysql数据库，执行以下SQL
```sql
change master to
master_host='192.168.56.111',master_user='slave01',master_password='010107',master_log_file='mysql-bin.000001',master_log_pos=584;
start slave;
```
4. 查看从库数据库状态
```sql
show slave status;
```