while true; do
  if ! /home/cloudflare-dns-updater -c "${CDU_CONFIGURATION_FILE}"; then
    echo "ERROR: cloudflare-dns-updater failed."
  fi
  echo "---"
  sleep "${CDU_EXEC_INTERVAL}"
done
