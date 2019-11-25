var fRequest = new XMLHttpRequest();
var fUrl = "files";
 
fRequest.open("GET", fUrl, true); 
fRequest.addEventListener("readystatechange", () => {
    if(fRequest.readyState === 4 && fRequest.status === 200) {       
		document.getElementById("filesRoot").innerHTML = fRequest.responseText;
    }
});
fRequest.send();