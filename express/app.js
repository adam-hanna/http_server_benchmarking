var express = require('express');
var app = express();

app.set('view engine', 'jade');

app.get('/', function(req, res) {
    res.render('index', { name: 'World' } );
});

app.listen(3000, function() {
    console.log('Listening on port 3000');
});
