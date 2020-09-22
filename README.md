# kafka

## step1 
- 启动zookeeper  
`$ bin/zookeeper-server-start.sh config/zookeeper.properties &`  
- 关闭zookeeper  
`$ bin/zookeeper-server-stop.sh`   
- 启动kafka  
`$ bin/kafka-server-start.sh config/server.properties &`
- 关闭kafka  
`$ bin/kafka-server-stop.sh`
- 创建topic  
`$ bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test`
- 查看所有topic  
`$ bin/kafka-topics.sh --list --bootstrap-server localhost:9092`
- 查看某个topic具体信息   
`$ bin/kafka-topics.sh --describe --bootstrap-server localhost:9092 --topic test`
- 删除topic (可直接删除的前提：delete.topic.enable=true)  
`$ bin/kafka-topics.sh --delete --bootstrap-server localhost:9092 --topic test`
- 生产消息  
`$ bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test`
- 消费消息  
`$ bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning`
