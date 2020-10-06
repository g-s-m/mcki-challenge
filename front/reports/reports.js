const fs = require('fs')

function dtoReport(report) {
  return {
    messageId : report.id,
    messageState : report.state,
    messageType : report.payload.reportType,
    messageText : report.payload.message
  }
}

//todo: read reports from backend 
function getReports(handler) {
    let report = {}
      fs.readFile('reports.json', (err, data) => {
        if (err) throw err;
        handler(JSON.parse(data));
      })
}

module.exports = {
    dtoReport : dtoReport,
    getReports : getReports
}
