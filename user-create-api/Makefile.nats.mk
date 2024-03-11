NATS=localhost:4222
CREATE_USER_QUEUE_SUBJECT=user.create


nats-user-create:  ## Send a message to the localstack queue
	nats pub -s localhost:4222 user.create '{"name": "William haha", "password": "123", "role_id": 9, "email": "haha@teste.com"}'
