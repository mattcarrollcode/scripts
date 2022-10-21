import { Octokit, App } from "octokit"

const ACCESS_TOKEN = "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"

if (ACCESS_TOKEN.length == 0 ) {
  console.log("add access token")
  console.log("to generate one go here: https://github.com/settings/tokens/new?scopes=repo")
  process.exit(1)
}

const octokit = new Octokit({
  auth: ACCESS_TOKEN
})

let res = await octokit.request('GET /notifications', {per_page: 100, page: 1})
console.log(res)