const http = require('http');
const hostname = '127.0.0.1';
const port = 3000;

const server = http.createServer((req, res) => {

	res.statusCode = 200;
	res.setHeader('Content-Type', 'application/json');
	
	console.log(req.method + ' ' + req.url + ' HTTP/' + req.httpVersion);
    for (var property in req.headers) {
		if (req.headers.hasOwnProperty(property)) {
            console.log(property + ': ' + req.headers[property])
        }
    }
	req.on('data', chunk => {
		if(req.method == "POST"){
			console.log('A chunk of data has arrived: ', JSON.parse(chunk));
		}
	});
	
	res.end('Hello World');
});


server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});