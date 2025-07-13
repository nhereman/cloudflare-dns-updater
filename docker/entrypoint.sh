while true; do
  for config_file in ${CDU_CONFIGURATION_FILES}; do
    echo "INFO: Executing cloudflare-dns-updater with configuration file ${config_file}"
    if ! /home/cloudflare-dns-updater -c "${config_file}"; then
      echo "ERROR: cloudflare-dns-updater failed for configuration file: ${config_file}."
    fi
  done
  echo "---"
  sleep "${CDU_EXEC_INTERVAL}"
done
