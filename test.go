package main

import (
	_ "fmt"
)

func main() {
	var delimEnds = [...]string{
		3: `"`,
		9: "'",
		// Determined empirically by running the below in various browsers.
		// var div = document.createElement("DIV");
		// for (var i = 0; i < 0x10000; ++i) {
		//   div.innerHTML = "<span title=x" + String.fromCharCode(i) + "-bar>";
		//   if (div.getElementsByTagName("SPAN")[0].title.indexOf("bar") < 0)
		//     document.write("<p>U+" + i.toString(16));
		// }
		1: " \t\n\f\r>",
	}
	//fmt.Println(delimEnds)
	print(delimEnds)

}
