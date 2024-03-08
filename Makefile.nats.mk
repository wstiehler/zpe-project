NATS=localhost:4222
CREATE_PRODUCT_QUEUE_SUBJECT=user.create


nats-product-create:  ## Send a message to the localstack queue
	nats pub -s localhost:4222 user.create '{"name": "William Villani", "role": "admin", "email": "william@teste.com"}'
