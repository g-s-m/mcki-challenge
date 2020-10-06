const http = require('http')
const addr = process.env.BACKEND_ADDR
const b_port = process.env.BACKEND_PORT

function dtoReport(report) {
  return {
    messageId : report.id,
    messageState : report.state,
    messageType : report.payload.reportType,
    messageText : report.payload.message
  }
}

function getReports(handler) {
     http.get('http://'+addr+':'+b_port+'/reports', (res) => {
        const { statusCode } = res;
        const contentType = res.headers['content-type'];

        let error;
        if (statusCode !== 200) {
          error = new Error('Request Failed.\n' +
                            `Status Code: ${statusCode}`);
        } else if (!/^application\/json/.test(contentType)) {
          error = new Error('Invalid content-type.\n' +
                            `Expected application/json but received ${contentType}`);
        }
        if (error) {
          console.error(error.message);
          res.resume();
          return;
        }

        res.setEncoding('utf8');
        let rawData = '';
        res.on('data', (chunk) => { rawData += chunk; });
        res.on('end', () => {
          try {
            const parsedData = JSON.parse(rawData);
            handler(parsedData)
          } catch (e) {
            console.error(e.message);
          }
        }).on('error', (e) => {
          console.error(`Got error: ${e.message}`);
        })
      });
}

module.exports = {
    dtoReport : dtoReport,
    getReports : getReports
}
