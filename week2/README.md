##  Go 进阶训练营作业  第一周

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答：这个Error不应该抛给上层，如果Get查找的行不存在，通过布尔值告诉上层，所查找的值不存在。





