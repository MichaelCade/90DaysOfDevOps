# Enterprise Search extension

Elastic Enterprise Search is a suite of products for search applications backed by the Elastic Stack.

## Requirements

* 2 GB of free RAM, on top of the resources required by the other stack components and extensions.

Enterprise Search exposes the TCP port `3002` for its Web UI and API.

## Usage

### Generate an encryption key

Enterprise Search requires one or more [encryption keys][enterprisesearch-encryption] to be configured before the
initial startup. Failing to do so prevents the server from starting.

Encryption keys can contain any series of characters. Elastic recommends using 256-bit keys for optimal security.

Those encryption keys must be added manually to the [`config/enterprise-search.yml`][config-enterprisesearch] file. By
default, the list of encryption keys is empty and must be populated using one of the following formats:

```yaml
secret_management.encryption_keys:
  - my_first_encryption_key
  - my_second_encryption_key
  - ...
```

```yaml
secret_management.encryption_keys: [my_first_encryption_key, my_second_encryption_key, ...]
```

> :information_source: To generate a strong encryption key, for example using the AES-256 cipher, you can use the
> OpenSSL utility or any other online/offline tool of your choice:
>
> ```console
> $ openssl enc -aes-256 -P
>
> enter aes-256-cbc encryption password: <a strong password>
> Verifying - enter aes-256-cbc encryption password: <repeat your strong password>
> ...
>
> key=<generated AES key>
> ```

### Enable Elasticsearch's API key service

Enterprise Search requires Elasticsearch's built-in [API key service][es-security] to be enabled in order to start.
Unless Elasticsearch is configured to enable TLS on the HTTP interface (disabled by default), this service is disabled
by default.

To enable it, modify the Elasticsearch configuration file in [`elasticsearch/config/elasticsearch.yml`][config-es] and
add the following setting:

```yaml
xpack.security.authc.api_key.enabled: true
```

### Configure the Enterprise Search host in Kibana

Kibana acts as the [management interface][enterprisesearch-ui] to Enterprise Search.

To enable the management experience for Enterprise Search, modify the Kibana configuration file in
[`kibana/config/kibana.yml`][config-kbn] and add the following setting:

```yaml
enterpriseSearch.host: http://enterprise-search:3002
```

### Start the server

To include Enterprise Search in the stack, run Docker Compose from the root of the repository with an additional command
line argument referencing the `enterprise-search-compose.yml` file:

```console
$ docker-compose -f docker-compose.yml -f extensions/enterprise-search/enterprise-search-compose.yml up
```

Allow a few minutes for the stack to start, then open your web browser at the address <http://localhost:3002> to see the
Enterprise Search home page.

Enterprise Search is configured on first boot with the following default credentials:

* user: *enterprise_search*
* password: *changeme*

## Security

The Enterprise Search password is defined inside the Compose file via the `ENT_SEARCH_DEFAULT_PASSWORD` environment
variable. We highly recommend choosing a more secure password than the default one for security reasons.

To do so, change the value `ENT_SEARCH_DEFAULT_PASSWORD` environment variable inside the Compose file **before the first
boot**:

```yaml
enterprise-search:

  environment:
    ENT_SEARCH_DEFAULT_PASSWORD: {{some strong password}}
```

> :warning: The default Enterprise Search password can only be set during the initial boot. Once the password is
> persisted in Elasticsearch, it can only be changed via the Elasticsearch API.

For more information, please refer to [User Management and Security][enterprisesearch-security].

## Configuring Enterprise Search

The Enterprise Search configuration is stored in [`config/enterprise-search.yml`][config-enterprisesearch]. You can
modify this file using the [Default Enterprise Search configuration][enterprisesearch-config] as a reference.

You can also specify the options you want to override by setting environment variables inside the Compose file:

```yaml
enterprise-search:

  environment:
    ent_search.auth.source: standard
    worker.threads: '6'
```

Any change to the Enterprise Search configuration requires a restart of the Enterprise Search container:

```console
$ docker-compose -f docker-compose.yml -f extensions/enterprise-search/enterprise-search-compose.yml restart enterprise-search
```

Please refer to the following documentation page for more details about how to configure Enterprise Search inside a
Docker container: [Running Enterprise Search Using Docker][enterprisesearch-docker].

## See also

[Enterprise Search documentation][enterprisesearch-docs]

[config-enterprisesearch]: ./config/enterprise-search.yml

[enterprisesearch-encryption]: https://www.elastic.co/guide/en/enterprise-search/current/encryption-keys.html
[enterprisesearch-security]: https://www.elastic.co/guide/en/workplace-search/current/workplace-search-security.html
[enterprisesearch-config]: https://www.elastic.co/guide/en/enterprise-search/current/configuration.html
[enterprisesearch-docker]: https://www.elastic.co/guide/en/enterprise-search/current/docker.html
[enterprisesearch-docs]: https://www.elastic.co/guide/en/enterprise-search/current/index.html
[enterprisesearch-ui]: https://www.elastic.co/guide/en/enterprise-search/current/user-interfaces.html

[es-security]: https://www.elastic.co/guide/en/elasticsearch/reference/current/security-settings.html#api-key-service-settings
[config-es]: ../../elasticsearch/config/elasticsearch.yml
[config-kbn]: ../../kibana/config/kibana.yml
