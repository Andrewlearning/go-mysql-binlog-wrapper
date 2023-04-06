# go-mysql-binlog-wrapper
A simple demo for how to connect to db and convert binlog into json format

# preparation
1. Check config path for your local mysql
```
mysql --help --verbose | grep my.cnf
/etc/my.cnf /etc/mysql/my.cnf /usr/local/mysql/etc/my.cnf ~/.my.cnf 
```

2. Turn on mysql binlog by add config
- `sudo vim /etc/my.cnf`
```shell
[mysqld]
log-bin = mysql-bin # turn on binlog
binlog-format = ROW # choose row format
server_id = 1 # define mysql replication，this can not same with canal slaveId
```

3. Check whether mysql binlog is turned on.
```
mysql> show variables like 'log_bin';
+---------------+-------+
| Variable_name | Value |
+---------------+-------+
| log_bin       | ON    |
+---------------+-------+
1 row in set (0.01 sec)
```

4. Setup your local mysql config in `const.go`, then run the code

# Reference
- code base
  - https://github.com/go-mysql-org/go-mysql#replication
- how to turn on local mysql server binlog
  - https://blog.csdn.net/zhengyong15984285623/article/details/73335646
