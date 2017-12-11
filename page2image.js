const Screenshot =require('page2image').default;
const argv = require('minimist')(process.argv.slice(2));

const url = argv.url || 'https://www.baoquan.com';
const output = argv.output || `output.jpg`;

const screenshot = new Screenshot({
    //waitFor:'.copyright',
    waitUntil: 'networkidle',
    waitForFunction: () => {
    window.imageList = window.imageList || Array.from(document.getElementsByTagName('img'));
return window.imageList.length <= window.imageList.reduce((loaded, imageElm) => (
    imageElm.complete ? loaded + 1 : loaded
), 0);
},
viewportConfig: { width: 1440, height: 768  },
screenshotConfig: { fullPage: true, path: output ,quality: 100 },
});
screenshot
    .takeScreenshot(url)
    .then(process.exit);
