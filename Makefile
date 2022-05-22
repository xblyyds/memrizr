.PHONY: keypair migrate-create migrate-up migrate-down migrate-force

PWD = $(shell pwd)
ACCTPATH = $(PWD)/account
MPATH = $(ACCTPATH)/migrations
PORT = 5432

# Default number of migrations to execute up or down
N = 1

create-keypair:
	@echo "Creating an rsa 256 key pair"
	openssl genpkey -algorithm RSA -out $(ACCTPATH)/rsa_private_$(ENV).pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in $(ACCTPATH)/rsa_private_$(ENV).pem -pubout -out $(ACCTPATH)/rsa_public_$(ENV).pem

migrate-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir $(MPATH) -seq -digits 5 $(NAME)

# 第一个root为账号名称，第二个root为数据库名称
migrate-up:
	migrate -source file://$(MPATH) -database postgres://root:123456@localhost:$(PORT)/root?sslmode=disable up $(N)

migrate-down:
	migrate -source file://$(MPATH) -database postgres://root:123456@localhost:$(PORT)/root?sslmode=disable down $(N)

migrate-force:
	migrate -source file://$(MPATH) -database postgres://root:123456@localhost:$(PORT)/root?sslmode=disable force $(VERSION)