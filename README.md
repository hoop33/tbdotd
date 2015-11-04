# Tech Books Deals of the Day

> Shows deals from various sites (primarily book publishers)

This is the source code for the Tech Books Deals of the Day site -- that site is currently using <https://github.com/hoop33/techbooksdotd>, but will switch to this soon.

It's written in [Go](http://www.golang.org/) using the [Revel Framework](https://revel.github.io/).

## To Run

To run:

* Clone the repo into your go path
* Install Revel
* Install the Revel command
* Run the app

Here are the commands:

```
$ go get github.com/hoop33/tbdotd
$ go get github.com/revel/revel
$ go get github.com/revel/cmd/revel
$ revel run github.com/hoop33/tbdotd
```

Open a browser to <http://localhost:9000>.

### The Cache

The app caches the results to avoid hitting the various sources repeatedly, using the following approach:

* For Daily deals with a "publish" timestamp, expire 24 hours after publish time
* For Weekly deals with a "publish" timestamp, expire 7 days after publish time
* For all others, cache results after 8AM for the rest of the day

To clear the cache, add a URL parameter: `clearCache=true`

## Contributing

We love contributions. Fork the project, make the changes, push to your fork, and open pull requests.

## License

Tech Books Deals of the Day is released under the [MIT license](http://hoop33.mit-license.org/license).
