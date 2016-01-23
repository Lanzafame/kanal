# Kanal Specification
Author: Adrian Lanzafame
Date: 17/01/2015
Revision: 1
Status: Raw

## Overview
Kanal is a data aggregation pipeline for the collection, transformation and sorting of various types of business, development, user tracking metrics.

## Language
The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [RFC 2119](http://tools.ietf.org/html/rfc2119).

## Scenarios
### Small Startup #1
Mr CEO of Unicorns 'R' Us wants to have monitors placed around the office that show key stats about each business unit. These statistics are disparate not only in type but in source. The developers want to know about server logs/metrics, new pull requests, code reviews to-do, etc. The sales team want to know about the amount of revenue they have in the sales pipeline, reminders of what features they are pushing this month. The client engagement team want to know the number of outstanding support tickets, the number of tickets that are in progress/being pushed out in the next release. And of course Mr CEO wants to put motivational (subliminal) messages (propaganda) on them. 

Mr Dev is **order** to _own_ this project. He finds plenty of paid services that provide what Mr CEO is after but **frugality**... So he goes in search of some Open Source Software (OSS) that he can piece together into some piecemeal variant of [Numerics by Cynapse][numerics], [Vital Statistics][vital], and [Geckoboard][gecko]. He comes across analytics dashboards such as [Keen IO's Analytics Dashboards][keenio] but he needs something that can gather and sort the disparate data into a more pliable form. Then he sees Kanal on Hacker News and it is the perfect solution ;).

## Non Goals
This version of Kanal will _not_ support the following features:
  - Any form of unified API for the retrieval of data from persistence.

## Components

### Trösker (_translation_: Harvester)
Collects data from third-party API endpoints as defined by kanal.toml. Trösker SHALL then pass the data onto the Kanal.

For each third-party API, the following interface MUST be implemented.

Trösker interface:
```go
type Trosker interface {
  Auth(pubkey string, seckey string)
  Get(q Query) *Response
  Timeout(t time.Time) bool
}
```

> - How will multiple tröskers be run concurrently?
> - Can there be multiple clusters of tröskers running concurrently?
> - How does a trösker handle 3xx/4xx/5xx?
> - Specifically, how does a trösker handle 429/401|403s?
> - [HTTP Status Codes][http]
> - Needs more information.

> Take inspiration from this [codewalk][sharemem].
> And [this][marcio]

### Kanal
Kanal is the in-memory queue between Trösker and the multitude of Gondolas

> - This section needs a lot of work.
> - **???** May not need a _queue queue_ here but just a `go channel`.

### Gondola
Gondola SHALL take a data packet off the queue, analyse the contents to determine whether any transformations are required before passing to the specified persistence endpoint (persistence mappings are defined in kanal.toml).

> More work required.

### Luka (_translation_: Harbor/Port)
Luka SHALL generate an API endpoint for each of the configured persistences types taken from the kanal.toml file.

> Again more work required.

### Configuration (kanal.toml)
The following elements MUST be configured in the kanal.toml:
  - Third-party API endpoints
  - Mappings of data type to persistence endpoints
  - Persistence endpoints

[numerics]: http://cynapse.com/numerics/
[vital]: http://vitalstatistics.net.au
[gecko]: http://geckoboard.com
[keenio]: https://github.com/keen/dashboards
[http]: https://en.wikipedia.org/wiki/List_of_HTTP_status_codes
[sharemem]: https://golang.org/doc/codewalk/sharemem/
[marcio]: http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/