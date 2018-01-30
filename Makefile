start-docker: ## Starts the docker containers for local development.
	@echo Starting docker containers

	@if [ $(shell docker ps -a | grep -ci xserver-mysql) -eq 0 ]; then \
		echo starting xserver-mysql; \
		docker run --name xserver-mysql -p 65500:3306 -v ${PWD}/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=alltheday \
		-e MYSQL_USER=xserver_user -e MYSQL_PASSWORD=xserver_password -e MYSQL_DATABASE=xserver_dev_database -d mysql:5.7 > /dev/null; \
	elif [ $(shell docker ps | grep -ci xserver-mysql) -eq 0 ]; then \
		echo restarting xserver-mysql; \
		docker start xserver-mysql > /dev/null; \
	fi

	@if [ $(shell docker ps -a | grep -ci xserver-redis) -eq 0 ]; then \
		echo starting xserver-redis; \
		docker run --name xserver-redis -p 65501:6379 -d redis > /dev/null; \
	elif [ $(shell docker ps | grep -ci xserver-redis) -eq 0 ]; then \
		echo restarting xserver-redis; \
		docker start xserver-redis > /dev/null; \
	fi

stop-docker: ## Stops the docker containers for local development.
	@echo Stopping docker containers

	@if [ $(shell docker ps -a | grep -ci xserver-mysql) -eq 1 ]; then \
		echo stopping xserver-mysql; \
		docker stop xserver-mysql > /dev/null; \
	fi

	@if [ $(shell docker ps -a | grep -ci xserver-redis) -eq 1 ]; then \
		echo stopping xserver-redis; \
		docker stop xserver-redis > /dev/null; \
	fi

clean-docker: ## Deletes the docker containers for local development.
	@echo Removing docker containers

	@if [ $(shell docker ps -a | grep -ci xserver-mysql) -eq 1 ]; then \
		echo removing xserver-mysql; \
		docker stop xserver-mysql > /dev/null; \
		docker rm -v xserver-mysql > /dev/null; \
	fi

	@if [ $(shell docker ps -a | grep -ci xserver-redis) -eq 1 ]; then \
		echo removing xserver-redis; \
		docker stop xserver-redis > /dev/null; \
		docker rm -v xserver-redis > /dev/null; \
	fi

