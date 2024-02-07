.PHONY: diff
include .env
# atlas migrate
migrate:
	atlas schema apply \
          -u $(DB_URL) \
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
		  -u "postgresql://postgres:34AB5gA5636443FE4Egc3-cGE-4*DC-G@monorail.proxy.rlwy.net:26753/railway" \
		  --to file://schema.sql \
		  --dev-url "docker://postgres/15/dev"
