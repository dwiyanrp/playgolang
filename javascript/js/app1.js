let iyem = require('iyem')

// let slowSquareThread = iyem.create(function(){  
//     var i = 0; 
//     var n = 100000;
//        while (++i < n * n) {}
//        $.finish(i);
//    })
//    slowSquareThread.start()
//    slowSquareThread.onFinish(function(result,err){
//     console.log(result,err) 
//    })
//    function slowSquare(n){
//     var i = 0;  
//     while (++i < n * n) {}
//    }
//    slowSquare(100)

// function slowSquare(n){
//     var i = 0;  
//     while (++i < n * n) {}
// }
// slowSquare(100000)

// let funfun = iyem.create(function(i){
//     var i = 0

//     setTimeout(
//         function() {
//             if(i++ > 5){
//                 return;
//             }
//         }, 1000
//     )
//     $.finish(i);
// });

// funfun.start();
// funfun.onFinish(function(result,err){
//     console.log(result, err)
// })

let funfun = iyem.create(() => {
	var i = 0;
	setInterval(() => {
		if(++i > 10){
			$.finish(i);
		}
		console.log(i)
	}, 500)
});

funfun.start();