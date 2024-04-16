FROM bitnami/kafka:3.7.0-debian-12-r2
# Download MongoDB&reg; Connector for Apache Kafka https://www.confluent.io/hub/mongodb/kafka-connect-mongodb
# RUN mkdir -p /opt/bitnami/kafka/plugins && \
#     cd /opt/bitnami/kafka/plugins && \
#     curl --remote-name -o "mongo-kafka-connect-1.2.0-all.jar" --location https://search.maven.org/remotecontent?filepath=org/mongodb/kafka/mongo-kafka-connect/1.2.0/mongo-kafka-connect-1.2.0-all.jar
# ADD https://search.maven.org/remotecontent?filepath=org/mongodb/kafka/mongo-kafka-connect/1.2.0/mongo-kafka-connect-1.2.0-all.jar /bitnami/kafka/plugins/mongo-kafka-connect-1.2.0-all.jar

RUN mkdir -p /bitnami/kafka/plugins

# ADD https://search.maven.org/remotecontent?filepath=org/apache/kafka/connect-transforms/3.7.0/connect-transforms-3.7.0.jar ./bitnami/kafka/plugins/connect-transforms-3.7.0.jar

COPY /deployments/debezium-connector-mongodb ./opt/bitnami/kafka/libs

CMD /opt/bitnami/kafka/bin/connect-standalone.sh /bitnami/kafka/config/connect-standalone.properties
#  /bitnami/kafka/config/mongodb.properties