#### postgresql 创建用户和数据库
```psql
create user chy  with password 'welcome1';
create database galaxy owner chy;
grant all privileges on database galaxy to chy;
```