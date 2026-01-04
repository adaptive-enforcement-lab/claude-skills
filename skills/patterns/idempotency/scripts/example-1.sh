# Idempotent: Running twice produces same result
mkdir -p /tmp/mydir    # Creates dir if missing, no-op if exists

# Not idempotent: Running twice fails or creates duplicates
mkdir /tmp/mydir       # Fails if directory exists