# cellphones

## build

docker build --tag backend:v1 -f ./deployments/Dockerfile .

## run docker

docker run --name backend -d -p 30000:8080 backend:v1 run -s 1

## get list of images in local registry

curl -X GET "http://localhost:5001/v2/\_catalog"

## get tags of an image

curl -X GET "http://localhost:5001/v2/<image>/tags/list"

sha256:6bde8e70a73153b22430ef2f547f3eb651491dcde52ba7d10fc9258ee5c00aed

## delete image from local registry

### list all images in the repository

curl -X GET "http://localhost:5001/v2/<repository-name>/tags/list

### find image's digest

curl -sS -H "Accept: application/vnd.docker.distribution.manifest.v2+json" -o /dev/null -w "%header{Docker-Content-Digest}" http://localhost:5001/v2/<repository>/manifests/latest

### delete images's manifest

curl -X DELETE "http://localhost:5001/v2/<repository-name>/manifests/<tag>"

k label nodes kind-control-plane name=node1

###

MongoDB&reg; can be accessed on the following DNS name(s) and ports from within your cluster:

    mongodb-0.mongodb-headless.default.svc.cluster.local:27017
    mongodb-1.mongodb-headless.default.svc.cluster.local:27017

To get the root password run:

    export MONGODB_ROOT_PASSWORD=$(kubectl get secret --namespace default mongodb -o jsonpath="{.data.mongodb-root-password}" | base64 -d)

To get the password for "user1" run:

    export MONGODB_PASSWORD=$(kubectl get secret --namespace default mongodb -o jsonpath="{.data.mongodb-passwords}" | base64 -d | awk -F',' '{print $1}')

To get the password for "user2" run:

    export MONGODB_PASSWORD=$(kubectl get secret --namespace default mongodb -o jsonpath="{.data.mongodb-passwords}" | base64 -d | awk -F',' '{print $2}')

To connect to your database, create a MongoDB&reg; client container:

    kubectl run --namespace default mongodb-client --rm --tty -i --restart='Never' --env="MONGODB_ROOT_PASSWORD=$MONGODB_ROOT_PASSWORD" --image localhost:5001/bitnami/mongodb:7.0.3-debian-11-r6 --command -- bash

Then, run the following command:
mongosh admin --host "mongodb-0.mongodb-headless.default.svc.cluster.local:27017,mongodb-1.mongodb-headless.default.svc.cluster.local:27017" --authenticationDatabase admin -u $MONGODB_ROOT_USER -p $MONGODB_ROOT_PASSWORD

bootstrap.servers = {{ include "common.names.fullname" . }}-controller-0.{{ include "common.names.fullname" . }}-controller-headless.{{ include "common.names.namespace" . }}.svc.{{ .Values.clusterDomain }}:{{ .Values.service.ports.client }},{{ include "common.names.fullname" . }}-controller-1.{{ include "common.names.fullname" . }}-controller-headless.{{ include "common.names.namespace" . }}.svc.{{ .Values.clusterDomain }}:{{ .Values.service.ports.client }},{{ include "common.names.fullname" . }}-controller-2.{{ include "common.names.fullname" . }}-controller-headless.{{ include "common.names.namespace" . }}.svc.{{ .Values.clusterDomain }}:{{ .Values.service.ports.client }}

{{ include "common.names.fullname" . }}-controller-0.{{ include "common.names.fullname" . }}-controller-headless.{{ include "common.names.namespace" . }}.svc.{{ .Values.clusterDomain }}:{{ .Values.service.ports.client }}

###

** Please be patient while the chart is being deployed **

Kafka can be accessed by consumers via port 9092 on the following DNS name from within your cluster:

    kafka.default.svc.cluster.local

Each Kafka broker can be accessed by producers via port 9092 on the following DNS name(s) from within your cluster:

    kafka-controller-0.kafka-controller-headless.default.svc.cluster.local:9092
    kafka-controller-1.kafka-controller-headless.default.svc.cluster.local:9092
    kafka-controller-2.kafka-controller-headless.default.svc.cluster.local:9092

The CLIENT listener for Kafka client connections from within your cluster have been configured with the following security settings: - SASL authentication

To connect a client to your Kafka, you need to create the 'client.properties' configuration files with the content below:

security.protocol=SASL_PLAINTEXT
sasl.mechanism=SCRAM-SHA-256
sasl.jaas.config=org.apache.kafka.common.security.scram.ScramLoginModule required \
 username="user1" \
 password="$(kubectl get secret kafka-user-passwords --namespace default -o jsonpath='{.data.client-passwords}' | base64 -d | cut -d , -f 1)";

To create a pod that you can use as a Kafka client run the following commands:

    kubectl run kafka-client --restart='Never' --image localhost:5001/bitnami/kafka:3.7.0-debian-12-r2 --namespace default --command -- sleep infinity
    kubectl cp --namespace default /path/to/client.properties kafka-client:/tmp/client.properties
    kubectl exec --tty -i kafka-client --namespace default -- bash

    PRODUCER:
        kafka-console-producer.sh \
            --producer.config /tmp/client.properties \
            --broker-list kafka-controller-0.kafka-controller-headless.default.svc.cluster.local:9092,kafka-controller-1.kafka-controller-headless.default.svc.cluster.local:9092,kafka-controller-2.kafka-controller-headless.default.svc.cluster.local:9092 \
            --topic test

    CONSUMER:
        kafka-console-consumer.sh \
            --consumer.config /tmp/client.properties \
            --bootstrap-server kafka.default.svc.cluster.local:9092 \
            --topic test \
            --from-beginning

WARNING: There are "resources" sections in the chart not set. Using "resourcesPreset" is not recommended for production. For production installations, please set the following values according to your workload needs:

- controller.resources
  +info https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/

bootstrap.servers = {{ include "common.names.fullname" . }}-0.{{ include "common.names.fullname" . }}-headless.{{ include "common.names.namespace" . }}.svc.{{ .Values.clusterDomain }}:{{ .Values.service.port }}

WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

# See the License for the specific language governing permissions and

# limitations under the License.

# These are defaults. This file just demonstrates how to override some settings.

bootstrap.servers=localhost:9092

# The converters specify the format of data in Kafka and how to translate it into Connect data. Every Connect user will

# need to configure these based on the format they want their data in when loaded from or stored into Kafka

key.converter=org.apache.kafka.connect.json.JsonConverter
value.converter=org.apache.kafka.connect.json.JsonConverter

# Converter-specific settings can be passed in by prefixing the Converter's setting with the converter we want to apply

# it to

key.converter.schemas.enable=true
value.converter.schemas.enable=true

offset.storage.file.filename=/tmp/connect.offsets

# Flush much faster than normal, which is useful for testing/debugging

offset.flush.interval.ms=10000

# Set to a list of filesystem paths separated by commas (,) to enable class loading isolation for plugins

# (connectors, converters, transformations). The list should consist of top level directories that include

# any combination of:

# a) directories immediately containing jars with plugins and their dependencies

# b) uber-jars with plugins and their dependencies

# c) directories immediately containing the package directory structure of classes of plugins and their dependencies

# Note: symlinks will be followed to discover dependencies or plugins.

# Examples:

# plugin.path=/usr/local/share/java,/usr/local/share/kafka/plugins,/opt/connectors,

#plugin.path=

{
"name": "mongodb-connector",
"config": {
"connector.class": "io.debezium.connector.mongodb.MongoDbConnector",
"mongodb.connection.string": "mongodb://root:root@mongodb-0.mongodb-headless.default.svc.cluster.local:27017,mongodb-1.mongodb-headless.default.svc.cluster.local:27017/?replicaSet=rs0",
"topic.prefix": "stg",
}
}

curl -X POST -H "Content-Type: application/json" -d @create_source_connector.json localhost:8083/connectors

curl localhost:8083/connector-plugins

kafka-topics.sh --list --bootstrap-server kafka-controller-0.kafka-controller-headless.default.svc.cluster.local:9092

kafka-topics.sh --create --topic stg --partitions 1 --replication-factor 1 --bootstrap-server kafka-controller-0.kafka-controller-headless.default.svc.cluster.local:9092

kafka-console-consumer.sh --bootstrap-server kafka-controller-0.kafka-controller-headless.default.svc.cluster.local:9092 --topic mongodb.stg.blocks --from-beginning
