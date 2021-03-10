# Mysql

## 插入数据

insert into 表名 (字段名称1,字段名称2) values (值1,值2)

注意点:

*  *在插入数据的时候指定的字段名称的顺序不用和表中的字段名称的顺序一致*
  * *insert into stu (id, name) values (1, 'lnj');*
* *在插入数据的时候指定的取值顺序必须和指定的字段名称顺序一致*
  * *insert into stu (name, id) values ('zs', 2);*
* *如果插入数据时指定的取值顺序和表中的字段顺序是一致的, 那么可以不指定字段名称*
  * *insert into stu values (3, 'ls');*
* *我们可以通过values同时插入多条数据*
  * *insert into stu values (4, 'ww'), (5, 'zl');*



## 更新数据

*update 表名 set 字段名称=值 [where 条件];*

注意点:

* *如果在更新数据的时候没有指定条件, 那么就会更新整张表中的数据*
  * *update stu set score=77;*
* *如果在更新数据的时候指定了条件, 那么只会更新满足条件的数据*
  * *update stu set score=88 where name='ls';*
* *在指定条件的时候, 我们可以通过AND来指定多个条件, AND===&&*
  * *update stu set score=100 where name='lnj' AND id=5;*
* *在指定条件的时候, 我们可以通过OR来指定多个条件, OR===||*
  * *update stu set score=66 where name='zs' OR name='ww';*
* *在更新数据的时候是可以同时更新多个字段的*
  * *update stu set name='it666', score=33 where id=5;*



## 查询数据

*select 字段名称1, 字段名称2 from 表名 [where 条件];*

> *# 查询满足条件的数据*
>
> *select \* from stu where score > 60;*
>
> *select id, name from stu where score > 60;*
>
> *select \* from stu where score = 77 || score = 88;*
>
> *select \* from stu where score in (77, 88);*
>
> *select \* from stu where score BETWEEN 77 AND 88;*
>
> *select \* from stu where score IS NOT NULL;*
>
> *select \* from stu where score IS NULL;*



> *where支持的运算符*
>
> *=（等于）、!=（不等于）、<>（不等于）、<（小于）、<=（小于等于）、>（大于）、>=（大于等于）；*
>
> *IN(set)；固定的范围值*
>
> *BETWEEN…AND；值在什么范围*
>
> *IS NULL；（为空） IS NOT NULL（不为空）*
>
> *AND；与*
>
> *OR；或*
>
> *NOT；非*
>
> *LIKE: 模糊查询*

# 删除数据

*delete from 表名 [where 条件];*

> *# 删除满足条件的数据*
>
> *delete from stu where score > 60;*

> *$ 删除所有的数据*

> *delete from stu;*

# 模糊查询

`like` 

[注意的地方](https://www.cnblogs.com/huangliang-hb/p/10048666.html)