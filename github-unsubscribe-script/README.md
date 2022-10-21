# Github script to unsubscribe to all subscriptions

In large orginizations you can be autosubscribed to hundreds of thousands of notifications. There is no bulk or easy way to mass unsubsribe. Even the API to unsubscribe is not bulk. As such this script can take O(days) to run for acounts with ~million subs.

## Setup

### Pre-setup
* go to https://github.com/settings/tokens/new?scopes=repo to get an access token
* make sure to enable the `notification` scope
* add the access token to `ACCESS_TOKEN` in `src/main.ts`
* install nvm and node16

### install

* `npm install`
* `node main.js`

## notes
* this takes forever, you have to delete each sub individually (no bulk API)
  * both APIs used (list notification and delete notification) use the "core"
    API rate limit of 80 calls/min (1.3/s), can delete a max ~80/min or 5000/hr[0].
    Some active Github accounts have 100,000's of subs which would take O(days).
    I had over 100,000.
* uses [octokit](https://github.com/octokit/octokit.js#usage)
* Uses list notification API: https://docs.github.com/en/rest/activity/notifications#list-notifications-for-the-authenticated-user
  * this API don't respect the "per page" parameter and only returns up to
    50 subs per call
* Uses delete notification API: https://docs.github.com/en/rest/activity/notifications#delete-a-thread-subscription

[0](https://stackoverflow.com/questions/13394077/is-there-a-way-to-increase-the-api-rate-limit-or-to-bypass-it-altogether-for-git)
