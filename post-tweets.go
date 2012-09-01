package main

import "os"
import "fmt"
import "flag"
import "net/http"
import "io/ioutil"
import "encoding/json"
import "time"
import "bytes"

type Tweet struct {
	Id_Str string
	Text   string
}

func main() {
	var (
		since_id     string
		screen_name  string
		access_token string
		tweets       []Tweet
		err          error
	)

	flag.StringVar(&screen_name, "screen_name", "", "Twitter user name without '@'")
	flag.StringVar(&access_token, "access_token", "", "App.net access token")
	flag.StringVar(&since_id, "since_id", "1", "Tweet ID to start with")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage_string())
	}
	flag.Parse()

	if screen_name == "" || access_token == "" {
		flag.Usage()
		return
	}

	for {
		tweets, err = get_tweets(screen_name, since_id)
		if err == nil && len(tweets) > 0 {
			for i := len(tweets) - 1; i >= 0; i-- {
				tweet := tweets[i]
				if post_tweet(access_token, tweet) {
					since_id = tweet.Id_Str
				}
			}
		}

		time.Sleep(1 * time.Minute)
	}
}

func get_tweets(screen_name string, since_id string) (tweets []Tweet, err error) {
	var (
		response      *http.Response
		response_body []byte
	)

	uri := "http://api.twitter.com/1/statuses/user_timeline.json?screen_name=" + screen_name + "&include_rts=false&exclude_replies=true&since_id=" + since_id

	response, err = http.Get(uri)
	defer response.Body.Close()
	if err != nil {
		return tweets, err
	}
	response_body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return tweets, err
	}

	return tweets, json.Unmarshal(response_body, &tweets)
}

func post_tweet(access_token string, tweet Tweet) (success bool) {
	uri := "https://alpha-api.app.net/stream/0/posts?access_token=" + access_token
	post_body, _ := json.Marshal(map[string]string{"text": tweet.Text})

	response, err := http.Post(uri, "application/json", bytes.NewBuffer(post_body))
	defer response.Body.Close()

	if err != nil || response.StatusCode != 200 {
		return false
	}

	return true
}

func usage_string() string {
	return "Usage: " + os.Args[0] + " --screen_name=SCREEN_NAME " +
		"--access_token=ACCESS_TOKEN [--since_id=SINCE_ID]\n\n" +
		"Arguments:\n\t--screen_name=SCREEN_NAME\tTwitter user " +
		"name without '@'\n" + "\t--access_token=ACCESS_TOKEN\t" +
		"App.net access token\n" + "\t--since_id=SINCE_ID\t\t" +
		"only sync Tweets newer than SINCE_ID\n"
}
