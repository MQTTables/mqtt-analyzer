const request = new XMLHttpRequest();
const url = "loadall";
 
request.open("GET", url, true); 
request.addEventListener("readystatechange", () => {

    if(request.readyState === 4 && request.status === 200) {       
		document.getElementById('packetsRoot').innerHTML = request.responseText;
    }
});
request.send();