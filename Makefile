.PHONY: diff
include .env

# atlas migrate
migrate:
	atlas schema apply \
          -u ${DB_URL} \
          --to file://schema.sql \
          --dev-url "docker://postgres/15/dev"

# atlas diff
diff:
	atlas migrate diff $(NAME) \
		  --dir "file://migrations" \
		  --to file://schema.sql \
		  --dev-url "docker://postgres/15/dev"

#atlas rollback
rollback:
	atlas schema rollback \
		  -u ${DB_URL} \
		  --to file://schema.sql \
		  --dev-url "docker://postgres/15/dev"
