var express = require('express');
var router = express.Router();
const {dtoReport, getReports} = require('../reports/reports');

const b_addr = process.env.BACKEND_ADDR
const b_port = process.env.BACKEND_PORT
const http = require('http')

function reloadBackend() {
   console.log("reload")
   const options = {
      hostname: b_addr,
      port: b_port,
      path: '/reports/reload',
      method: 'POST',
      body: {}
   }
    const req = http.request(options, res => {
      res.on('data', d => {
        process.stdout.write(d)
      })
    })

    req.on('error', error => {
      console.error(error)
    })
    req.end()
}


function setReportState(id, state) {
    const data = JSON.stringify({
      ticketState: state
    })
    const options = {
      hostname: b_addr,
      port: b_port,
      path: `/reports/${id}`,
      method: 'PUT',
      headers : {
        'Content-Type': 'application/json',
        'Content-Length': data.length
      }
    }
    const req = http.request(options, res => {
      res.on('data', d => {
        process.stdout.write(d)
      })
    })

    req.on('error', error => {
      console.error(error)
    })

    req.write(data)
    req.end()
}

function resolveReport(id) {
    setReportState(id, 'RESOLVED')
}

function blockReport(id) {
    setReportState(id, 'BLOCKED')
}

/* GET home page. */
router.get('/', function(req, res, next) {
  getReports((report) => {
      let reportList = [];
      report.elements.forEach(function(item, i, arr) {
        reportList.push(dtoReport(item));
      });
    res.render('index', { 
      newListItems: reportList,
    });
  });
});

router.post('/block/:id', function(req, res, next) {
    blockReport(req.params.id);
    res.redirect('/');
});

router.post('/resolve/:id', function(req, res, next) {
    resolveReport(req.params.id);
    res.redirect('/');
});

router.get('/reload', function(req, res, next) {
    reloadBackend();
    res.redirect('/');
});


module.exports = router
