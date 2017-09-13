Redirector
==========

A simple webserver that only redirects traffic based on a YAML configuration
file.


## Configuration

A sample configuration file is available in config.yaml. Each value
can also be overridden with an environment variable.

```yaml
---
bind_address: 0.0.0.0:8080
redirects:
  - rules.yaml
```

The bind address is the address and port that the HTTP server should listen to.
The value can also be set by setting the `BIND_ADDRESS` environment variable.

The redirects array is evaluated in the order that files appear in the list.
Files can be local or remote (http and https only for now). The environment
variable `REDIRECT_SOURCE` can also be used, with multiple values separated
by commas.


## Redirect Rules

Each redirect rule has a name, a match pattern, and one or more targets.

To redirect blahblahblah.com (with an optional leading www) to
somethingdifferent.com using a 301 redirect you would use
```yaml
- name: Blah Test
  pattern: ^(www\.)?blahblahblah.com\/
  targets:
  - target: https://somethingdifferent.com/
    code: 301
```

### By Source IP

You can get more complex and redirect based on the source IP address (as
reported in the X-Forwareded-For header) too:
```yaml
- name: TV
  pattern: ^tv\.university\.edu/
  targets:
  - target: https://contour.university.edu
    code: 302
    when:
      source:
        - 10.0.0.0/8
        - 192.168.1.0/24
  - target: https://google.com
    code: 302
```

The above example essentially says redirect tv.university.edu to
contour.university.edu when inside one of the two source IP ranges, otherwise
redirect to google.com. Whenever multiple targets are used you should use
302 redirects to avoid browsers caching the redirects.

### Substitutions

Because everything is stored as regular expressions, substitutions are also
possible:
```yaml
- name: All to SSL
  pattern: ^(www\.)?(?P<name>[a-zA-Z0-9\-]+).com\/
  targets:
  - target: https://${name}.university.edu/
    code: 301
```

### Fall Back

Regular expressions are evaluated in the order they appear in the configuration
file. You should consider having some sort of catch all resource at the bottom
to prevent dead end 404s.

### Reloading

Redirect rules can be reloaded without restarting the application by sending
an HTTP GET request to `/reload`.
