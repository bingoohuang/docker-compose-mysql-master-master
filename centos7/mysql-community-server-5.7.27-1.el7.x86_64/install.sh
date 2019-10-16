#!/usr/bin/env bash

yum install -y libaio*
cp my.cnf /etc/my.cnf
rpm -ivh *.rpm --nodeps --force
chown mysql:mysql /var/lib/mysql
systemctl daemon-reload
systemctl enable mysqld
systemctl start mysqld
grep 'temporary password' /var/lib/mysql/mysqld.log
password=$(grep -oP 'temporary password(.*): \K(\S+)' /var/lib/mysql/mysqld.log)
newpwd=$(grep -oP 'RootPassword\s*=\s*"\K[^"]+' ../mysqlclusterinit/config.toml)
mysqladmin --user=root --password="$password" password "$newpwd"