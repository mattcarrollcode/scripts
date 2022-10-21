import { Octokit, App } from "octokit"

const ACCESS_TOKEN = ""

if (ACCESS_TOKEN.length == 0 ) {
  console.log("add access token")
  console.log("to generate one go here: https://github.com/settings/tokens/new?scopes=repo")
  process.exit(1)
}

const octokit = new Octokit({
  auth: ACCESS_TOKEN
})

console.log('downloading subscription data...')
let page = 1
let res = {data: [1]}

while ( res.data.length > 0 ) {
  let res = await octokit.request('GET /notifications', {page: page})
  if (res.status != 200) {
    console.log(res)
    console.log("couldn't get sub data")
    process.exit(1)
  }
  for ( const sub of res.data ) {
    let res = await octokit.request(`DELETE /notifications/threads/${sub.id}/subscription`, {
      thread_id: sub.id
    })
    if (res.status != 204) {
      console.log(res)
      console.log("couldn't delete sub")
      process.exit(1)
    }
  }
  page++
  console.log(`deleted ${page*50} subs`)
}

console.log(`done.`)
process.exit(0)
