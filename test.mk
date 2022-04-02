test_postgres_backup:
	./.build/rkc backups generate backup \
		--definition=cmd/backups/generate/test_data/examples/postgres.yaml \
		--template postgres

test_postgres_backup_k8s:
	./.build/rkc backups generate backup \
		--definition=cmd/backups/generate/test_data/examples/postgres.yaml \
		--template postgres \
		--kubernetes \
		--gpg-key-path cmd/backups/generate/test_data/examples/gpg.key

test_postgres_backup_k8s_sealed_secret:
	./.build/rkc backups generate backup \
		--definition=cmd/backups/generate/test_data/examples/postgres.yaml \
		--template postgres \
		--kubernetes \
		--gpg-key-path cmd/backups/generate/test_data/examples/valid-sealed-secret.yaml
