package main

func screens(i int) string {
	if i == 1 {
		return `
		<html window-icon="./sciter.png">
			<head>
			</head>
			<body>
				<h1> This is second page </h1>			
			</body>		
		</html>
	`
	}

	return `
		<html window-icon="./sciter.png">
			<head>
			</head>
			<body>
				<h1> No Html Files Need Any More </h1>
				<button #myname> Change page </button>
			</body>
			<script type="text/tiscript">			
				self#myname.on("click",function(){							
					view.changePage()			
					view.msgbox("Callling for Page Change")
				})
			</script>
		</html>
	`
}
