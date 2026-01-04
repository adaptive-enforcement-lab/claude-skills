MARKER=".completed-$RUN_ID"
[ -f "$MARKER" ] && exit 0
# Do work...
touch "$MARKER"