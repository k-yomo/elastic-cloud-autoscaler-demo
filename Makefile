

.PHONY: setup
setup:
	make check-required-env
	cd terraform && terraform init
	cd terraform && terraform apply
	make set-env-vars

.PHONY: cleanup
cleanup:
	make check-required-env
	cd terraform && terraform destroy

.PHONY: check-required-env
check-required-env:
ifndef EC_API_KEY
$(error env variable EC_API_KEY is not set)
endif

.PHONY: set-env-vars
set-env-vars:
	echo "DEPLOYMENT_ID=$(shell cd terraform && terraform output demo_deployment_id)" > .env
	echo "ELASTICSEARCH_CLOUD_ID=$(shell cd terraform && terraform output demo_elasticsearch_cloud_id)" >> .env
	echo "ELASTICSEARCH_URL=$(shell cd terraform && terraform output demo_elasticsearch_url)" >> .env
	echo "ELASTICSEARCH_USERNAME=$(shell cd terraform && terraform output demo_elasticsearch_username)" >> .env
	echo "ELASTICSEARCH_PASSWORD=$(shell cd terraform && terraform output demo_elasticsearch_password)" >> .env
	echo "MONITORING_ELASTICSEARCH_CLOUD_ID=$(shell cd terraform && terraform output demo_monitoring_elasticsearch_cloud_id)" >> .env
	echo "MONITORING_ELASTICSEARCH_USERNAME=$(shell cd terraform && terraform output demo_monitoring_elasticsearch_username)" >> .env
	echo "MONITORING_ELASTICSEARCH_PASSWORD=$(shell cd terraform && terraform output demo_monitoring_elasticsearch_password)" >> .env

.PHONY: index-products
index-products:
	cd scripts/index_products && gzip -d products.csv.gz > products.csv
	cd scripts/index_products && go run .

.PHONY: run
run:
	go run .

.PHONY: loadtest
loadtest:
	 cd k6 && k6 run script.js
