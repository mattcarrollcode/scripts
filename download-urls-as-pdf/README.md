# Download URLs as PDF

Note: This script scrolls down each page before downloading a PDF to ensure all assets are loaded (e.g. lazily loaded images on a page). This significantly increased page save times.

## Setup

### Node.js
1. Install nvm: https://github.com/nvm-sh/nvm#installing-and-updating
1. install node: `nvm install --lts`
1. instal depedencies: `cd` to this dir, `npm install`

### Input JSON
1. Create a JSON file that has the following form:
   ```json
   [
       { 
           url: "www.example.com",
           title: "file-name-of-pdf",
           folder: "name-of-folder" 
       },
       ...
   ]
   ```
1. Create all the folders specified in the JSON file
1. Add the filepath to the value of the `JSON_FILE_PATH` variable in `index.js`

## Usage
1. `node index.js`
1. [Optional] Login to any websites if the content is behind a login
1. Press any key to continue
1. Output will be in the folder you specified in the JSON with the filename `${title}.pdf`. You must create the folders before running this script.


## Testing
1. Add the to the value of the `JSON_FILE_PATH` variable in `index.js` to `test-input.json`
1. Run the script: `node index.js`
1. Observe a file called `example.pdf` that has the contents of the example.com webpage in it.
