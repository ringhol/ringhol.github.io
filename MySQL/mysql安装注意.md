---
dg-home: false
dg-publish: true
---
## Centos7安装mysql
```bash
wget http://dev.mysql.com/get/mysql57-community-release-el7-8.noarch.rpm
```
```bash
rpm --import https://repo.mysql.com/RPM-GPG-KEY-mysql-2022
```
```bash
yum install mysql-community-server
```
```bash
grep 'temporary password' /var/log/mysqld.log
```
修改密码策略以使用简单密码
```sql
set global validate_password_policy=0;
```
密码最短长度
```sql
set global validate_password_length=4;
```
## 初始化
>mysqld --initialize

## 修改密码
>alter user USER() identified by 'password';
>alter user 'root'@'localhost' password expire never;
>flush privileges;

mariadb修改密码
```bash
mysql_secure_installation
```
## 开启远程登录
```sql
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'passwd'
```

## 新建用户指定数据库
```sql
grant all privileges on student.* to test3@localhost identified by ’123456′;
```