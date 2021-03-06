TODO: Intro

## New Features!

### CLI/Options: Add `--tag` flag and `tags` option to set test-wide tags (#553)

You can now specify any number of tags on the command line using the `--tag NAME=VALUE` flag. You can also use the `tags` option to the set tags in the code.

The specified tags will be applied across all metrics. However if you have set a tag with the same name on a request, check or custom metric in the code that tag value will have precedence.

Thanks to @antekresic for their work on this!

**Docs**: [Test wide tags](https://docs.k6.io/v1.0/docs/tags-and-groups#section-test-wide-tags) and [Options](https://docs.k6.io/v1.0/docs/options#section-available-options)

### k6/http: Support for HTTP NTLM Authentication (#556)

```js
import http from "k6/http";
import { check } from "k6";

export default function() {
    // Passing username and password as part of URL plus the auth option will authenticate using HTTP Digest authentication
    let res = http.get("http://user:passwd@example.com/path", {auth: "ntlm"});

    // Verify response
    check(res, {
        "status is 200": (r) => r.status === 200
    });
}
```

**Docs**: [HTTP Params](http://k6.readme.io/docs/params-k6http)

### HAR converter: Add support for correlating JSON values (#516)

There is now support for correlating JSON values in recordings, replacing recorded request values with references to the previous response.

Thanks to @cyberw for their work on this!

### InfluxDB collector: Add support for sending certain sample tags as fields (#585)

Since InfluxDB indexes tags, highly variable information like `vu`, `iter` or even `url` may lead to high memory usage. The InfluxDB documentation [recommends](https://docs.influxdata.com/influxdb/v1.5/concepts/schema_and_data_layout/#encouraged-schema-design) to use fields in that case, which is what k6 does now. There is a new `INFLUXDB_TAGS_AS_FIELDS` option (`collectors.influxdb.tagsAsFields` in the global k6 JSON config) that specifies which of the tags k6 emits will be sent as fields to InfluxDB. By default that's only `url` (but not `name`), `vu` and `iter` (if enabled).

Thanks to @danron for their work on this!


## UX

* Clearer error message when using `open` function outside init context (#563)
* Better error message when a script or module can't be found (#565). Thanks to @antekresic for their work on this!

## Internals

* Removed all httpbin.org usage in tests, now a local transient HTTP server is used instead (#555). Thanks to @mccutchen for the great [go-httpbin](https://github.com/mccutchen/go-httpbin) library!
* Fixed various data races and enabled automated testing with `-race` (#564)

## Bugs
* Archive: archives generated on Windows can now run on *nix and vice versa. (#566)
