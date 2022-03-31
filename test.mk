test_postgres_backup:
	./.build/rkc backups generate backup \\
		--definition=cmd/backups/generate/test_data/examples/postgres.yaml
