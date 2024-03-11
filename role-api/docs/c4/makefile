has_java:=$(shell which java)

setup:
	npm i -g c4builder

new-project:
	c4builder new &&
	cd project &&
	c4builder

config:
	cd zpe-systems && npm run config

watch:
	cd zpe-systems && npm run watch

build:
	cd zpe-systems && npm run build

install:
	@if [ "${has_java}" = "" ]; then\
        echo "install java before setup";\
		exit 1;\
    fi
	npm install

config-standalone:
	cd zpe-systems && c4builder config

watch-standalone:
	cd zpe-systems && c4builder site --watch

build-standalone:
	cd zpe-systems && c4builder