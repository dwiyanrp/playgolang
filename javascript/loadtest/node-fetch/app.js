const fetch = require('node-fetch');
const interval = require('interval-promise')

var url = 'http://localhost:3000/'
var method = 'GET'
var totalUsers = 1

function start() {
    for (var i = 0; i < totalUsers; i++) {
        spawnUser()
	}
}

function spawnUser() {
	interval(async() => {
		await getData(url)
	}, 0)
}

const getData = async url => {
	try {
	  const response = await fetch(url);
	  const text = await response.text();
	  console.log(text);
	} catch (error) {
	  console.log(error);
	}
};

start()