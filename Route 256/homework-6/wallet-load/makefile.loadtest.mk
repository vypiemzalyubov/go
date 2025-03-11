loadtest-images:
	docker pull grafana/grafana:latest
	docker pull victoriametrics/victoria-metrics:latest
	docker pull prom/node-exporter:latest
	docker pull opensearchproject/opensearch:2.4.0
	docker pull graylog/graylog:5.0
	docker pull gitlab-registry.ozon.dev/qa/classroom-14/students/ws-perfomance-testing:stable

loadtest: dir=$dir
loadtest:
	docker run \
		--rm -it \
		--network wallet_ompnw \
		-v $(CURDIR)/loadtests/${dir}:/var/loadtest \
		gitlab-registry.ozon.dev/qa/classroom-14/students/ws-perfomance-testing:stable \
		'yandex-tank -f load.yaml'

loadtest-http:
	make loadtest dir=http-pool

loadtest-grpc:
	make loadtest dir=grpc-pool

loadtest-http-grpc:
	make loadtest dir=http-grpc-pools
