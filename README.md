# post-tweets

`post-tweets` checks your Twitter account every minute for new tweets, and
posts them to your [App.net](https://app.net) account.

## Installation

In order to compile an executable `post-tweets` binary, make sure that you
have checked out the source
(`git checkout https://github.com/KlausTrainer/post-tweets.git`) and that
you have a [Go compiler](http://golang.org/doc/install) available. Having
`post-tweets.go` in your current shell working directory, building it is as
simple as entering the following command:

	go build post-tweets.go

When successful, you should have a binary named `post-tweets`. A neat feature
of it is that (like every compiled package produced by the Go compiler) it is
statically linked and does not depend on any shared libraries at runtime. You
therefore can just go ahead and copy it to for instance some server and run it
there, provided that you have the same type of processor architecture and
operating system there.

## Usage

If you run `post-tweets` without argments or e.g. `post-tweets --help`, you
will get a brief usage description, like this one: 

	./post-tweets --screen_name=SCREEN_NAME --access_token=ACCESS_TOKEN [--since_id=SINCE_ID]

A sample usage of `post-tweets` would be to invoke it like this:

	./post-tweets --screen_name=KlausTrainer --access_token=AQAAAAAAAHwBjIwUzmYWt0l4X89yg_6-Np_oAjT3bdFcx5cObVZSeY-1mcsL3PfdwxEU_qqagOcSbQM6eI5Yu3q3arbd9-v63B --since_id=241575522692366336

Make sure that you enter your appropriate values for `SCREEN_NAME`,
`ACCESS_TOKEN`, and `SINCE_ID`. The easy part of it is `SCREEN_NAME`, which is
your Twitter user name with the '@'-character omitted (e.g., instead of
`@KlausTrainer` just put `KlausTrainer`).

The optional `SINCE_ID` value allows you to make sure that only tweets
newer than a certain one are posted to App.net. In order to get a tweet's ID
that you can use as `SINCE_ID` value, you can just follow the time link in its
top-right border on the Twitter website. The tweet ID is always the number
after the last slash in a tweet's URI. For instance the ID of
[this tweet](https://twitter.com/KlausTrainer/status/241575522692366336) is
`241575522692366336`.

Regarding the `ACCESS_TOKEN` value: In order to keep it short here (read:
*in order to maintain my laziness*), I won't elaborate on how to get an
App.net access token that allows `post-tweets` to post tweets to
your App.net account. Feel free to send a pull requests!
