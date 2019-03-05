package main

import (
	"fmt"
	"net/http"
	"os"
)

var page = `<html lang="en"><head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<!--  -->
		<meta property="og:title" content="OMG Is Dean Available Right Now?">
		<meta property="og:description" content="The question you've been asking yourself.">
		<!-- -->
		<title>OMG Is Dean Available Right Now?</title>
		<link href="https://fonts.googleapis.com/css?family=Lato:900,300" rel="stylesheet" type="text/css">
		<style>
			body {
				font-family: 'Lato', sans-serif;
				text-shadow: 2px 2px #222;
				color: #EBEBEB;
				background-color: #15171A;
				background-repeat: repeat;
			}

			#main-div {
				margin-top: 5em;
			}

			.center {
				margin-left: auto;
				margin-right: auto;
				width: 15em;
				text-align: center;
			}

			#question {
				font-size: 20pt;
				font-weight: 300;
			}

			#answer {
				margin-top: 3em;
				font-size: 50pt;
				font-weight: 900;
			}
		</style>
	</head>
	<body>
		<div id="main-div">
			<div id="question" class="center">OMG Is Dean Available Right Now?</div>
			<div id="answer" class="center">YES</div>
		</div>
	</body>
</html>`

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, page)
	})
	http.ListenAndServe(":"+port, nil)
}
