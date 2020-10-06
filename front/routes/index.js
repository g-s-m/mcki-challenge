var express = require('express');
var router = express.Router();

const {dtoReport, getReports} = require('../reports/reports');

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

module.exports = router;
