import puppeteer from 'puppeteer';
import * as fs from 'fs';
import * as path from 'path';

// Expected JSON File form:
// [
//     { 
//         url: "www.example.com",
//         title: "file-name-of-pdf",
//         folder: "name-of-folder" 
//     },
//     ...
// ]
const JSON_FILE_PATH = "test-input.json"

const PAGE_WIDTH = "1200"
const PAGE_HEIGHT = "800"

async function main() {
    if (JSON_FILE_PATH == "") {
        console.log("Add the filepath to the value of the `JSON_FILE_PATH` variable in `index.js`")
        process.exit(1)
    }
    console.log("Launching browser...")
    const browser = await puppeteer.launch({
        headless: false,
        args: [`--window-size=${PAGE_WIDTH},${PAGE_HEIGHT}`],
        defaultViewport: {
            width: parseInt(PAGE_WIDTH, 10),
            height: parseInt(PAGE_HEIGHT, 10)
        }
    });
    const pages = await browser.pages()
    const page = pages[0] // Get first tab in open browser

    // In case what you need to download is behind a login form
    console.log("Login in the browser and then press any key to continue...")
    await keypress()
    console.log("Key pressed.")

    let rawdata = fs.readFileSync(JSON_FILE_PATH);
    let webpages = JSON.parse(rawdata);
    for (const webpage of webpages) {
        console.log(`Downloading ${webpage.url}`);

        // Navigate to page, wait until all network traffic stops
        await page.goto(webpage.url, { waitUntil: 'networkidle2', networkIdleTimeout: 5000 });
        // Scroll through the page to ensure all content loads
        await autoScroll(page);
        // Save PDF
        const filename = `${webpage.title}.pdf`
        const filePath = path.join(webpage.folder, filename);
        const pdfConfig = {
            path: filePath, // Saves file to this location
            format: 'A4',
            width: `${PAGE_WIDTH}px`,
            height: `${PAGE_HEIGHT}px`
        };
        await page.pdf(pdfConfig);
    }
    await browser.close();
    console.log('Done.')
    process.exit(0)
}


const keypress = async () => {
    process.stdin.setRawMode(true)
    return new Promise(resolve => process.stdin.once('data', () => {
        process.stdin.setRawMode(false)
        resolve()
    }))
}

async function autoScroll(page) {
    await page.evaluate(async () => {
        await new Promise((resolve) => {
            var totalHeight = 0;
            var distance = 100;
            var timer = setInterval(() => {
                var scrollHeight = document.body.scrollHeight;
                window.scrollBy(0, distance);
                totalHeight += distance;

                if (totalHeight >= scrollHeight - window.innerHeight) {
                    clearInterval(timer);
                    resolve();
                }
            }, 100);
        });
    });
}

main()
