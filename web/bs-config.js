var proxy = require('http-proxy-middleware');

var helloProxy = proxy('/hello', {
    target: 'http://localhost:8080'
});

module.exports = {
    server: {
        middleware: {
            1: helloProxy
        }
    },
    files: [
        './dist/*'
    ],
    watchOptions: {
        awaitWriteFinish: true
    }
};
