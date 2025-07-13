# cloudflare-dns-updater

The tool `cloudflare-dns-updater` allows to automatically set your public IP address on a cloudflare DNS Record.

## Usage

```bash
./cloudflare-dns-updater -C <configuration file>
```

If the option `-C <configuration file>` is not provided, the tool will look for it in `<HOME DIRECTORY>/.config/cloudflare-dns-updater/config.json`.


## Configuration

The tool requires the following information from the configuration (configuration file or environment variables):

- `domain name`: The domain name of the DNS Record to update. It is used to validate that the DNS record being updated is the correct one.
- `DNS Record ID`: The ID of the DNS record to update.
- `e-mail address`: The email address of the cloudflare account.
- `Cloudflare API key`: The API key of the cloudflare account. 
- `Cloudflare zone ID`: The ID of the zone of the cloudflare account.

### Configuration file

The configuration file should be in json format and look like that: 

```json
{
  "domain": "<domain name>",
  "record": "<DNS Record ID>",
  "email": "<e-mail address>",
  "api_key": "<Cloudflare API key>",
  "zone": "<Cloudflare zone ID>"
}
```

If any of these entries is configured using environment variable, it can be removed from the configuration file. If everything is configured using environment variables, the configuration file should still exist and look like that:

```json
{}
```
### Environment variables

The following environment variables can be used:

- `CDU_DOMAIN_NAME`: domain name
- `CDU_DNS_RECORD_ID`: DNS Record ID
- `CDU_CF_EMAIL`: email address
- `CDU_CF_API_KEY`: Cloudflare API key
- `CDU_CF_ZONE_ID`: Cloudflare zone ID


## Docker image

A docker image is available to run this tool on multiple DNS records every 60 seconds (by default).

The name of the image is: `nhereman/cloudflare-dns-updater:1.0.0`

To start it execute the following command:

```bash
docker exec --rm -d nhereman/cloudflare-dns-updater:1.0.0 -v <path to configuration file>:/home/config.json
```

You can defines: the following environment variables can be defined in the docker container:

- `CDU_EXEC_INTERVAL`: Change the interval between two executions of the tool. Default: `60`
- `CDU_CONFIGURATION_FILES`: List of the configuration files to use separated by spaces (one by DNS record). Default: `/home/config.json`

## Helm chart

An helm chart is available to deploy the tool on kubernetes.

In order to use it, go in the helm subdirectory and create the file `cloudflare-dns-updater/values.yaml`. This file should be filed this way with your value:

```yaml
cloudflare:
  zone: # Cloudflare zone ID
  records: # List of DNS records to update
    - id: # DNS record ID
      name: # DNS record domain name
    - id:
      name:
  secret:
    email: # Cloudflare account email address in base64
    apiKey: # Cloudflare account api key in base64
interval: # Interval between two executions of the tool. Default to 60 if not provided
```

Then execute the following command to deploy it:

```bash
helm upgrade --install cloudflare-dns-updater cloudflare-dns-updater
```
