has_java:=$(shell which java)

setup:
	npm i -g c4builder

new-project:
	c4builder new &&
	cd project &&
	c4builder

config:
	cd alerts-publisher && npm run config

watch:
	cd alerts-publisher && npm run watch

build:
	cd alerts-publisher && npm run build

install:
	@if [ "${has_java}" = "" ]; then\
        echo "install java before setup";\
		exit 1;\
    fi
	npm install

config-standalone:
	cd alerts-publisher && c4builder config

watch-standalone:
	cd alerts-publisher && c4builder site --watch

build-standalone:
	cd alerts-publisher && c4builder