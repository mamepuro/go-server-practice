# これは何か
DBの管理をする。一旦モノレポ構成だが、規模が大きくなったらリポジトリを切っても良いかもしれない

## How to use

以下のコマンドを叩く

```
docker-compose up
sql-migrate up -env local -config ./dbconfig.yml
```