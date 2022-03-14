# kafka

## 1、目录
- kafka kafka启动脚本

## 备注
- 1、本来想着，写个shell脚本一键启动，但是环境变量有点问题，懒得整了，主要目标还是代码对kafka进行操作。
- 2、先简单启动kafka然后使用kafka-go包对其进行主题创建，生产以及消费
- 3、遇到的问题，kafka在M1芯片的Linux版本存在一些问题，基于自己的学习情况，先跳过kafka，转而使用RabbitMQ