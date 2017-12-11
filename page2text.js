const argv = require('minimist')(process.argv.slice(2));
var fs = require('fs');
const puppeteer = require('puppeteer');

const url = argv.url || 'https://www.baoquan.com';
const outtxt = argv.outtxt || `out.txt`;

puppeteer.launch().then(async browser => {
    const page = await browser.newPage();
    await page.goto(url);
    await page.content().then(function(value) {
        fs.writeFile(outtxt, value, {flag: 'a'}, function (err) {
        });
    });
    await page.close();
    await browser.close();
    //process.exit(0);
   // await page.takeScreenshot
}).then(process.exit(0))








