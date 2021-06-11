'use strict'

const { PurgeCSS } = require('purgecss')

const express = require("express")
const multer  = require('multer')
const upload = multer()

const app = express()

app.post("/", upload.fields([{name: "html"}, {name: "css"}]), (req, res) => {
  const result = new PurgeCSS().purge({
    content: [
      {
        raw: req.files['html'][0].buffer.toString(),
        extension: 'html'
      }
    ],
    css: [
      {
        raw: req.files['css'][0].buffer.toString()
      }
    ]
  })
  result.then(result => {
    res.send(result[0].css)
  }).catch(error => {
    res.send(error)
  })
})

const PORT = process.env.PORT || 8080

app.listen(PORT, () => {
  console.log(`App listening on port ${PORT}`);
})

module.exports = app;
